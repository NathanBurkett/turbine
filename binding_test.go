package turbine_test

import (
	"github.com/nathanburkett/turbine"
	"testing"
)

func TestNewBinding(t *testing.T) {
	type args struct {
		name       string
		resolution interface{}
	}
	type assertion struct {
		isSingleton bool
		isFactory   bool
	}
	tests := []struct {
		name string
		args args
		want assertion
	}{
		{
			name: "Type Singleton",
			args: args{
				name:       "foo",
				resolution: "bar",
			},
			want: assertion{
				isSingleton: true,
				isFactory: false,
			},
		},
		{
			name: "Type Factory",
			args: args{
				name: "foo",
				resolution: func(c *turbine.Container) interface{} {
					return "bar"
				},
			},
			want: assertion{
				isSingleton: false,
				isFactory: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := turbine.NewBinding(tt.args.name, tt.args.resolution); got.IsSingleton() != tt.want.isSingleton || got.IsFactory() != tt.want.isFactory {
				t.Errorf("NewBinding() = %v, want isSingleton %v and isFactory %v", got, tt.want.isSingleton, tt.want.isFactory)
			}
		})
	}
}

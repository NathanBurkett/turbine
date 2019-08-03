package turbine_test

import (
	"github.com/nathanburkett/turbine"
	"reflect"
	"testing"
)

func TestContainer_Set(t *testing.T) {
	type args struct {
		name string
		item interface{}
	}
	tests := []struct {
		name    string
		c       *turbine.Container
		args    []args
		wantErr bool
	}{
		{
			name: "Handles setting successfully",
			c:    &turbine.Container{},
			args: []args{
				{
					name: "foo",
					item: "bar",
				},
				{
					name: "bar",
					item: "baz",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			for _, arg := range tt.args {
				val := tt.c.Set(arg.name, arg.item)

				if val != nil {
					err = val
				}
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Container.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContainer_Get(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name     string
		c        *turbine.Container
		args     args
		wantItem interface{}
		wantOk   bool
	}{
		{
			name: "Handles getting successfully",
			c: turbine.New(false, map[string]interface{}{
				"foo": "bar",
			}),
			args: args{
				name: "foo",
			},
			wantItem: "bar",
			wantOk:   true,
		},
		{
			name: "Handles getting successfully but item does not exist",
			c: turbine.New(false, map[string]interface{}{
				"foo": "bar",
			}),
			args: args{
				name: "bar",
			},
			wantItem: nil,
			wantOk:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotItem, gotOk := tt.c.Get(tt.args.name)
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("Container.Get() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Container.Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

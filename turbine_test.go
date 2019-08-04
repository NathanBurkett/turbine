package turbine_test

import (
	"github.com/nathanburkett/turbine"
	"reflect"
	"testing"
)

func TestContainer_IsStrict(t *testing.T) {
	tests := []struct {
		name string
		c    *turbine.Container
		want bool
	}{
		{
			name: "strict is false",
			c:    turbine.New(false, turbine.BindingMap{}),
			want: false,
		},
		{
			name: "strict is true",
			c:    turbine.New(true, turbine.BindingMap{}),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := tt.c.IsStrict(); result != tt.want {
				t.Errorf("Container.IsStrict() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestContainer_Has(t *testing.T) {
	type args struct {
		name string
	}

	tests := []struct {
		name string
		c    *turbine.Container
		args args
		want bool
	}{
		{
			name: "Has item should be true",
			c: turbine.New(false, turbine.BindingMap{
				"foo": {
					Name: "foo",
					Resolution: func (c *turbine.Container) interface{} {
						return "foo"
					},
				},
			}),
			args: args{
				name: "foo",
			},
			want: true,
		},
		{
			name: "Container with no items and Has(name) call should be false",
			c:    turbine.New(false, turbine.BindingMap{}),
			args: args{
				name: "foo",
			},
			want: false,
		},
		{
			name: "Container with items but no item with name should be false",
			c: turbine.New(false, turbine.BindingMap{
				"foo": {
					Name: "foo",
					Resolution: "bar",
				},
			}),
			args: args{
				"bar",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val := tt.c.Has(tt.args.name); val != tt.want {
				t.Errorf("Container.Has(\"%s\") = %v, want %v", tt.args.name, val, tt.want)
			}
		})
	}
}

func TestContainer_Set(t *testing.T) {
	type args struct {
		b turbine.Binding
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
					b: turbine.Binding{
						Name: "foo",
						Resolution: "bar",
					},
				},
				{
					b: turbine.Binding{
						Name: "foo",
						Resolution: func (c *turbine.Container) interface{} {
							return "bar"
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error

			for _, arg := range tt.args {
				val := tt.c.Set(arg.b)

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
			c: turbine.New(false, turbine.BindingMap{
				"foo": turbine.NewBinding("foo", "bar"),
			}),
			args: args{
				name: "foo",
			},
			wantItem: "bar",
			wantOk:   true,
		},
		{
			name: "Handles getting but item does not exist",
			c: turbine.New(false, turbine.BindingMap{
				"foo": turbine.NewBinding("foo", "bar"),
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

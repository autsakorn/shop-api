package utils

import (
	"reflect"
	"testing"
)

func TestTransformQueryGetAll(t *testing.T) {
	type args struct {
		queryStr string
	}
	tests := []struct {
		name      string
		args      args
		wantQuery map[string]string
		wantErr   bool
	}{
		{
			"Base case",
			args{"Attribute:Value"},
			map[string]string{"Attribute": "Value"},
			false,
		},
		{
			"Key with dot, Object__Attribute",
			args{"Object.Attribute:Value"},
			map[string]string{"Object__Attribute": "Value"},
			false,
		},
		{
			"Error case",
			args{"Attribute:Value,Detail"},
			map[string]string{"Attribute": "Value"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuery, err := TransformQueryGetAll(tt.args.queryStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransformQueryGetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotQuery, tt.wantQuery) {
				t.Errorf("TransformQueryGetAll() = %v, want %v", gotQuery, tt.wantQuery)
			}
		})
	}
}

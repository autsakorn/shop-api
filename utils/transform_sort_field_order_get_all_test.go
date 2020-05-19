package utils

import (
	"reflect"
	"testing"
)

func TestTransFormSortFieldOrderGetAll(t *testing.T) {
	type args struct {
		sortbyStr string
		orderStr  string
	}
	tests := []struct {
		name        string
		args        args
		wantOrderby []string
		wantErr     bool
	}{
		{"Pass case", args{"", ""}, []string{}, false},
		{"Pass case desc", args{"Attr1", "desc"}, []string{"-Attr1"}, false},
		{"Pass case asc", args{"Attr1", "asc"}, []string{"Attr1"}, false},
		{"Error invalid order", args{"Attr1", "Asc"}, []string{}, true},
		{"Error invalid order", args{"Attr1,Attr2", "asc,Desc"}, []string{}, true},
		{"Pass case 2 field, 1 order", args{"Attr1,Attr2", "asc"}, []string{"Attr1", "Attr2"}, false},
		{"Pass case 2 field, 1 order", args{"Attr1,Attr2", "desc"}, []string{"-Attr1", "-Attr2"}, false},
		{"Error case 2 field, 1 order", args{"Attr1,Attr2", "esc"}, []string{}, true},
		{"Error case 2 field, 3 order", args{"Attr1,Attr2", "asc, asc,asc"}, []string{}, true},
		{"Error case 3 field, 2 order", args{"Attr1,Attr2,Attr3", "asc,asc"}, []string{}, true},
		{"Error case 0 field, 1 order", args{"", "asc"}, []string{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOrderby, err := TransFormSortFieldOrderGetAll(tt.args.sortbyStr, tt.args.orderStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransFormSortFieldOrderGetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOrderby, tt.wantOrderby) {
				t.Errorf("TransFormSortFieldOrderGetAll() = %v, want %v", gotOrderby, tt.wantOrderby)
			}
		})
	}
}

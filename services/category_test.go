package services

import (
	"errors"
	"reflect"
	"shop-api/models"
	"shop-api/storage"
	"shop-api/storage/mock"
	"shop-api/types"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewCategoryService(t *testing.T) {
	tests := []struct {
		name  string
		wantS CategoryService
	}{
		{"Base case", CategoryService{Storage: storage.NewStorage()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := NewCategoryService(); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("NewCategoryService() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestCategoryService_Add(t *testing.T) {
	type fields struct {
		Storage storage.Storage
	}
	type args struct {
		input *types.InputAddCategory
	}
	type mockResponse struct {
		ID  int64
		Err error
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		mockStatus       int32
		mockResponse     mockResponse
		wantResponseCode int
		wantErr          bool
	}{
		{
			"Base case",
			fields{Storage: storage.Storage{}},
			args{
				input: &types.InputAddCategory{
					Name:   "Name",
					Detail: "Detail",
					Status: "Active",
				},
			},
			1,
			mockResponse{1, nil},
			types.ResponseCode["CreatedSuccess"],
			false,
		},
		{
			"Invalid status case",
			fields{Storage: storage.Storage{}},
			args{
				input: &types.InputAddCategory{
					Name:   "Name",
					Detail: "Detail",
					Status: "active",
				},
			},
			-1,
			mockResponse{1, nil},
			types.ResponseCode["BadRequest"],
			true,
		},
		{
			"Storage Add Fail",
			fields{Storage: storage.Storage{}},
			args{
				input: &types.InputAddCategory{
					Name:   "Name",
					Detail: "Detail",
					Status: "Inactive",
				},
			},
			0,
			mockResponse{0, errors.New("Add Fail")},
			types.ResponseCode["BadRequest"],
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mackCategory := mock.NewMockCategory(ctrl)
			mackCategory.EXPECT().
				Add(&models.Category{
					Name:   tt.args.input.Name,
					Detail: tt.args.input.Detail,
					Status: tt.mockStatus,
				}).
				AnyTimes().
				Return(tt.mockResponse.ID, tt.mockResponse.Err)
			tt.fields.Storage.Category = mackCategory
			s := CategoryService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, err := s.Add(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponseCode != tt.wantResponseCode {
				t.Errorf("CategoryService.Add() = %v, want %v", gotResponseCode, tt.wantResponseCode)
			}
		})
	}
}

func TestCategoryService_Delete(t *testing.T) {
	type fields struct {
		Storage storage.Storage
	}
	type args struct {
		id int64
	}
	type mockResponse struct {
		Num int64
		Err error
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		mockResponse     mockResponse
		wantResponseCode int
		wantErr          bool
	}{
		{
			"Base case",
			fields{Storage: storage.Storage{}},
			args{1},
			mockResponse{1, nil},
			types.ResponseCode["Success"],
			false,
		},
		{
			"Not found case",
			fields{Storage: storage.Storage{}},
			args{100},
			mockResponse{0, nil},
			types.ResponseCode["BadRequest"],
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mackCategory := mock.NewMockCategory(ctrl)
			mackCategory.EXPECT().
				Delete(&models.Category{
					ID: tt.args.id,
				}).
				AnyTimes().
				Return(tt.mockResponse.Num, tt.mockResponse.Err)
			tt.fields.Storage.Category = mackCategory
			s := CategoryService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, err := s.Delete(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponseCode != tt.wantResponseCode {
				t.Errorf("CategoryService.Delete() = %v, want %v", gotResponseCode, tt.wantResponseCode)
			}
		})
	}
}

func TestCategoryService_GetByID(t *testing.T) {
	type fields struct {
		Storage storage.Storage
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantResponseCode int
		wantResult       types.OutputCategory
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CategoryService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, gotResult, err := s.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponseCode != tt.wantResponseCode {
				t.Errorf("CategoryService.GetByID() gotResponseCode = %v, want %v", gotResponseCode, tt.wantResponseCode)
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("CategoryService.GetByID() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestCategoryService_GetAll(t *testing.T) {
	type fields struct {
		Storage storage.Storage
	}
	type args struct {
		query  map[string]string
		fields []string
		sortby []string
		order  []string
		offset int64
		limit  int64
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantResponseCode int
		wantResults      []types.OutputCategory
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CategoryService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, gotResults, err := s.GetAll(tt.args.query, tt.args.fields, tt.args.sortby, tt.args.order, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponseCode != tt.wantResponseCode {
				t.Errorf("CategoryService.GetAll() gotResponseCode = %v, want %v", gotResponseCode, tt.wantResponseCode)
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("CategoryService.GetAll() gotResults = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func TestCategoryService_UpdateByID(t *testing.T) {
	type fields struct {
		Storage storage.Storage
	}
	type args struct {
		id       int64
		category *types.InputUpdateCategory
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantResponseCode int
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CategoryService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, err := s.UpdateByID(tt.args.id, tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.UpdateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponseCode != tt.wantResponseCode {
				t.Errorf("CategoryService.UpdateByID() = %v, want %v", gotResponseCode, tt.wantResponseCode)
			}
		})
	}
}

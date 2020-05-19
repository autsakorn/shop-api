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
		id  int64
		err error
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		mockStatus       int32
		mockResponse     mockResponse
		wantResponseCode int
		wantID           int64
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
			1,
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
			0,
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
			0,
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
				Return(tt.mockResponse.id, tt.mockResponse.err)
			tt.fields.Storage.Category = mackCategory
			s := CategoryService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, gotID, err := s.Add(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponseCode != tt.wantResponseCode {
				t.Errorf("CategoryService.Add() gotResponseCode = %v, want %v", gotResponseCode, tt.wantResponseCode)
			}
			if gotID != tt.wantID {
				t.Errorf("CategoryService.Add() gotID = %v, want %v", gotID, tt.wantID)
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
		num int64
		err error
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
				Return(tt.mockResponse.num, tt.mockResponse.err)
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
	type mockResponse struct {
		result models.Category
		err    error
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		mockResponse     mockResponse
		wantResponseCode int
		wantResult       types.OutputCategory
		wantErr          bool
	}{
		{
			"Status Active",
			fields{Storage: storage.Storage{}}, args{1},
			mockResponse{models.Category{Name: "Name", Detail: "Detail", Status: 1}, nil},
			types.ResponseCode["Success"],
			types.OutputCategory{Name: "Name", Detail: "Detail", StatusRes: "Active"},
			false,
		},
		{
			"Status Inactive",
			fields{Storage: storage.Storage{}}, args{2},
			mockResponse{models.Category{Name: "Name", Detail: "Detail", Status: 0}, nil},
			types.ResponseCode["Success"],
			types.OutputCategory{Name: "Name", Detail: "Detail", StatusRes: "Inactive"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mackCategory := mock.NewMockCategory(ctrl)
			mackCategory.EXPECT().
				GetByID(tt.args.id).
				AnyTimes().
				Return(tt.mockResponse.result, tt.mockResponse.err)
			tt.fields.Storage.Category = mackCategory
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
		order  []string
		offset int64
		limit  int64
	}
	type mockResponse struct {
		result []models.Category
		err    error
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		mockResponse     mockResponse
		wantResponseCode int
		wantResults      []types.OutputCategory
		wantErr          bool
	}{
		{
			"Status Active",
			fields{Storage: storage.Storage{}},
			args{},
			mockResponse{
				[]models.Category{
					{Name: "Name", Detail: "Detail", Status: 1},
				},
				nil,
			},
			types.ResponseCode["Success"],
			[]types.OutputCategory{
				{Name: "Name", Detail: "Detail", StatusRes: "Active"},
			},
			false,
		},
		{
			"Status Inactive",
			fields{Storage: storage.Storage{}},
			args{query: map[string]string{"Name": "Mas"}},
			mockResponse{
				[]models.Category{
					{Name: "Name", Detail: "Detail", Status: 0},
				},
				nil,
			},
			types.ResponseCode["Success"],
			[]types.OutputCategory{
				{Name: "Name", Detail: "Detail", StatusRes: "Inactive"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mackCategory := mock.NewMockCategory(ctrl)
			mackCategory.EXPECT().
				GetAll(tt.args.query, tt.args.order, tt.args.offset, tt.args.limit).
				AnyTimes().
				Return(tt.mockResponse.result, tt.mockResponse.err)
			tt.fields.Storage.Category = mackCategory
			s := CategoryService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, gotResults, err := s.GetAll(tt.args.query, tt.args.order, tt.args.offset, tt.args.limit)
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
	type mockResponse struct {
		num int64
		err error
	}
	type args struct {
		id       int64
		category *types.InputUpdateCategory
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
			args{id: 1, category: &types.InputUpdateCategory{Name: "Name"}},
			mockResponse{num: 1, err: nil},
			types.ResponseCode["Success"],
			false,
		},
		{
			"Not found ID",
			fields{Storage: storage.Storage{}},
			args{id: 2, category: &types.InputUpdateCategory{Name: "Name"}},
			mockResponse{num: 0, err: nil},
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
				UpdateByID(&models.Category{
					Name: tt.args.category.Name,
					ID:   tt.args.id,
				}).
				AnyTimes().
				Return(tt.mockResponse.num, tt.mockResponse.err)
			tt.fields.Storage.Category = mackCategory
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

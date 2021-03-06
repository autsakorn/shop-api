package services

import (
	"context"
	"errors"
	"reflect"
	"shop-api/helper"
	ormmock "shop-api/helper/orm_mock"
	"shop-api/models"
	"shop-api/storage"
	categorymock "shop-api/storage/category_mock"
	"shop-api/types"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
)

func TestNewCategoryService(t *testing.T) {
	tests := []struct {
		name  string
		wantS CategoryService
	}{
		{"Base case", CategoryService{Storage: storage.NewStorage(), Orm: helper.NewOrm()}},
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
		name         string
		fields       fields
		args         args
		mockStatus   int32
		mockResponse mockResponse
		wantID       int64
		wantErr      bool
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
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ctx := context.Background()
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()

			// Mock storage method
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockCategory := categorymock.NewMockCategory(ctrl)
			mockCategory.EXPECT().
				Add(
					ormer,
					&models.Category{
						Name:   tt.args.input.Name,
						Detail: tt.args.input.Detail,
						Status: tt.mockStatus,
					}).
				AnyTimes().
				Return(tt.mockResponse.id, tt.mockResponse.err)

			// Set properties serivce
			tt.fields.Storage.Category = mockCategory
			s := CategoryService{
				Storage: tt.fields.Storage,
				Orm:     ormMocked,
			}
			gotID, err := s.Add(ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
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
		name         string
		fields       fields
		args         args
		mockResponse mockResponse
		wantErr      bool
	}{
		{
			"Base case",
			fields{Storage: storage.Storage{}},
			args{1},
			mockResponse{1, nil},
			false,
		},
		{
			"Not found case",
			fields{Storage: storage.Storage{}},
			args{100},
			mockResponse{0, nil},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ctx := context.Background()
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()

			// Mock storage method
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockCategory := categorymock.NewMockCategory(ctrl)
			mockCategory.EXPECT().
				Delete(
					ormer,
					&models.Category{
						ID: tt.args.id,
					}).
				AnyTimes().
				Return(tt.mockResponse.num, tt.mockResponse.err)

			// Set properties serivce
			tt.fields.Storage.Category = mockCategory
			s := CategoryService{
				Storage: tt.fields.Storage,
				Orm:     ormMocked,
			}
			err := s.Delete(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
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
		name         string
		fields       fields
		args         args
		mockResponse mockResponse
		wantResult   types.OutputCategory
		wantErr      bool
	}{
		{
			"Status Active",
			fields{Storage: storage.Storage{}}, args{1},
			mockResponse{models.Category{Name: "Name", Detail: "Detail", Status: 1}, nil},
			types.OutputCategory{Name: "Name", Detail: "Detail", StatusRes: "Active"},
			false,
		},
		{
			"Status Inactive",
			fields{Storage: storage.Storage{}}, args{2},
			mockResponse{models.Category{Name: "Name", Detail: "Detail", Status: 0}, nil},
			types.OutputCategory{Name: "Name", Detail: "Detail", StatusRes: "Inactive"},
			false,
		},
		{
			"Not found ID",
			fields{Storage: storage.Storage{}}, args{3},
			mockResponse{models.Category{}, errors.New("No Row")},
			types.OutputCategory{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ctx := context.Background()
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()

			// Mock storage method
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockCategory := categorymock.NewMockCategory(ctrl)
			mockCategory.EXPECT().
				GetByID(ormer, tt.args.id).
				AnyTimes().
				Return(tt.mockResponse.result, tt.mockResponse.err)

			// Set properties serivce
			tt.fields.Storage.Category = mockCategory
			s := CategoryService{
				Storage: tt.fields.Storage,
				Orm:     ormMocked,
			}
			gotResult, err := s.GetByID(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
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
		name         string
		fields       fields
		args         args
		mockResponse mockResponse
		wantResults  []types.OutputCategory
		wantErr      bool
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
			[]types.OutputCategory{
				{Name: "Name", Detail: "Detail", StatusRes: "Active"},
			},
			false,
		},
		{
			"Status Inactive",
			fields{Storage: storage.Storage{}},
			args{query: map[string]string{"Name": "Mas", "Status": "Inactive"}},
			mockResponse{
				[]models.Category{
					{Name: "Name", Detail: "Detail", Status: 0},
				},
				nil,
			},
			[]types.OutputCategory{
				{Name: "Name", Detail: "Detail", StatusRes: "Inactive"},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ctx := context.Background()
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()

			// Mock storage method
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockCategory := categorymock.NewMockCategory(ctrl)
			mockCategory.EXPECT().
				GetAll(ormer, tt.args.query, tt.args.order, tt.args.offset, tt.args.limit).
				AnyTimes().
				Return(tt.mockResponse.result, tt.mockResponse.err)

			// Set properties serivce
			tt.fields.Storage.Category = mockCategory
			s := CategoryService{
				Storage: tt.fields.Storage,
				Orm:     ormMocked,
			}
			gotResults, err := s.GetAll(ctx, tt.args.query, tt.args.order, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
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
		num    int64
		result models.Category
		err    error
	}
	type args struct {
		id       int64
		category *types.InputUpdateCategory
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockStatus   int32
		mockResponse mockResponse
		wantErr      bool
	}{
		{
			"Base case",
			fields{Storage: storage.Storage{}},
			args{id: 1, category: &types.InputUpdateCategory{Name: "Name", Status: "Active"}},
			1,
			mockResponse{num: 1, result: models.Category{ID: 1, CreatedAt: time.Now()}, err: nil},
			false,
		},
		{
			"Invalid status",
			fields{Storage: storage.Storage{}},
			args{id: 1, category: &types.InputUpdateCategory{Name: "Name", Status: "active"}},
			1,
			mockResponse{num: 1, result: models.Category{ID: 1, CreatedAt: time.Now()}, err: nil},
			true,
		},
		{
			"Not found ID",
			fields{Storage: storage.Storage{}},
			args{id: 2, category: &types.InputUpdateCategory{Name: "Name", Status: "Active"}},
			1,
			mockResponse{num: 0, result: models.Category{}, err: nil},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ctx := context.Background()
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()

			// Mock storage method
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockCategory := categorymock.NewMockCategory(ctrl)
			mockCategory.EXPECT().
				GetByID(ormer, tt.args.id).
				AnyTimes().
				Return(tt.mockResponse.result, tt.mockResponse.err)
			mockCategory.EXPECT().
				UpdateByID(
					ormer,
					&models.Category{
						Name:      tt.args.category.Name,
						ID:        tt.mockResponse.result.ID,
						CreatedAt: tt.mockResponse.result.CreatedAt,
						Status:    tt.mockStatus,
					}).
				AnyTimes().
				Return(tt.mockResponse.num, tt.mockResponse.err)

			// Set properties serivce
			tt.fields.Storage.Category = mockCategory
			s := CategoryService{
				Storage: tt.fields.Storage,
				Orm:     ormMocked,
			}
			err := s.UpdateByID(ctx, tt.args.id, tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("CategoryService.UpdateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

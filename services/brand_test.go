package services

import (
	"context"
	"errors"
	"reflect"
	"shop-api/helper"
	ormmock "shop-api/helper/orm_mock"
	"shop-api/models"
	"shop-api/storage"
	brandmock "shop-api/storage/brand_mock"
	"shop-api/types"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestNewBrandService(t *testing.T) {
	tests := []struct {
		name  string
		wantS BrandService
	}{
		{"Base case", BrandService{storage.NewStorage(), helper.NewOrm()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := NewBrandService(); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("NewBrandService() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestBrandService_Add(t *testing.T) {
	type fields struct {
		Storage storage.Storage
		Orm     helper.OrmInterface
	}
	type args struct {
		ctx   context.Context
		input *types.InputAddBrand
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
		wantID       int64
		wantErr      bool
	}{
		{
			"Base case",
			fields{},
			args{ctx: context.Background(), input: &types.InputAddBrand{Title: "Title", Slug: "Slug"}},
			mockResponse{1, nil},
			1,
			false,
		},
		{
			"Fail case",
			fields{},
			args{ctx: context.Background(), input: &types.InputAddBrand{Title: "Title", Slug: "Slug"}},
			mockResponse{0, errors.New("Error")},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()

			// Mock storage method
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockBrand := brandmock.NewMockBrand(ctrl)
			mockBrand.EXPECT().
				Add(
					ormer,
					&models.Brand{
						Title: tt.args.input.Title,
						Slug:  tt.args.input.Slug,
					}).
				AnyTimes().
				Return(tt.mockResponse.num, tt.mockResponse.err)

			// Set properties serivce
			tt.fields.Storage.Brand = mockBrand
			tt.fields.Orm = ormMocked

			s := BrandService{
				Storage: tt.fields.Storage,
				Orm:     tt.fields.Orm,
			}
			gotID, err := s.Add(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("BrandService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotID != tt.wantID {
				t.Errorf("BrandService.Add() = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}

func TestBrandService_Delete(t *testing.T) {
	type fields struct {
		Storage storage.Storage
		Orm     helper.OrmInterface
	}
	type args struct {
		ctx context.Context
		id  int64
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
		{"Base case", fields{}, args{context.Background(), 1}, mockResponse{1, nil}, false},
		{"Fail case", fields{}, args{context.Background(), 2}, mockResponse{0, nil}, true},
		{"Fail case", fields{}, args{context.Background(), 3}, mockResponse{3, errors.New("Error")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()

			// Mock storage method
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockBrand := brandmock.NewMockBrand(ctrl)
			mockBrand.EXPECT().
				Delete(
					ormer,
					&models.Brand{
						ID: tt.args.id,
					}).
				AnyTimes().
				Return(tt.mockResponse.num, tt.mockResponse.err)

			// Set properties serivce
			tt.fields.Storage.Brand = mockBrand
			tt.fields.Orm = ormMocked

			s := BrandService{
				Storage: tt.fields.Storage,
				Orm:     tt.fields.Orm,
			}
			if err := s.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("BrandService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBrandService_GetByID(t *testing.T) {
	type fields struct {
		Storage storage.Storage
		Orm     helper.OrmInterface
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	type mockResponse struct {
		result models.Brand
		err    error
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockResponse mockResponse
		wantResult   types.OutputBrand
		wantErr      bool
	}{
		{
			"Base case",
			fields{},
			args{context.Background(), 1},
			mockResponse{models.Brand{Title: "title"}, nil},
			types.OutputBrand{Title: "title"},
			false,
		},
		{
			"Fail case",
			fields{},
			args{context.Background(), 1},
			mockResponse{models.Brand{Title: "title"}, errors.New("Error")},
			types.OutputBrand{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()

			// Mock storage method
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockBrand := brandmock.NewMockBrand(ctrl)
			mockBrand.EXPECT().
				GetByID(
					ormer,
					tt.args.id,
				).
				AnyTimes().
				Return(tt.mockResponse.result, tt.mockResponse.err)

			// Set properties serivce
			tt.fields.Storage.Brand = mockBrand
			tt.fields.Orm = ormMocked

			s := BrandService{
				Storage: tt.fields.Storage,
				Orm:     tt.fields.Orm,
			}
			gotResult, err := s.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("BrandService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("BrandService.GetByID() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestBrandService_GetAll(t *testing.T) {
	type fields struct {
		Storage storage.Storage
		Orm     helper.OrmInterface
	}
	type args struct {
		ctx    context.Context
		query  map[string]string
		order  []string
		offset int64
		limit  int64
	}
	type mockResponse struct {
		result []models.Brand
		err    error
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockResponse mockResponse
		wantResults  []types.OutputBrand
		wantErr      bool
	}{
		{
			"Base case",
			fields{},
			args{context.Background(), map[string]string{}, make([]string, 0), 10, 10},
			mockResponse{[]models.Brand{
				{Title: "Title"},
			}, nil},
			[]types.OutputBrand{
				{Title: "Title"},
			},
			false,
		},
		{
			"Fail case",
			fields{},
			args{context.Background(), map[string]string{}, make([]string, 0), 10, 10},
			mockResponse{[]models.Brand{{}}, errors.New("error")},
			[]types.OutputBrand{{}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()

			// Mock storage method
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockBrand := brandmock.NewMockBrand(ctrl)
			mockBrand.EXPECT().
				GetAll(
					ormer,
					tt.args.query,
					tt.args.order,
					tt.args.offset,
					tt.args.limit,
				).
				AnyTimes().
				Return(tt.mockResponse.result, tt.mockResponse.err)

			// Set properties serivce
			tt.fields.Storage.Brand = mockBrand
			tt.fields.Orm = ormMocked

			s := BrandService{
				Storage: tt.fields.Storage,
				Orm:     tt.fields.Orm,
			}
			gotResults, err := s.GetAll(tt.args.ctx, tt.args.query, tt.args.order, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("BrandService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("BrandService.GetAll() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func TestBrandService_UpdateByID(t *testing.T) {
	type fields struct {
		Storage storage.Storage
		Orm     helper.OrmInterface
	}
	type args struct {
		ctx   context.Context
		id    int64
		input *types.InputUpdateBrand
	}
	type mockResponse struct {
		result models.Brand
		err    error
		num    int64
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
			fields{},
			args{
				context.Background(),
				1,
				&types.InputUpdateBrand{Title: "New Title"},
			},
			mockResponse{models.Brand{Title: "Title", CreatedAt: time.Now(), UpdatedAt: time.Now()}, nil, 1},
			false,
		},
		{
			"Fail case",
			fields{},
			args{
				context.Background(),
				2,
				&types.InputUpdateBrand{Title: "New Title"},
			},
			mockResponse{models.Brand{}, nil, 0},
			true,
		},
		{
			"Fail db error case",
			fields{},
			args{
				context.Background(),
				3,
				&types.InputUpdateBrand{Title: "New Title"},
			},
			mockResponse{models.Brand{Title: "Title"}, errors.New("Error"), 0},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()

			// Mock storage method
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockBrand := brandmock.NewMockBrand(ctrl)
			mockBrand.EXPECT().
				GetByID(
					ormer,
					tt.args.id,
				).
				AnyTimes().
				Return(tt.mockResponse.result, tt.mockResponse.err)

			mockBrand.EXPECT().
				UpdateByID(
					ormer,
					&models.Brand{
						Title:     tt.args.input.Title,
						Slug:      tt.args.input.Slug,
						CreatedAt: tt.mockResponse.result.CreatedAt,
						UpdatedAt: tt.mockResponse.result.UpdatedAt,
					},
				).
				AnyTimes().
				Return(tt.mockResponse.num, tt.mockResponse.err)

			// Set properties serivce
			tt.fields.Storage.Brand = mockBrand
			tt.fields.Orm = ormMocked

			s := BrandService{
				Storage: tt.fields.Storage,
				Orm:     tt.fields.Orm,
			}
			if err := s.UpdateByID(tt.args.ctx, tt.args.id, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("BrandService.UpdateByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

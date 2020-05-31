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
	categorymock "shop-api/storage/category_mock"
	requestmock "shop-api/storage/request_mock"
	"shop-api/types"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewRequestService(t *testing.T) {
	tests := []struct {
		name  string
		wantS RequestService
	}{
		{
			"Base case", RequestService{Storage: storage.NewStorage(), Orm: helper.NewOrm()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := NewRequestService(); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("NewRequestService() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestRequestService_Add(t *testing.T) {
	type fields struct {
		Storage storage.Storage
		Orm     helper.OrmInterface
	}
	type args struct {
		ctx   context.Context
		input *types.InputAddRequest
	}
	type mockResponseGetCategory struct {
		categoryResponse models.Category
		err              error
	}
	type mockResponseAddCategory struct {
		num int64
		err error
	}
	type mockResponseGetBrand struct {
		brandResponse models.Brand
		err           error
	}
	type mockResponseAddBrand struct {
		num int64
		err error
	}
	type mockResponseAddRequest struct {
		num int64
		err error
	}
	tests := []struct {
		name                    string
		fields                  fields
		args                    args
		mockResponseGetCategory mockResponseGetCategory
		mockResponseAddCategory mockResponseAddCategory
		mockResponseGetBrand    mockResponseGetBrand
		mockResponseAddBrand    mockResponseAddBrand
		mockResponseAddRequest  mockResponseAddRequest
		mockBrandID             int64
		mockCategoryID          int64
		wantID                  int64
		wantErr                 bool
	}{
		{
			"Base case",
			fields{},
			args{
				ctx:   context.Background(),
				input: &types.InputAddRequest{Title: "title", Brand: "brand", Category: "category"},
			},
			mockResponseGetCategory{categoryResponse: models.Category{ID: 1}, err: nil},
			mockResponseAddCategory{},
			mockResponseGetBrand{brandResponse: models.Brand{ID: 1}, err: nil},
			mockResponseAddBrand{},
			mockResponseAddRequest{1, nil},
			1,
			1,
			1,
			false,
		},
		{
			"Base case add new brand and new category",
			fields{},
			args{
				ctx:   context.Background(),
				input: &types.InputAddRequest{Title: "title", Brand: "brand", Category: "category"},
			},
			mockResponseGetCategory{categoryResponse: models.Category{ID: 0}, err: nil},
			mockResponseAddCategory{1, nil},
			mockResponseGetBrand{brandResponse: models.Brand{ID: 0}, err: nil},
			mockResponseAddBrand{1, nil},
			mockResponseAddRequest{1, nil},
			1,
			1,
			1,
			false,
		},
		{
			"Fail",
			fields{},
			args{
				ctx:   context.Background(),
				input: &types.InputAddRequest{Title: "title", Brand: "brand", Category: "category"},
			},
			mockResponseGetCategory{categoryResponse: models.Category{ID: 0}, err: nil},
			mockResponseAddCategory{1, nil},
			mockResponseGetBrand{brandResponse: models.Brand{ID: 0}, err: nil},
			mockResponseAddBrand{1, nil},
			mockResponseAddRequest{1, errors.New("Erro")},
			1,
			1,
			1,
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
				GetByTitle(
					ormer,
					tt.args.input.Brand,
				).
				AnyTimes().
				Return(tt.mockResponseGetBrand.brandResponse, tt.mockResponseGetBrand.err)

			mockBrand.EXPECT().
				Add(
					ormer,
					&models.Brand{
						Title: tt.args.input.Brand,
					},
				).
				AnyTimes().
				Return(tt.mockResponseAddBrand.num, tt.mockResponseAddBrand.err)

			mockCategory := categorymock.NewMockCategory(ctrl)
			mockCategory.EXPECT().
				GetByName(
					ormer,
					tt.args.input.Category,
				).
				AnyTimes().
				Return(tt.mockResponseGetCategory.categoryResponse, tt.mockResponseGetCategory.err)
			mockCategory.EXPECT().
				Add(
					ormer,
					&models.Category{
						Name: tt.args.input.Category,
					},
				).
				AnyTimes().
				Return(tt.mockResponseAddCategory.num, tt.mockResponseAddCategory.err)

			mockRequest := requestmock.NewMockRequest(ctrl)
			mockRequest.EXPECT().
				Add(
					ormer,
					&models.Request{
						Title:    tt.args.input.Title,
						Brand:    &models.Brand{ID: tt.mockBrandID},
						Category: &models.Category{ID: tt.mockCategoryID},
					}).
				AnyTimes().
				Return(tt.mockResponseAddRequest.num, tt.mockResponseAddRequest.err)
			// Set properties serivce
			tt.fields.Storage.Brand = mockBrand
			tt.fields.Storage.Category = mockCategory
			tt.fields.Storage.Request = mockRequest
			tt.fields.Orm = ormMocked

			s := RequestService{
				Storage: tt.fields.Storage,
				Orm:     tt.fields.Orm,
			}
			gotID, err := s.Add(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotID != tt.wantID {
				t.Errorf("RequestService.Add() = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}

func TestRequestService_GetByID(t *testing.T) {
	type fields struct {
		Storage storage.Storage
		Orm     helper.OrmInterface
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult types.OutputRequest
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := RequestService{
				Storage: tt.fields.Storage,
				Orm:     tt.fields.Orm,
			}
			gotResult, err := s.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("RequestService.GetByID() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestRequestService_GetAll(t *testing.T) {
	type fields struct {
		Storage storage.Storage
		Orm     helper.OrmInterface
	}
	type args struct {
		ctx     context.Context
		query   map[string]string
		request []string
		offset  int64
		limit   int64
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantResults []types.OutputRequest
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := RequestService{
				Storage: tt.fields.Storage,
				Orm:     tt.fields.Orm,
			}
			gotResults, err := s.GetAll(tt.args.ctx, tt.args.query, tt.args.request, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("RequestService.GetAll() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func TestRequestService_UpdateByID(t *testing.T) {
	type fields struct {
		Storage storage.Storage
		Orm     helper.OrmInterface
	}
	type args struct {
		ctx   context.Context
		id    int64
		input *types.InputUpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := RequestService{
				Storage: tt.fields.Storage,
				Orm:     tt.fields.Orm,
			}
			if err := s.UpdateByID(tt.args.ctx, tt.args.id, tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("RequestService.UpdateByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

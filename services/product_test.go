package services

import (
	"context"
	"errors"
	"reflect"
	"shop-api/helper"
	ormmock "shop-api/helper/orm_mock"
	"shop-api/models"
	"shop-api/storage"
	productmock "shop-api/storage/product_mock"
	"shop-api/types"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
)

func TestNewProductService(t *testing.T) {
	tests := []struct {
		name   string
		wantPs ProductService
	}{
		{"Base case", ProductService{Storage: storage.NewStorage(), Orm: helper.NewOrm()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPs := NewProductService(); !reflect.DeepEqual(gotPs, tt.wantPs) {
				t.Errorf("NewProductService() = %v, want %v", gotPs, tt.wantPs)
			}
		})
	}
}

func TestProductService_Add(t *testing.T) {
	type fields struct {
		Storage storage.Storage
	}
	type args struct {
		product types.InputAddProduct
	}
	type mockResponse struct {
		id  int64
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
			fields{Storage: storage.Storage{}},
			args{types.InputAddProduct{Name: "Name", Detail: "Detail"}},
			mockResponse{1, nil},
			1,
			false,
		},
		{
			"Bad request",
			fields{Storage: storage.Storage{}},
			args{types.InputAddProduct{Name: "Name", Detail: "Detail"}},
			mockResponse{0, errors.New("Fail")},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()
			ctx := context.Background()

			// Mock method storage
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProduct := productmock.NewMockProduct(ctrl)
			mockProduct.EXPECT().
				Add(
					ormer,
					&models.Product{
						Name:     tt.args.product.Name,
						Detail:   tt.args.product.Detail,
						Category: &models.Category{},
					}).
				AnyTimes().
				Return(tt.mockResponse.id, tt.mockResponse.err)

			// Set properties service
			tt.fields.Storage.Product = mockProduct

			ps := ProductService{
				Storage: tt.fields.Storage,
				Orm:     ormMocked,
			}
			gotID, err := ps.Add(ctx, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotID != tt.wantID {
				t.Errorf("ProductService.Add() gotID = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}

func TestProductService_Delete(t *testing.T) {
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
		{"Base case", fields{Storage: storage.Storage{}}, args{1}, mockResponse{1, nil}, false},
		{"Not found ID", fields{Storage: storage.Storage{}}, args{2}, mockResponse{0, errors.New("Not Found")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()
			ctx := context.Background()

			// Mock method storage
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProduct := productmock.NewMockProduct(ctrl)
			mockProduct.EXPECT().
				Delete(
					ormer,
					&models.Product{
						ID: tt.args.id,
					}).
				AnyTimes().
				Return(tt.mockResponse.num, tt.mockResponse.err)

			// Set properties service
			tt.fields.Storage.Product = mockProduct

			ps := ProductService{
				Storage: tt.fields.Storage,
				Orm:     ormMocked,
			}
			err := ps.Delete(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProductService_GetByID(t *testing.T) {
	type fields struct {
		Storage storage.Storage
	}
	type args struct {
		id int64
	}
	type mockResponse struct {
		product models.Product
		err     error
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockResponse mockResponse
		wantResult   types.OutputProduct
		wantErr      bool
	}{
		{
			"Base case",
			fields{storage.Storage{}},
			args{1},
			mockResponse{models.Product{Name: "Name", CreatedAt: time.Now()}, nil},
			types.OutputProduct{Name: "Name"},
			false,
		},
		{
			"Base case",
			fields{storage.Storage{}},
			args{1},
			mockResponse{models.Product{Name: "Name", CreatedAt: time.Now(), UpdatedAt: time.Now()}, nil},
			types.OutputProduct{Name: "Name"},
			false,
		},
		{
			"No row found",
			fields{storage.Storage{}},
			args{100},
			mockResponse{models.Product{}, errors.New("No row found")},
			types.OutputProduct{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()
			ctx := context.Background()

			// Mock method storage
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProduct := productmock.NewMockProduct(ctrl)
			mockProduct.EXPECT().
				GetByID(ormer, tt.args.id).
				AnyTimes().
				Return(tt.mockResponse.product, tt.mockResponse.err)

			// Set properties service
			tt.fields.Storage.Product = mockProduct
			ps := ProductService{
				Storage: tt.fields.Storage,
				Orm:     ormMocked,
			}
			gotResult, err := ps.GetByID(ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ProductService.GetByID() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestProductService_GetAll(t *testing.T) {
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
		result []models.Product
		err    error
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		mockResponse mockResponse
		wantResults  []types.OutputProduct
		wantErr      bool
	}{
		{
			"Base case",
			fields{storage.Storage{}},
			args{map[string]string{}, []string{}, 0, 0},
			mockResponse{[]models.Product{{Name: "Name", CreatedAt: time.Now()}}, nil},
			[]types.OutputProduct{{Name: "Name"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()
			ctx := context.Background()

			// Mock method storage
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProduct := productmock.NewMockProduct(ctrl)
			mockProduct.EXPECT().
				GetAll(ormer, tt.args.query, tt.args.order, tt.args.offset, tt.args.limit).
				AnyTimes().
				Return(tt.mockResponse.result, tt.mockResponse.err)

			// Set properties service
			tt.fields.Storage.Product = mockProduct
			ps := ProductService{
				Storage: tt.fields.Storage,
				Orm:     ormMocked,
			}
			gotResults, err := ps.GetAll(ctx, tt.args.query, tt.args.order, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("ProductService.GetAll() gotResults = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func TestProductService_UpdateByID(t *testing.T) {
	type fields struct {
		Storage storage.Storage
	}
	type args struct {
		id      int64
		product *types.InputUpdateProduct
	}
	type thirdPartyResponse struct {
		num     int64
		product models.Product
		err     error
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		thirdPartyResponse thirdPartyResponse
		wantErr            bool
	}{
		{
			"Base case",
			fields{storage.Storage{}},
			args{1, &types.InputUpdateProduct{}},
			thirdPartyResponse{1, models.Product{ID: 1, Category: &models.Category{ID: 1}}, nil},
			false,
		},
		{
			"Not Found ID",
			fields{storage.Storage{}},
			args{1, &types.InputUpdateProduct{}},
			thirdPartyResponse{0, models.Product{ID: 1, Category: &models.Category{ID: 1}}, errors.New("Not Found")},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()
			ctx := context.Background()

			// Mock method storage
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockProduct := productmock.NewMockProduct(ctrl)
			mockProduct.EXPECT().
				GetByID(ormer, tt.args.id).
				AnyTimes().
				Return(tt.thirdPartyResponse.product, tt.thirdPartyResponse.err)
			mockProduct.EXPECT().
				UpdateByID(
					ormer,
					&models.Product{
						ID:        tt.thirdPartyResponse.product.ID,
						Name:      tt.args.product.Name,
						Detail:    tt.args.product.Detail,
						Brand:     tt.args.product.Brand,
						Model:     tt.args.product.Model,
						Cost:      tt.args.product.Cost,
						Price:     tt.args.product.Price,
						Stock:     tt.args.product.Stock,
						CreatedAt: tt.thirdPartyResponse.product.CreatedAt,
						Category:  &models.Category{ID: tt.args.product.Category.ID},
					}).
				AnyTimes().
				Return(tt.thirdPartyResponse.num, tt.thirdPartyResponse.err)

			// Set properties service
			tt.fields.Storage.Product = mockProduct
			ps := ProductService{
				Storage: tt.fields.Storage,
				Orm:     ormMocked,
			}
			err := ps.UpdateByID(ctx, tt.args.id, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.UpdateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

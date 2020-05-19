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

func TestNewProductService(t *testing.T) {
	tests := []struct {
		name   string
		wantPs ProductService
	}{
		{"Base case", ProductService{Storage: storage.NewStorage()}},
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
		name             string
		fields           fields
		args             args
		mockResponse     mockResponse
		wantResponseCode int
		wantID           int64
		wantErr          bool
	}{
		{
			"Base case",
			fields{Storage: storage.Storage{}},
			args{types.InputAddProduct{Name: "Name", Detail: "Detail"}},
			mockResponse{1, nil},
			types.ResponseCode["CreatedSuccess"],
			1,
			false,
		},
		{
			"Bad request",
			fields{Storage: storage.Storage{}},
			args{types.InputAddProduct{Name: "Name", Detail: "Detail"}},
			mockResponse{0, errors.New("Fail")},
			types.ResponseCode["BadRequest"],
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mackProduct := mock.NewMockProduct(ctrl)
			mackProduct.EXPECT().
				Add(&models.Product{
					Name:     tt.args.product.Name,
					Detail:   tt.args.product.Detail,
					Category: &models.Category{},
				}).
				AnyTimes().
				Return(tt.mockResponse.id, tt.mockResponse.err)
			tt.fields.Storage.Product = mackProduct

			ps := ProductService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, gotID, err := ps.Add(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponseCode != tt.wantResponseCode {
				t.Errorf("ProductService.Add() gotResponseCode = %v, want %v", gotResponseCode, tt.wantResponseCode)
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
			ps := ProductService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, err := ps.Delete(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponseCode != tt.wantResponseCode {
				t.Errorf("ProductService.Delete() = %v, want %v", gotResponseCode, tt.wantResponseCode)
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
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantResponseCode int
		wantResult       types.OutputProduct
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := ProductService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, gotResult, err := ps.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponseCode != tt.wantResponseCode {
				t.Errorf("ProductService.GetByID() gotResponseCode = %v, want %v", gotResponseCode, tt.wantResponseCode)
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
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantResponseCode int
		wantResults      []types.OutputProduct
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := ProductService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, gotResults, err := ps.GetAll(tt.args.query, tt.args.order, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponseCode != tt.wantResponseCode {
				t.Errorf("ProductService.GetAll() gotResponseCode = %v, want %v", gotResponseCode, tt.wantResponseCode)
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
			ps := ProductService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, err := ps.UpdateByID(tt.args.id, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.UpdateByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponseCode != tt.wantResponseCode {
				t.Errorf("ProductService.UpdateByID() = %v, want %v", gotResponseCode, tt.wantResponseCode)
			}
		})
	}
}

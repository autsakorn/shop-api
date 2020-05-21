package services

import (
	"errors"
	"reflect"
	"shop-api/config"
	"shop-api/helper"
	"shop-api/models"
	"shop-api/storage"
	"shop-api/storage/mock"
	"shop-api/types"
	"testing"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
)

func init() {
	env, _ := config.FromEnv()
	orm.RegisterDataBase("default", env.Driver, env.Sqlconn)
}
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
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mackProduct := mock.NewMockProduct(ctrl)
			ormer := helper.NewOrm(false)
			mackProduct.EXPECT().
				Add(
					ormer,
					&models.Product{
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
			gotResponseCode, gotID, err := ps.Add(ormer, tt.args.product)
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
		{"Base case", fields{Storage: storage.Storage{}}, args{1}, mockResponse{1, nil}, types.ResponseCode["Success"], false},
		{"Not found ID", fields{Storage: storage.Storage{}}, args{2}, mockResponse{0, errors.New("Not Found")}, types.ResponseCode["BadRequest"], true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mackProduct := mock.NewMockProduct(ctrl)
			ormer := helper.NewOrm(false)
			mackProduct.EXPECT().
				Delete(
					ormer,
					&models.Product{
						ID: tt.args.id,
					}).
				AnyTimes().
				Return(tt.mockResponse.num, tt.mockResponse.err)
			tt.fields.Storage.Product = mackProduct

			ps := ProductService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, err := ps.Delete(ormer, tt.args.id)
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
	type mockResponse struct {
		product models.Product
		err     error
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		mockResponse     mockResponse
		wantResponseCode int
		wantResult       types.OutputProduct
		wantErr          bool
	}{
		{
			"Base case",
			fields{storage.Storage{}},
			args{1},
			mockResponse{models.Product{Name: "Name", CreatedAt: time.Now()}, nil},
			types.ResponseCode["Success"],
			types.OutputProduct{Name: "Name"},
			false,
		},
		{
			"Base case",
			fields{storage.Storage{}},
			args{1},
			mockResponse{models.Product{Name: "Name", CreatedAt: time.Now(), UpdatedAt: time.Now()}, nil},
			types.ResponseCode["Success"],
			types.OutputProduct{Name: "Name"},
			false,
		},
		{
			"No row found",
			fields{storage.Storage{}},
			args{100},
			mockResponse{models.Product{}, errors.New("No row found")},
			types.ResponseCode["BadRequest"],
			types.OutputProduct{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mackProduct := mock.NewMockProduct(ctrl)
			ormer := helper.NewOrm(false)
			mackProduct.EXPECT().
				GetByID(ormer, tt.args.id).
				AnyTimes().
				Return(tt.mockResponse.product, tt.mockResponse.err)
			tt.fields.Storage.Product = mackProduct
			ps := ProductService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, gotResult, err := ps.GetByID(ormer, tt.args.id)
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
	type mockResponse struct {
		result []models.Product
		err    error
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		mockResponse     mockResponse
		wantResponseCode int
		wantResults      []types.OutputProduct
		wantErr          bool
	}{
		{
			"Base case",
			fields{storage.Storage{}},
			args{map[string]string{}, []string{}, 0, 0},
			mockResponse{[]models.Product{{Name: "Name", CreatedAt: time.Now()}}, nil},
			types.ResponseCode["Success"],
			[]types.OutputProduct{{Name: "Name"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mackProduct := mock.NewMockProduct(ctrl)
			ormer := helper.NewOrm(false)
			mackProduct.EXPECT().
				GetAll(ormer, tt.args.query, tt.args.order, tt.args.offset, tt.args.limit).
				AnyTimes().
				Return(tt.mockResponse.result, tt.mockResponse.err)
			tt.fields.Storage.Product = mackProduct
			ps := ProductService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, gotResults, err := ps.GetAll(ormer, tt.args.query, tt.args.order, tt.args.offset, tt.args.limit)
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
		wantResponseCode   int
		wantErr            bool
	}{
		{
			"Base case",
			fields{storage.Storage{}},
			args{1, &types.InputUpdateProduct{}},
			thirdPartyResponse{1, models.Product{ID: 1, Category: &models.Category{ID: 1}}, nil},
			types.ResponseCode["Success"],
			false,
		},
		{
			"Not Found ID",
			fields{storage.Storage{}},
			args{1, &types.InputUpdateProduct{}},
			thirdPartyResponse{0, models.Product{ID: 1, Category: &models.Category{ID: 1}}, errors.New("Not Found")},
			types.ResponseCode["BadRequest"],
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mackProduct := mock.NewMockProduct(ctrl)
			ormer := helper.NewOrm(false)
			mackProduct.EXPECT().
				GetByID(ormer, tt.args.id).
				AnyTimes().
				Return(tt.thirdPartyResponse.product, tt.thirdPartyResponse.err)
			mackProduct.EXPECT().
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
						Quantity:  tt.args.product.Quantity,
						CreatedAt: tt.thirdPartyResponse.product.CreatedAt,
						Category:  &models.Category{ID: tt.args.product.Category.ID},
					}).
				AnyTimes().
				Return(tt.thirdPartyResponse.num, tt.thirdPartyResponse.err)
			tt.fields.Storage.Product = mackProduct
			ps := ProductService{
				Storage: tt.fields.Storage,
			}
			gotResponseCode, err := ps.UpdateByID(ormer, tt.args.id, tt.args.product)
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

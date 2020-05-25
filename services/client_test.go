package services

import (
	"context"
	"reflect"
	"shop-api/helper"
	ormmock "shop-api/helper/orm_mock"
	"shop-api/storage"
	clientmock "shop-api/storage/client_mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewClientService(t *testing.T) {
	tests := []struct {
		name string
		want ClientService
	}{
		{
			"Base case",
			ClientService{Storage: storage.NewStorage(), Orm: helper.NewOrm()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientService_VerifyXApiKey(t *testing.T) {
	type fields struct {
		Storage storage.Storage
		Orm     helper.OrmInterface
	}
	type args struct {
		ctx     context.Context
		xAPIKey string
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
		wantErr      bool
	}{
		{
			"Base case",
			fields{storage.Storage{}, ormmock.NewOrmMock()},
			args{
				ctx:     context.Background(),
				xAPIKey: "QWERTYUIO",
			},
			mockResponse{1, nil},
			false,
		},
		{
			"Not found x_api_key",
			fields{storage.Storage{}, ormmock.NewOrmMock()},
			args{
				ctx:     context.Background(),
				xAPIKey: "QWERTYUIO",
			},
			mockResponse{0, nil},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock ormer
			ormMocked := ormmock.OrmMock{}
			ormer := ormMocked.NewOrms()

			// Mock storage
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockClient := clientmock.NewMockClient(ctrl)
			status := int32(1)
			mockClient.EXPECT().
				GetByXApiKeyAndStatus(ormer, tt.args.xAPIKey, status).
				AnyTimes().
				Return(tt.mockResponse.id, tt.mockResponse.err)
			tt.fields.Storage.Client = mockClient

			// Set properties service
			s := ClientService{
				Storage: tt.fields.Storage,
				Orm:     tt.fields.Orm,
			}
			if err := s.VerifyXApiKey(tt.args.ctx, tt.args.xAPIKey); (err != nil) != tt.wantErr {
				t.Errorf("ClientService.VerifyXApiKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

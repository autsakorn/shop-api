package services

import (
	"context"
	"errors"
	"fmt"
	"shop-api/helper"
	"shop-api/storage"
)

// Client represents all possible actions available for client services
type Client interface {
	VerifyXApiKey(context.Context, string) (int64, error)
}

// ClientService defines properties
type ClientService struct {
	Storage storage.Storage
	Orm     helper.OrmInterface
}

// NewClientService func return ClientService object
func NewClientService() ClientService {
	s := ClientService{}
	s.Storage = storage.NewStorage()
	s.Orm = helper.NewOrm()
	return s
}

// VerifyXApiKey service for retrieve client BY XApiKey
func (s ClientService) VerifyXApiKey(ctx context.Context, xAPIKey string) (err error) {
	ormer := s.Orm.NewOrms()
	status := int32(1) // Declare status = 1, This service find by x_api_key and status is Active
	id, err := s.Storage.Client.GetByXApiKeyAndStatus(ormer, xAPIKey, status)
	fmt.Println("id", id)
	if err != nil || id < 1 {
		err = errors.New("Invalid X-API-KEY")
		return
	}
	return
}

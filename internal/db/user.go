package db

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/anil1226/go-employee/internal/service/user"
	uuid "github.com/satori/go.uuid"
)

func (d *Database) GetUser(ctx context.Context, id string) (user.User, error) {
	contClient, err := d.GetContainerClient(containerUser)
	if err != nil {
		return user.User{}, err
	}
	pk := azcosmos.NewPartitionKeyString(id)
	// Read an item
	itemResponse, err := contClient.ReadItem(ctx, pk, id, nil)
	if err != nil {
		return user.User{}, err
	}
	var itemResponseBody user.User
	err = json.Unmarshal(itemResponse.Value, &itemResponseBody)
	if err != nil {
		return user.User{}, err
	}
	return itemResponseBody, nil

}

func (d *Database) CreateUser(ctx context.Context, user user.User) error {
	contClient, err := d.GetContainerClient(containerUser)
	if err != nil {
		return err
	}
	user.ID = uuid.NewV4().String()
	pk := azcosmos.NewPartitionKeyString(user.ID)
	marshalled, err := json.Marshal(user)
	if err != nil {
		return err
	}
	_, err = contClient.CreateItem(ctx, pk, marshalled, nil)
	if err != nil {
		return err
	}
	return nil
}

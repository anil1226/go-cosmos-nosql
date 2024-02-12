package db

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/anil1226/go-employee/internal/models"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func (d *Database) GetUser(ctx context.Context, user models.User) (models.User, error) {
	contClient, err := d.GetContainerClient(containerUser)
	if err != nil {
		return models.User{}, err
	}
	pk := azcosmos.NewPartitionKeyString(user.Name)

	query := "SELECT * FROM users c WHERE c.name = @name"

	queryOptions := azcosmos.QueryOptions{
		QueryParameters: []azcosmos.QueryParameter{
			{Name: "@name", Value: user.Name},
		},
	}

	pager := contClient.NewQueryItemsPager(query, pk, &queryOptions)
	// Read an item
	// itemResponse, err := contClient.ReadItem(ctx, pk, name, nil)
	// if err != nil {
	// 	return models.User{}, err
	// }
	var itemResponseBody models.User
	for pager.More() {
		response, err := pager.NextPage(ctx)
		if err != nil {
			return models.User{}, err
		}

		for _, bytes := range response.Items {

			err := json.Unmarshal(bytes, &itemResponseBody)
			if err != nil {
				return models.User{}, err
			}

		}
	}

	if len(strings.Trim(itemResponseBody.Name, "")) == 0 {
		return models.User{}, errors.New("user does not exist")
	}

	if bcrypt.CompareHashAndPassword([]byte(itemResponseBody.Password), []byte(user.Password)) == nil {
		return itemResponseBody, nil
	}
	return models.User{}, errors.New("password does not match")

}

func (d *Database) CreateUser(ctx context.Context, user models.User) error {
	contClient, err := d.GetContainerClient(containerUser)
	if err != nil {
		return err
	}
	user.ID = uuid.NewV4().String()
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		return err
	}
	pk := azcosmos.NewPartitionKeyString(user.Name)
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

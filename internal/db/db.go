package db

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/anil1226/go-employee/internal/employee"
)

type Database struct {
	Client *azcosmos.ContainerClient
}

const (
	cosmosDbEndpoint = "https://goprac.documents.azure.com:443/"
	cosmosDbKey      = "XZ3lSTbdfT5mqARLFEXYyqDwzanTRkErGi5bnTLePpar0FqQ5f7AX07L6zdsbXfexqK7idQHBSyOACDbIlBcpg=="
	dbName           = "goprac"
	containerName    = "employees"
)

func NewDatabase() (*Database, error) {
	cred, err := azcosmos.NewKeyCredential(cosmosDbKey)
	if err != nil {
		return nil, err
	}
	client, err := azcosmos.NewClientWithKey(cosmosDbEndpoint, cred, nil)
	if err != nil {
		return nil, err
	}

	// database := azcosmos.DatabaseProperties{ID: dbName}
	// _, err = client.CreateDatabase(context.Background(), database, nil)
	// if err != nil {
	// 	return nil, err
	// }
	database, err := client.NewDatabase(dbName)
	if err != nil {
		return nil, err
	}

	container, err := database.NewContainer(containerName)
	if err != nil {
		return nil, err
	}
	return &Database{
		Client: container,
	}, nil
}

func (d *Database) ReadItem() (*employee.Employee, error) {
	pk := azcosmos.NewPartitionKeyString("1")
	id := "1"

	// Read an item
	itemResponse, err := d.Client.ReadItem(context.Background(), pk, id, nil)
	if err != nil {
		return nil, err
	}

	var itemResponseBody employee.Employee
	err = json.Unmarshal(itemResponse.Value, &itemResponseBody)
	if err != nil {
		return nil, err
	}

	return &itemResponseBody, nil

}

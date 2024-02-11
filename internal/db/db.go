package db

import (
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type Database struct {
	Client *azcosmos.ContainerClient
}

const (
	cosmosDbEndpoint = "https://goprac.documents.azure.com:443/"
	cosmosDbKey      = "zqlVKJaKYP7jhvwPAfX91vWaXRYcvdJAVSRBkEgieC4JRWigIGDJ7ESRRBQWYn13uEeyCPiE1McNACDb6hiHmQ=="
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

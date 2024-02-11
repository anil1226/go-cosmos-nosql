package db

import (
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/anil1226/go-employee/config"
)

type Database struct {
	Client *azcosmos.DatabaseClient
}

var (
	cosmosDbEndpoint = config.GetEnvKey("cosmosDbEndpoint")
	cosmosDbKey      = config.GetEnvKey("cosmosDbKey")
	dbName           = config.GetEnvKey("dbName")
	containerEmp     = config.GetEnvKey("containerEmp")
	containerUser    = config.GetEnvKey("containerUser")
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

	// container, err := database.NewContainer(containerName)
	// if err != nil {
	// 	return nil, err
	// }
	return &Database{
		Client: database,
	}, nil
}

package db

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/anil1226/go-employee/internal/service/employee"
	uuid "github.com/satori/go.uuid"
)

func (d *Database) GetEmployee(ctx context.Context, id string) (employee.Employee, error) {
	contClient, err := d.GetContainerClient(containerEmp)
	if err != nil {
		return employee.Employee{}, err
	}
	pk := azcosmos.NewPartitionKeyString(id)
	// Read an item
	itemResponse, err := contClient.ReadItem(ctx, pk, id, nil)
	if err != nil {
		return employee.Employee{}, err
	}
	var itemResponseBody employee.Employee
	err = json.Unmarshal(itemResponse.Value, &itemResponseBody)
	if err != nil {
		return employee.Employee{}, err
	}
	return itemResponseBody, nil

}

func (d *Database) CreateEmployee(ctx context.Context, emp employee.Employee) error {
	contClient, err := d.GetContainerClient(containerEmp)
	if err != nil {
		return err
	}
	emp.ID = uuid.NewV4().String()
	pk := azcosmos.NewPartitionKeyString(emp.ID)
	marshalled, err := json.Marshal(emp)
	if err != nil {
		return err
	}
	_, err = contClient.CreateItem(ctx, pk, marshalled, nil)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) UpdateEmployee(ctx context.Context, emp employee.Employee) error {
	contClient, err := d.GetContainerClient(containerEmp)
	if err != nil {
		return err
	}
	pk := azcosmos.NewPartitionKeyString(emp.ID)
	marshalled, err := json.Marshal(emp)
	if err != nil {
		return err
	}
	_, err = contClient.ReplaceItem(ctx, pk, emp.ID, marshalled, nil)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteEmployee(ctx context.Context, id string) error {
	contClient, err := d.GetContainerClient(containerEmp)
	if err != nil {
		return err
	}
	pk := azcosmos.NewPartitionKeyString(id)
	_, err = contClient.DeleteItem(ctx, pk, id, nil)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) GetContainerClient(name string) (*azcosmos.ContainerClient, error) {
	return d.Client.NewContainer(name)

}

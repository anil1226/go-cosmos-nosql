package db

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/anil1226/go-employee/internal/employee"
	uuid "github.com/satori/go.uuid"
)

func (d *Database) GetEmployee(ctx context.Context, id string) (employee.Employee, error) {
	pk := azcosmos.NewPartitionKeyString(id)
	// Read an item
	itemResponse, err := d.Client.ReadItem(ctx, pk, id, nil)
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
	emp.ID = uuid.NewV4().String()
	pk := azcosmos.NewPartitionKeyString(emp.ID)
	marshalled, err := json.Marshal(emp)
	if err != nil {
		return err
	}
	_, err = d.Client.CreateItem(ctx, pk, marshalled, nil)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) UpdateEmployee(ctx context.Context, emp employee.Employee) error {
	pk := azcosmos.NewPartitionKeyString(emp.ID)
	marshalled, err := json.Marshal(emp)
	if err != nil {
		return err
	}
	_, err = d.Client.ReplaceItem(ctx, pk, emp.ID, marshalled, nil)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteEmployee(ctx context.Context, id string) error {
	pk := azcosmos.NewPartitionKeyString(id)
	_, err := d.Client.DeleteItem(ctx, pk, id, nil)
	if err != nil {
		return err
	}
	return nil
}

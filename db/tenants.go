package db

import "fmt"

type Tenant struct {
	ID       int
	Name     string
	Password string
}

type CreateTenantRequest struct {
	Name     string
	Password string
}

func (db DB) CreateTenant(request CreateTenantRequest) error {
	_, err := db.db.NamedExec(
		`INSERT INTO tenants (name, password) VALUES (:name, :password)`,
		map[string]interface{}{
			"name":     request.Name,
			"password": request.Password,
		},
	)

	return err
}

func (db DB) CreateTenantDB(request CreateTenantRequest) error {
	_, err := db.db.Exec(
		fmt.Sprintf("CREATE DATABASE %s WITH TEMPLATE = tenant OWNER = goblogs", request.Name),
	)

	return err
}

type DeleteTenantRequest struct {
	ByName string
}

func (db DB) DeleteTenant(request DeleteTenantRequest) error {
	_, err := db.db.NamedExec(
		"DELETE FROM tenants WHERE name = :name",
		map[string]interface{}{
			"name": request.ByName,
		},
	)

	return err
}

type GetTenantRequest struct {
	ByName string
}

func (db DB) GetTenant(request GetTenantRequest) (*Tenant, error) {
	var tenant Tenant

	if err := db.db.Get(&tenant, "SELECT * FROM tenants WHERE name = $1", request.ByName); err != nil {
		return nil, err
	}

	return &tenant, nil
}

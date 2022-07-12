package main

import "database/sql"

//	Fetch seed data from DB
func fetchSeedData(db *sql.DB) error {

	//	Fetch roles
	rows, err := db.Query("SELECT id, name FROM roles")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var role Role
		if err := rows.Scan(&role.ID, &role.Name); err != nil {
			return err
		}
		roles = append(roles, role)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	//	Fetch operations
	rows, err = db.Query("SELECT id, type FROM operations")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var operation Operation
		if err := rows.Scan(&operation.ID, &operation.Type); err != nil {
			return err
		}
		operations = append(operations, operation)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	//	Fetch objects
	rows, err = db.Query("SELECT id, name FROM objects")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var object Object
		if err := rows.Scan(&object.ID, &object.Name); err != nil {
			return err
		}
		objects = append(objects, object)
	}

	//	Fetch permissions
	rows, err = db.Query("SELECT id, object_id, operation_id FROM permissions")
	if err != nil {
		return err
	}

	for rows.Next() {
		var permission Permission
		if err := rows.Scan(&permission.ID, &permission.ObjectID, &permission.OperationID); err != nil {
			return err
		}

		permissions = append(permissions, permission)
	}

	//	Populate permissions from objects and operations
	for _, permission := range permissions {
		object, err := getObjectByID(permission.ObjectID)
		if err != nil {
			return err
		}
		permission.Object = object

		operation, err := getOperationByID(permission.OperationID)
		if err != nil {
			return err
		}
		permission.Operation = operation
	}
	return nil
}

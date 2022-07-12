package main

import "fmt"

//	Get operation by type
func getOperationByType(type_ string) (Operation, error) {
	for _, operation := range operations {
		if operation.Type == type_ {
			return operation, nil
		}
	}
	return Operation{}, fmt.Errorf("Operation not found")
}

//	Get role by name
func getRoleByName(name string) (Role, error) {
	for _, role := range roles {
		if role.Name == name {
			return role, nil
		}
	}
	return Role{}, fmt.Errorf("Role not found")
}

//	Get object by name
func getObjectByName(name string) (Object, error) {
	for _, object := range objects {
		if object.Name == name {
			return object, nil
		}
	}
	return Object{}, fmt.Errorf("Object not found")
}

//	Get object by ID
func getObjectByID(id string) (Object, error) {
	for _, object := range objects {
		if object.ID == id {
			return object, nil
		}
	}
	return Object{}, fmt.Errorf("Object not found")
}

//	Get operation by ID
func getOperationByID(id string) (Operation, error) {
	for _, operation := range operations {
		if operation.ID == id {
			return operation, nil
		}
	}
	return Operation{}, fmt.Errorf("Operation not found")
}

//	Get permission by object and operation
func getPermission(object Object, operation Operation) (Permission, error) {
	for _, permission := range permissions {
		if permission.ObjectID == object.ID && permission.OperationID == operation.ID {
			return permission, nil
		}
	}
	return Permission{}, fmt.Errorf("Permission not found")
}

//	Add permissions for role from objects and operations
func addPermissions(role *Role) error {
	switch role.Name {

	case "admin":
		for _, object := range objects {
			for _, operation := range operations {
				permission, err := getPermission(object, operation)
				if err != nil {
					return err
				}
				role.Permissions = append(role.Permissions, Permission{
					ID:        permission.ID,
					Object:    object,
					Operation: operation,
				})
			}
		}

	case "editor":
		for _, object := range objects {
			var operations []string
			switch object.Name {
			case "entities", "integrations":
				operations = append(operations, []string{"insert"}...)
			case "roles":
				operations = append(operations, []string{"insert", "view"}...)
			default:
				operations = append(operations, []string{"insert", "view", "update", "delete"}...)
			}

			for _, operation := range operations {
				operation, err := getOperationByType(operation)
				if err != nil {
					return err
				}
				permission, err := getPermission(object, operation)
				if err != nil {
					return err
				}
				role.Permissions = append(role.Permissions, Permission{
					ID:        permission.ID,
					Object:    object,
					Operation: operation,
				})
			}
		}
	case "viewer":
		viewOp, _ := getOperationByType("view")
		for _, object := range objects {
			permission, err := getPermission(object, viewOp)
			if err != nil {
				return err
			}
			role.Permissions = append(role.Permissions, Permission{
				ID:        permission.ID,
				Object:    object,
				Operation: viewOp,
			})
		}
	}

	/* 	//	Load permissions
	   	for i := 0; i < len(role.Permissions); i++ {
	   		for j := 0; j < len(permissions); j++ {
	   			if permissions[j].ObjectID == role.Permissions[i].Object.ID &&
	   				permissions[j].Operation.ID == role.Permissions[i].OperationID {
	   				role.Permissions[i].ID = permissions[j].ID
	   				permissions[j].Object = role.Permissions[i].Object
	   				permissions[j].Operation = role.Permissions[i].Operation
	   			}
	   		}
	   	}
	*/
	return nil
}

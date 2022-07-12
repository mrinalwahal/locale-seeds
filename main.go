package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgrespassword"
	dbname   = "postgres"

	sqlStatement = `INSERT INTO role_has_permission (role_id, permission_id) VALUES ($1, $2)`
)

var (
	roles       []Role
	operations  []Operation
	objects     []Object
	permissions []Permission
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//	Populate our arrays
	err = fetchSeedData(db)
	if err != nil {
		panic(err)
	}

	//	Populate our permissions
	for i := 0; i < len(roles); i++ {
		err = addPermissions(&roles[i])
		if err != nil {
			panic(err)
		}
	}

	/* 	//	Print our permissions
	   	for _, item := range permissions {
	   		fmt.Println("ID: ", item.ID)
	   		fmt.Println("ObjectID: ", item.ObjectID)
	   		fmt.Println("Object: ", item.Object.ID)
	   		fmt.Println("OperationID: ", item.OperationID)
	   	}
	*/

	//	Seed the DB with permissions for every role
	for _, role := range roles {
		for _, permission := range role.Permissions {

			//	Print permission for object and operation
			fmt.Println(role.Name + " : " + fmt.Sprintf("%+v", permission.Object.Name) + " " + fmt.Sprintf("%+v", permission.Operation.Type) + " " + fmt.Sprintf("%+v", permission.ID))

			_, err = db.Exec(sqlStatement, role.ID, permission.ID)
			if err != nil {
				panic(err)
			}

		}
	}
}

func printRoles() {

	for _, role := range roles {
		fmt.Println("------")
		fmt.Println("Role: " + fmt.Sprintf("%+v", role.Name))
		fmt.Println("------")

		for _, permission := range role.Permissions {

			//	Print permission for object and operation
			fmt.Println(fmt.Sprintf("%+v", permission.Object.Name) + " " + fmt.Sprintf("%+v", permission.Operation.Type))
		}

		fmt.Println()
	}

}

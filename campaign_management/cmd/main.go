package main

import (
	"campaing_management/pckg/datastore"
	"campaing_management/pckg/routers"
	"fmt"
)

func main() {
	fmt.Println("Started on port :8080")
	datastore.InitialMigration()
	routers.InitializeRouter()
}

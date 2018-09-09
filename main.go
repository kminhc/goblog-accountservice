package main

import (
	"fmt"

	"github.com/kminhc/goblog/accountservice/dbclient"

	"github.com/kminhc/goblog/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient() {
	service.DBCLient = &dbclient.BoltClient{}
	service.DBCLient.OpenBoltDb()
	service.DBCLient.Seed()
}

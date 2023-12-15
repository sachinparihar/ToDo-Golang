package main

import (
	"htmx/model"
	"htmx/routes"
)

func main() {

	err := model.ConnectToMongoDB()
	if err != nil {
		panic(err)
	}
	routes.SetupAndRun()
}

package main

import (
	"github.com/hafif/echoFramework/db"
	"github.com/hafif/echoFramework/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))

}

package main

import (
	db "echo-framework-maul/db"
	routes "echo-framework-maul/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))
}

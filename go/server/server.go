package server

import "recipe-app/go/database"

func Start() {
	database.SetDBConnection(database.NewDBOptions())

	router := setupRouter()

	router.Run(":8090")
}

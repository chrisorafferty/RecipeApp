package server

func Start() {
	router := setupRouter()

	router.Run(":8090")
}

### Setup
- Install go
- Install docker

### Running go server in docker
- `docker build --tag recipe-app .`
- `docker run -d --name recipe-app -p 8090:8090 recipe-app` (port is important, exposes container port on machine port). -d runs in background
- visit <serverip>:8090/api/recipe

### Running go server locally
- Run `go run main.go`
- visit localhost:8090/api/recipe
## Setup
- Install go
- Install docker
- Install postgres 18
- Set the password as postgres
- Create a postgres database called recipe_app

## Running database migrations
### Initial setup
- go run go/migrations/*.go init
- go run *.go up

### Perform migration (run every git pull)
- go run *.go up

## Running go server in docker
- `docker build --tag recipe-app .`
- `docker run -d --name recipe-app -p 8090:8090 recipe-app` (port is important, exposes container port on machine port). -d runs in background
- visit <serverip>:8090/api/recipe

## Running go server locally
- Run `go run go/main.go`
- visit localhost:8090/api/recipe

## Running frontend locally
- Run `npm run dev`
- visit localhost:5173

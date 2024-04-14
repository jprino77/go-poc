package main

import "github.com/jprino77/go-poc/cmd/web"

func main() {

	deps := web.InitDependencies()
	router := web.InitRoutes(deps)
	web.InitServer(router)
}

package main

import "github.com/indrabay/helloibe-api/pkg"

func main() {
	server := pkg.StartServer()

	server.Run(":9999")
}

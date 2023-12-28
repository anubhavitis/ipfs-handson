package main

import "fileverse-test/src/server"

func main() {
	r := server.GetServer()

	// Run the server
	r.Run(":8080")

}

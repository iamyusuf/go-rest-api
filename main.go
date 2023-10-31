package main

var (
	GET    = "GET"
	POST   = "POST"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

func main() {
	server := NewAPIServer(":8081")
	server.Run()
}

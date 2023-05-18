package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", "GET", server.AddMiddleware(HandleRoot, Logging()))
	server.Handle("/home", "GET", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Handle("/create", "POST", server.AddMiddleware(PostRequest, Logging()))
	server.Handle("/create/user", "POST", server.AddMiddleware(UserPostRequest, Logging()))
	server.Listen()
}

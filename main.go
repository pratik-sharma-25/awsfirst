package main

import "net/http"

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /greet", greetUser)
	http.ListenAndServe(":7080", router)
}

func greetUser(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("Welcome user!!"))
}

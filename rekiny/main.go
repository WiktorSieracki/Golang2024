package main

import (
	"fmt"
	"net/http"
	"rekiny/routes"
)


func main() {
	data, err := Load_data()
	if err != nil {
		fmt.Println(err)
		return
	}
	mux := http.NewServeMux()
    mux.HandleFunc("GET /get", routes.Get(&data))
    mux.HandleFunc("POST /post", routes.Post(&data))
    mux.HandleFunc("PUT /update", routes.Update(&data))
    mux.HandleFunc("DELETE /delete", routes.Delete(&data))

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mux)
}
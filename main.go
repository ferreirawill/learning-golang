package main

import ("fmt"
		"net/http"
		"log")

func main(){
	fmt.Println("Initing API")

	server:= Server{}

	server.initalizeRoutes()
	server.initDatabase()
	
	
	if err := http.ListenAndServe(":5000",server); err != nil {
		log.Fatalf("Could not listem on port 5000 %v",err)
	}
}
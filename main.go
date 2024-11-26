package main

import (
	"fmt"
	"log"
	"net/http"
	"gorm_emp/controller"
	"gorm_emp/database"

	"github.com/gorilla/mux"
)

func main(){
    //  db initialization
	db ,error:= database.InitDB()
	if error!=nil{
		log.Fatalf("Db initialisation failed...")
	}

	router:= mux.NewRouter()
	router.HandleFunc("/api/emp", controller.CreateEmp(db)).Methods("POST")
    router.HandleFunc("/api/emp", controller.GetEmp(db)).Methods("GET")
    router.HandleFunc("/api/emp/{id}", controller.GetbyId(db)).Methods("GET")
    router.HandleFunc("/api/emp/{id}", controller.UpdateId(db)).Methods("PUT")
    router.HandleFunc("/api/emp/{id}", controller.DeleteId(db)).Methods("DELETE")

	
	fmt.Println("Server starting on: 8085")
	log.Fatal(http.ListenAndServe(":8085", router))
}
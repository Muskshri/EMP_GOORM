package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"gorm_emp/model"
	"gorm_emp/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateEmp(db *gorm.DB) http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	 //unmarshelling -->json to struct
	 var emp model.Emp
	 err:= json.NewDecoder(r.Body).Decode(&emp)
	 if err!= nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
			return
	 }
	 //createServer bhejna h 
	 emp, err = service.CreateService(emp, db)
	 if err!= nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	 }
	 //marshal karna h--> string to json
	 w.WriteHeader(http.StatusCreated)
     json.NewEncoder(w).Encode(emp)
	}
}


func GetEmp(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
     //unmarshal
	 var emp []model.Emp
	 //get the data
     emp, err:= service.GetService(db)
	 if err!= nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	 }
    //unmarshal
	w.WriteHeader(http.StatusOK)
     json.NewEncoder(w).Encode(emp)

	}
}

func GetbyId(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
        param:= mux.Vars(r)
		id, _:= strconv.Atoi(param["id"])
        
		//get from service
		var emp model.Emp
		emp, err:= service.GetbyIdService(id, db)
        if err!= nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		 }

		 w.WriteHeader(http.StatusOK)
		 json.NewEncoder(w).Encode(emp)
	}
}

func UpdateId(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		param:= mux.Vars(r)
		id, _:= strconv.Atoi(param["id"])

		var emp model.Emp
		err := json.NewDecoder(r.Body).Decode(&emp)
		if err!= nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
				return
		 }

		emp, err = service.UpdateIdService(id, emp, db)
		if err!= nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		 }
        
		 w.WriteHeader(http.StatusCreated)
		 json.NewEncoder(w).Encode(emp)
		
		//id , emp,db
	}
}

func DeleteId(db *gorm.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		param:= mux.Vars(r)
		id, _:= strconv.Atoi(param["id"])

		err:= service.DeleteService(id,db)
		if err!= nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
    

}

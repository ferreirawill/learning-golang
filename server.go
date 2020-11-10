package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"gitlab.com/Ferreira.will/ormapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



type Server struct{
	DB *gorm.DB
    http.Handler
}




func (s *Server)initalizeRoutes() {
	server := new(Server)
	
	router := http.NewServeMux()
	router.Handle("/user",http.HandlerFunc(s.handleUsers))

	server.Handler = router

	s.Handler = server

}

func (s *Server)handleUsers(w http.ResponseWriter,r *http.Request){
	p :=models.Person{}
	switch r.Method {
		case http.MethodGet:

			usrList,_ := p.ReadAll(s.DB) 
			jsonList,_ := p.PersonListToJsonList(usrList)
			fmt.Fprint(w,jsonList)
		
		case http.MethodPost:
		
			body,_ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body,&p)
			p.Create(s.DB)
			w.Header().Set("Content-Type","application/json")
			jsonBody,_ := json.Marshal(p)
			fmt.Fprint(w,bytes.NewBuffer(jsonBody))
		
		
		case http.MethodPut:
		
			fmt.Fprint(w,"hello")
		
		case http.MethodDelete:
			body,_ := ioutil.ReadAll(r.Body)
			json.Unmarshal(body,&p)
			p.Delete(s.DB)
			w.Header().Set("Content-Type","application/json")
			jsonBody,_ := json.Marshal(p)
			fmt.Fprint(w,bytes.NewBuffer(jsonBody))
		}


}


func (s *Server) initDatabase(){
	var err error
	
	if err = godotenv.Load(); err !=nil{
		log.Println("can't open .env file")
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/shanghai",
	os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),os.Getenv("DB_PORT")) 

	s.DB, err = gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatalf("Nao abriu o db")
	}
	

	if(!s.DB.Migrator().HasTable(&models.Person{})){
		s.DB.Migrator().CreateTable(&models.Person{})
	}
	if(!s.DB.Migrator().HasTable(&models.Employment{})){
		s.DB.Migrator().CreateTable(&models.Employment{})
	}
	if(!s.DB.Migrator().HasTable(&models.Account{})){
		s.DB.Migrator().CreateTable(&models.Account{})
	}
	fmt.Println(s.DB)

}
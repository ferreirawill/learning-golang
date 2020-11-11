package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
	"gitlab.com/Ferreira.will/ormapi/api/auth"
	"gitlab.com/Ferreira.will/ormapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



type Server struct{
	DB *gorm.DB
	RedisClient *redis.Client
	http.Handler

}





func (s *Server)initRedis(){

	dsn := os.Getenv("REDIS_DSN")

	if len(dsn) == 0{
		dsn = "localhost:6379"
		os.Setenv("REDIS_DSN",dsn)
	}

	s.RedisClient = redis.NewClient(&redis.Options{
		Addr: dsn,
	})
	
	_,err :=s.RedisClient.Ping().Result()

	if err != nil {
		panic(err)
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


func (s *Server)initalizeRoutes() {
	server := new(Server)
	router := http.NewServeMux()
	router.Handle("/user",http.HandlerFunc(s.handleUsers))
	router.Handle("/login",http.HandlerFunc(s.handleLogin))
	router.Handle("/register",http.HandlerFunc(s.handleRegister))
	server.Handler = router

	s.Handler = server

}


func (s *Server) handleLogin(w http.ResponseWriter,r *http.Request){
	var err error
	
	if r.Method != http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w)
		return
		//log.Fatalf("handleLogin - Method not valid  | Expected: %v , Received: %v",http.MethodPost,r.Method)
	}

	if r.Body == nil{
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w)
		return
		//log.Fatalf("handleLogin - Body is nil")
	}

	u := models.User{}
	body,err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body,&u)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w)
		return
	}


	if u.Username != "admin@admin.com.br" || u.Password != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w)
		return
	 }

	var jwtToken *models.TokenDetails
	jwtToken,err = auth.CreateToken("a00a666e-a8f0-4efd-98f8-10c45eb22df5")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w)
		return
	}

	saveErr := auth.CreateAuth("a00a666e-a8f0-4efd-98f8-10c45eb22df5",jwtToken,s.RedisClient)

	if saveErr !=nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w)
		return
	}


	tokens,_ := json.Marshal(jwtToken)


	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,bytes.NewBuffer(tokens))
}

func (s *Server) handleRegister(w http.ResponseWriter,r *http.Request){	
	if r.Method != http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w)
	}
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

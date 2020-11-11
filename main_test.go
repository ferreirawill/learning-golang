package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"gitlab.com/Ferreira.will/ormapi/models"
)



func TestLogin(t *testing.T) {

	s := Server{}

	s.initalizeRoutes()
	s.initDatabase()
	s.initRedis()


	t.Run("Assert Method Fail", func(t *testing.T){
	request,_ := http.NewRequest(http.MethodGet,fmt.Sprintf("/login"),nil)
	response := httptest.NewRecorder()

	s.ServeHTTP(response,request)

	if response.Code != http.StatusMethodNotAllowed{
		t.Errorf("Wrong response | Got: %v  , Want: %v",response.Code,http.StatusMethodNotAllowed)
	}
	})

	
	t.Run("Assert Post Login", func(t *testing.T){
		u := models.User{Username:"admin@admin.com.br",Password: "admin"}
		userMarshal,_ := json.Marshal(u)
		request,_ := http.NewRequest(http.MethodPost,fmt.Sprintf("/login"),bytes.NewBuffer(userMarshal))
		response := httptest.NewRecorder()
		
		s.ServeHTTP(response,request)
		if response.Code != http.StatusOK{
			t.Errorf("Wrong response | Got: %v  , Want: %v",response.Code,http.StatusOK)
		}
		})
}










func TestGet(t *testing.T) {

	s := Server{}

	s.initalizeRoutes()
	s.initDatabase()
	
	request,_ := http.NewRequest(http.MethodGet,fmt.Sprintf("/user"),nil)
	response := httptest.NewRecorder()

	s.ServeHTTP(response,request)

	if response.Body.String() != "William"{
		t.Errorf("Wrong response %s",response.Body.String())
	}
}

func TestPost(t *testing.T)  {
	s := Server{}

	s.initalizeRoutes()
	s.initDatabase()
	
	p := models.Person{Name: "william",Age:20,Document: "12321231"}
	personMarshal,_ := json.Marshal(p)
	request,_ := http.NewRequest(http.MethodPost,fmt.Sprintf("/user"),bytes.NewBuffer(personMarshal))
	response := httptest.NewRecorder()
	
	s.ServeHTTP(response,request)
	
	newU := models.Person{}
	jsonBytes,_ := ioutil.ReadAll(response.Body)
	json.Unmarshal(jsonBytes,&newU)

	fmt.Println(newU)
	if p.Name != newU.Name{
		t.Errorf("Valor recebido diferente do esperado| Recebido: %v  Esperado: %v",newU.Name,p.Name)
	}
	
}


func TestPut(t *testing.T)  {
	s := Server{}

	s.initalizeRoutes()
	s.initDatabase()
	
	p := models.Person{Name: "william",Age:20,Document: "12321231"}
	personMarshal,_ := json.Marshal(p)
	request,_ := http.NewRequest(http.MethodPost,fmt.Sprintf("/user"),bytes.NewBuffer(personMarshal))
	response := httptest.NewRecorder()
	
	s.ServeHTTP(response,request)
	
	newU := models.Person{}
	jsonBytes,_ := ioutil.ReadAll(response.Body)
	json.Unmarshal(jsonBytes,&newU)

	fmt.Println(newU)
	if p.Name != newU.Name{
		t.Errorf("Valor recebido diferente do esperado| Recebido: %v  Esperado: %v",newU.Name,p.Name)
	}
	
}

func TestDelete(t *testing.T)  {
	s := Server{}

	s.initalizeRoutes()
	s.initDatabase()
	
	p :=models.Person{Name: "Mauricio",Age: 21,Document:"12339210"}
	personMarshal,_ := json.Marshal(p)
	request,_ := http.NewRequest(http.MethodDelete,fmt.Sprintf("/user"),bytes.NewBuffer(personMarshal))
	response := httptest.NewRecorder()
	
	s.ServeHTTP(response,request)
		
}

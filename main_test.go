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
	
	u := models.User{Name: "william",Age:20,Employment: "Developer"}
	userMarshal,_ := json.Marshal(u)
	request,_ := http.NewRequest(http.MethodPost,fmt.Sprintf("/user"),bytes.NewBuffer(userMarshal))
	response := httptest.NewRecorder()
	
	s.ServeHTTP(response,request)
	
	newU := models.User{}
	jsonBytes,_ := ioutil.ReadAll(response.Body)
	json.Unmarshal(jsonBytes,&newU)

	fmt.Println(newU)
	if u != newU{
		t.Errorf("Valor recebido diferente do esperado| Recebido: %v  Esperado: %v",newU,u)
	}
	
}


func TestPut(t *testing.T)  {
	s := Server{}

	s.initalizeRoutes()
	s.initDatabase()
	
	u := models.User{Name: "william",Age:20,Employment: "Developer"}
	userMarshal,_ := json.Marshal(u)
	request,_ := http.NewRequest(http.MethodPost,fmt.Sprintf("/user"),bytes.NewBuffer(userMarshal))
	response := httptest.NewRecorder()
	
	s.ServeHTTP(response,request)
	
	newU := models.User{}
	jsonBytes,_ := ioutil.ReadAll(response.Body)
	json.Unmarshal(jsonBytes,&newU)

	fmt.Println(newU)
	if u != newU{
		t.Errorf("Valor recebido diferente do esperado| Recebido: %v  Esperado: %v",newU,u)
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

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

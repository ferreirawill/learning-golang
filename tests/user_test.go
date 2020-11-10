package tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"encoding/json"
	"gitlab.com/Ferreira.will/ormapi/models"
)

func TestCreateUser(t *testing.T) {

	u:= OpenAndReadFile(t)

	u.CreateUser()
	fmt.Println(u.Name)
}



func OpenAndReadFile(t *testing.T) models.User{
	t.Helper()
	file, err := os.OpenFile("sample.json",os.O_RDONLY,0666)
	defer file.Close()
	
	if err != nil {
		t.Errorf("Erro ao abrir arquivo")
	}

	user := models.User{}

	jsonBytes,_ := ioutil.ReadAll(file)


	err = json.Unmarshal(jsonBytes, &user)
	
	return user
}


func assertEquals(t *testing.T,got,want string){
	t.Helper()
	if got != want {
		t.Errorf("Got: %q | Want: %q",got,want)
	}
}
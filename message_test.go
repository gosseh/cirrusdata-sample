package main

import (
    "testing"
	"encoding/json"
	"os"
	"io/ioutil"
)

func TestStuff(t *testing.T) {
	//Initialize our config struct
	configfile, _ := os.Open("config.json")
	defer configfile.Close()
	err := json.NewDecoder(configfile).Decode(&config)
	check(err)
	
	validate_getData_0(t)  
	validate_getData_1(t)
	validate_getData_2(t) 
	
	validate_verifyCred_0(t)
	validate_verifyCred_1(t)
	validate_verifyCred_2(t)
	
	validate_readMessage_0(t)
	validate_readMessage_1(t)
	validate_readMessage_2(t)
	
	validate_createMessage_0(t)
	validate_createMessage_1(t)
	validate_createMessage_2(t)
}

func validate_getData_0(t *testing.T) {
	result0, result1 := getData("https://reqres.in/api/users", "rachel.howell@reqres.in", "")
	expected0 := true
	expected1 := User{12, "rachel.howell@reqres.in", "Rachel", "Howell" , "https://reqres.in/img/faces/12-image.jpg"}
	
	if result0 != expected0 || result1 != expected1 {
		t.Errorf("Expected result %v and %v, but got %v and %v", expected0, expected1, result0, result1)
	}
}

func validate_getData_1(t *testing.T) {
	result0, result1 := getData("https://reqres.in/api/users", "qres.in", "")
	expected0 := false
	expected1 := User{}
	
	if result0 != expected0 || result1 != expected1 {
		t.Errorf("Expected result %v and %v, but got %v and %v", expected0, expected1, result0, result1)
	}
}

func validate_getData_2(t *testing.T) {
	result0, result1 := getData("https://reqres.in/api/users", "janet.weaver@reqres.in", "")
	expected0 := true
	expected1 := User{2, "janet.weaver@reqres.in", "Janet", "Weaver" , "https://reqres.in/img/faces/2-image.jpg"}
	
	if result0 != expected0 || result1 != expected1 {
		t.Errorf("Expected result %v and %v, but got %v and %v", expected0, expected1, result0, result1)
	}
}

func validate_verifyCred_0(t *testing.T) {
	users := []User{User{2, "janet.weaver@reqres.in", "Janet", "Weaver" , "https://reqres.in/img/faces/2-image.jpg"}, 
						User{3, "roberto@umich.edu", "Rec", "s" , ""}, 
						User{2, "roberto@hotmail.com", "", "" , ""}}
	
	result0, result1 := verifyCred(users, "janet.weaver@reqres.in", "password")
	expected0 := true
	expected1 := User{2, "janet.weaver@reqres.in", "Janet", "Weaver" , "https://reqres.in/img/faces/2-image.jpg"}
	
	if result0 != expected0 || result1 != expected1 {
		t.Errorf("Expected result %v and %v, but got %v and %v", expected0, expected1, result0, result1)
	}
}

func validate_verifyCred_1(t *testing.T) {
	users := []User{User{2, "janet.weaver@reqres.in", "Janet", "Weaver" , "https://reqres.in/img/faces/2-image.jpg"}, 
						User{3, "roberto@umich.edu", "Rec", "s" , ""}, 
						User{2, "roberto@hotmail.com", "", "" , ""}}
	
	result0, result1 := verifyCred(users, "roberto@hotmail.com", "password")
	expected0 := true
	expected1 := User{2, "roberto@hotmail.com", "", "" , ""}
	
	if result0 != expected0 || result1 != expected1 {
		t.Errorf("Expected result %v and %v, but got %v and %v", expected0, expected1, result0, result1)
	}
}

func validate_verifyCred_2(t *testing.T) {
	users := []User{User{2, "janet.weaver@reqres.in", "Janet", "Weaver" , "https://reqres.in/img/faces/2-image.jpg"}, 
						User{3, "roberto@umich.edu", "Rec", "s" , ""}, 
						User{}}
	
	result0, result1 := verifyCred(users, "unknown", "password")
	expected0 := false
	expected1 := User{}
	
	if result0 != expected0 || result1 != expected1 {
		t.Errorf("Expected result %v and %v, but got %v and %v", expected0, expected1, result0, result1)
	}
}

func validate_readMessage_0(t *testing.T) {
	message_path := "messages/5"
	message := "testestest"
	ioutil.WriteFile(message_path, []byte(message), 0644)
	result0, result1, result2 := readMessage(Req{"charles.morris@reqres.in", "password", message})
	expected0 := true
	expected1 := 200
	expected2 := message
	
	os.Remove(message_path)
	
	
	if result0 != expected0 && result1 != expected1 && result2 != expected2 {
		t.Errorf("Expected result %v and %v and %v, but got %v and %v and %v", expected0, expected1, expected2, result0, result1, result2)
	}
}

func validate_readMessage_1(t *testing.T) {
	message_path := "messages/5"
	message := "testestest"
	ioutil.WriteFile(message_path, []byte(message), 0644)
	result0, result1, result2 := readMessage(Req{"tracey.ramos@reqres.in", "password", message})
	expected0 := false
	expected1 := 404
	expected2 := "MESSAGE DOES NOT EXIST"
	
	os.Remove(message_path)
	
	
	if result0 != expected0 && result1 != expected1 && result2 != expected2 {
		t.Errorf("Expected result %v and %v and %v, but got %v and %v and %v", expected0, expected1, expected2, result0, result1, result2)
	}
}

func validate_readMessage_2(t *testing.T) {
	message_path := "messages/5"
	message := "testestest"
	ioutil.WriteFile(message_path, []byte(message), 0644)
	result0, result1, result2 := readMessage(Req{"traceymos@reqres.in", "password", message})
	expected0 := false
	expected1 := 401
	expected2 := "INVALID CREDENTIALS"
	
	os.Remove(message_path)
	
	
	if result0 != expected0 && result1 != expected1 && result2 != expected2 {
		t.Errorf("Expected result %v and %v and %v, but got %v and %v and %v", expected0, expected1, expected2, result0, result1, result2)
	}
}

func validate_createMessage_0(t *testing.T) {
	message_path := "messages/5"
	message := "testestest"
	result0, result1, result2 := createMessage(Req{"tracey.ramos@reqres.in", "password", message})
	expected0 := true
	expected1 := 203
	expected2 := message
	
	os.Remove(message_path)
	
	
	if result0 != expected0 && result1 != expected1 && result2 != expected2 {
		t.Errorf("Expected result %v and %v and %v, but got %v and %v and %v", expected0, expected1, expected2, result0, result1, result2)
	}
}

func validate_createMessage_1(t *testing.T) {
	message_path := "messages/5"
	message := "testestest"
	result0, result1, result2 := createMessage(Req{"trey.ramos@reqres.in", "password", message})
	expected0 := false
	expected1 := 401
	expected2 := "INVALID CREDENTIALS"
	
	os.Remove(message_path)
	
	
	if result0 != expected0 && result1 != expected1 && result2 != expected2 {
		t.Errorf("Expected result %v and %v and %v, but got %v and %v and %v", expected0, expected1, expected2, result0, result1, result2)
	}
}

func validate_createMessage_2(t *testing.T) {
	message_path := "messages/5"
	message := "testedddddddddddddddddddddddddddddddsaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaadwdwadaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaadswwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwstest"
	result0, result1, result2 := createMessage(Req{"tracey.ramos@reqres.in", "password", message})
	expected0 := false
	expected1 := 413
	expected2 := "MESSAGE TOO LONG"
	
	os.Remove(message_path)
	
	
	if result0 != expected0 && result1 != expected1 && result2 != expected2 {
		t.Errorf("Expected result %v and %v and %v, but got %v and %v and %v", expected0, expected1, expected2, result0, result1, result2)
	}
}


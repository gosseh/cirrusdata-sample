package main
	
import (
    "fmt"
	"encoding/json"
	"net/http"
	"strconv"
	"os"
	"io/ioutil"
	"log"
)


// Struct for configuration from config.json, this file can be modified to change these values
type Configuration struct{
	APIurl string
	MSGdir string
}
// Global variable for persistent access of configuration
var config Configuration

// Struct to hold requests to our server
type Req struct{
	Username string
	Password string
	Message  string
}

// Struct for JSON of API
type Page struct {
	Page	    int
    Per_page    int
	Total_pages int
	Data        []User
}

// Struct for user JSON
type User struct {
	ID         int
    Email      string
	First_name string
	Last_name  string
	Avatar     string
}

// Generic error handler
func check(err error) {
    if err != nil {
        panic(err)
    }
}

// Searches for the username in the slice of all users, can be changed later for SQL calls
func verifyCred(all_users []User, username string, password string) (bool, User) {
	for _, user := range all_users {
		if user.Email == username {
			return true, user
		}
	}
	
	return false, User{}
}

// Gets all the users from the API and searches for the credentials, can be changed later for SQL calls
func getData(url string, username string, password string) (bool, User){
	all_users := []User{}
	
	// Get number of pages from first request
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	page_info := Page{}
    err = json.NewDecoder(resp.Body).Decode(&page_info)
    check(err)
	
	// Iterate through all pages 
	total_pages := page_info.Total_pages
    for i := 1; i <= total_pages; i++ {
        url_param := url + "?page=" + strconv.Itoa(i)
		
		resp, err := http.Get(url_param)
		check(err)
		defer resp.Body.Close()
		page_info := Page{}
		err = json.NewDecoder(resp.Body).Decode(&page_info)
		check(err)
		
		all_users = append(all_users, page_info.Data...)
    }
	

	return verifyCred(all_users, username, password)
}

// Handle the parsing of GET requests
func readMessage(req Req) (success bool, code int, message string){
	found, user := getData(config.APIurl, req.Username, req.Password)
	if found != true{
		return false, 401, "INVALID CREDENTIALS"
	}
	
	message_path := config.MSGdir + strconv.Itoa(user.ID)

	// Check if user has a message file
	if _, err := os.Stat(message_path); os.IsNotExist(err) {
		return false, 404, "MESSAGE DOES NOT EXIST"
	}
	
	data, err := ioutil.ReadFile(message_path)
    check(err)
	message = string(data)
	
	check(err)
	fmt.Println("Reading File " + message_path + " ...")
	return true, 200, message
}

// Handle the parsing of POST requests
func createMessage(req Req) (success bool, code int, message string){
	found, user := getData(config.APIurl, req.Username, req.Password)
	if found != true{
		return false, 401, "INVALID CREDENTIALS"
	}
	if len(req.Message) > 256{
		return false, 413, "MESSAGE TOO LONG"
	}
	
	message_path := config.MSGdir + strconv.Itoa(user.ID)
	
	err := ioutil.WriteFile(message_path, []byte(req.Message), 0644)
	check(err)
	fmt.Println("Writing File " + message_path + " ...")
	return true, 201, req.Message
}

func getReq(w http.ResponseWriter, r *http.Request) {
	req := Req{}
	err := json.NewDecoder(r.Body).Decode(&req)
    check(err)
	
    w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "GET":
		success, code, message := readMessage(req)
		body := "{\"message\": \"" + message + "\"}"
		if !success {
			w.WriteHeader(code)
			w.Write([]byte(body))
			break
		}
        w.WriteHeader(code)
        w.Write([]byte(body))
    case "POST":
		success, code, message := createMessage(req)
		body := "{\"message\": \"" + message + "\"}"
		if !success {
			w.WriteHeader(code)
			w.Write([]byte(body))
			break
		}
        w.WriteHeader(code)
        w.Write([]byte(body))
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "Error. Bad request."}`))
    }
}

func handleRequests() {
    http.HandleFunc("/", getReq)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	//Initialize our config struct
	configfile, _ := os.Open("config.json")
	defer configfile.Close()
	err := json.NewDecoder(configfile).Decode(&config)
	check(err)
	
	handleRequests()
}




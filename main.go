package main

import (
	"encoding/json"
	"io"
	"net/http"
	"path/filepath"
)

func main() {
	//initialize router

	//start server

}

//setup up router function
//initalize a new chi router
//set up the middle wares
//Logger
//Recoverer
//cors
// set cors option for origin
// set cors option for the methods

// registor routes

//create a function to registor all of the routes

//create routes for each of the end points
//assign method to handler functions

//create handler for each endpoint
//create the CRUD function for each of the endpoints

///Contact

var contactFilePath = filepath.Join("data", "contact.json")

func UpdateContact(w http.ResponseWriter, r *http.Request) {

	var contact Contact

	if err := ReadJson(contactFilePath, &contact); err != nil {
		http.Error(w, "Error Reading contact from contact.json", http.StatusInternalServerError)
		return
	}

	var newContact Contact

	if err := json.NewDecoder(r.Body).Decode(&newContact); err != nil {
		http.Error(w, "Error decoding Contact from request", http.StatusBadRequest)
		return
	}
	contact = newContact

	if err := WriteJson(contactFilePath, &contact); err != nil {
		http.Error(w, "Error writing new contact to contact.json", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newContact)
	return
}

func GetContacts(w http.ResponseWriter, r *http.Request) {

}

///Jobs

///Projects

//Util function to Read and write Json

func ReadJson(filepath string, target interface{}) error {
	file, err := os.open(filepath)
	if err != nil {
		return err
	}

	defer file.close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, target)
}

func WriteJson(filepath string, data interface{}) error {
	file, err := os.open(filepath)
	if err != nil {
		return err
	}

	defer file.close()

	jsonData, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		return err
	}

	_, err := file.write(file, jsonData)

}

//create the models for the each sections of the resumes

type Resume struct {
	Contact  Contact
	Jobs     []Jobs
	Projects []Projects
	Skills   []string
}

type Contact struct {
	FirstName string
	LastName  string
	Email     string
	Github    string
}

type Jobs struct {
	Company      string
	Duration     string
	Task         []string
	Title        string
	Technologies []string
}

type Projects struct {
	Name        string
	Description string
	Stack       string
	Purpose     string
	Github      string
}

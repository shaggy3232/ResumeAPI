package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/shaggy3232/ResumeAPI/models"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", ResumeHandler)
	r.HandleFunc("/contact", ContactHandler).Methods("GET")
	r.HandleFunc("/projects", ProjectHandler).Methods("GET")
	r.HandleFunc("/work", WorkHandler).Methods("GET")

	http.Handle("/", r)

	http.ListenAndServe(":8000", r)
}

func ResumeHandler(w http.ResponseWriter, r *http.Request) {
	Resume, err := GetFullResume()
	if err != nil {
		fmt.Println("could not get resume from json")
		return
	}

	encode(w, r, http.StatusOK, Resume)

}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	Contact, err := GetContact()
	if err != nil {
		fmt.Println("could not get contact from the resume")
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusOK, Contact)
}

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := GetProjects()

	if err != nil {
		fmt.Println("could not get projects")
		encode(w, r, http.StatusInternalServerError, err)
		return

	}

	encode(w, r, http.StatusOK, projects)
}

func WorkHandler(w http.ResponseWriter, r *http.Request) {
	works, err := GetWork()

	if err != nil {
		fmt.Println("could not get works from resume")
		encode(w, r, http.StatusInternalServerError, err)
		return
	}

	encode(w, r, http.StatusOK, works)
}

//get the server running and get http requests

//create functions to handle the http requests

//handlers the requests are to parset models and return json data

// write functions that will get the data from the json files and return that data

// create the models that the json data gets parsed into

func GetFullResume() (models.Resume, error) {
	var resume models.Resume
	jsonFile, err := os.ReadFile("./data/resume.json")

	if err != nil {
		fmt.Println("could not read resume.json")
		return resume, err
	}

	json.Unmarshal(jsonFile, &resume)

	return resume, nil

}

func GetContact() (models.Contact, error) {
	var resume models.Resume
	jsonFile, err := os.ReadFile("./data/resume.json")

	if err != nil {
		fmt.Println("could not read resume.json")
		return resume.Contact, err
	}

	json.Unmarshal(jsonFile, &resume)

	return resume.Contact, nil
}

func GetProjects() ([]models.Project, error) {
	var resume models.Resume
	jsonFile, err := os.ReadFile("./data/resume.json")

	if err != nil {
		fmt.Println("could not read resume.json")
		return resume.Projects, err
	}

	json.Unmarshal(jsonFile, &resume)

	return resume.Projects, nil
}

func GetWork() ([]models.Work, error) {
	var resume models.Resume
	jsonFile, err := os.ReadFile("./data/resume.json")

	if err != nil {
		fmt.Println("could not read resume.json")
		return resume.Work, err
	}

	json.Unmarshal(jsonFile, &resume)

	return resume.Work, nil
}

func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("Encode json: %w", err)
	}

	return nil
}

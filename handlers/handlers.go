package handlers

import (
	"encoding/json"
	"fmt"
	"goweb/db"
	"goweb/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json") // HEADERS para respuestas Json
	// rw.Header().Set("Content-Type", "text/xml") // HEADERS para respuestas XML
	// Para respuestas yaml no es necesario setear los headers

	db.Conenct()
	users := models.ListUsers()
	db.Close()
	// output, _ := xml.Marshal(users) // salida XML
	// output, _ := yaml.Marshal(users) // salida yaml
	output, _ := json.Marshal(users) // salida Json
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	// Obtener id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	db.Conenct()
	user := models.GetUser(userId)
	db.Close()
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	// Obtener registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Conenct()
		user.Save()
		db.Close()
	}

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	// Obtener registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Conenct()
		user.Save()
		db.Close()
	}

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	// Obtener id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	db.Conenct()
	user := models.GetUser(userId)
	user.Delete()
	db.Close()

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

package main

import (
	"goweb/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Rutas
	mux := mux.NewRouter()
	//Endpoints
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", mux))
}

/* func dbMethodsTest() {
	db.Conenct()

	db.Ping()
	fmt.Println(db.ExistsTable("users"))

	db.CreateTable(models.UserSchema, "users")
	fmt.Println(models.ListUsers())
	db.TruncateTable("users")

	user := models.CreateUser("jane smith", "654321", "jane.s@mail.com")
	fmt.Println(user)

	users := models.ListUsers()
	fmt.Println(users)

	user := models.GetUser(2)
	fmt.Println(user)

	user.Username = "juan perez"
	user.Save()
	fmt.Println(models.ListUsers())

	user.Delete()

	db.Close()
}
*/

package main

import (
	"goweb/db"
)

func main() {
	db.Conenct()

	// db.Ping()
	// fmt.Println(db.ExistsTable("users"))

	// db.CreateTable(models.UserSchema, "users")
	// fmt.Println(models.ListUsers())
	// db.TruncateTable("users")

	// user := models.CreateUser("jane smith", "654321", "jane.s@mail.com")
	// fmt.Println(user)

	// users := models.ListUsers()
	// fmt.Println(users)

	// user := models.GetUser(2)
	// fmt.Println(user)

	// user.Username = "juan perez"
	// user.Save()
	// fmt.Println(models.ListUsers())

	// user.Delete()

	db.Close()
}

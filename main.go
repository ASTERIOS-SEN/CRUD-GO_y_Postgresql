package main

import (
	//"fmt"
	"fmt"
	"gopgsql/db"
	"gopgsql/models"
)


func main() {
	db.Connect()
	//fmt.Println(db.ExistsTable("users"))
	//db.CreateTable(models.UserSchema,"users")
	//user := models.CreateUser("Danny", "Data123", "sevilla@gmail.com")
	//fmt.Println(user)
	//db.TruncateTable("users")
	//db.TruncateTable("users")
	user := models.GetUser(5)
	fmt.Println(user)
	/*user.Usename = "Juan"
	user.Password = "juan123"
	user.Email = "juan@gmail.com"
	user.Save()*/

	user.Delete()

	

	fmt.Println(models.ListUsers())
	//db.Ping()
	db.Close()
}
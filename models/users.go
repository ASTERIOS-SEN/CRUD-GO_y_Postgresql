package models

import (
	//"database/sql"
	"gopgsql/db"
	//"os/user"
)

type User struct {
	Id       int64
	Usename  string
	Password string
	Email    string
}

type Users []User

const UserSchema string = `create table users(
	id serial primary key,
	username varchar(30) not null,
	password varchar(100) not null,
	email varchar(50),
	create_data timestamp default current_timestamp)`

//Construir usuario
func NewUser(username, password, email string) *User {
	user := &User{Usename: username, Password: password, Email: email}
	return user
}

//Crear usuario e insertar bd
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.insert()
	return user
}

//Insertar Registro
func (user *User) insert() {
	sql := "INSERT INTO users(username, password, email) values ($1, $2, $3)"
	result, _ := db.Exec(sql, user.Usename, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

//Listar todos los registros
func ListUsers() Users {
	sql := "select id, username, password, email from users"
	users := Users{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Usename, &user.Password, &user.Email) //Revisar el fallo de variables methodo no espeficado
		users = append(users, user)
	}

	return users	
}

//Obtener un registro 
func GetUser(id int) *User {
	user := NewUser("","","")

	sql :=  "select id, username, password, email from users where id=$1"
	rows, _:= db.Query(sql, id)

	for rows.Next(){	
		rows.Scan(&user.Id, &user.Usename, &user.Password, &user.Email)
	}

	return user
}

//Actualizar registro
func (user *User) update(){
	sql := "UPDATE users set username=$1, password=$2, email=$3 where id=$4"
	db.Exec(sql, user.Usename,user.Password,user.Email,user.Id)
}

//Guardar o editar un registro
func (user *User) Save(){
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

//Eliminar un registro
func (user *User) Delete(){
	sql :=	"delete from users where id=$1"
	db.Exec(sql, user.Id)
}




/*func (user *User) insert() {
sql := "Insert into users (username, password, email) values ($1, $2, $3)"
result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
user.id, _ = result.LastInsertId()
}*/

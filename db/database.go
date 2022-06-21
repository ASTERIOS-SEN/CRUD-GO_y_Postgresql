package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Data123"
	dbname   = "blog_db"
)

//Guarda la conexion
var db *sql.DB

//Realiza conexion
func Connect() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	connection, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")

	db = connection
}

//Cierra la conexion
func Close() {
	db.Close()
}

//Verificar la conexion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//Verifica si una tabla ya existe

func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SELECT * FROM information_schema.tables WHERE table_schema = '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return rows.Next()
}

//Reniciar el registro de una tabla 
func TruncateTable(tableName string){
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	db.Exec(sql)
}

//Polimorfismo a Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

//Polimorfismo a Query
func Query(query string, args ...interface{}) (*sql.Rows, error){
	rows, err := db.Query(query, args...)
	if err != nil{
		fmt.Println(err)
	}

	return rows, err
}

//Crea una tabla
func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

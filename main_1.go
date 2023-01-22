package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"text/template"
	"log"
	"net/http"
	"os"
)

// define a user model
type User struct {
	Id    int
	Name  string
	Email string
	Password string
}

// load .env file
func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// connect to the database and return it as an object
func dbConn() (db *sql.DB) {
	// pass the db credentials into variables
	host := goDotEnvVariable("DBHOST")
	port := goDotEnvVariable("DBPORT")
	dbUser := goDotEnvVariable("DBUSER")
	dbPass := goDotEnvVariable("DBPASS")
	dbname := goDotEnvVariable("DBNAME")
	// create a connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, dbUser, dbPass, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

//define a template
var tmpl = template.Must(template.ParseGlob("forms/*"))

//define an Index function that includes the http write and read parameters
func Index(w http.ResponseWriter, r *http.Request) {
	//connect to the database
	db := dbConn()
	 // Query the database and return all rows from the user table
	rows, err := db.Query(`SELECT "id", "email", "name" FROM public."users"`)
	  //Handle any errors
	CheckError(err)
		//Define and populate a User struct from the returned data from the query
	usr := User{}
		//The list of Users that will be passed to the html template
	res := []User{}
		//Loop through each row and populate a User
	for rows.Next() {
		var id int
		var email, name string
		err = rows.Scan(&id, &email, &name)
		CheckError(err)
		usr.Id = id
		usr.Email = email
		usr.Name = name
		res = append(res, usr)
	}
		//write to the Index template that will range through the User struct displaying a field for the returned data
	tmpl.ExecuteTemplate(w, "Index", res)
	//close the database connection
	defer db.Close()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}


func main() {
	//Provide address server will be provided on
	log.Println("Server started on: http://localhost:8080")
	//Create and serve a route for the Index function
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
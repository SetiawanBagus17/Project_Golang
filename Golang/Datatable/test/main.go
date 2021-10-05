package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/ddynamic/godatatables"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	env, _ := godotenv.Read("test/.env")

	if env["DATABASE_URL"] == "" {
		env["DATABASE_URL"] = "root:password@tcp(127.0.0.1:3306)/demo1?parseTime=true&charset=utf8mb4,utf8"
	}

	godotenv.Write(env, "test/.env")

	godotenv.Load("test/.env")

	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Println(err)
	}

	tmpl := template.Must(template.ParseFiles("test/test.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		godatatables.DataTables(w, r, db, "employee", "", "",
			godatatables.Column{Name: "id", Display: "Id"},
			godatatables.Column{Name: "emp_name", Display: "Name"},
			godatatables.Column{Name: "salary", Display: "Salary"},
			godatatables.Column{Name: "gender", Display: "Gender"},
			godatatables.Column{Name: "email", Display: "Email"})
	})

	http.ListenAndServe(":8080", nil)
}

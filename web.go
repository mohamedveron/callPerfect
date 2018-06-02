package main

import (
    "fmt"
    "html/template"
    "strings"
    "database/sql"
    "log"
    "net/http"
    "strconv"

    _ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "01117042116vero"
    dbName := "negro"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return db
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {

        var flag = checkUser(r)
        if flag{
            t, _ := template.ParseFiles("index.gtpl")
            t.Execute(w, nil)
        }else{
             t, _ := template.ParseFiles("login.gtpl")
             t.Execute(w, nil)
        }
        
    }
}

func Insert(r *http.Request) {
    db := dbConn()
    if r.Method == "POST" {
        id, err := strconv.Atoi(r.Form.Get("id"))
        address := r.Form["address"][0]
        name := r.Form["username"][0]
        password := r.Form["password"][0]
        insForm, err := db.Prepare("INSERT INTO student(id, address, name, password) VALUES(?,?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(id, address, name, password)
        log.Println("INSERT: Name: ")
        
    }
    defer db.Close()
}

func register(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("form/register.gtpl")
        t.Execute(w, nil)
       // tmpl.ExecuteTemplate(w, "register", "res")
    } else {
        r.ParseForm()
        
        // insert into db
        Insert(r)
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    }
}

func checkUser(r *http.Request) bool{
    db := dbConn()
    r.ParseForm()
    // logic part of log in
    name := r.Form["username"][0]
    password := r.Form["password"][0]
    fmt.Println("name:", r.Form["username"])
    fmt.Println("pass:", r.Form["password"])
    var flag bool
    err := db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM student WHERE name = ? and password = ? ", name, password).Scan(&flag)
    if err != nil {
        panic(err.Error())
    }
    
    return flag
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func main() {
    http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/login", login)
    http.HandleFunc("/register", register)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
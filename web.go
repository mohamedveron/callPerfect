package main

import (
    "fmt"
    "database/sql"
    "log"
    "net/http"
    "strconv"
    "encoding/json"
    "github.com/gorilla/mux"

    _ "github.com/go-sql-driver/mysql"
)

type Person struct {
    ID        string   `json:"id,omitempty"`
    Name      string   `json:"name,omitempty"`
    Password  string   `json:"password,omitempty"`
    Address   string   `json:"address,omitempty"`
}

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

func login(w http.ResponseWriter, r *http.Request) {
   var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
     if checkUser(person){
        json.NewEncoder(w).Encode(person)
     }else{
        json.NewEncoder(w).Encode("{ok}") 
     }
    
}

func Insert(person Person) {
    db := dbConn()

        id, err := strconv.Atoi(person.ID)
        address := person.Address
        name := person.Name
        password := person.Password
        insForm, err := db.Prepare("INSERT INTO student(id, address, name, password) VALUES(?,?,?,?)")
        if err != nil {
            panic(err.Error())
        }
        insForm.Exec(id, address, name, password)
        log.Println("INSERT: Name: ")
        
    defer db.Close()
}

func register(w http.ResponseWriter, r *http.Request) {
    //params := mux.Vars(r)
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
     Insert(person)
    json.NewEncoder(w).Encode(person)
}

func checkUser(person Person) bool{
    db := dbConn()
    // logic part of log in
    name := person.Name
    password := person.Password
    fmt.Println("name:", name)
    fmt.Println("pass:", password)
    var flag bool
    err := db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM student WHERE name = ? and password = ? ", name, password).Scan(&flag)
    if err != nil {
        panic(err.Error())
    }
    
    return flag
}


func main() {
    router := mux.NewRouter()
    router.HandleFunc("/login", login).Methods("POST")
    router.HandleFunc("/register", register).Methods("POST")
    err := http.ListenAndServe(":9090", router) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
    
}
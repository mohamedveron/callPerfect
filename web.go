package main

import (
<<<<<<< HEAD
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

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func login(w http.ResponseWriter, r *http.Request) {
    setupResponse(&w, r)
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
    setupResponse(&w, r)
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
    
=======
	"log"

       "github.com/gin-gonic/gin"
	   "github.com/jinzhu/gorm"
	   "net/http"
     "time"
     "github.com/gin-contrib/cors"
     "golang.org/x/crypto/bcrypt"
	   _ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
 // companyModel describes a companyModel type
 companyModel struct {
  gorm.Model
  Company_name     string `json:"companyName" binding:"required"`
  Contact_name     string `json:"contactName" binding:"required"`
  Mobile     string `json:"mobile" binding:"required"`
  Email     string `json:"email" binding:"required"`
  Password     string `json:"password" binding:"required"`
  Address1     string `json:"address1" binding:"required"`
  Address2     string `json:"address2" binding:"required"`
  City     string `json:"city" binding:"required"`
  State     string `json:"state" binding:"required"`
  Zip     string `json:"zip" binding:"required"`
  HearAboutUs     string `json:"hearAboutUs" binding:"required"`
  TotalPhones     int `json:"totalPhones,string,omitempty"`
  Card_name     string `json:"card_name" binding:"required"`
  Card_number     string `json:"card_number" binding:"required"`
  Card_zip    string `json:"card_zip" binding:"required"`
  ExpiredDate     time.Time `json:"expiredDate" binding:"required"`
  SecurityDate     time.Time `json:"securityDate" binding:"required"`
  TermsAccepted     int `json:"termsAccepted" binding:"required"`


 }

// transformedTodo represents a formatted todo
 loginModel struct {
  Email     string `json:"email"`
  Password  string   `json:"password"`
 }

 registerModel struct{
   Company companyModel
 }

 packageModel struct{
    gorm.Model
    Type string `json:"type" binding:"required"`
    Price int   `json:"price,string,omitempty"`

    options  []optionModel `gorm:"many2many:package_options;"`

 }

 optionModel struct{
   gorm.Model
    Content string  `json:"content" binding:"required"`
 }

 feedBackModel struct{
   gorm.Model
    Content string  `json:"content" binding:"required"`
    UserID  string `json:"id" binding:"required"`
 }

)

var db *gorm.DB

func init() {
 //open a db connection
 var err error
 db, err = gorm.Open("mysql", "root:01117042116vero@/callperfect?charset=utf8&parseTime=True&loc=Local")
 if err != nil {
  panic("failed to connect database")
 }

//Migrate the schema
 db.AutoMigrate(&companyModel{})
 db.AutoMigrate(&feedBackModel{})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// createTodo add a new todo
func register(c *gin.Context) {
  var json registerModel

  err := c.BindJSON(&json)

  if err == nil {	

        json.Company.Password, _ = HashPassword(json.Company.Password)
        
        db.Save(&json.Company)		
				c.JSON(http.StatusOK, gin.H{"status": "you are signed up"})
			
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}


 c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "company item created successfully!"})

}

func login(c *gin.Context) {
  var json loginModel
  var com companyModel

  c.BindJSON(&json)

  db.Find(&com, "email = ?", json.Email)
  
  match := CheckPasswordHash(json.Password, com.Password)

  if !match {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No company found!"})
      return
 }


 c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item found successfully!"})

}

// createTodo add a new todo
func addFeedBack(c *gin.Context) {
  var json feedBackModel

  err := c.BindJSON(&json)

  if err == nil {	

        
        db.Save(&json)		
				c.JSON(http.StatusOK, gin.H{"status": "your feedback submited"})
			
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}


 c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "feedback item created successfully!"})

}

func getFeedBack(c *gin.Context) {
  var json feedBackModel
  var feeds []feedBackModel

  err := c.BindJSON(&json)

  if err == nil {	
        db.Find(&feeds, "user_id= ?", json.UserID)	
				c.JSON(http.StatusOK, gin.H{"status": feeds})
			
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}


 c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "feedback item created successfully!"})

}

func main() {
router := gin.Default()

  // handle cors problem
  router.Use(cors.Default())

v1 := router.Group("/api/v1/company")
 {
  v1.POST("/register", register)
  v1.POST("/login", login)
  v1.POST("/feedBack", addFeedBack)
  v1.POST("/getFeedBack", getFeedBack)
 }
 router.Run(":9090")
>>>>>>> password_hashing
}
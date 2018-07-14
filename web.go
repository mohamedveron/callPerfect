package main

import (

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

func main() {
router := gin.Default()

  // handle cors problem
  router.Use(cors.Default())

v1 := router.Group("/api/v1/company")
 {
  v1.POST("/register", register)
  v1.POST("/login", login)
 }
 router.Run(":9090")
}
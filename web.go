package main

import (
       "github.com/gin-gonic/gin"
	   "github.com/jinzhu/gorm"
	   "net/http"
     "log"
	   _ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
 // companyModel describes a companyModel type
 companyModel struct {
  gorm.Model
  Company_name     string `json:"name" binding:"required"`
  Contact_name     string `json:"contact" binding:"required"`
  Mobile     string `json:"mobile" binding:"required"`

 }

// transformedTodo represents a formatted todo
 transformedTodo struct {
  ID        uint   `json:"id"`
  Title     string `json:"title"`
  Completed bool   `json:"completed"`
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

// createTodo add a new todo
func register(c *gin.Context) {
  var json companyModel

  err := c.BindJSON(&json)

  if err == nil {	

        db.Save(&json)		
				c.JSON(http.StatusOK, gin.H{"status": "you are signed up"})
			
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}


 c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "company item created successfully!"})

}

func login(c *gin.Context) {
  var json companyModel
  var com companyModel

  c.BindJSON(&json)
  db.First(&com, "company_name = ?", json.Company_name)

  if com.Company_name == "" {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No company found!"})
      return
 }


 c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item found successfully!"})

}

func main() {
router := gin.Default()

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	// authorized := router.Group("/")
	// // per group middleware! in this case we use the custom created
	// // AuthRequired() middleware just in the "authorized" group.
	// authorized.Use(AuthRequired())
	// {
	// 	authorized.POST("/login", loginEndpoint)
	// 	authorized.POST("/submit", submitEndpoint)
	// 	authorized.POST("/read", readEndpoint)
	// }

v1 := router.Group("/api/v1/company")
 {
  v1.POST("/register", register)
  v1.POST("/login", login)
 }
 router.Run(":9090")
}
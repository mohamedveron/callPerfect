package main

import (
       "github.com/gin-gonic/gin"
	   "github.com/jinzhu/gorm"
	   "net/http"
     "time"
     "github.com/gin-contrib/cors"
     "golang.org/x/crypto/bcrypt"
	   _ "github.com/jinzhu/gorm/dialects/mysql"
     "log"
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

 PackageModel struct{
    gorm.Model
    Type string `json:"type" binding:"required"`
    Price int   `json:"price,string,omitempty"`
    Options  []OptionModel `json:"options" binding:"required"`
 }

 OptionModel struct{
   gorm.Model
    Content string  `json:"content" binding:"required"`
 }

 PackageOptionModel struct{
   gorm.Model
    PackageID uint
    OptionID uint
 }


 feedBackModel struct{
   gorm.Model
    Content string  `json:"content" binding:"required"`
    UserID  string `json:"id" binding:"required"`
    
 }

 SliderModel struct{
   gorm.Model
    ImagePath string  `json:"path" binding:"required"`
    Link  string `json:"link" binding:"required"`
    Deleted int   `json:"deleted,string,omitempty"`
    Content string  `json:"content" binding:"required"`
    Title string  `json:"title" binding:"required"`
    
 }

  FeatureModel struct{
   gorm.Model
    ImagePath string  `json:"path" binding:"omitempty"`
    Link  string `json:"link" binding:"omitempty"`
    Deleted int   `json:"deleted,string,omitempty"`
    Activted int   `json:"activated,string,omitempty"`
    Content string   `json:"content,omitempty"`
    Title string  `json:"title" binding:"required"`
    
 }

 SubscriberModel struct{
   gorm.Model
    Email     string `json:"email" binding:"required"`
    Activated int   `json:"activated,string,omitempty"`   
 }

 ContactUsModel struct{
   gorm.Model
    Mobile1     string `json:"mobile1" binding:"required"`
    Mobile2     string `json:"mobile2,omitempty"` 
    Address1     string `json:"address1" binding:"required"`
    Address2     string `json:"address2,omitempty"`
    Facebook     string `json:"facebook,omitempty"`
    Twitter     string `json:"twitter,omitempty"` 
    Skype     string `json:"skype,omitempty"` 
    Linkedin     string `json:"linkedin,omitempty"` 
    Youtube     string `json:"youtube,omitempty"`
    Deleted int   `json:"deleted,string,omitempty"`    
 }

 ProductModel struct{
   gorm.Model
    Name string   `json:"content,omitempty"`
    Price     float32 `json:"price,string,omitempty"`
    Activted int   `json:"activated,string,omitempty"`
    Type int  `json:"int" binding:"required"`
    
 }

 ProductImages struct{
   gorm.Model
   ImagePath string  `json:"path" binding:"omitempty"`
   ProjectID int  `json:"int" binding:"required"`
 }

 ProductTypes struct{
   gorm.Model
   Name string   `json:"content,omitempty"`
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
 db.AutoMigrate(&PackageModel{})
 db.AutoMigrate(&OptionModel{})
 db.AutoMigrate(&PackageOptionModel{})
 db.AutoMigrate(&SliderModel{})
 db.AutoMigrate(&FeatureModel{})
 db.AutoMigrate(&SubscriberModel{})
 db.AutoMigrate(&ContactUsModel{})
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

// create add a new feedback
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

func addPackageAndOptions(c *gin.Context){
  var json PackageModel
  var pkg PackageModel
  var pid uint

    c.Bind(&json)
    db.Find(&pkg, "type = ?", json.Type)

    // if package exist don't insert
    if pkg.Type == json.Type{
        pid = pkg.ID
    }else{
      db.Save(&json)
      pid = json.ID
    }
    

  for o,v := range json.Options{

    var option OptionModel
    var packageOptions PackageOptionModel
    packageOptions.PackageID = pid

    db.Find(&option, "content = ?", v.Content)

    //check if option exists
    if option.Content == v.Content{
        log.Println(o)
    }else{
       option.Content = v.Content
      db.Save(&option)
    }
   
    packageOptions.OptionID = option.ID
    db.Save(&packageOptions)

  } 

  c.JSON(http.StatusOK, gin.H{"status": "your package and it's options submited"})

}

type Row struct {
    x string
    y string

}

// get active packages and its options
func getPackageAndOptions(c *gin.Context){
  rows, err := db.Table("package_models").Select("package_models.price, option_models.content").Joins("join package_option_models on package_option_models.package_id = package_models.id").Joins("join option_models on package_option_models.option_id = option_models.id").Rows()
     
     log.Println(err)
      for rows.Next() {
         var row Row

        if err := rows.Scan(&row.x, &row.y); err != nil {
            // do something with error
        } else {
            log.Println(row)
        }
      }
}

// create add a new slider
func addSlider(c *gin.Context) {
  var json SliderModel

  err := c.BindJSON(&json)

  if err == nil {	

        db.Save(&json)		
				c.JSON(http.StatusOK, gin.H{"status": "your Slider submited"})
			
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

}

// get all active sliders
func getActivedSliders(c *gin.Context) {
  var sliders []SliderModel
            
        db.Find(&sliders, "deleted = 0")	

  c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "sliders": sliders})
}


// create add a new Feature
func addFeature(c *gin.Context) {
  var json FeatureModel

  err := c.BindJSON(&json)

  if err == nil {	

        db.Save(&json)		
				c.JSON(http.StatusOK, gin.H{"status": "your Feature submited"})
			
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

}

// get all active features
func getActivedFeatures(c *gin.Context) {
  var features []FeatureModel
            
        db.Find(&features, "deleted = 0 and activted = 1")	

  c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "features": features})
}

// create add a new Suscribe
func addSuscriber(c *gin.Context) {
  var json SubscriberModel

  err := c.BindJSON(&json)

  if err == nil {	

        db.Save(&json)		
				c.JSON(http.StatusOK, gin.H{"status": "your Suscribe submited"})
			
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

}

// get all active features
func getActiveSuscribers(c *gin.Context) {
  var Suscribers []SubscriberModel
            
  db.Find(&Suscribers, " activated = 1")	

  c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "Suscribers": Suscribers})
}

// create add a new Suscribe
func addContactUs(c *gin.Context) {
  var json ContactUsModel

  err := c.BindJSON(&json)

  if err == nil {	

        db.Save(&json)		
				c.JSON(http.StatusOK, gin.H{"status": "your ContactUs submited"})
			
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

}

// get last active contact us
func getActiveContactUs(c *gin.Context) {
  var contactus ContactUsModel
            
  db.Last(&contactus, " deleted = 0")	

  c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "Contactus": contactus})
}

// create add a new Product
func addProduct(c *gin.Context) {
  var json ProductModel

  err := c.BindJSON(&json)

  if err == nil {	

        db.Save(&json)		
				c.JSON(http.StatusOK, gin.H{"status": "your Product submited"})
			
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

}

// get active products
func getActiveProducts(c *gin.Context) {
  var products []ProductModel
            
  db.Find(&products, " activated = 0")	

  c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "products": products})
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
  v1.POST("/addPackage", addPackageAndOptions)
  v1.POST("/addSlider", addSlider)
  v1.GET("/getActivedSliders", getActivedSliders)
  v1.POST("/addFeature", addFeature)
  v1.GET("/getActivedFeatures", getActivedFeatures)
  v1.POST("/addSuscriber", addSuscriber)
  v1.GET("/getActiveSuscribers", getActiveSuscribers)
  v1.POST("/addContactUs", addContactUs)
  v1.GET("/getActiveContactUs", getActiveContactUs)
  v1.GET("/getPackages", getPackageAndOptions)
 }
 router.Run(":9090")
}


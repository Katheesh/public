package main
//gl-api
//gfaaGUBrSRuvFQdK

import (
	"fmt"
	"log"
	"context"
	"net/http"
	"os"
  "time"
  //"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/danilopolani/gocialite.v1"
	"github.com/gin-gonic/gin"

	// Mongodb connection 
	"go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/mongo"
)

// database and collection names are statically declared
const database, collection = "go-mongo-practice", "user"

var gocial = gocialite.NewDispatcher()

func main() {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.POST("/signup", SignUp)
	}
	router.Run("127.0.0.1:8090")
}

type User struct {
  Username string `form:"username"`
  Password string `form:"password"`
  Name string `form:"name"`
  Email string `form:"email"`
  Avatar string `form:"avatar"`
  Provider string `form:"provider"`
  AccessId string `form:"access_id"`
}

func SignUp(c *gin.Context) {
	
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://gl-api:gfaaGUBrSRuvFQdK@cluster0.v7znf.mongodb.net/test?retryWrites=true&w=majority"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
            log.Fatal(err)
    }else{
      fmt.Printf("connection successfull...\n")
    } 

    var user User

    if c.ShouldBind(&user) == nil {

      data := &User{
      	Username: user.Username,
        Password: user.Password,
        Name: user.Name,
        Email: user.Email,
        Avatar: user.Avatar,
        Provider: user.Provider,
        AccessId: user.AccessId,
      }

      collection := client.Database("test").Collection("users")
      insertResult, err := collection.InsertOne(context.TODO(), data)

      fmt.Print(insertResult)
      c.JSON(http.StatusOK, data)

      if err != nil {
        log.Fatal(err)
      }

    }

    /*token, err := CreateToken(253)
    if err != nil {
      c.JSON(http.StatusUnprocessableEntity, err.Error())
      return
    }*/
  	//compare the user from the request, with the one we defined:
  	/*if user.username != u.username || user.password != u.password {
     	c.JSON(http.StatusUnauthorized, "Please provide valid login details")
     	return
  	}*/

}

func CreateToken(userId uint64) (string, error) {
  	var err error
  	//Creating Access Token
  	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
  	atClaims := jwt.MapClaims{}
  	atClaims["authorized"] = true
  	atClaims["user_id"] = userId
  	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
  	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
  	if err != nil {
     	return "", err
  	}
  	return token, nil
}

func Login(c *gin.Context) {
  
  client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://gl-api:gfaaGUBrSRuvFQdK@cluster0.v7znf.mongodb.net/test?retryWrites=true&w=majority"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
            log.Fatal(err)
    }else{
      fmt.Printf("connection successfull...\n")
    }

    data := &User{
      //Username: "username",
      Password: "password",
      //Name: "",
      Email: "",
      //Avatar: "",
      //Provider: "github",
      //AccessId: "49123454322",
    }

    

    collection := client.Database("test").Collection("users")
    insertResult, err := collection.InsertOne(context.TODO(), data)

    fmt.Print(insertResult)

    token, err := CreateToken(253)
    if err != nil {
      c.JSON(http.StatusUnprocessableEntity, err.Error())
      return
    }
    
    c.JSON(http.StatusOK, token)

  if err != nil {
    log.Fatal(err)
  }

    var u User
    if err := c.ShouldBindJSON(&u); err != nil {
      c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
      return
    }
    //compare the user from the request, with the one we defined:
    /*if user.username != u.username || user.password != u.password {
      c.JSON(http.StatusUnauthorized, "Please provide valid login details")
      return
    }*/

}

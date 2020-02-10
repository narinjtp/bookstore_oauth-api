package app

import (
	"github.com/gin-gonic/gin"
	"github.com/narinjtp/bookstore_oauth-api/src/domain/access_token"
	"github.com/narinjtp/bookstore_oauth-api/src/http"
	"github.com/narinjtp/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)
func StartApplication(){

	////Redis
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	//
	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	//// Output: PONG <nil>
	//redis.NewPool()
	atService := access_token.NewService(db.NewDbRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}

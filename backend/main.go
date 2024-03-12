package main

import (
	"fmt"
	"log"
	"mlss/api"
	"time"

	pg "mlss/db/postgres"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	DB     *pg.Database
}

const webPort = ":8080"

// 之後要要反轉注入什麼一次性的 singleton
func NewServer(db *pg.Database) *Server {
	r := gin.New()
	r.Use(gin.Recovery())
	s := &Server{
		router: r,
		DB:     db,
	}

	s.routesV2()
	return s
}

// 回傳的 Type 是 *gin.Engine
func (s *Server) routesV2() {

	s.router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "authorization", "Referer"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour, // pre-flight request cache
	}))
	s.router.GET("/ping", func(c *gin.Context) {
		c.String(200, " pong")
	})
	s.router.POST("/signin", api.SignInHandler(s.DB))
	s.router.POST("/register", api.RegisterHandler(s.DB))
}
func main() {

	// db
	db, err := pg.NewDB()
	if err != nil {
		log.Printf("here is err:%+v", err) // handle err
	}
	defer db.Close()

	// 設定路由
	s := NewServer(db)
	s.router.Run(webPort)

	// 已啟動提示
	fmt.Println("Server is running on port " + webPort)

}

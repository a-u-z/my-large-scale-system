package api

import (
	"fmt"
	"log"
	pg "mlss/db/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// RegisterRequest struct represents the structure of the registration request data
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(db *pg.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodPost {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
			return
		}

		var requestBody RegisterRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Basic validation
		if requestBody.Username == "" || requestBody.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
			return
		}

		// Perform user registration logic (e.g., store user in the database)
		if err := registerUser(db, requestBody); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
			return
		}

		// Registration successful
		c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
	}
}

func registerUser(d *pg.Database, user RegisterRequest) error {

	// Check if the username already exists
	exists, err := d.FindUserByUsername(user.Username)
	if err != nil { // 處理資料庫查詢錯誤
		return err
	} else if exists { // 處理當 username 已存在的情況
		return fmt.Errorf("user exists")
	}

	// 處理當 username 尚未存在的情況
	d.InsertUser(user.Username, user.Password)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

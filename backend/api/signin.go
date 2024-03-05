package api

import (
	"bytes"
	"encoding/json"
	"io"
	pg "mlss/db/postgres"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignInHandler(db *pg.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 讀取請求的 Body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// 解析 JSON 資料
		var user User
		err = json.NewDecoder(bytes.NewReader(body)).Decode(&user)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}

		// 驗證使用者輸入
		if user.Username == "" || user.Password == "" {
			c.JSON(400, gin.H{"error": "Invalid username or password"})
			return
		}

		// 訪問資料庫
		db.QueryData()
		exist, err := db.FindOne(user.Username, user.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": "Internal server error"})
			return
		}
		if !exist {
			c.JSON(401, gin.H{"error": "Incorrect password"})
			return
		}

		// 驗證成功
		c.JSON(200, gin.H{"message": "Logged in successfully"})
	}
}

package api

import (
	"bytes"
	"encoding/json"
	"io"
	pg "mlss/db/postgres"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignInHandler(db *pg.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := jaegercfg.Configuration{
			ServiceName: "backend", // 对其发起请求的的调用链，叫什么服务
			Sampler: &jaegercfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &jaegercfg.ReporterConfig{
				LogSpans:           true,
				LocalAgentHostPort: "jaeger-agent:6831",
			},
		}

		jLogger := jaegerlog.StdLogger
		tracer, closer, err := cfg.NewTracer(
			jaegercfg.Logger(jLogger),
		)

		defer closer.Close()
		if err != nil {
		}

		// 创建第一个 span A
		parentSpan := tracer.StartSpan("A")
		defer parentSpan.Finish()

		B(tracer, parentSpan)
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
func B(tracer opentracing.Tracer, parentSpan opentracing.Span) {
	// 继承上下文关系，创建子 span
	childSpan := tracer.StartSpan(
		"B",
		opentracing.ChildOf(parentSpan.Context()),
	)
	time.Sleep(time.Second * 5)
	defer childSpan.Finish() // 可手动调用 Finish()
}

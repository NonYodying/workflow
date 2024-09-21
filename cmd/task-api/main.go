package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/NonYodying/workflow/internal/auth"
	"github.com/NonYodying/workflow/internal/item"
	"github.com/NonYodying/workflow/internal/mylog"
	"github.com/NonYodying/workflow/internal/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// POST 	/items
// GET 		/items?status=xxxxx
// PATCH	/items/:id

// GET 		/items/:id
// PUT		/items/:id
// DELETE 	/items/:id

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// FOO: Bar
	fmt.Println("FOO: ", os.Getenv("FOO"))
	fmt.Println("TEST: ", os.Getenv("TEST"))

	// Connect database
	db, err := gorm.Open(
		postgres.Open(
			//"postgres://postgres:password@localhost:5432/task",
			os.Getenv("DATABASE_URL"),
		),
	)
	if err != nil {
		log.Panic(err)
	}

	// Controller
	controller := item.NewController(db)

	// Router
	r := gin.Default()

	config := cors.DefaultConfig()
	// frontend URL
	config.AllowOrigins = []string{
		"http://localhost:8000",
		"http://127.0.0.1:8000",
	}
	r.Use(cors.New(config))
	r.Use(mylog.Logger())
	// r.Use(Logger2())
	r.GET("/version", func(c *gin.Context) {
		version, err := GetLatestDBVersion(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"version": version})
	})

	r.GET("/test", func(ctx *gin.Context) {
		fmt.Println("---- Test -----")
		// nil, false
		value, hasValue := ctx.Get("example3")
		log.Printf("default example? = %v, %T \n", value, value)
		if hasValue {
			log.Printf("example = %v, %T \n", value, value)
		} else {
			log.Println("example does not exists")
		}
		// for i := 0; i < 4; i++ {
		// 	fmt.Println(i)
		// 	time.Sleep(1 * time.Second)
		// }
		ctx.JSON(201, "test response")
	})
	userController := user.NewController(db, os.Getenv("JWT_SECRET"))
	r.POST("/login", userController.Login)

	// Register router
	items := r.Group("/items")
	// items.Use(mylog.Logger2())

	items.Use(auth.Guard(os.Getenv("JWT_SECRET")))
	{
		items.POST("", controller.CreateItem)
		items.GET("", controller.FindItems)
		items.GET("/:id", controller.GetLatest)
		items.PUT("/:id", controller.UpdateLatestItem)
		items.PATCH("/:id", controller.UpdateLatestItemStatus)
		items.DELETE("/:id", controller.DeleteLatestItem)

		items.PUT("/id/:id", controller.UpdateItem)
	}

	// Start server
	// if err := r.Run(); err != nil {
	// 	log.Panic(err)
	// }
	// endless.DefaultHammerTime = 10 * time.Second
	// if err := endless.ListenAndServe(":8080", r); err != nil {
	// 	log.Panic(err.Error())

	// }

	srv := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: r.Handler(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

type GooseDBVersion struct {
	ID        int
	VersionID int
	IsApplied bool
	Tstamp    string
}

// TableName overrides the table name used by User to `profiles`
func (GooseDBVersion) TableName() string {
	return "goose_db_version"
}

// GetLatestDBVersion returns the latest applied version from the goose_db_version table.
func GetLatestDBVersion(db *gorm.DB) (int, error) {
	var version GooseDBVersion

	// Query to get the latest version applied
	err := db.Order("version_id desc").Where("is_applied = ?", true).First(&version).Error
	if err != nil {
		return 0, err
	}

	return version.VersionID, nil
}

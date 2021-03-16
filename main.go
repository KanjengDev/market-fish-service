package main

import (
	"database/sql"
	"fmt"
	"log"
	"market-fish-service/auth"
	"market-fish-service/handler"
	"market-fish-service/helper"
	"market-fish-service/inventory"
	"market-fish-service/migration"
	"market-fish-service/user"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
	Client   *sql.DB
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//dsn local
	dsn := "root:@tcp(127.0.0.1:3306)/marketfish?charset=utf8mb4&parseTime=True&loc=Local"

	// dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	Database = db
	Client, _ = db.DB()
	autoCreate := os.Getenv("DB_MIGRATE")
	if autoCreate == "true" {
		fmt.Println("Drop and recreateing all tables....")
		migration.AutoMigrate(db)
		fmt.Println("All table recreated successfully ...")
	}

	userRepository := user.NewRepository(db)
	inventoryRepository := inventory.NewRepository(db)

	authService := auth.NewService()
	userService := user.NewService(userRepository)
	inventoryService := inventory.NewService(inventoryRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	inventoryHandler := handler.NewInventoryHandler(inventoryService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images", "./images")
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.GET("/inventory", inventoryHandler.GetInventory)

	api.POST("/inventory", authMiddleware(authService, userService), inventoryHandler.CreateCampaign)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Bearer token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized Token", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := uint(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}

}

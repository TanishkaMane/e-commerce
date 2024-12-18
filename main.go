package main

import (
	"fmt"
	"log"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tanishkamane/ecommerce/controllers"
	"github.com/tanishkamane/ecommerce/database"
	"github.com/tanishkamane/ecommerce/middleware"
	"github.com/tanishkamane/ecommerce/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loe=ading env file : ", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	Secret_key := os.Getenv("SECRET_KEY")
	fmt.Println("Secret key : ", Secret_key)

	app := controllers.NewApplication(database.GetMongoCollection(database.Client, "Products"), database.GetMongoCollection(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("addyToCart", app.AddToCart())
	router.GET("removeItem", app.RemoveItem())
	router.GET("cartCheckout", app.BuyFromCart())
	router.GET("instantBuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}

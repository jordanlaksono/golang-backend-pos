package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/pkg"
	"github.com/jordanlaksono/golang-backend-pos.git/routes"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
)

func main() {
	/**
	* ========================
	*  Setup Application
	* ========================
	 */

	db := setupDatabase()
	app := setupApp()

	/**
	* ========================
	* Initialize All Route
	* ========================
	 */

	routes.NewRouteUser(db, app)

	/**
	* ========================
	*  Listening Server Port
	* ========================
	 */

	err := app.Run(":" + pkg.GodotEnv("PORT"))

	if err != nil {
		defer logrus.Error("Server is not running ")
		logrus.Fatal(err)
	}
}

/**
* ========================
* Database Setup
* ========================
 */

func setupDatabase() *gorm.DB {

	dsn := "host=localhost user=postgres password=aero1996 dbname=ark_weekone port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Database connection failed")
		logrus.Fatal(err)
		return nil
	}

	//  Initialize all model for auto migration here
	err = db.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		defer logrus.Info("Database migration failed")
		logrus.Fatal(err)
		return nil
	}

	return db
}

/**
* ========================
* Application Setup
* ========================
 */

func setupApp() *gin.Engine {

	app := gin.Default()

	if pkg.GodotEnv("GO_ENV") != "development" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize all middleware here
	app.Use(helmet.Default())
	app.Use(gzip.Gzip(gzip.BestCompression))
	app.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"},
		AllowHeaders:    []string{"Content-Type", "Authorization", "Accept-Encoding"},
	}))

	return app
}

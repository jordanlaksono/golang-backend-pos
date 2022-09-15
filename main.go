package main

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jordanlaksono/golang-backend-pos.git/models"
	"github.com/jordanlaksono/golang-backend-pos.git/routes"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	routes.NewRouteRole(db, app)
	routes.NewRouteMenu(db, app)
	routes.NewRouteMenuRole(db, app)
	routes.NewRouteKaryawan(db, app)
	routes.NewRouteSatuan(db, app)
	routes.NewRouteLokasi(db, app)
	routes.NewRouteKelompok(db, app)
	routes.NewRouteTop(db, app)
	routes.NewRouteSupplier(db, app)
	routes.NewRouteBarang(db, app)
	routes.NewRoutePelanggan(db, app)
	routes.NewRoutePenjualan(db, app)
	/**
	* ========================
	*  Listening Server Port
	* ========================
	 */

	//err := app.Run(":" + os.Getenv("PORT"))
	err := app.Run(":8080")

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

	//dbUser := os.Getenv("DB_USER")
	//dbPass := os.Getenv("DB_PASSWORD")
	//dbHost := os.Getenv("DB_HOST")
	//dbName := os.Getenv("DB_NAME")
	//dbPort := os.Getenv("DB_PORT")

	dsn := "host=localhost user=postgres password=aero1996 dbname=ark_weekone port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName)
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

	// if pkg.GodotEnv("GO_ENV") != "development" {
	// 	gin.SetMode(gin.ReleaseMode)
	// } else {
	gin.SetMode(gin.DebugMode)
	//}

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

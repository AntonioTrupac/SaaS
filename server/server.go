package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AntonioTrupac/hannaWebshop/directives"
	"github.com/AntonioTrupac/hannaWebshop/graph/resolver"
	"github.com/AntonioTrupac/hannaWebshop/middleware"
	"github.com/AntonioTrupac/hannaWebshop/model"
	authService "github.com/AntonioTrupac/hannaWebshop/service/auth"
	moodService "github.com/AntonioTrupac/hannaWebshop/service/moods"
	productService "github.com/AntonioTrupac/hannaWebshop/service/products"
	userService "github.com/AntonioTrupac/hannaWebshop/service/users"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	logDB "github.com/AntonioTrupac/hannaWebshop/loggers"
	logging "github.com/sirupsen/logrus"
)

var database *gorm.DB

func initLog() {
	logging.SetOutput(os.Stdout)
	logging.SetFormatter(&logging.JSONFormatter{})
	logLevel, err := logging.ParseLevel(os.Getenv("LOG_LEVEL"))

	if err != nil {
		logLevel = logging.InfoLevel
	}

	logging.SetLevel(logLevel)
}

func main() {
	err := godotenv.Load(".env")

	initLog()

	if err != nil {
		logging.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		logging.Panic("Cannot read port from environment!")
	}

	initDB()
	//sqlDb, _ := database.DB()
	//defer func(sqlDb *sql.DB) {
	//	err := sqlDb.Close()
	//	if err != nil {
	//		return
	//	}
	//}(sqlDb)

	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware)

	productsService := productService.NewProducts(database)
	usersService := userService.NewUsers(database)
	moodsService := moodService.NewMoods(database)
	authServices := authService.NewAuth(database, usersService)

	c := generated.Config{Resolvers: resolver.NewResolver(usersService, productsService, moodsService, authServices)}
	c.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	fmt.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal("", http.ListenAndServe(":"+port, router))
}

func initDB() {
	var err error
	dsn := os.Getenv("DSN")

	if dsn == "" {
		logging.Fatalf("DSN string is empty")
	}

	dbSql, err := sql.Open("mysql", dsn)

	if err != nil {
		logging.Error("Something went wrong with the DSN")
	}

	database, err = gorm.Open(mysql.New(mysql.Config{
		Conn: dbSql,
	}), &gorm.Config{
		Logger:                                   logDB.LogGORM(),
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Println(err.Error())
		logging.Panic("FAILED TO CONNECT TO DB")
	}

	err = migrate(database)

	if err != nil {
		fmt.Println(err)
		logging.Panic("MODELS NOT ADDED")
	}
}

func migrate(db *gorm.DB) error {
	err := db.Debug().AutoMigrate(&model.User{}, &model.Address{}, &model.Product{}, &model.Image{}, &model.Category{}, &model.Mood{}, &model.MoodType{}, &model.UserAuth{})
	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AntonioTrupac/hannaWebshop/graph/resolver"
	logging "github.com/AntonioTrupac/hannaWebshop/loggers"
	"github.com/AntonioTrupac/hannaWebshop/model"
	productService "github.com/AntonioTrupac/hannaWebshop/service/products"
	userService "github.com/AntonioTrupac/hannaWebshop/service/users"
	"github.com/joho/godotenv"
	"go.uber.org/zap/zapcore"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	logDB "github.com/AntonioTrupac/hannaWebshop/loggers"
)

var database *gorm.DB

func main() {
	err := godotenv.Load(".env")

	myLog := logging.New(
		logging.WithDebug(),
		logging.WithCaller(),
		logging.WithLevel(zapcore.InfoLevel),
	)

	defer myLog.Log.Sync()

	if err != nil {
		// logger.Fatal("Error loading .env file")
		myLog.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		myLog.Panic("Cannot read port from environment!")
	}

	initDB()

	productsService := productService.NewProducts(database)
	usersService := userService.NewUsers(database)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver.NewResolver(usersService, productsService)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	myLog.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	myLog.Fatalf("", http.ListenAndServe(":"+port, nil))
}

func initDB() {
	var err error
	dsn := os.Getenv("DSN")

	if dsn == "" {
		log.Fatalf("DSN string is empty")
	}

	dbSql, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Print("ERRROR")
	}

	database, err = gorm.Open(mysql.New(mysql.Config{
		Conn: dbSql,
	}), &gorm.Config{
		Logger:                                   logDB.LogGORM(),
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Println(err.Error())

		panic("FAILED TO CONNECT TO DB")
	}

	err = migrate(database)

	if err != nil {
		fmt.Println(err)
		panic("MODELS NOT ADDED")
	}
}

func migrate(db *gorm.DB) error {
	err := db.Debug().AutoMigrate(&model.User{}, &model.Address{}, &model.Product{}, &model.Image{}, &model.Category{})
	if err != nil {
		return err
	}

	return nil
}

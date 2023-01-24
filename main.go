package main

import (
	"bank/handler"
	"bank/logs"
	"bank/repository"
	"bank/service"
	"fmt"
	_ "fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initializeTimeZone()
	initializeConfig()
	initializeDatabase()
	db := initializeDatabase()

	customerRepository := repository.NewCustomerRepositoryDB(db)
	customer_service := service.NewCustomerService(customerRepository)
	customerHandler := handler.InitializeCustomerHandler(customer_service)

	account_repository_db := repository.InitializeAccountRepositoryDB(db)
	account_service := service.InitializeAccountService(account_repository_db)
	account_handler := handler.NewAccountHandler(account_service)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customer_id:[0-9]+}/accounts", account_handler.GetAccounts).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/accounts", account_handler.NewAccount).Methods(http.MethodPost)

	logs.Info("Banking service started at port " + viper.GetString("app.port"))
	http.ListenAndServe(":8000", router)

}

func initializeConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initializeTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initializeDatabase() *sqlx.DB {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

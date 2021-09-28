package main

import (
	"fmt"
	"log"

	_handlerUser "bansosman/app/controllers/users"
	_routes "bansosman/app/routes"
	_servBooks "bansosman/bussiness/users"
	_repoUsers "bansosman/drivers/mysql/users"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(status string) *gorm.DB {
	db := "kampusmerdeka"
	connectionString := fmt.Sprintf("root:@tcp(127.0.0.1:3306)/%s?parseTime=True", db)

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&_repoUsers.RecordUsers{},
	)

	return DB
}

func main() {
	db := InitDB("")
	e := echo.New()

	// factory of domain
	usersRepo := _repoUsers.NewRepoMysql(db)
	usersServe := _servBooks.NewServe(usersRepo, nil)
	userHandler := _handlerUser.NewHandler(usersServe)
	// initial of routes
	routesInit := _routes.HandlerRoute{
		UsersHandler: *userHandler,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8000"))
}

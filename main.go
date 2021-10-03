package main

import (
	//!handler

	_handlerApbn "bansosman/app/controllers/apbn"
	_handlerDaerah "bansosman/app/controllers/daerah"
	_handlerUser "bansosman/app/controllers/users"

	//!middleware
	mid "bansosman/app/middleware"

	//!routes
	_routes "bansosman/app/routes"

	//!service
	_servApn "bansosman/bussiness/apbn"
	_servedaerah "bansosman/bussiness/daerah"
	_servUser "bansosman/bussiness/users"

	//!Repository
	_repoApbn "bansosman/drivers/mysql/apbn"
	_repoDerahs "bansosman/drivers/mysql/daerah"

	_repoUsers "bansosman/drivers/mysql/users"

	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func init() {
// 	viper.SetConfigFile(`./app/json.json`)
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// }

func InitDB(status string) *gorm.DB {
	db := "bansosman"
	connectionString := fmt.Sprintf("root:@tcp(127.0.0.1:3306)/%s?parseTime=True", db)

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&_repoUsers.Users{},
		&_repoApbn.Apbns{},
		&_repoDerahs.Daerahs{},
		// &_repoKab.Kabs{},
	)
	return DB
}

func main() {
	db := InitDB("")
	e := echo.New()
	mid.LogMiddlewareInit(e)
	jwSecret := "qwerty12345"
	jwtint := 2
	configJWT := mid.ConfigJwt{
		SecretJwT: jwSecret,
		Expired:   int64(jwtint),
	}

	//* factory of domain
	//user
	usersRepo := _repoUsers.NewRepoMysql(db)
	usersServe := _servUser.NewService(usersRepo, &configJWT)
	userHandler := _handlerUser.NewHandler(usersServe)
	// ?admin
	// ?apbn
	apbnRepo := _repoApbn.NewRepoMysql(db)
	apbnserve := _servApn.NewService(apbnRepo)
	apbnHandler := _handlerApbn.NewHandler(apbnserve)

	//?daerah
	daerahRepo := _repoDerahs.NewRepoMysql(db)
	daerahServe := _servedaerah.NewServe(daerahRepo)
	daerahHandler := _handlerDaerah.Newhandler(daerahServe)

	//* initial of routes
	routesInit := _routes.HandlerRoute{

		JwtMiddleware: configJWT.Init(),
		UsersHandler:  *userHandler,
		Apbnhandler:   *apbnHandler,
		Daerahhandler: *daerahHandler,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8000"))
}

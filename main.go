package main

import (
	//!handler

	_handlerAdmin "bansosman/app/controllers/admin"
	_handlerApbn "bansosman/app/controllers/apbn"
	_handlerDaerah "bansosman/app/controllers/daerah"
	_handlerUser "bansosman/app/controllers/users"

	//!middleware
	mid "bansosman/app/middleware"

	//!routes
	_routes "bansosman/app/routes"

	//!service
	_ServAdmin "bansosman/bussiness/admin"
	_servApn "bansosman/bussiness/apbn"
	_servedaerah "bansosman/bussiness/daerah"
	_servUser "bansosman/bussiness/users"

	//!Repository
	_repoAdmin "bansosman/drivers/mysql/admin"
	_repoApbn "bansosman/drivers/mysql/apbn"
	_repoDerahs "bansosman/drivers/mysql/daerah"
	_repoUsers "bansosman/drivers/mysql/users"
	_GeoRepo "bansosman/drivers/thirdparty/ipapi"

	//mysql
	_dbDriver "bansosman/drivers/mysql"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigName("docker")
	viper.AddConfigPath("./app/config/")
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_repoUsers.Users{},
		&_repoApbn.Apbns{},
		&_repoDerahs.Daerahs{},
		&_repoAdmin.Admins{},
		// &_repoKab.Kabs{},
	)
	roles := []_repoAdmin.Roles{{ID: 1, Name: "Owner"}, {ID: 2, Name: "Admin"}}
	db.Create(&roles)
}

func main() {

	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitDB()
	dbMigrate(db)
	e := echo.New()
	mid.LogMiddlewareInit(e)
	configJWT := mid.ConfigJwt{
		SecretJwT: viper.GetString(`jwt.secret`),
		Expired:   int64(viper.GetInt(`jwt.expired`)),
	}

	//* factory of domain
	//user
	usersRepo := _repoUsers.NewRepoMysql(db)
	usersServe := _servUser.NewService(usersRepo, &configJWT)
	userHandler := _handlerUser.NewHandler(usersServe)
	//?api
	GeoRepo := _GeoRepo.NewIpAPI()
	// ?admin
	adminRepo := _repoAdmin.NewMySQLRepository(db)
	adminServe := _ServAdmin.NewadminService(adminRepo, &configJWT)
	adminHandler := _handlerAdmin.NewUserController(adminServe)
	// ?apbn
	apbnRepo := _repoApbn.NewRepoMysql(db)
	apbnserve := _servApn.NewService(apbnRepo)
	apbnHandler := _handlerApbn.NewHandler(apbnserve)

	//?daerah
	daerahRepo := _repoDerahs.NewRepoMysql(db)
	daerahServe := _servedaerah.NewServe(daerahRepo, GeoRepo)
	daerahHandler := _handlerDaerah.Newhandler(daerahServe)

	//* initial of routes
	routesInit := _routes.HandlerRoute{

		JwtMiddleware:   configJWT.Init(),
		AdminController: *adminHandler,
		UsersHandler:    *userHandler,
		Apbnhandler:     *apbnHandler,
		Daerahhandler:   *daerahHandler,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8080"))
}

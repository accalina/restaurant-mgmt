package main

import (
	"github.com/accalina/restaurant-mgmt/app"
)

// @title RESTAURANT-MGMT API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	app.LoadAppConfig()
	app.New().Serve()
}

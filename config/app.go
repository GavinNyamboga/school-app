package config

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"school_app/handler"
	"school_app/model"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Connect(config *Config) {

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	//users
	a.Get("/users", a.HandleRequest(handler.GetUsers))
	a.Get("/users/{id}", a.HandleRequest(handler.GetUserById))
	a.Post("/users", a.HandleRequest(handler.CreateUser))
	a.Put("users/{id}", a.HandleRequest(handler.UpdateUser))
	a.Delete("users/{id}", a.HandleRequest(handler.DeleteUser))

	//schools
	a.Get("/schools", a.HandleRequest(handler.GetSchools))
	a.Get("/schools/{id}", a.HandleRequest(handler.GetSchoolById))
	a.Post("/schools", a.HandleRequest(handler.CreateSchool))
	a.Put("/schools/{id}", a.HandleRequest(handler.UpdateSchool))
	a.Delete("/schools/{id}", a.HandleRequest(handler.DeleteSchool))
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

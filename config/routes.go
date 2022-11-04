package config

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"school_app/handler"
)

func (a *App) SetRouters() {
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

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) HandleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}

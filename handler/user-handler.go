package handler

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/mail"
	"school_app/model"
	"school_app/utils"
	"strconv"
)

func GetUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := model.GetUsers(db)
	utils.RespondJSON(w, http.StatusOK, users)
}

func GetUserById(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	user := getUserOr404(db, ID, w)
	if user == nil {
		return
	}
	utils.RespondJSON(w, http.StatusOK, user)

}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	CreateUser := &model.User{}

	utils.ParseBody(r, CreateUser)
	//hash password
	rawPwd := CreateUser.Password
	hash, _ := HashPassword(rawPwd)
	CreateUser.Password = hash

	//validate email
	if !EmailIsValid(CreateUser.Email) {
		utils.RespondError(w, http.StatusBadRequest, "Email is invalid")
		return
	}

	CreateUser.Base = model.CreateModelBase(CreateUser.Base)
	u := CreateUser.CreateUser(db)
	utils.RespondJSON(w, http.StatusOK, u)
}
func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateUser = &model.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	userDetails := model.GetUser(db, ID)
	if updateUser.Password != "" {
		hash, _ := HashPassword(updateUser.Password)
		userDetails.Password = hash
	}
	person := updateUser.Person

	if person.FirstName != "" && person.FirstName != person.FirstName {
		userDetails.Person.FirstName = person.FirstName
	}
	if person.LastName != "" && person.LastName != userDetails.Person.LastName {
		userDetails.Person.LastName = person.LastName
	}

	if updateUser.Email != "" && updateUser.Email != userDetails.Email {
		userDetails.Email = updateUser.Email
	}

	db.Save(&userDetails)
	utils.RespondJSON(w, http.StatusOK, userDetails)
}
func DeleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	user := model.DeleteUser(db, ID)
	utils.RespondJSON(w, http.StatusOK, user)
}

// getUserOr404 gets a task instance if exists, or respond the 404 error otherwise
func getUserOr404(db *gorm.DB, id int64, w http.ResponseWriter) *model.User {
	user := model.User{}
	if err := db.First(&user, id).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &user
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func EmailIsValid(email string) bool {
	_, err := mail.ParseAddressList(email)
	return err == nil
}

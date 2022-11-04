package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"school_app/model"
	"school_app/utils"
	"strconv"
)

func GetSchools(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var schools []model.School
	db.Find(&schools)
	utils.RespondJSON(w, http.StatusOK, schools)
}

func GetSchoolById(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	schoolId := vars["id"]
	ID, err := strconv.ParseInt(schoolId, 0, 0)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	school := getSchoolOr404(db, ID, w)
	if school == nil {
		return
	}
	utils.RespondJSON(w, http.StatusOK, school)
}

func CreateSchool(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	school := model.School{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&school); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	//school.SimpleModelBase = model.CreateSimpleModelBase(school.SimpleModelBase)

	if err := db.Save(&school).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondJSON(w, http.StatusCreated, school)
}

func UpdateSchool(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	schoolId := vars["id"]
	ID, err := strconv.ParseInt(schoolId, 0, 0)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	school := getSchoolOr404(db, ID, w)
	if school == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&school); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Save(&school).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, school)

}

func DeleteSchool(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	schoolId := vars["id"]
	ID, err := strconv.ParseInt(schoolId, 0, 0)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	school := getSchoolOr404(db, ID, w)
	if school == nil {
		return
	}

	if err := db.Delete(&school).Error; err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, nil)
}

// getSchoolOr404 gets a task instance if exists, or respond the 404 error otherwise
func getSchoolOr404(db *gorm.DB, id int64, w http.ResponseWriter) *model.School {
	school := model.School{}
	if err := db.First(&school, id).Error; err != nil {
		utils.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &school
}

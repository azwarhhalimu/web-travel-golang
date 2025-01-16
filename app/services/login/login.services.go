package login

import (
	"web_traveler/app/config"
	"web_traveler/app/model"

	"golang.org/x/crypto/bcrypt"
)

func DoLogin(username string, password string) (status string, data_loginx model.TblLogin) {
	var data_login model.TblLogin
	cek := config.DB.Where("username=?", username).First(&data_login)
	if cek.RowsAffected > 0 {
		err := bcrypt.CompareHashAndPassword([]byte(data_login.Password), []byte(password))
		if err == nil {
			return "success", data_login
		}
		return "password_salah", model.TblLogin{}
	}
	return "username_not_found", model.TblLogin{}

}

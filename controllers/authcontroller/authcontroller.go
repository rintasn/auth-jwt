package authcontroller

import (
	"FinalTestLogin/config"
	"FinalTestLogin/helper"
	"FinalTestLogin/models"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// ambil data user berdasarkan username
	var user models.User
	if err := models.DB.Table("tbusers").Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// cek apakah password valid
	//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
	//	response := map[string]string{"message": err.Error()}
	//	helper.ResponseJSON(w, http.StatusBadRequest, response)
	//	return
	//}

	if user.Password != userInput.Password {
		response := map[string]string{"message": "Password Salah"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 5)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "FinalTestLogin",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// Mendeklarasikan algoritma yang akan digunakan signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// set token yang ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Login Berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)
	return
}

func Register(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan json
	//var userInput models.User
	//decoder := json.NewDecoder(r.Body)
	//if err := decoder.Decode(&userInput); err != nil {
	//	response := map[string]string{"message": err.Error()}
	//	helper.ResponseJSON(w, http.StatusBadRequest, response)
	//	return
	//}
	//defer r.Body.Close()
	//
	//// hash pass menggunakan bcrypt
	//hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	//userInput.Password = string(hashPassword)
	//
	//// insert ke database
	//if err := models.DB.Table("tbusers").Create(&userInput).Error; err != nil {
	//	response := map[string]string{"message": err.Error()}
	//	helper.ResponseJSON(w, http.StatusInternalServerError, response)
	//	return
	//}
	//
	//response := map[string]string{"message": "success"}
	//helper.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// hapus token yang ada di cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "logout berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

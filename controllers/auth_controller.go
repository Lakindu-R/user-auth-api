package controllers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"time"
	"user-auth-api/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)
var users []models.User
var userID = 1
var secretKey = []byte("mysecretkey")
// @Summary Register a new user
// @Description Create user with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body object true "User Data"
// @Success 200 {string} string "User registered"
// @Router /register [post]

// Register Function
func Register(w http.ResponseWriter,r *http.Request){
	var input struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Username == "" || input.Password ==""{
		http.Error(w,"Invalid input",http.StatusBadRequest)
		return
	}

	hash, _:= bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{
		ID: userID,
		Username: input.Username,
		Password: string(hash),
	}
	userID++
	users = append(users, user)
	w.Write([]byte("User Registerd"))
}
// Login Function
// @Summary Login user
// @Description Authenticate user and return JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body object true "Login Data"
// @Success 200 {object} map[string]string
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	for _, user := range users {
		if user.Username == input.Username {

			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
			if err != nil {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}

			claims := jwt.MapClaims{
				"username": user.Username,
				"exp":      time.Now().Add(time.Hour).Unix(),
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenStr, _ := token.SignedString(secretKey)

			json.NewEncoder(w).Encode(map[string]string{
				"token": tokenStr,
			})
			return
		}
	}

	http.Error(w, "User not found", http.StatusUnauthorized)
}
// protected function
// @Summary Protected route
// @Description Requires JWT token
// @Tags Auth
// @Produce plain
// @Security BearerAuth
// @Success 200 {string} string "Protected data"
// @Router /protected [get]
func Protected(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	fmt.Println("HEADER:", authHeader)

	if authHeader == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	tokenStr := authHeader[len("Bearer "):]
	

	_, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Protected data"))
}
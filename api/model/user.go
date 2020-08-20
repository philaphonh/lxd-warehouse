package model

import "github.com/dgrijalva/jwt-go"

// User type
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Image     string `json:"imageUrl"`
}

// UserClean type
type UserClean struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Image     string `json:"imageUrl"`
}

// UserCleanRaw type
type UserCleanRaw struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Role      int    `json:"role"`
	Image     string `json:"imageUrl"`
}

// UserRole type
type UserRole struct {
	ID       int    `json:"id"`
	RoleName string `json:"roleName"`
}

// Auth type
type Auth struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	ImageURL string `json:"imageUrl"`
	Token    string `json:"token"`
}

// AuthClaim type
type AuthClaim struct {
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

package models

import (
    "github.com/jinzhu/gorm"
)

// User représente le modèle de l'utilisateur dans la base de données
type User struct {
    gorm.Model
    Username     string    `gorm:"unique;not null"`
    Email        string    `gorm:"unique;not null"`
    PasswordHash string    `gorm:"not null"`
}

type LoginCredentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
	Email    string `json:"email"`
}


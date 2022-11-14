package database

import "github.com/Eliwelton-The-Espada/go-expert/apis/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

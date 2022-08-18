package services

import "github.com/DevProblems/sarang-gin-mongo-apis/models"

type UserService interface {
    CreateUser(*models.User) error
    GetUser(*string) (*models.User, error)
    GetAll() ([]*models.User, error)
    UpdateUser(*models.User) error
    DeleteUser(*string) error
}
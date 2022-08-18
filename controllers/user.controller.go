package controllers

import (
    "net/http"
    "github.com/DevProblems/sarang-gin-mongo-apis/models"
    "github.com/DevProblems/sarang-gin-mongo-apis/services"
    "github.com/gin-gonic/gin"
)

type UserController struct {
    UserService services.UserService
}
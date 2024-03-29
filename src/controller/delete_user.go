package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/golang/users-api/src/configuration/exception"
	"github.com/odanaraujo/golang/users-api/src/configuration/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) DeleteUser(ctx *gin.Context) {

	id := ctx.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		logger.Error("Error trying to validate id", err, zap.String(
			"Journey", "DeleteUser"))
		errMessage := exception.BadRequestException("user id is not a valid id")
		ctx.JSON(errMessage.Code, errMessage)
		return
	}

	if err := uc.service.DeleteUser(id); err != nil {
		logger.Error("Error trying delete user", err, zap.String(
			"Journey", "DeleteUser"))

		ctx.JSON(err.Code, err)
		return
	}

	ctx.Status(http.StatusOK)
}

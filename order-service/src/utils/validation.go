package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidationResponse(ctx *gin.Context, err error) {
	// Check if the error is a validation error
	if errors, ok := err.(validator.ValidationErrors); ok {
		var errorMsgs []string
		// Iterate over validation errors and construct custom error messages
		for _, e := range errors {
			errorMsgs = append(errorMsgs, fmt.Sprintf("Validation error for field %s: %s", e.Field(), e.Tag()))
		}
		// Return custom error message
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMsgs})
		return
	}
	// Return generic error message if the error is not a validation error
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

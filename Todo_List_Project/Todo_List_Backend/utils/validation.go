package utils

import (
	"fmt"
	"regexp"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateAplanumSpace(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	return regex.MatchString(fl.Field().String())
}

func ValidateBlackList(fl validator.FieldLevel) bool {
	blacklist := []string{"'", "--", "<script>", "DROP TABLE", "SELECT * FROM", "DELETE FROM", "UPDATE", "INSERT INTO", "ALTER TABLE", "CREATE TABLE", "DROP DATABASE", "CREATE DATABASE", "TRUNCATE TABLE", "ALTER DATABASE", "DROP INDEX", "DROP VIEW", "DROP TRIGGER", "DROP FUNCTION", "DROP PROCEDURE", "DROP USER", "DROP ROLE", "DROP SCHEMA"}
	for _, word := range blacklist {
		if strings.Contains(fl.Field().String(), word) {
			return false
		}
	}
	return true
}


func ValidateInput(ctx *gin.Context, input interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("alphanumericAndSpace", ValidateAplanumSpace)
	validate.RegisterValidation("blacklist", ValidateBlackList)
	err := validate.Struct(input)
	if err != nil {
		fmt.Println(err.Error())

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}

		// Print out every errors after validation
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(string(err.Error()))
		}

		return err
	}

	return nil
}
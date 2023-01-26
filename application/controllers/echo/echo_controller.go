package controller_echo

import (
	"csv-wrapper/domain/services/echo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func EchoCsv(c *gin.Context) {
	records, err := service_echo.GetCsvFromFormAsString(c)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	result, err := service_echo.GetCsvLinesInStringArray(records)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	typeOfResult := c.Query("type")
	if typeOfResult == "flatten" {
		c.IndentedJSON(http.StatusOK, strings.Join(result, ","))
		return
	} else {
		c.IndentedJSON(http.StatusOK, result)
		return
	}
}

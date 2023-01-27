package service_echo

import (
	"encoding/csv"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func GetCsvFromFormAsString(c *gin.Context) ([][]string, error) {
	fileFromForm, err := c.FormFile("file")

	if err != nil {
		return nil, err
	}

	file, err := fileFromForm.Open()

	if err != nil {
		return nil, err
	}

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		return nil, err
	}

	return records, nil
}

func GetCsvLinesInStringArray(records [][]string) ([]string, error) {
	linesQuantities := len(records)
	result := make([]string, linesQuantities)
	for index, column := range records {
		if linesQuantities != len(column) {
			return nil, errors.New("invalid size for matrix")
		}
		for _, item := range column {
			_, err := strconv.Atoi(item)
			if err != nil {
				return nil, errors.New("matrix does not contain only integers")
			}
		}
		result[index] = strings.Join(column, ",")
	}
	return result, nil
}

func SumValuesFromSpreadsheet(records []string) int {
	result := 0
	for _, item := range records {
		for _, itemSplit := range strings.Split(item, ",") {
			itemConvertedToInteger, _ := strconv.Atoi(itemSplit)
			result += itemConvertedToInteger
		}
	}
	return result
}

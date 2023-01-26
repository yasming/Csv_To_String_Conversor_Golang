package controller_echo_test

import (
	"bytes"
	"csv-wrapper/infra/http/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestReturnInvalidTypeOfRequest(t *testing.T) {
	router, responseDecoder := setupRoutes()
	req, _ := http.NewRequest("POST", "/echo", nil)
	router.ServeHTTP(responseDecoder, req)
	assert.Equal(t, 422, responseDecoder.Code)
	assert.Equal(t, "{\n    \"error\": \"request Content-Type isn't multipart/form-data\"\n}", responseDecoder.Body.String())
}

func TestReturnInvalidSizeOfSpreadsheet(t *testing.T) {
	router, responseDecoder := setupRoutes()
	body, multipartNewWriter := setupFileBody("./spreadsheets_tests/invalid_size.csv")
	req, _ := http.NewRequest("POST", "/echo", body)
	req.Header.Add("Content-Type", multipartNewWriter.FormDataContentType())
	router.ServeHTTP(responseDecoder, req)
	assert.Equal(t, 422, responseDecoder.Code)
	assert.Equal(t, "{\n    \"error\": \"invalid size for matrix\"\n}", responseDecoder.Body.String())
}

func TestReturnInvalidValuesInSpreadsheet(t *testing.T) {
	router, responseDecoder := setupRoutes()
	body, multipartNewWriter := setupFileBody("./spreadsheets_tests/invalid_matrix.csv")
	req, _ := http.NewRequest("POST", "/echo", body)
	req.Header.Add("Content-Type", multipartNewWriter.FormDataContentType())
	router.ServeHTTP(responseDecoder, req)
	assert.Equal(t, 422, responseDecoder.Code)
	assert.Equal(t, "{\n    \"error\": \"matrix does not contain only integers\"\n}", responseDecoder.Body.String())
}

func TestReturnArrayFromSpreadsheetRecords(t *testing.T) {
	router, responseDecoder := setupRoutes()
	body, multipartNewWriter := setupFileBody("./spreadsheets_tests/valid.csv")
	req, _ := http.NewRequest("POST", "/echo", body)
	req.Header.Add("Content-Type", multipartNewWriter.FormDataContentType())
	router.ServeHTTP(responseDecoder, req)
	assert.Equal(t, 200, responseDecoder.Code)
	assert.Equal(t, "[\n    \"1,2\",\n    \"3,4\"\n]", responseDecoder.Body.String())
}

func TestReturnOnlyOneLineForTypeFlattenFromSpreadsheetRecords(t *testing.T) {
	router, responseDecoder := setupRoutes()
	body, multipartNewWriter := setupFileBody("./spreadsheets_tests/valid.csv")
	req, _ := http.NewRequest("POST", "/echo?type=flatten", body)
	req.Header.Add("Content-Type", multipartNewWriter.FormDataContentType())
	router.ServeHTTP(responseDecoder, req)
	assert.Equal(t, 200, responseDecoder.Code)
	assert.Equal(t, "\"1,2,3,4\"", responseDecoder.Body.String())
}

func setupFileBody(fileName string) (*bytes.Buffer, *multipart.Writer) {
	buffer := new(bytes.Buffer)
	multipartNewWriter := multipart.NewWriter(buffer)
	fw, _ := multipartNewWriter.CreateFormFile("file", fileName)
	fd, _ := os.Open(fileName)
	defer fd.Close()
	_, _ = io.Copy(fw, fd)
	multipartNewWriter.Close()

	return buffer, multipartNewWriter
}

func setupRoutes() (*gin.Engine, *httptest.ResponseRecorder) {
	router := routes.AddRoutes()
	responseDecoder := httptest.NewRecorder()
	return router, responseDecoder
}

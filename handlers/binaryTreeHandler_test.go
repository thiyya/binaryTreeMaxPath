package handlers

import (
	"binaryTreeMaxPath/constants"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Handler(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/max-path", func(ctx *gin.Context) {
		HandleBinaryTreeMaxPath(ctx)
	})

	req, _ := http.NewRequest("POST", "/max-path", strings.NewReader(constants.TestCase1))
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	actual := w.Code
	expected := http.StatusOK
	assert.Equal(expected, actual, "Response must be OK")

	p, _ := ioutil.ReadAll(w.Body)
	assert.Contains(string(p), "\"MaxPathValue\": 18", "Test Case 1 result must be 18")

	req, _ = http.NewRequest("POST", "/max-path", strings.NewReader(constants.TestCase2))
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	actual = w.Code
	expected = http.StatusOK
	assert.Equal(expected, actual, "Response must be OK")

	p, _ = ioutil.ReadAll(w.Body)
	assert.Contains(string(p), "\"MaxPathValue\": 6", "Test Case 2 result must be 6")

	req, _ = http.NewRequest("POST", "/max-path", strings.NewReader(constants.TestCase3))
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	actual = w.Code
	expected = http.StatusOK
	assert.Equal(expected, actual, "Response must be OK")

	p, _ = ioutil.ReadAll(w.Body)
	assert.Contains(string(p), "\"MaxPathValue\": 114", "Test Case 3 result must be 114")

	req, _ = http.NewRequest("POST", "/max-path", strings.NewReader(constants.AdditionalTestCase4))
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	actual = w.Code
	expected = http.StatusOK
	assert.Equal(expected, actual, "Response must be OK")

	p, _ = ioutil.ReadAll(w.Body)
	assert.Contains(string(p), "\"MaxPathValue\": 29", "Additional test case result must be 29")

	req, _ = http.NewRequest("POST", "/max-path", strings.NewReader("aar"))
	req.Header.Add("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	p, _ = ioutil.ReadAll(w.Body)
	assert.NotContains(string(p), "MaxPathValue", "It must return binding error")

}

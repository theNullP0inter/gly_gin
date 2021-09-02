package gin

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/googly/controller"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/service"
)

func TestHandleHttpError(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	e := &controller.HttpError{
		Code:    400,
		Message: "mock",
		Errors: map[string]string{
			"foo": "bar",
		},
	}
	HandleHttpError(c, e)

}

func TestBaseGinController(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	l := new(logger.MockGooglyLogger)

	res := map[string]string{"foo": "bar"}
	b_res, _ := json.Marshal(gin.H{
		"foo": "bar",
	})

	con := NewBaseGinController(l)

	con.HttpResponse(c, res, 200)
	assert.Equal(t, 200, w.Code)
	b, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, b_res, b)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	con.HttpReplySuccess(c, res)
	assert.Equal(t, 200, w.Code)
	b, _ = ioutil.ReadAll(w.Body)
	assert.Equal(t, b_res, b)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	con.HttpReplyGinBindError(c, nil)
	assert.Equal(t, 422, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	con.HttpReplyGinNotFoundError(c, nil)
	assert.Equal(t, 404, w.Code)

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	serr := &service.ServiceError{
		Code: 500,
	}
	con.HttpReplyServiceError(c, serr)
	assert.Equal(t, 500, w.Code)

}

func TestNewBaseGinController(t *testing.T) {
	// Testing for panic
	l := new(logger.MockGooglyLogger)
	_ = NewBaseGinController(l)
}

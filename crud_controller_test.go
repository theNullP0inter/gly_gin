package gin

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/theNullP0inter/googly/logger"
	"github.com/theNullP0inter/googly/service"
)

type MockSerializer struct {
	Foo string `json:"foo"`
}

func TestBaseGinCrudControllerCreate(t *testing.T) {
	l := logger.NewGooglyLogger()
	s := new(service.MockCrudService)
	h := new(MockGinQueryParametersHydrator)

	con := NewBaseGinCrudController(
		l, s, h,
		new(MockSerializer), new(MockSerializer),
		new(MockSerializer), new(MockSerializer),
	)

	// Success
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(`{"foo":"bar"}`)))
	s.On("Create", mock.Anything).Return(new(MockSerializer), nil)
	con.Create(c)
	assert.Equal(t, 200, w.Code)
	s.AssertExpectations(t)

	// Wrong Params
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(``)))
	s.On("Create", mock.Anything).Return(new(MockSerializer), nil)
	con.Create(c)
	assert.Equal(t, 422, w.Code)

}

func TestBaseGinCrudControllerGet(t *testing.T) {
	l := logger.NewGooglyLogger()
	s := new(service.MockCrudService)
	h := new(MockGinQueryParametersHydrator)

	con := NewBaseGinCrudController(
		l, s, h,
		new(MockSerializer), new(MockSerializer),
		new(MockSerializer), new(MockSerializer),
	)

	id := "1"

	// Success
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: id})
	s.On("GetItem", mock.Anything).Return(new(MockSerializer), nil)
	con.Get(c)
	assert.Equal(t, 200, w.Code)
	s.AssertExpectations(t)

}

func TestBaseGinCrudControllerList(t *testing.T) {
	l := logger.NewGooglyLogger()
	s := new(service.MockCrudService)
	h := new(MockGinQueryParametersHydrator)

	con := NewBaseGinCrudController(
		l, s, h,
		new(MockSerializer), new(MockSerializer),
		new(MockSerializer), new(MockSerializer),
	)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)

	s.On("GetList", mock.Anything).Return(new(MockSerializer), nil)
	h.On("Hydrate", mock.Anything).Return(mock.Anything, nil)
	con.List(c)
	assert.Equal(t, 200, w.Code)
	s.AssertExpectations(t)
	h.AssertExpectations(t)

}

func TestBaseGinCrudControllerUpdate(t *testing.T) {
	l := logger.NewGooglyLogger()
	s := new(service.MockCrudService)
	h := new(MockGinQueryParametersHydrator)

	con := NewBaseGinCrudController(
		l, s, h,
		new(MockSerializer), new(MockSerializer),
		new(MockSerializer), new(MockSerializer),
	)

	id := "1"

	// Success
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPut, "/", bytes.NewBuffer([]byte(`{"foo":"bar"}`)))
	c.Params = append(c.Params, gin.Param{Key: "id", Value: id})
	s.On("Update", mock.AnythingOfType("string"), mock.Anything).Return(new(MockSerializer), nil)
	con.Update(c)
	assert.Equal(t, 200, w.Code)
	s.AssertExpectations(t)

	// Wrong Params
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPut, "/", bytes.NewBuffer([]byte(``)))
	c.Params = append(c.Params, gin.Param{Key: "id", Value: id})
	s.On("Update", mock.AnythingOfType("string"), mock.Anything).Return(new(MockSerializer), nil)
	con.Update(c)
	assert.Equal(t, 422, w.Code)
	s.AssertExpectations(t)

}

func TestBaseGinCrudControllerDelete(t *testing.T) {
	l := logger.NewGooglyLogger()
	s := new(service.MockCrudService)
	h := new(MockGinQueryParametersHydrator)

	con := NewBaseGinCrudController(
		l, s, h,
		new(MockSerializer), new(MockSerializer),
		new(MockSerializer), new(MockSerializer),
	)

	id := "1"

	// Success
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: id})
	s.On("Delete", mock.Anything).Return(new(MockSerializer), nil)
	con.Delete(c)
	assert.Equal(t, 200, w.Code)
	s.AssertExpectations(t)

}

package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/theNullP0inter/googly/controller"
)

type MockGinQueryParameters struct {
}

type MockGinQueryParametersHydrator struct {
	mock.Mock
}

func (q *MockGinQueryParametersHydrator) Hydrate(context *gin.Context) (controller.QueryParameters, error) {
	q.Called()
	return new(MockGinQueryParameters), nil
}

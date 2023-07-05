package controller

import (
	"net/http"
	"testing"

	"github.com/accalina/restaurant-mgmt/app/model"
	mocksservice "github.com/accalina/restaurant-mgmt/mocks/app/service"
	mockssharedservice "github.com/accalina/restaurant-mgmt/mocks/pkg/shared/service"

	// mocksconfig "github.com/accalina/restaurant-mgmt/mocks/pkg/configuration"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
)

func TestMenuController(t *testing.T) {
	menuService := &mocksservice.MenuService{}
	menuService.Mock.On("GetAllMenu", mock.Anything, mock.Anything).Return([]model.MenuResponse{}, model.Meta{}, nil)

	service := &mockssharedservice.Service{}
	service.On("Menu").Return(menuService)

	menuController := NewMenuController(service)
	// Create a new empty fiber.Ctx object
	fastHTTPReq := &fasthttp.Request{}
	fastHTTPResp := &fasthttp.Response{}
	fastHTTPCtx := &fasthttp.RequestCtx{
		Request:  *fastHTTPReq,
		Response: *fastHTTPResp,
	}
	ctx := fiber.New().AcquireCtx(fastHTTPCtx)

	// Set the method and route for the context
	ctx.Method(http.MethodGet)
	ctx.Path("/menu")
	err := menuController.getAllMenu(ctx)
	assert.NoError(t, err)

}

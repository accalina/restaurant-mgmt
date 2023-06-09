package controller

import (
	"fmt"
	"net/http"
	"testing"

	mocksconfig "github.com/accalina/restaurant-mgmt/mocks/configuration"
	mocksservice "github.com/accalina/restaurant-mgmt/mocks/service"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
)

var config = &mocksconfig.Config{Mock: mock.Mock{}}
var menuService = &mocksservice.MenuService{Mock: mock.Mock{}}
var menuController = MenuController{menuService, config}

func TestMenuController(t *testing.T) {
	menuService.Mock.On("GetAllMenu", mock.Anything, mock.Anything).Return([]model.MenuResponse{}, model.Meta{}, nil)
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
	fmt.Println(err)

}

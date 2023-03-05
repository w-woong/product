package route

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	commondto "github.com/w-woong/common/dto"
	"github.com/w-woong/common/middlewares"
	"github.com/w-woong/product/delivery"
	"github.com/w-woong/product/port"
)

func ProductRoute(router *mux.Router, conf commondto.ConfigHttp, usc port.ProductUsc) *delivery.ProductHttpHandler {

	handler := delivery.NewProductHttpHandler(time.Duration(conf.Timeout)*time.Second, usc)

	router.HandleFunc("/v1/product/{id}",
		middlewares.AuthBearerHandler(handler.HandleFindProduct, conf.BearerToken),
	).Methods(http.MethodGet)

	router.HandleFunc("/v1/product",
		middlewares.AuthBearerHandler(handler.AddProductHandle, conf.BearerToken)).Methods(http.MethodPost)
	return handler
}

func ProductGroupRoute(router *mux.Router, conf commondto.ConfigHttp, usc port.GroupUsc) *delivery.GroupHttpHandler {

	handler := delivery.NewGroupHttpHandler(time.Duration(conf.Timeout)*time.Second, usc)

	router.HandleFunc("/v1/product/group",
		middlewares.AuthBearerHandler(handler.AddGroupHandle, conf.BearerToken)).Methods(http.MethodPost)
	router.HandleFunc("/v1/product/group/{id}",
		middlewares.AuthBearerHandler(handler.HandleFindGroup, conf.BearerToken),
	).Methods(http.MethodGet)

	return handler
}

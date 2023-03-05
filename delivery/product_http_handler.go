package delivery

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-wonk/si/v2"
	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/product/dto"
	"github.com/w-woong/product/port"
)

type ProductHttpHandler struct {
	timeout time.Duration
	usc     port.ProductUsc
}

func NewProductHttpHandler(timeout time.Duration, usc port.ProductUsc) *ProductHttpHandler {
	return &ProductHttpHandler{
		timeout: timeout,
		usc:     usc,
	}
}

func (d *ProductHttpHandler) AddProductHandle(w http.ResponseWriter, r *http.Request) {
	var o dto.Product
	reqBody := common.HttpBody{
		Document: &o,
	}
	if err := si.DecodeJson(&reqBody, r.Body); err != nil {
		common.HttpError(w, http.StatusBadRequest)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
	rowsAffected, err := d.usc.AddProduct(r.Context(), o)
	if err != nil {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	if rowsAffected != 1 {
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error("product was not added", logger.UrlField(r.URL.String()))
		return
	}

	if err := common.HttpBodyOK.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}

func (d *ProductHttpHandler) HandleFindProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	product, err := d.usc.FindProduct(r.Context(), id)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			if err := common.HttpBodyRecordNotFound.EncodeTo(w); err != nil {
				logger.Error(err.Error(), logger.UrlField(r.URL.String()))
			}
			return
		}
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	resBody := common.HttpBody{
		Status:   http.StatusOK,
		Count:    1,
		Document: &product,
	}

	if err := resBody.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}

func (d *ProductHttpHandler) HandleFindProductsByGroupID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	products, err := d.usc.FindProductsByGroupID(r.Context(), id)
	if err != nil {
		if errors.Is(err, common.ErrRecordNotFound) {
			if err := common.HttpBodyRecordNotFound.EncodeTo(w); err != nil {
				logger.Error(err.Error(), logger.UrlField(r.URL.String()))
			}
			return
		}
		common.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	resBody := common.HttpBody{
		Status:    http.StatusOK,
		Count:     len(products),
		Documents: &products,
	}

	if err := resBody.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}

package delivery

import (
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	"github.com/w-woong/common/logger"
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

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

type GroupHttpHandler struct {
	timeout time.Duration
	usc     port.GroupUsc
}

func NewGroupHttpHandler(timeout time.Duration, usc port.GroupUsc) *GroupHttpHandler {
	return &GroupHttpHandler{
		timeout: timeout,
		usc:     usc,
	}
}

func (d *GroupHttpHandler) HandleFindGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	group, err := d.usc.FindGroup(r.Context(), id)
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
		Document: &group,
	}

	if err := resBody.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}

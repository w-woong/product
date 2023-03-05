package delivery

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-wonk/si/v2"
	"github.com/gorilla/mux"
	commondto "github.com/w-woong/common/dto"
	"github.com/w-woong/common/logger"
	"github.com/w-woong/product/dto"
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

func (d *GroupHttpHandler) AddGroupHandle(w http.ResponseWriter, r *http.Request) {
	var o dto.Group
	reqBody := commondto.HttpBody{
		Document: &o,
	}
	if err := si.DecodeJson(&reqBody, r.Body); err != nil {
		commondto.HttpError(w, http.StatusBadRequest)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
	rowsAffected, err := d.usc.AddGroup(r.Context(), o)
	if err != nil {
		commondto.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	if rowsAffected != 1 {
		commondto.HttpError(w, http.StatusInternalServerError)
		logger.Error("group was not added", logger.UrlField(r.URL.String()))
		return
	}

	if err := commondto.HttpBodyOK.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}

func (d *GroupHttpHandler) HandleFindGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	group, err := d.usc.FindGroup(r.Context(), id)
	if err != nil {
		if errors.Is(err, commondto.ErrRecordNotFound) {
			if err := commondto.HttpBodyRecordNotFound.EncodeTo(w); err != nil {
				logger.Error(err.Error(), logger.UrlField(r.URL.String()))
			}
			return
		}
		commondto.HttpError(w, http.StatusInternalServerError)
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}

	resBody := commondto.HttpBody{
		Status:   http.StatusOK,
		Count:    1,
		Document: &group,
	}

	if err := resBody.EncodeTo(w); err != nil {
		logger.Error(err.Error(), logger.UrlField(r.URL.String()))
		return
	}
}

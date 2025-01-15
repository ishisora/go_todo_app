package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/ishisora/go_todo_app/entity"
)

type GetTask struct {
	Service   GetTaskService
	Validator *validator.Validate
}

func (gt *GetTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	t, err := gt.Service.GetTask(ctx, entity.TaskID(id))
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := task{
		ID:     t.ID,
		Title:  t.Title,
		Status: t.Status,
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}

package getredis

import (
	dtredis "golangredis/models/dtRedis"
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

type getData interface {
	GetDataRedis(req dtredis.DataSet) (resp dtredis.RespGetData, err error)
}

type Handler struct {
	getData getData
	render  *renderer.Render
}

func NewHandler(getData getData, render *renderer.Render) *Handler {
	return &Handler{
		getData: getData,
		render:  render,
	}
}

func (h *Handler) GetDataRedis(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	key := param.Get("key")

	request := dtredis.DataSet{
		Key: key,
	}

	value, err := h.getData.GetDataRedis(request)
	if err != nil {
		log.Println("Failed to get value in redis: ", err)
		h.render.JSON(w, http.StatusNotFound, &dtredis.ResponseData{
			Message: "Failed to get data",
			Data:    http.StatusNotFound,
		})
		return
	}

	h.render.JSON(w, http.StatusOK, &dtredis.ResponseData{
		Message: "Success",
		Data:    value,
	})
}

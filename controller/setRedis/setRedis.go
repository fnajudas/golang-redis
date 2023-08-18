package setredis

import (
	"fmt"
	dtredis "golangredis/models/dtRedis"
	"log"
	"net/http"
)

const (
	value = "this is value"
	key   = "this is key"
)

type getData interface {
	GetDataRedis(req dtredis.DataSet) (resp dtredis.RespGetData, err error)
}

type setData interface {
	SetData(request dtredis.DataSet) error
}

type SetRedis struct {
	setData setData
	getData getData
}

func NewSetRedis(setData setData, getData getData) *SetRedis {
	return &SetRedis{
		setData: setData,
		getData: getData,
	}
}

func (s *SetRedis) SetData(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	value := param.Get("value")
	key := param.Get("key")

	request := dtredis.DataSet{
		Value: value,
		Key:   key,
	}

	err := s.setData.SetData(request)
	if err != nil {
		log.Println("Failed to set value in Redis:", err)
		return
	}

	getValue, err := s.getData.GetDataRedis(request)
	if err != nil {
		log.Println("Failed to get value from Redis:", err)
		return
	}

	fmt.Println("Value set in Redis successfully:", getValue.Key)
}

package dtredis

type DataSet struct {
	Value string `json:"value"`
	Key   string `json:"key"`
}

type ResponseSetRedis struct {
	Value string `json:"value"`
}

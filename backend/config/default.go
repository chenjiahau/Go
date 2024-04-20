package config

type DefaultResponseFormat struct {
	Data		interface{} `json:"data"`
	Error 	interface{} `json:"error"`
}
package cache

type ICache interface {
	Set(id int32, value interface{})
	Get(id int32) interface{}
}

package main

var store = map[string]interface{}{}

func putValue(key string, value interface{}) {
	store[key] = value
}

func getValue(key string) interface{} {
	return store[key]
}

func getStore() map[string]interface{} {
	return store
}

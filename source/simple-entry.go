package main

type SimpleEntry struct {
	key   string
	value interface{}
}

func NewSimpleEntry(key string, value interface{}) *SimpleEntry {
	this := new(SimpleEntry)
	this.key = key
	this.value = value
	return this
}

func (se *SimpleEntry) GetKey() string {
	return se.key
}

func (se *SimpleEntry) GetValue() interface{} {
	return se.value
}

func (se *SimpleEntry) GetStorageSize() int32 {
	return 1
}

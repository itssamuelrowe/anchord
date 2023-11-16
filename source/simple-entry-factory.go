package main

type SimpleEntryFactory struct {
}

func NewSimpleEntryFactory() *SimpleEntryFactory {
	this := new(SimpleEntryFactory)
	return this
}

func (sef *SimpleEntryFactory) CreateEntry(key string, value interface{}) Entry {
	return NewSimpleEntry(key, value)
}

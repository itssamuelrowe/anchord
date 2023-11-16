package main

type Collection interface {
	PutEntry(key string, value interface{})
	RemoveEntry(key string)
}

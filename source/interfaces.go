package main

type Entry interface {
	GetKey() string
	GetValue() interface{}
	GetStorageSize() int32
}

// Entry Factory

type EntryFactory interface {
	CreateEntry(key string, value interface{}) Entry
}

type DataStore interface {
	GetEntries() []Entry
	GetTotalStorageSize() int32
	GetSize() int
	PutEntry(entry Entry)
	RemoveEntry(key string)
	GetStorageLimit() int32
	Clear()
}

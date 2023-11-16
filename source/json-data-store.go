package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonDataStore struct {
	path string
	file *os.File
}

func NewJsonDataStore(path string) *JsonDataStore {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	return &JsonDataStore{
		path,
		file,
	}
}

func (ds *JsonDataStore) GetTotalStorageSize() int32 {
	return 0
}

func (ds *JsonDataStore) GetSize() int {
	return -1
}

type Record struct {
	Type int8   `json:"type"`
	Key  string `json:"key"`
}

type InsertRecord struct {
	Record
	Value interface{} `json:"value"`
}

type DeleteRecord struct {
	Record
}

func (ds *JsonDataStore) PutEntry(entry Entry) {
	record := InsertRecord{
		Record: Record{
			Key:  entry.GetKey(),
			Type: 1,
		},
		Value: entry.GetValue(),
	}
	content, _ := json.Marshal(record)
	_, err := ds.file.Write(content)
	if err != nil {
		fmt.Print(err)
	}
	ds.file.WriteString("\n")
}

func (ds *JsonDataStore) RemoveEntry(key string) {
	record := DeleteRecord{
		Record{
			Key:  key,
			Type: 2,
		},
	}
	content, _ := json.Marshal(record)
	ds.file.Write(content)
}

func (ds *JsonDataStore) GetEntries() []Entry {
	var result []Entry
	return result
}

func (ds *JsonDataStore) Clear() {
	os.Truncate(ds.path, 0)
}

func (ds *JsonDataStore) GetStorageLimit() int32 {
	return 100
}

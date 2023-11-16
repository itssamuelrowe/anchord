package main

import (
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

type RBTDataStore struct {
	tree *rbt.Tree
}

func NewRBTDataStore() *RBTDataStore {
	this := new(RBTDataStore)
	this.tree = rbt.NewWithStringComparator()

	return this
}

func (ds *RBTDataStore) GetTotalStorageSize() int32 {
	return int32(ds.tree.Size())
}

func (ds *RBTDataStore) GetSize() int {
	return ds.tree.Size()
}

func (ds *RBTDataStore) PutEntry(entry Entry) {
	ds.tree.Put(entry.GetKey(), entry)
}

func (ds *RBTDataStore) RemoveEntry(key string) {
	ds.tree.Remove(key)
}

func (ds *RBTDataStore) GetEntries() []Entry {
	var result []Entry
	iterator := ds.tree.Iterator()
	for iterator.Next() {
		entry := iterator.Value()
		result = append(result, entry.(Entry))
	}
	return result
}

func (ds *RBTDataStore) Clear() {
	ds.tree.Clear()
}

func (ds *RBTDataStore) GetStorageLimit() int32 {
	return 10
}

package main

type LSMTCollection struct {
	entryFactory   EntryFactory
	temporaryStore DataStore
	recoveryStore  DataStore
	permanentStore DataStore
}

func NewLSMTCollection(
	entryFactory EntryFactory,
	temporaryStore DataStore,
	recoveryStore DataStore,
	permanentStore DataStore,
) *LSMTCollection {
	this := new(LSMTCollection)
	this.entryFactory = entryFactory
	this.temporaryStore = temporaryStore
	this.recoveryStore = recoveryStore
	this.permanentStore = permanentStore

	return this
}

func (c *LSMTCollection) PutEntry(key string, value interface{}) {
	entry := c.entryFactory.CreateEntry(key, value)
	newTDSSize := c.temporaryStore.GetTotalStorageSize() + entry.GetStorageSize()
	if newTDSSize <= c.temporaryStore.GetStorageLimit() {
		c.temporaryStore.PutEntry(entry)
		c.recoveryStore.PutEntry(entry)
	} else {
		// Flush temporary storage to permanent store.
		entries := c.temporaryStore.GetEntries()
		for _, entry := range entries {
			c.permanentStore.PutEntry(entry)
		}
		c.temporaryStore.Clear()
		c.recoveryStore.Clear()
	}
}

func (c *LSMTCollection) Flush() {
	// Flush temporary storage to permanent store.
	entries := c.temporaryStore.GetEntries()
	for _, entry := range entries {
		c.permanentStore.PutEntry(entry)
	}
	c.temporaryStore.Clear()
}

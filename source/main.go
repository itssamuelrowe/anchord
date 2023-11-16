package main

func main() {
	factory := NewSimpleEntryFactory()
	temporaryStore := NewRBTDataStore()
	recoveryStore := NewJsonDataStore("./data/recovery.json")
	permanentStore := NewJsonDataStore("./data/permanent.json")

	collection := NewLSMTCollection(
		factory,
		temporaryStore,
		recoveryStore,
		permanentStore,
	)
	collection.PutEntry("Everest", "8,848")
	collection.PutEntry("K2", "8,611m")
	collection.PutEntry("Kangchenjunga", "8,586m")
	collection.PutEntry("Lhotse", "8,516m")
	collection.PutEntry("Makalu", "8,485m")
	collection.PutEntry("Cho Oyu", "8,188m")
	collection.PutEntry("Dhaulagiri", "8,167m")
	collection.PutEntry("Manaslu", "8,163m")
	collection.PutEntry("Nanga Parbat", "8,126m")
	collection.PutEntry("Annapurna", "8,091m")
	collection.PutEntry("Gasherbrum I", "8,080m")
	collection.PutEntry("Broad Peak", "8,051m")
	collection.PutEntry("Gasherbrum II", "8,035m")
	collection.PutEntry("Shishapangma", "8,027m")

	recoveryStore.file.Close()
	permanentStore.file.Close()
}

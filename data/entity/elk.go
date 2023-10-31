package entity

type ElasticIndexConfig struct {
	NumberOfReplicas int `json:"number_of_replicas"`
	NumberOfShards   int `json:"number_of_shards"`
}

type ElasticIndex struct {
	Health       string `json:"health"`
	Status       string `json:"status"`
	Index        string `json:"index"`
	Uuid         string `json:"uuid"`
	Pri          string `json:"pri"`
	Rep          string `json:"rep"`
	DocsCount    string `json:"docs.count"`
	DocsDeleted  string `json:"docs.deleted"`
	StoreSize    string `json:"store.size"`
	PriStoreSize string `json:"pri.store.size"`
}

type ElasticAliasAdd struct {
	Add ElasticAliasPair `json:"add"`
}

type ElasticAliasPair struct {
	Index string `json:"index"`
	Alias string `json:"alias"`
}

package elk_service

import (
	"github.com/lowl11/boost/data/entity"
	"github.com/lowl11/boost/internal/helpers/elk_parser"
)

type bindAliasRequest struct {
	Actions []entity.ElasticAliasAdd `json:"actions"`
}

type searchResult struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`

	Hits searchHits `json:"hits"`
}

type searchHits struct {
	Total struct {
		Value    int    `json:"value"`
		Relation string `json:"relation"`
	} `json:"total"`
	MaxScore float64         `json:"max_score"`
	Hits     []searchHitItem `json:"hits"`
}

type searchHitItem struct {
	Index  string         `json:"_index"`
	Type   string         `json:"_type"`
	Id     string         `json:"_id"`
	Score  float64        `json:"_score"`
	Source map[string]any `json:"_source"`
}

type createIndexRequest struct {
	Settings indexSettings `json:"settings"`
	Mappings indexMappings `json:"mappings"`
}

type indexSettings struct {
	Index entity.ElasticIndexConfig `json:"index"`
}

type indexMappings struct {
	Properties map[string]elk_parser.MappingField `json:"properties"`
}

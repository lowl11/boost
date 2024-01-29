package elastic

import (
	"github.com/google/uuid"
	"github.com/lowl11/boost/data/entity"
	"time"
)

type Document struct {
	ID        uuid.UUID `json:"id" elk:"id"`
	CreatedAt time.Time `json:"created_at" elk:"created_at"`
}

func NewDocument(customID ...uuid.UUID) Document {
	document := Document{
		CreatedAt: time.Now(),
	}
	if len(customID) > 0 {
		document.ID = customID[0]
	} else {
		document.ID = uuid.New()
	}
	return document
}

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
	Properties map[string]mappingField `json:"properties"`
}

package domain

import "time"

// PlantCategory Model
type PlantCategory struct {
	ID                   int64     `json:"id" db:"id"`
	LargeClassification  string    `json:"large_classification" db:"large_classification"`
	MiddleClassification string    `json:"middle_classification" db:"middle_classification"`
	SmallClassification  string    `json:"small_classification" db:"small_classification"`
	Thesaurus            string    `json:"thesaurus" db:"thesaurus"`
	HarvestSite          string    `json:"harvest_site" db:"harvest_site"`
	AttributeItem        string    `json:"attribute_item" db:"attribute_item"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
}

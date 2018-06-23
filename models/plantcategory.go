package models

import "time"

// PlantCategory 植物の種類に関わるモデル
type PlantCategory struct {
	ID                   int64     `db:"id, primarykey, autoincrement" json:"id"`
	LargeClassification  string    `db:"large_classification" json:"large_classification"`
	MiddleClassification string    `db:"middle_classification" json:"middle_classification"`
	SmallClassification  string    `db:"small_classification" json:"small_classification"`
	Thesaurus            string    `db:"thesaurus" json:"thesaurus"`
	HarvestSite          string    `db:"harvest_site" json:"harvest_site"`
	AttributeItem        string    `db:"attribute_item" json:"attribute_item"`
	CreatedAt            time.Time `db:"created_at" json:"created_at"`
	UpdatedAt            time.Time `db:"updated_at" json:"updated_at"`
}

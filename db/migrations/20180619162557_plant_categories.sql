
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `plant_categories` (
  `id` int(11) UNSIGNED NOT NULL,
  `large_classification` varchar(255) NOT NULL COMMENT '大分類',
  `middle_classification` varchar(255) NOT NULL COMMENT '中分類',
  `small_classification` varchar(255) NOT NULL COMMENT '小分類',
  `thesaurus` varchar(255) DEFAULT NULL COMMENT 'シーソラス(別名)',
  `harvest_site` varchar(255) DEFAULT NULL COMMENT '収穫部位',
  `attribute_item` varchar(255) DEFAULT NULL COMMENT '属性項目',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='観測植物を管理するテーブル';
ALTER TABLE `plant_categories`
  ADD PRIMARY KEY (`id`);
  
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE plant_categories;

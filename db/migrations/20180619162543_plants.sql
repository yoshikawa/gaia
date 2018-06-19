
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE plants (
  `id` int(11) UNSIGNED NOT NULL,
  `plant_categorie_id` int(11) UNSIGNED NOT NULL COMMENT '観測植物の分類ID',
  `variety` varchar(255) NOT NULL COMMENT '品種等',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='観測植物を管理するテーブル';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE plants;

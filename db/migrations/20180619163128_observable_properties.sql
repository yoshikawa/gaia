
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `observable_properties` (
  `id` int(11) UNSIGNED NOT NULL COMMENT '管理用',
  `observable_property_number` varchar(255) NOT NULL COMMENT '計測項目のID',
  `system` varchar(5) NOT NULL COMMENT '植物体の場所',
  `classification` varchar(255) NOT NULL COMMENT '観測の分類',
  `classification_en` varchar(255) NOT NULL COMMENT '観測の分類(英語)',
  `observation_property` varchar(255) NOT NULL COMMENT '物質量・現象',
  `observation_property_en` varchar(255) NOT NULL COMMENT '物質量・現象(英語)',
  `unit_of_measure` varchar(20) DEFAULT NULL COMMENT '単位',
  `display_unit` varchar(20) DEFAULT NULL COMMENT '表示単位',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '管理用',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '管理用'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='計測項目の情報を保存したテーブル';
ALTER TABLE `observable_properties`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `observable_property_id` (`observable_property_number`);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE observable_properties;
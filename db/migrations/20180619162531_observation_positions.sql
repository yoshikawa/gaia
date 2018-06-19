
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE observation_positions (
  `id` int(11) UNSIGNED NOT NULL COMMENT '管理用',
  `field_id` int(11) UNSIGNED NOT NULL COMMENT '圃場管理のID',
  `plant_id` int(11) UNSIGNED DEFAULT NULL COMMENT '植物種別リストのID',
  `name` varchar(255) NOT NULL COMMENT '登録名',
  `point` geometry NOT NULL COMMENT '緯度経度',
  `altitude` double(5,1) DEFAULT NULL COMMENT '標高(m)',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '管理用',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '管理用',
  `deleted` datetime DEFAULT NULL COMMENT '論理削除(管理用)'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='観測位置を管理するテーブル';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE observation_positions;

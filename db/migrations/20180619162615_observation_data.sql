
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `observation_data` (
  `id` int(11) UNSIGNED NOT NULL COMMENT '管理用',
  `observation_position_id` int(11) UNSIGNED NOT NULL COMMENT '観測点のID',
  `sensor_id` int(11) UNSIGNED NOT NULL COMMENT 'センサの番号',
  `value` double(10,4) NOT NULL COMMENT '値',
  `datetime` datetime NOT NULL COMMENT '観測日時',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '管理用',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '管理用'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='観測データを保存するテーブル';
ALTER TABLE `observation_data`
  ADD PRIMARY KEY (`id`),
  ADD KEY `observation_position_id` (`observation_position_id`),
  ADD KEY `observation_data_ibfk_2` (`sensor_id`);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE observation_data;

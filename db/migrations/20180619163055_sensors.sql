
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `sensors` (
  `id` int(11) UNSIGNED NOT NULL COMMENT '管理用',
  `sensing_device_id` int(11) UNSIGNED NOT NULL COMMENT '機器種別のID',
  `observable_property_id` int(11) UNSIGNED NOT NULL COMMENT '計測項目のID',
  `individual_id` int(11) UNSIGNED DEFAULT NULL COMMENT '個体情報のID',
  `sensor_number` varchar(50) NOT NULL COMMENT 'センサのID',
  `observation_condition` double(4,1) NOT NULL COMMENT '計測条件(m)',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '管理用',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '管理用',
  `deleted` datetime DEFAULT NULL COMMENT '管理用(論理削除)'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='センサ情報を管理するテーブル';
ALTER TABLE `sensors`
  ADD PRIMARY KEY (`id`),
  ADD KEY `sensors_ibfk_2` (`individual_id`),
  ADD KEY `sensors_ibfk_3` (`observable_property_id`),
  ADD KEY `sensor_device_id` (`sensing_device_id`),
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理用';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE sensors;

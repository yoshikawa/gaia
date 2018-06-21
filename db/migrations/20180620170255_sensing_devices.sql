
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `sensing_devices` (
  `id` int(11) UNSIGNED NOT NULL COMMENT '管理用',
  `observation_position_id` int(11) UNSIGNED NOT NULL COMMENT '観測位置のID',
  `vendor_id` char(16) NOT NULL COMMENT 'サービス提供者の番号',
  `device_type` varchar(50) NOT NULL COMMENT 'デバイスの種別ID',
  `serial_number` varchar(50) NOT NULL COMMENT 'シリアルの番号',
  `name` varchar(255) NOT NULL COMMENT '製品名',
  `transmission_distance` smallint(1) UNSIGNED NOT NULL COMMENT '伝送距離(m)',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '管理用',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '管理用',
  `deleted` datetime DEFAULT NULL COMMENT '論理削除(管理用)'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='機器を管理するテーブル';
ALTER TABLE `sensing_devices`
  ADD PRIMARY KEY (`id`),
  ADD KEY `observation_position_id` (`observation_position_id`) USING BTREE,
  ADD KEY `sensor_devices_ibfk_2` (`vendor_id`);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE `sensing_devices`;
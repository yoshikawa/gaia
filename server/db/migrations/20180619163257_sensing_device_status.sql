
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `sensing_device_status` (
  `id` int(11) UNSIGNED NOT NULL COMMENT '管理用',
  `sensing_device_id` int(11) UNSIGNED NOT NULL COMMENT '機器管理のID',
  `battery` smallint(1) UNSIGNED DEFAULT NULL COMMENT 'バッテリー残量',
  `status` smallint(1) UNSIGNED NOT NULL COMMENT '機器の状態',
  `datetime` datetime NOT NULL COMMENT '観測日時',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '管理用',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '管理用'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='機器の状態を管理するテーブル';
ALTER TABLE `sensing_device_status`
  ADD PRIMARY KEY (`id`),
  ADD KEY `sensor_device_status_ibfk_1` (`sensing_device_id`),
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理用';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE sensing_device_status;

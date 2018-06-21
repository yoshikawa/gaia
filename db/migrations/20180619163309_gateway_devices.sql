
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `gateway_devices` (
  `id` int(11) UNSIGNED NOT NULL COMMENT '管理用',
  `vendor_id` char(16) NOT NULL COMMENT 'サービス提供者の番号',
  `device_type` varchar(50) NOT NULL COMMENT '機器の種類',
  `serial_number` varchar(50) NOT NULL COMMENT 'シリアル番号',
  `point` geometry NOT NULL COMMENT '緯度経度(GPS)',
  `name` varchar(255) NOT NULL COMMENT '製品名',
  `transmission_distance` smallint(1) UNSIGNED NOT NULL COMMENT '伝送距離',
  `authentication_name` char(16) NOT NULL COMMENT 'サーバ認証ID',
  `password` varchar(255) NOT NULL COMMENT 'パスワード',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '管理用',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '管理用'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ネットワーク構成する機器を管理するテーブル';
ALTER TABLE `gateway_devices`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `authentication_name` (`authentication_name`),
  ADD KEY `vendor_id` (`vendor_id`);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE gateway_devices;
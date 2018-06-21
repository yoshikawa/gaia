
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `individuals` (
  `id` int(11) UNSIGNED NOT NULL COMMENT '管理用',
  `short_name` varchar(255) DEFAULT NULL COMMENT '略称名',
  `long_name` varchar(255) NOT NULL COMMENT '正式名称',
  `model_number` varchar(255) NOT NULL COMMENT '型番',
  `intended_application` varchar(255) NOT NULL COMMENT '使用用途',
  `sensor_type` varchar(255) NOT NULL COMMENT 'センサー種別',
  `minimum_observation_range` float(6,2) NOT NULL COMMENT '計測限界の最小値',
  `maximum_observation_range` float(6,2) NOT NULL COMMENT '計測限界の最大値',
  `code_of_observation_range` varchar(10) NOT NULL COMMENT '計測限界の記号',
  `minimum_observation_accuracy_of_measured_value` float(8,4) NOT NULL COMMENT '最小精度',
  `maximum_observation_accuracy_of_measured_value` float(8,4) NOT NULL COMMENT '最大精度',
  `code_of_observation_accuracy` varchar(10) NOT NULL COMMENT '観察精度の記号',
  `observation_resolution_of_measured_value` float(3,2) NOT NULL COMMENT 'センサーの解像度',
  `code_of_observation_resolution` varchar(10) NOT NULL COMMENT '観察精度の記号',
  `observation_interval_strict` tinyint(1) DEFAULT '0' COMMENT '計測間隔の厳密性',
  `observation_timing_strict` tinyint(1) DEFAULT '0' COMMENT '計測時間の厳密性',
  `manufacturer` varchar(255) DEFAULT NULL COMMENT 'メーカ',
  `location_name` varchar(255) DEFAULT NULL COMMENT '販売場所',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '管理用',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '管理用'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='個体情報を管理するテーブル';
ALTER TABLE `individuals`
  ADD PRIMARY KEY (`id`);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE individuals;

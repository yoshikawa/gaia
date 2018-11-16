
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE fields (
  `id` int(11) UNSIGNED NOT NULL COMMENT '管理用',
  `organization_id` int(11) UNSIGNED NOT NULL COMMENT 'グループ番号',
  `name` varchar(255) NOT NULL COMMENT '登録名',
  `area` geometry NOT NULL COMMENT '圃場範囲',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '管理用',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '管理用',
  `deleted` datetime DEFAULT NULL COMMENT '論理削除(管理用)'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='圃場情報の管理するテーブル';
ALTER TABLE `fields`
  ADD PRIMARY KEY (`id`),
  ADD KEY `organization_id` (`organization_id`),
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理用';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE fields;

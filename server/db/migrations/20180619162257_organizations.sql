
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE organizations (
  `id` int(11) UNSIGNED NOT NULL COMMENT '管理用',
  `name` varchar(255) NOT NULL COMMENT '名前',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '管理用',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '管理用'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
ALTER TABLE `organizations`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`),
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理用';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE organizations;

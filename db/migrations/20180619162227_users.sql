
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE users (
  `id` int(11) UNSIGNED NOT NULL COMMENT '管理用',
  `organization_id` int(11) UNSIGNED NOT NULL COMMENT 'グループ番号',
  `name` varchar(255) NOT NULL COMMENT '名前',
  `email` varchar(255) NOT NULL COMMENT '登録メールアドレス',
  `password` varchar(255) NOT NULL COMMENT '登録パスワード',
  `country` varchar(20) NOT NULL COMMENT '国',
  `administrator` tinyint(1) NOT NULL DEFAULT '0' COMMENT '登録者の管理権限の有無',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '管理用',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '管理用'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ユーザを管理するテーブル';
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `organization_id` (`organization_id`);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE users;
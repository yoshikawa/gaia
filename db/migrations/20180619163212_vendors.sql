
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `vendors` (
  `id` char(16) NOT NULL COMMENT 'サービス提供者の番号',
  `name` varchar(255) NOT NULL COMMENT '提供者名',
  `organization_name` varchar(255) NOT NULL COMMENT '組織名',
  `voice` varchar(20) NOT NULL COMMENT '電話番号',
  `facsimile` varchar(20) DEFAULT NULL COMMENT 'FAX',
  `delivary_point` varchar(255) NOT NULL COMMENT '番地',
  `city` varchar(255) NOT NULL COMMENT '都道府県、市町村',
  `postal_code` char(8) NOT NULL COMMENT '郵便番号',
  `country` varchar(20) NOT NULL COMMENT '国',
  `electronic_mail_address` varchar(255) NOT NULL COMMENT '電子メールアドレス',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '管理用',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '管理用'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ベンダの情報を管理するテーブル';
ALTER TABLE `vendors`
  ADD PRIMARY KEY (`id`);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE vendors;

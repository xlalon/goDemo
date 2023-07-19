
-- CREATE TABLE IF NOT EXISTS `go_demo`.`chains` (
-- 	`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
-- 	`code` varchar(64) NOT NULL COMMENT '公链编号',
-- 	`name` varchar(64) NOT NULL COMMENT '公链名称',
-- 	`status` varchar(64) NOT NULL COMMENT '状态',
-- 	`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
-- 	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
-- 	`deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
-- 	PRIMARY KEY (`id`),
-- 	UNIQUE KEY `code` (`code`)
-- ) ENGINE = InnoDB CHARSET = utf8mb4 COMMENT '公链表'


-- CREATE TABLE IF NOT EXISTS `go_demo`.`assets` (
-- 	`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
-- 	`code` varchar(64) NOT NULL COMMENT '资产编号',
-- 	`name` varchar(64) NOT NULL COMMENT '资产名称',
-- 	`chain_code` varchar(64) NOT NULL COMMENT '公链编号',
-- 	`identity` varchar(64) NOT NULL COMMENT '资产身份证',
-- 	`precision` int NOT NULL COMMENT '资产精度',
-- 	`status` varchar(64) NOT NULL COMMENT '状态',
-- 	`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
-- 	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
-- 	`deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
-- 	PRIMARY KEY (`id`),
-- 	UNIQUE KEY `chain_identity` (`chain_code`, `identity`)
-- ) ENGINE = InnoDB CHARSET = utf8mb4 COMMENT '资产信息表'


-- CREATE TABLE IF NOT EXISTS `go_demo`.`asset_settings` (
-- 	`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
-- 	`chain_code` varchar(64) NOT NULL COMMENT '公链编号',
-- 	`asset_code` varchar(64) NOT NULL COMMENT '资产编号',
-- 	`min_deposit_amount` decimal(21, 9) UNSIGNED NOT NULL COMMENT '最小充值数量',
-- 	`withdraw_fee` decimal(21, 9) UNSIGNED NOT NULL COMMENT '提现手续费',
-- 	`to_hot_threshold` decimal(21, 9) UNSIGNED NOT NULL COMMENT '充值到热阈值',
-- 	`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
-- 	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
-- 	`deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
-- 	PRIMARY KEY (`id`),
-- 	UNIQUE KEY `chain_asset_code` (`chain_code`, `asset_code`)
-- ) ENGINE = InnoDB CHARSET = utf8mb4 COMMENT '资产设置表'


-- CREATE TABLE IF NOT EXISTS `go_demo`.`accounts` (
-- 	`id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
-- 	`chain` varchar(64) NOT NULL COMMENT '公链编号',
-- 	`address` varchar(256) NOT NULL COMMENT '地址',
-- 	`memo` varchar(64) NOT NULL COMMENT 'Memo',
-- 	`status` varchar(64) NOT NULL COMMENT '状态',
-- 	`version` int NOT NULL DEFAULT '0' COMMENT '乐观锁版本号',
-- 	`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
-- 	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
-- 	`deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
-- 	PRIMARY KEY (`id`),
-- 	UNIQUE KEY `chain_address_memo` (`chain`, `address`, `memo`)
-- ) ENGINE = InnoDB CHARSET = utf8mb4 COMMENT '地址表'
-- TBPARTITION BY mod_hash(chain) TBPARTITIONS 10


-- CREATE TABLE IF NOT EXISTS `go_demo`.`deposits` (
--   `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
--   `chain` varchar(64) NOT NULL COMMENT '公链编号',
--   `tx_hash` varchar(256) NOT NULL COMMENT 'tx hash',
--   `v_out` int NOT NULL DEFAULT '0' COMMENT 'v_out',
--   `sender` varchar(128) NOT NULL DEFAULT '' COMMENT '发送地址',
--   `recipient` varchar(128) NOT NULL COMMENT '接收地址 ',
--   `memo` varchar(64) NOT NULL DEFAULT '' COMMENT 'memo',
--   `asset` varchar(64) NOT NULL COMMENT '资产编号',
--   `amount` decimal(21,9) unsigned NOT NULL COMMENT '数量',
--   `height` int NOT NULL DEFAULT '0' COMMENT '交易发生的高度',
--   `status` varchar(24) NOT NULL DEFAULT '' COMMENT '状态',
--   `comment` varchar(256) NOT NULL DEFAULT '' COMMENT '备注',
--   `version` int NOT NULL DEFAULT '0' COMMENT '乐观锁版本号',
--   `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--   `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
--   PRIMARY KEY (`id`),
--   UNIQUE KEY `chain_asset_tx` (`chain`,`tx_hash`,`v_out`,`receiver`,`memo`,`asset`,`amount`),
--   KEY `tx_hash` (`tx_hash`)
--   ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='充值表'
-- TBPARTITION BY mod_hash(tx_hash) TBPARTITIONS 10


-- CREATE TABLE IF NOT EXISTS `go_demo`.`income_cursors` (
-- `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
-- `chain_code` varchar(64) NOT NULL COMMENT '公链编号',
-- `height` int NOT NULL DEFAULT 0 COMMENT '区块高度',
-- `address` varchar(256) NOT NULL DEFAULT '' COMMENT '地址',
-- `label` varchar(64) NOT NULL COMMENT '账号标签',
-- `tx_hash` varchar(256) NOT NULL DEFAULT '' COMMENT 'tx hash',
-- `direction` varchar(24) NOT NULL DEFAULT 'ASC' COMMENT '方向',
-- `index` int NOT NULL DEFAULT 0 COMMENT '账号index',
-- `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
-- `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
-- `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
-- PRIMARY KEY (`id`)
-- ) ENGINE = InnoDB CHARSET = utf8mb4 COMMENT '入账扫描游标表'

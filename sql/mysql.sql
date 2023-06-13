
/*Table structure for table `greet_order_payment` */

DROP TABLE IF EXISTS `greet_order_payment`;

CREATE TABLE `greet_order_payment` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `order_id` int NOT NULL DEFAULT '0' COMMENT '订单id|补款单id',
    `member_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
    `business_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '业务类型',
    `order_sn` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '销巴订单号|补款单号，未加索引，不能做条件字句',
    `payment_status` smallint NOT NULL DEFAULT '0' COMMENT '支付单状态1：已创建待支付；2：已支付；3：支付失败；4:支付异常,需要人工介入;',
    `payment_sn` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付订单号（调用支付服务创建支付时返回）',
    `payment_amount` decimal(10,2) DEFAULT NULL COMMENT '支付单支付金额',
    `payment_type` tinyint NOT NULL DEFAULT '0' COMMENT '支付方式： 1：银联支付 2：支付宝支付 3：微信支付 4.个人余额 5.小巴余额',
    `payment_params` json DEFAULT NULL COMMENT '创建支付订单时的入参',
    `payment_result` json DEFAULT NULL COMMENT '创建支付订单时的响应',
    `deleted_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_order_id` (`order_id`),
    KEY `idx_member_id` (`member_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单支付订单信息'


/*Table structure for table `greet_order_payment_notice` */

DROP TABLE IF EXISTS `greet_order_payment_notice`;

CREATE TABLE `greet_order_payment_notice` (
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `payment_serial_number` VARCHAR(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '支付订单号',
    `third_notice` JSON DEFAULT NULL COMMENT '第三方通知结果',
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    `created_at` TIMESTAMP NULL DEFAULT NULL,
    `updated_at` TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_payment_serial_number` (`payment_serial_number`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='订单支付第三方通知结果'




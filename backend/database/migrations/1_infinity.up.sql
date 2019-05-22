CREATE TABLE `users`(
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `email` VARCHAR(100) NOT NULL,
    `name` VARCHAR(100) DEFAULT NULL,
    `password` VARCHAR(250) DEFAULT NULL,
    `designation` VARCHAR(100) NOT NULL,
    `emp_id` VARCHAR(100) NOT NULL,
    `user_type` ENUM('USER', 'ADMIN') NOT NULL DEFAULT 'USER',
    `user_status` TINYINT(1) NOT NULL,
    `created_at` TIMESTAMP NULL,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY(`id`),
    UNIQUE KEY `email`(`email`),
    KEY `deleted_at`(`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = latin1;

INSERT
INTO
    `users`(`id`, `email`, `name`, `password`, `designation`, `emp_id`, `user_type`, `user_status`)
VALUES(1, "su@admin.com", "Default Generic Admin", "$2a$04$FddCOm6PXR7Rs524ISWcIev1iotTj76BvK1FKjzken4DlHpd5fm8S", "SuperUser", "EMP001", "ADMIN", TRUE),(2, "user@user.com", "Default Generic User", "$2a$04$C092gPZaSfkuvL57jCdhy.4fcTIzWPvT0BJWcTvAObRTJiyYNcCtO", "GenericUser", "EMP002", "USER", TRUE);

CREATE TABLE `courses`(
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `description` VARCHAR(300) NULL,
    `created_at` TIMESTAMP NULL,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY(`id`),
    KEY `deleted_at`(`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = latin1;

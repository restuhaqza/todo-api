
-- SQL MIGRATIONS FOR DEVCODE TODO API by Restu Haqqi Muzakir

CREATE DATABASE IF NOT EXISTS `restu_haqqi_muzakir`;

USE `restu_haqqi_muzakir`;

CREATE TABLE IF NOT EXISTS activity_group(
    `id` int(11) NOT NULL auto_increment PRIMARY KEY,
    `title` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT NOW(),
    `updated_at` DATETIME NOT NULL DEFAULT NOW(),
    `deleted_at` DATETIME NULL
);

CREATE TABLE IF NOT EXISTS todo(
 `id` int(11) NOT NULL auto_increment PRIMARY KEY,
 `activity_group_id` int(11),
 `title` varchar(255) NOT NULL,
 `is_active` bool,
 `priority` varchar(16),
 `created_at` DATETIME NOT NULL DEFAULT NOW(),
 `updated_at` DATETIME NOT NULL DEFAULT NOW(),
 `deleted_at` DATETIME NULL,

 CONSTRAINT FK_Todo_ActivityGroup FOREIGN KEY (activity_group_id)
 REFERENCES activity_group(id)
);
/*
SQLyog Community v13.1.1 (64 bit)
MySQL - 10.3.16-MariaDB : Database - ottotestjunior
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`ottotestjunior` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `ottotestjunior`;

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(128) NOT NULL,
  `email` varchar(128) NOT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `group_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;

/*Data for the table `users` */

insert  into `users`(`id`,`username`,`email`,`created_at`,`updated_at`,`deleted_at`,`password`,`group_id`) values 
(2,'juandaantonpa','juandaantonius08@gmail.com','2022-06-22','2022-06-22 03:18:37.374',NULL,'$2a$14$l4cS6SLxQQ1YJiarnTQx3uiZ83jFhjb5Fd75Zaaj7iBCfvmTR3uMm',0),
(10,'melanibasaro','melanibasaro06@gmail.com','2022-06-23','2022-06-23 08:15:48.951',NULL,'$2a$14$UfKxVzRmJ6ODl5QNQb1/euAbbZAKANxcbMABu7vHOMQQM9vKNXh5C',0);

/*Table structure for table `wallets` */

DROP TABLE IF EXISTS `wallets`;

CREATE TABLE `wallets` (
  `wallet_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `no_phone` varchar(64) DEFAULT NULL,
  `amount` int(11) DEFAULT NULL,
  `amount_locked` int(11) DEFAULT NULL,
  `is_locked` int(11) DEFAULT NULL,
  `description` varchar(256) DEFAULT NULL,
  `created_at` date DEFAULT NULL,
  `updated_at` date DEFAULT NULL,
  `deleted_at` date DEFAULT NULL,
  PRIMARY KEY (`wallet_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `wallets_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

/*Data for the table `wallets` */

insert  into `wallets`(`wallet_id`,`user_id`,`no_phone`,`amount`,`amount_locked`,`is_locked`,`description`,`created_at`,`updated_at`,`deleted_at`) values 
(1,10,'',0,0,0,NULL,'2022-06-23','2022-06-23',NULL);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

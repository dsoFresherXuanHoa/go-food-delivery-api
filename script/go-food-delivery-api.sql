-- MySQL dump 10.13  Distrib 8.0.34, for Win64 (x86_64)
--
-- Host: localhost    Database: go_food_delivery_dev
-- ------------------------------------------------------
-- Server version	8.0.34

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `accounts`
--

DROP TABLE IF EXISTS `accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `accounts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `email` varchar(191) NOT NULL,
  `secret_code` bigint NOT NULL,
  `employee_id` bigint unsigned NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `employee_id` (`employee_id`),
  KEY `idx_accounts_deleted_at` (`deleted_at`),
  KEY `fk_roles_accounts` (`role_id`),
  CONSTRAINT `fk_employees_account` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`),
  CONSTRAINT `fk_roles_accounts` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES (1,'2023-11-03 09:26:22.525','2023-11-03 09:26:22.525',NULL,'dsoFresherXuanHoa','$2a$05$oLLG/OC6hDtE8lcC3NUhke3LIm34XMpbXDpvZd96oyCvtGX5GwIh6','dso.intern.xuanhoa@gmail.com',2248,1,1),(2,'2023-11-03 09:27:37.726','2023-11-03 09:27:37.726',NULL,'dsoInternXuanHoa','$2a$05$TjI1083YKkLAk7L2tTvZk.MgC9OvlrP/2aHCSujAS7zmMjmaFgqAy','training.dso.xuanhoa@gmail.com',6433,2,2);
/*!40000 ALTER TABLE `accounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `areas`
--

DROP TABLE IF EXISTS `areas`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `areas` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `pos` longtext NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_areas_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `areas`
--

LOCK TABLES `areas` WRITE;
/*!40000 ALTER TABLE `areas` DISABLE KEYS */;
INSERT INTO `areas` VALUES (1,'2023-11-03 09:35:15.965','2023-11-03 09:35:15.965',NULL,'F1A1'),(2,'2023-11-03 09:35:20.588','2023-11-03 09:35:20.588',NULL,'F1A2'),(3,'2023-11-03 09:35:25.808','2023-11-03 09:35:25.808',NULL,'F1A3'),(4,'2023-11-03 09:35:29.706','2023-11-03 09:35:29.706',NULL,'F1A4'),(5,'2023-11-03 09:35:44.194','2023-11-03 09:35:44.194',NULL,'F2A1'),(6,'2023-11-03 09:35:48.077','2023-11-03 09:35:48.077',NULL,'F2A2'),(7,'2023-11-03 09:35:52.229','2023-11-03 09:35:52.229',NULL,'F2A3'),(8,'2023-11-03 09:35:56.081','2023-11-03 09:35:56.081',NULL,'F2A4');
/*!40000 ALTER TABLE `areas` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `bills`
--

DROP TABLE IF EXISTS `bills`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `bills` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  `quantity` bigint NOT NULL,
  `order_id` bigint unsigned NOT NULL,
  `product_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_bills_deleted_at` (`deleted_at`),
  KEY `fk_orders_bills` (`order_id`),
  KEY `fk_products_bills` (`product_id`),
  CONSTRAINT `fk_orders_bills` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  CONSTRAINT `fk_products_bills` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bills`
--

LOCK TABLES `bills` WRITE;
/*!40000 ALTER TABLE `bills` DISABLE KEYS */;
INSERT INTO `bills` VALUES (1,'2023-11-03 09:43:55.725','2023-11-03 09:43:55.725',NULL,0,1,1,2),(2,'2023-11-03 09:43:55.814','2023-11-03 09:43:55.814',NULL,0,2,1,1),(3,'2023-11-03 09:44:17.203','2023-11-03 09:44:17.203',NULL,0,1,2,3),(4,'2023-11-03 09:44:17.269','2023-11-03 09:44:17.269',NULL,0,3,2,1);
/*!40000 ALTER TABLE `bills` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext NOT NULL,
  `description` longtext NOT NULL,
  `thumb` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_categories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'2023-11-03 09:33:44.394','2023-11-03 09:33:44.394',NULL,'Món Ăn Chính','Món Ăn Chính, Thơm Ngon Bổ Dưỡng!',NULL),(2,'2023-11-03 09:33:54.489','2023-11-03 09:33:54.489',NULL,'Món Ăn Nhẹ','Món Ăn Nhẹ, Thơm Ngon Bổ Dưỡng!',NULL),(3,'2023-11-03 09:34:06.131','2023-11-03 09:34:06.131',NULL,'Nước Giải Khát','Nước Giải Khát, Thơm Ngon Bổ Dưỡng!',NULL),(4,'2023-11-03 09:34:32.106','2023-11-03 09:34:32.106',NULL,'Điểm Tâm','Điểm Tâm, Thơm Ngon Bổ Dưỡng!',NULL);
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `discounts`
--

DROP TABLE IF EXISTS `discounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `discounts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `min_quantity` bigint NOT NULL DEFAULT '0',
  `discount_percent` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_discounts_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discounts`
--

LOCK TABLES `discounts` WRITE;
/*!40000 ALTER TABLE `discounts` DISABLE KEYS */;
INSERT INTO `discounts` VALUES (1,'2023-11-03 09:40:24.543','2023-11-03 09:40:24.543',NULL,5,10),(2,'2023-11-03 09:40:58.911','2023-11-03 09:40:58.911',NULL,10,10),(3,'2023-11-03 09:41:40.100','2023-11-03 09:41:40.100',NULL,10,15),(4,'2023-11-03 09:42:32.026','2023-11-03 09:42:32.026',NULL,6,5);
/*!40000 ALTER TABLE `discounts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `employees`
--

DROP TABLE IF EXISTS `employees`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `employees` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `full_name` longtext NOT NULL,
  `tel` varchar(191) NOT NULL,
  `gender` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `tel` (`tel`),
  KEY `idx_employees_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employees`
--

LOCK TABLES `employees` WRITE;
/*!40000 ALTER TABLE `employees` DISABLE KEYS */;
INSERT INTO `employees` VALUES (1,'2023-11-03 09:26:22.476','2023-11-03 09:26:22.476',NULL,'Le Xuan Hoa','0364015071',1),(2,'2023-11-03 09:27:37.667','2023-11-03 09:27:37.667',NULL,'Xuan Hoa Le','0356415806',1);
/*!40000 ALTER TABLE `employees` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `note` longtext,
  `status` tinyint(1) DEFAULT '0',
  `accepted` tinyint(1) DEFAULT '0',
  `rejected` tinyint(1) DEFAULT '0',
  `compensate` tinyint(1) DEFAULT '0',
  `employee_id` bigint unsigned NOT NULL,
  `table_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_orders_deleted_at` (`deleted_at`),
  KEY `fk_employees_orders` (`employee_id`),
  KEY `fk_tables_orders` (`table_id`),
  CONSTRAINT `fk_employees_orders` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`),
  CONSTRAINT `fk_tables_orders` FOREIGN KEY (`table_id`) REFERENCES `tables` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,'2023-11-03 09:43:55.671','2023-11-03 09:44:22.180',NULL,'Không Cay Không Lấy Tiền',0,1,0,0,2,1),(2,'2023-11-03 09:44:17.171','2023-11-03 09:44:26.736',NULL,'Ít Cay Thôi Nha',0,1,0,0,2,2);
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext NOT NULL,
  `description` longtext NOT NULL,
  `price` double NOT NULL,
  `thumb` longtext,
  `reorder_level` bigint DEFAULT '0',
  `stock_amount` bigint DEFAULT '0',
  `category_id` bigint unsigned NOT NULL,
  `discount_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`),
  KEY `fk_discounts_products` (`discount_id`),
  KEY `fk_categories_products` (`category_id`),
  CONSTRAINT `fk_categories_products` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`),
  CONSTRAINT `fk_discounts_products` FOREIGN KEY (`discount_id`) REFERENCES `discounts` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'2023-11-03 09:40:27.247','2023-11-03 09:44:17.305',NULL,'Cơm Chiên Dương Châu','Cơm Chiên Dương Châu',30000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1698979227/go_food_delivery_dev/hyo6vjxfprpfwqxxiylp.png',5,195,1,1),(2,'2023-11-03 09:40:59.740','2023-11-03 09:43:55.758',NULL,'Cơm Gà Xối Mỡ','Cơm Gà Xối Mỡ',35000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1698979260/go_food_delivery_dev/fep1cukus4zfhwvpgswd.png',1,299,1,2),(3,'2023-11-03 09:41:41.048','2023-11-03 09:44:17.244',NULL,'Gà Xào Xả Ớt','Cơm Gà Xối Mỡ',35000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1698979301/go_food_delivery_dev/jwhb3kwqahleozhe7wzv.png',1,199,1,3),(4,'2023-11-03 09:42:32.774','2023-11-03 09:42:32.774',NULL,'Pessi','Pessi',15000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1698979353/go_food_delivery_dev/dv82jv5syy2wuxjcyco9.jpg',0,200,3,4);
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reset_passwords`
--

DROP TABLE IF EXISTS `reset_passwords`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reset_passwords` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `managerId` bigint NOT NULL,
  `username` longtext NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_reset_passwords_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reset_passwords`
--

LOCK TABLES `reset_passwords` WRITE;
/*!40000 ALTER TABLE `reset_passwords` DISABLE KEYS */;
/*!40000 ALTER TABLE `reset_passwords` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'2023-11-03 09:24:35.322','2023-11-03 09:24:35.322',NULL,'Manager'),(2,'2023-11-03 09:24:40.631','2023-11-03 09:24:40.631',NULL,'Waiter'),(3,'2023-11-03 09:24:48.328','2023-11-03 09:24:48.328',NULL,'Chief');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tables`
--

DROP TABLE IF EXISTS `tables`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tables` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `normal_table` tinyint(1) DEFAULT '1',
  `available` tinyint(1) DEFAULT '1',
  `capacity` bigint NOT NULL,
  `base_price` double DEFAULT '20000',
  `employee_id` bigint unsigned NOT NULL,
  `area_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tables_deleted_at` (`deleted_at`),
  KEY `fk_employees_tables` (`employee_id`),
  KEY `fk_areas_tables` (`area_id`),
  CONSTRAINT `fk_areas_tables` FOREIGN KEY (`area_id`) REFERENCES `areas` (`id`),
  CONSTRAINT `fk_employees_tables` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tables`
--

LOCK TABLES `tables` WRITE;
/*!40000 ALTER TABLE `tables` DISABLE KEYS */;
INSERT INTO `tables` VALUES (1,'2023-11-03 09:36:55.817','2023-11-03 09:43:55.881',NULL,1,0,10,25000,2,2),(2,'2023-11-03 09:36:58.756','2023-11-03 09:44:17.335',NULL,1,0,10,25000,2,2),(3,'2023-11-03 09:37:03.675','2023-11-03 09:37:03.675',NULL,1,1,10,25000,2,1),(4,'2023-11-03 09:37:09.161','2023-11-03 09:37:09.161',NULL,1,1,10,25000,2,1),(5,'2023-11-03 09:37:20.975','2023-11-03 09:37:20.975',NULL,0,1,20,35000,2,1),(6,'2023-11-03 09:37:21.685','2023-11-03 09:37:21.685',NULL,0,1,20,35000,2,1),(7,'2023-11-03 09:37:27.457','2023-11-03 09:37:27.457',NULL,0,1,20,35000,1,1),(8,'2023-11-03 09:37:28.058','2023-11-03 09:37:28.058',NULL,0,1,20,35000,1,1);
/*!40000 ALTER TABLE `tables` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-11-03 10:03:14

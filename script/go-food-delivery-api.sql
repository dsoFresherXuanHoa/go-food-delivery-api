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
INSERT INTO `accounts` VALUES (1,'2023-10-20 20:37:03.955','2023-10-20 20:37:03.955',NULL,'dsoFresherXuanHoa','$2a$05$e8uVp3P2fHhisI8U4BzXVulRJ51TXvnES7yjcJdhtj2UaY8NoDFIy','dso.intern.xuanhoa@gmail.com',8594,1,1),(2,'2023-10-20 20:37:44.697','2023-10-20 22:58:22.570',NULL,'dsoInternXuanHoa','$2a$05$RWPB7i5wHB5fLQK4l6/2Cex0Zi7o4dCiCiP.S07QIA4y1o62zwEVC','training.dso.xuanhoa@gmail.com',2868,2,2);
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `areas`
--

LOCK TABLES `areas` WRITE;
/*!40000 ALTER TABLE `areas` DISABLE KEYS */;
INSERT INTO `areas` VALUES (1,'2023-10-20 20:52:35.062','2023-10-20 20:52:35.062',NULL,'F1A1'),(2,'2023-10-20 20:52:40.163','2023-10-20 20:52:40.163',NULL,'F1A2');
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
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bills`
--

LOCK TABLES `bills` WRITE;
/*!40000 ALTER TABLE `bills` DISABLE KEYS */;
INSERT INTO `bills` VALUES (1,'2023-10-20 21:08:52.239','2023-10-20 21:08:52.239',NULL,0,1,1,2),(2,'2023-10-20 21:08:52.305','2023-10-20 21:08:52.305',NULL,0,2,1,1),(3,'2023-10-20 21:09:46.973','2023-10-20 21:09:46.973',NULL,0,1,2,2),(9,'2023-10-20 22:02:01.962','2023-10-20 22:02:01.962',NULL,0,1,5,2),(10,'2023-10-20 22:02:02.006','2023-10-20 22:02:02.006',NULL,0,2,5,1),(11,'2023-10-20 22:02:41.165','2023-10-20 22:02:41.165',NULL,0,1,6,2),(12,'2023-10-20 22:02:41.221','2023-10-20 22:02:41.221',NULL,0,2,6,1),(13,'2023-10-21 09:05:39.163','2023-10-21 09:05:39.163',NULL,0,1,7,2),(14,'2023-10-21 09:05:39.252','2023-10-21 09:05:39.252',NULL,0,2,7,1);
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
INSERT INTO `categories` VALUES (1,'2023-10-20 20:39:32.393','2023-10-20 20:39:32.393',NULL,'Thức Ăn Nhanh','Thức Ăn Nhanh, Thơm Ngon Bổ Dưỡng!',NULL),(2,'2023-10-20 20:39:42.025','2023-10-20 20:39:42.025',NULL,'Bữa Chính','Bữa Chính, Thơm Ngon Bổ Dưỡng!',NULL),(3,'2023-10-20 20:39:51.760','2023-10-20 20:39:51.760',NULL,'Tráng Miệng','Tráng Miệng, Thơm Ngon Bổ Dưỡng!',NULL),(4,'2023-10-20 20:40:01.178','2023-10-20 20:40:01.178',NULL,'Trái Cây','Trái Cây, Thơm Ngon Bổ Dưỡng!',NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discounts`
--

LOCK TABLES `discounts` WRITE;
/*!40000 ALTER TABLE `discounts` DISABLE KEYS */;
INSERT INTO `discounts` VALUES (1,'2023-10-20 20:48:17.531','2023-10-20 20:48:17.531',NULL,0,10),(2,'2023-10-20 20:48:50.433','2023-10-20 20:48:50.433',NULL,0,10),(3,'2023-10-20 20:51:07.641','2023-10-20 20:51:07.641',NULL,10,10);
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
INSERT INTO `employees` VALUES (1,'2023-10-20 20:37:03.878','2023-10-20 20:37:03.878',NULL,'Xuan Hoa Le','0364015071',1),(2,'2023-10-20 20:37:44.634','2023-10-20 20:37:44.634',NULL,'Le Xuan Hoa','0356415806',1);
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
  KEY `fk_tables_orders` (`table_id`),
  KEY `fk_employees_orders` (`employee_id`),
  CONSTRAINT `fk_employees_orders` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`),
  CONSTRAINT `fk_tables_orders` FOREIGN KEY (`table_id`) REFERENCES `tables` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,'2023-10-20 21:08:52.203','2023-10-20 21:35:59.555',NULL,'Cho nhiều tỏi phi',1,0,1,0,2,1),(2,'2023-10-20 21:09:46.928','2023-10-20 21:36:15.514',NULL,'Cho cay lên',1,0,0,0,2,2),(5,'2023-10-20 22:02:01.903','2023-10-20 22:02:41.048',NULL,'Cho cay lên xíu nữa',0,0,0,1,2,2),(6,'2023-10-20 22:02:41.136','2023-10-20 22:02:41.136',NULL,'Cho cay lên xíu nữa',0,0,0,0,2,2),(7,'2023-10-21 09:05:39.117','2023-10-21 09:05:39.117',NULL,'Cho cay lên xíu nữa',0,0,0,0,1,2);
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
  KEY `fk_categories_products` (`category_id`),
  KEY `fk_discounts_products` (`discount_id`),
  CONSTRAINT `fk_categories_products` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`),
  CONSTRAINT `fk_discounts_products` FOREIGN KEY (`discount_id`) REFERENCES `discounts` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'2023-10-20 20:48:20.299','2023-10-21 09:05:39.295',NULL,'Cơm Gà Xối Mỡ','Cơm Gà Xối Mỡ',30000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1697809701/go_food_delivery_dev/f6icnsumk4l3mrftgfy7.png',14,186,3,1),(2,'2023-10-20 20:48:51.238','2023-10-21 09:05:39.229',NULL,'Cơm Chiên Dương Châu','Cơm Chiên Dương Châu',25000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1697809732/go_food_delivery_dev/o6rztmnlfd0uzevlcaid.png',7,193,3,2),(3,'2023-10-20 20:51:11.448','2023-10-20 20:51:11.448',NULL,'Pepse','Pepse',10000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1697809872/go_food_delivery_dev/rusyf65dxmmpnccryti0.jpg',0,500,4,3);
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reset_passwords`
--

LOCK TABLES `reset_passwords` WRITE;
/*!40000 ALTER TABLE `reset_passwords` DISABLE KEYS */;
INSERT INTO `reset_passwords` VALUES (7,'2023-10-20 22:58:22.595','2023-10-20 22:58:22.595',NULL,1,'dsoInternXuanHoa');
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
INSERT INTO `roles` VALUES (1,'2023-10-20 20:36:04.712','2023-10-20 20:36:04.712',NULL,'Manager'),(2,'2023-10-20 20:36:12.037','2023-10-20 20:36:12.037',NULL,'Waiter'),(3,'2023-10-20 20:36:18.977','2023-10-20 20:36:18.977',NULL,'Chief');
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
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tables`
--

LOCK TABLES `tables` WRITE;
/*!40000 ALTER TABLE `tables` DISABLE KEYS */;
INSERT INTO `tables` VALUES (1,'2023-10-20 20:53:50.456','2023-10-20 21:08:52.366',NULL,0,0,10,20000,2,1),(2,'2023-10-20 20:54:06.572','2023-10-21 09:05:39.333',NULL,0,0,10,20000,2,2),(3,'2023-10-20 20:54:16.562','2023-10-20 20:54:16.562',NULL,1,1,20,50000,2,2),(4,'2023-10-20 20:54:21.617','2023-10-20 20:54:21.617',NULL,1,1,20,50000,1,1);
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

-- Dump completed on 2023-10-21  9:59:29

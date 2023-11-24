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
  `password` longtext NOT NULL,
  `email` varchar(191) NOT NULL,
  `secret_code` bigint NOT NULL,
  `employee_id` bigint unsigned NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  UNIQUE KEY `employee_id` (`employee_id`),
  KEY `idx_accounts_deleted_at` (`deleted_at`),
  KEY `fk_roles_accounts` (`role_id`),
  CONSTRAINT `fk_employees_account` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`),
  CONSTRAINT `fk_roles_accounts` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES (1,'2023-11-24 09:38:16.874','2023-11-24 09:38:16.874',NULL,'$2a$05$4mr2xY1UjYzMavS7tLQfteW6cddPlrArIhuZ5OqYwzWPh85yrmh6S','dso.fresher.xuanhoa@gmail.com',9689,1,1),(2,'2023-11-24 09:56:32.975','2023-11-24 09:57:17.729','2023-11-24 09:57:39.521','$2a$05$1LMfCzK1aBJZAPr4bTPsxegF20mVa41uPLGsMqQcaiqquIFQU51ga','dso.junior.xuanhoa@gmail.com',3804,2,2),(3,'2023-11-24 10:22:53.055','2023-11-24 10:22:53.055',NULL,'$2a$05$KHZLh7pgxwRzF/rFaBcXnebD7AkamDmRb4UCmIvaXMsG3to.YCz9a','dso.intern.xuanhoa@gmail.com',6100,3,2);
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
INSERT INTO `areas` VALUES (1,'2023-11-24 09:39:33.132','2023-11-24 09:39:33.132',NULL,'F1A1'),(2,'2023-11-24 09:39:37.401','2023-11-24 09:39:37.401',NULL,'F1A2');
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
  `compensate` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_bills_deleted_at` (`deleted_at`),
  KEY `fk_orders_bills` (`order_id`),
  KEY `fk_products_bills` (`product_id`),
  CONSTRAINT `fk_orders_bills` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  CONSTRAINT `fk_products_bills` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bills`
--

LOCK TABLES `bills` WRITE;
/*!40000 ALTER TABLE `bills` DISABLE KEYS */;
INSERT INTO `bills` VALUES (1,'2023-11-24 09:52:09.453','2023-11-24 10:38:37.440',NULL,0,5,1,3,1),(2,'2023-11-24 09:52:09.552','2023-11-24 09:52:09.552',NULL,0,5,1,2,0),(3,'2023-11-24 09:52:09.608','2023-11-24 09:52:09.608',NULL,0,5,1,1,0),(4,'2023-11-24 09:52:22.093','2023-11-24 09:52:22.093',NULL,0,1,2,3,0),(5,'2023-11-24 09:52:22.143','2023-11-24 09:52:22.143',NULL,0,1,2,2,0),(6,'2023-11-24 09:52:22.199','2023-11-24 09:52:22.199',NULL,0,1,2,1,0),(7,'2023-11-24 10:26:16.663','2023-11-24 10:26:16.663',NULL,0,1,3,3,0),(8,'2023-11-24 10:26:16.778','2023-11-24 10:26:16.778',NULL,0,1,3,2,0),(9,'2023-11-24 10:26:16.834','2023-11-24 10:26:16.834',NULL,0,1,3,1,0),(10,'2023-11-24 10:38:37.474','2023-11-24 10:38:37.474',NULL,0,5,1,3,0);
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
  PRIMARY KEY (`id`),
  KEY `idx_categories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'2023-11-24 09:43:55.454','2023-11-24 09:43:55.454',NULL,'Món Chính','Món Chính, Thơm Ngon Bổ Dưỡng!'),(2,'2023-11-24 09:44:06.276','2023-11-24 09:44:06.276',NULL,'Món Ăn Nhẹ','Món Ăn Nhẹ, Thơm Ngon Bổ Dưỡng!'),(3,'2023-11-24 09:44:16.171','2023-11-24 09:44:16.171',NULL,'Điểm Tâm','Điểm Tâm, Thơm Ngon Bổ Dưỡng!'),(4,'2023-11-24 09:44:28.540','2023-11-24 09:44:28.540',NULL,'Tráng Miệng','Tráng Miệng, Thơm Ngon Bổ Dưỡng!'),(5,'2023-11-24 10:33:46.264','2023-11-24 10:33:46.264',NULL,'Bún','Bún'),(6,'2023-11-24 10:33:52.217','2023-11-24 10:33:52.217',NULL,'Cơm','Cơm');
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
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discounts`
--

LOCK TABLES `discounts` WRITE;
/*!40000 ALTER TABLE `discounts` DISABLE KEYS */;
INSERT INTO `discounts` VALUES (1,'2023-11-24 09:45:00.114','2023-11-24 10:05:38.511',NULL,5,5),(2,'2023-11-24 09:45:26.256','2023-11-24 09:45:26.256',NULL,5,3),(3,'2023-11-24 09:46:01.779','2023-11-24 09:46:01.779',NULL,5,5),(4,'2023-11-24 10:06:47.933','2023-11-24 10:06:47.933',NULL,10,5),(5,'2023-11-24 10:13:01.534','2023-11-24 10:13:01.534',NULL,5,10);
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
  `tel` longtext NOT NULL,
  `gender` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_employees_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employees`
--

LOCK TABLES `employees` WRITE;
/*!40000 ALTER TABLE `employees` DISABLE KEYS */;
INSERT INTO `employees` VALUES (1,'2023-11-24 09:38:16.817','2023-11-24 09:38:16.817',NULL,'Lê Xuân Hòa','0364015071',0),(2,'2023-11-24 09:56:32.925','2023-11-24 09:56:32.925','2023-11-24 09:57:39.459','Xuân Hòa Lê','0364015071',0),(3,'2023-11-24 10:22:53.005','2023-11-24 10:22:53.005',NULL,'Lê Xuân Hòa','0364015071',0);
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
  `employee_id` bigint unsigned NOT NULL,
  `table_id` bigint unsigned NOT NULL,
  `include_table_cost` tinyint(1) DEFAULT NULL,
  `total_price` bigint DEFAULT NULL,
  `reason` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_orders_deleted_at` (`deleted_at`),
  KEY `fk_tables_orders` (`table_id`),
  KEY `fk_employees_orders` (`employee_id`),
  CONSTRAINT `fk_employees_orders` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`),
  CONSTRAINT `fk_tables_orders` FOREIGN KEY (`table_id`) REFERENCES `tables` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,'2023-11-24 09:52:09.389','2023-11-24 11:23:14.106',NULL,'Không Cho Hành Lá',1,0,1,1,1,1,501500,'\"Không đủ nguyên liệu làm món Cơm gà\"'),(2,'2023-11-24 09:52:22.044','2023-11-24 10:53:48.082',NULL,'Không Cho Hành Lá',0,1,0,1,1,0,100000,NULL),(3,'2023-11-24 10:26:16.615','2023-11-24 10:26:16.864',NULL,'Không Cho Hành Lá',0,0,0,3,2,1,125000,NULL);
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
  `uint` longtext NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`),
  KEY `fk_categories_products` (`category_id`),
  KEY `fk_discounts_products` (`discount_id`),
  CONSTRAINT `fk_categories_products` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`),
  CONSTRAINT `fk_discounts_products` FOREIGN KEY (`discount_id`) REFERENCES `discounts` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'2023-11-24 09:45:03.108','2023-11-24 10:05:38.876','2023-11-24 10:05:44.861','Cơm Chiên Dương Châu','Cơm Chiên Dương Châu',35000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700793904/go_food_delivery_dev/y2vptqctr8g1k44ghedh.jpg',6,194,1,1,'Dĩa'),(2,'2023-11-24 09:45:27.128','2023-11-24 10:26:16.804',NULL,'Cơm Chả Cá','Cơm Chả Cá',30000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700793928/go_food_delivery_dev/vaadv6prshbofdg7ehwx.jpg',7,193,1,2,'Dĩa'),(3,'2023-11-24 09:46:02.543','2023-11-24 10:26:16.739',NULL,'Bún Hà Nội','Bún Hà Nội',35000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700793964/go_food_delivery_dev/wlxvgqrtzwg5p1accbzn.jpg',7,193,1,3,'Dĩa'),(4,'2023-11-24 10:06:50.836','2023-11-24 10:06:50.836',NULL,'Bánh Mỳ Chả Cá','Bánh Mỳ Chả Cá',20000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700795212/go_food_delivery_dev/iifixvddkldcqkswtoaz.jpg',0,300,3,4,''),(5,'2023-11-24 10:13:04.279','2023-11-24 10:13:04.279',NULL,'Bánh Mì Thịt Nguội','Bánh Mì Thịt Nguội',30000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700795586/go_food_delivery_dev/zwhipk4prpmy2whztrre.jpg',0,300,2,5,'');
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
  `email` longtext NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_reset_passwords_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reset_passwords`
--

LOCK TABLES `reset_passwords` WRITE;
/*!40000 ALTER TABLE `reset_passwords` DISABLE KEYS */;
INSERT INTO `reset_passwords` VALUES (1,'2023-11-24 09:57:17.785','2023-11-24 09:57:17.785',NULL,1,'dso.intern.xuanhoa@gmail.com');
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
INSERT INTO `roles` VALUES (1,'2023-11-24 09:37:44.499','2023-11-24 09:37:44.499',NULL,'Adminstrator'),(2,'2023-11-24 09:37:54.405','2023-11-24 09:37:54.405',NULL,'Waiter'),(3,'2023-11-24 09:38:02.399','2023-11-24 09:38:02.399',NULL,'Chief');
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tables`
--

LOCK TABLES `tables` WRITE;
/*!40000 ALTER TABLE `tables` DISABLE KEYS */;
INSERT INTO `tables` VALUES (1,'2023-11-24 09:40:29.769','2023-11-24 11:23:14.138',NULL,1,1,10,20000,1,1),(2,'2023-11-24 09:40:49.871','2023-11-24 10:26:16.890',NULL,0,0,15,25000,1,1);
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

-- Dump completed on 2023-11-24 13:33:48

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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES (1,'2023-11-10 08:27:25.709','2023-11-10 08:45:22.063',NULL,'$2a$05$ij9HKt39eYXdmqTiYvoqyu72BJmSVjlnf8F9xYMO6TGdcvCWtUib6','training.dso.xuanhoa@gmail.com',7648,1,2),(2,'2023-11-10 08:37:10.240','2023-11-10 08:37:10.240',NULL,'$2a$05$RTSdwKT25.Hdg5PjykT6tODeSqDj0GuC/xGrPloB097GKognbGODq','dso.intern.xuanhoa@gmail.com',2441,2,1);
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
INSERT INTO `areas` VALUES (1,'2023-11-10 09:31:04.321','2023-11-10 09:31:04.321',NULL,'F1A1'),(2,'2023-11-10 09:31:08.786','2023-11-10 09:31:08.786',NULL,'F1A2'),(3,'2023-11-10 09:31:12.917','2023-11-10 09:31:12.917',NULL,'F1A3'),(4,'2023-11-10 09:31:16.639','2023-11-10 09:31:16.639',NULL,'F1A4'),(5,'2023-11-10 09:31:21.856','2023-11-10 09:31:21.856',NULL,'F2A1'),(6,'2023-11-10 09:31:26.583','2023-11-10 09:31:26.583',NULL,'F2A2'),(7,'2023-11-10 09:31:30.052','2023-11-10 09:31:30.052',NULL,'F2A3'),(8,'2023-11-10 09:31:33.202','2023-11-10 09:31:33.202',NULL,'F2A4');
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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bills`
--

LOCK TABLES `bills` WRITE;
/*!40000 ALTER TABLE `bills` DISABLE KEYS */;
INSERT INTO `bills` VALUES (1,'2023-11-10 10:39:59.720','2023-11-10 11:27:28.975',NULL,1,2,1,2,0),(2,'2023-11-10 10:39:59.774','2023-11-10 10:39:59.774',NULL,0,2,1,2,0),(3,'2023-11-10 10:40:24.175','2023-11-10 10:40:24.175',NULL,0,1,2,3,0),(4,'2023-11-10 10:40:24.327','2023-11-10 10:40:24.327',NULL,0,3,2,1,0),(5,'2023-11-10 11:28:05.170','2023-11-10 11:28:05.170',NULL,0,1,3,3,0),(6,'2023-11-10 11:28:05.225','2023-11-10 11:28:05.225',NULL,0,3,3,1,0),(7,'2023-11-10 11:28:06.648','2023-11-10 11:28:06.648',NULL,0,1,4,3,0),(8,'2023-11-10 11:28:06.737','2023-11-10 11:28:06.737',NULL,0,3,4,1,0);
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
INSERT INTO `categories` VALUES (1,'2023-11-10 08:48:04.942','2023-11-10 08:48:04.942',NULL,'Món Chính','Món Chính, Thơm Ngon Bổ Dưỡng!',NULL),(2,'2023-11-10 08:48:27.831','2023-11-10 08:48:27.831',NULL,'Món Ăn Nhẹ','Món Ăn Nhẹ, Thơm Ngon Bổ Dưỡng!',NULL),(3,'2023-11-10 08:48:39.542','2023-11-10 08:48:39.542',NULL,'Nước Giải Khát','Nước Giải Khát, Thơm Ngon Bổ Dưỡng!',NULL),(4,'2023-11-10 08:48:59.118','2023-11-10 08:48:59.118',NULL,'Điểm Tâm','Điểm Tâm, Thơm Ngon Bổ Dưỡng!',NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discounts`
--

LOCK TABLES `discounts` WRITE;
/*!40000 ALTER TABLE `discounts` DISABLE KEYS */;
INSERT INTO `discounts` VALUES (4,'2023-11-10 09:27:10.339','2023-11-10 09:27:10.339',NULL,3,10),(5,'2023-11-10 09:28:05.359','2023-11-10 09:28:05.359',NULL,2,5),(6,'2023-11-10 09:29:31.633','2023-11-10 09:29:31.633',NULL,5,3),(11,'2023-11-10 09:48:13.629','2023-11-10 09:48:13.629',NULL,5,3);
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
INSERT INTO `employees` VALUES (1,'2023-11-10 08:27:25.673','2023-11-10 08:27:25.673',NULL,'Le Xuan Hoa','0356415806',0),(2,'2023-11-10 08:37:10.142','2023-11-10 08:37:10.142',NULL,'Le Xuan Hoa','0364015071',0);
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
  `refundable` tinyint(1) NOT NULL,
  `reason` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_orders_deleted_at` (`deleted_at`),
  KEY `fk_tables_orders` (`table_id`),
  KEY `fk_employees_orders` (`employee_id`),
  CONSTRAINT `fk_employees_orders` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`),
  CONSTRAINT `fk_tables_orders` FOREIGN KEY (`table_id`) REFERENCES `tables` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,'2023-11-10 10:39:59.678','2023-11-10 11:21:36.326',NULL,'Chỉ Nước Giải Khát Thôi',0,0,1,1,1,1,NULL),(2,'2023-11-10 10:40:24.107','2023-11-10 11:24:07.907',NULL,'Không Cho Hành Lá',1,1,0,1,1,0,NULL),(3,'2023-11-10 11:28:05.129','2023-11-10 11:28:27.144',NULL,'Không Cho Hành Lá',1,1,0,1,1,0,NULL),(4,'2023-11-10 11:28:06.605','2023-11-10 11:28:42.250',NULL,'Không Cho Hành Lá',1,1,0,1,1,0,NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'2023-11-10 09:27:13.638','2023-11-10 11:28:06.763',NULL,'Cơm Chiên Dương Châu','Cơm Chiên Dương Châu',30000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1699583234/go_food_delivery_dev/yc5k4f4i46vvlwuakbkf.png',11,189,1,4,'Dĩa'),(2,'2023-11-10 09:28:06.516','2023-11-10 10:39:59.800',NULL,'Cơm Gà Xối Mỡ','Cơm Gà Xối Mỡ',35000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1699583286/go_food_delivery_dev/evt4k2ktu0pmmdvq4auz.png',9,201,1,5,'Dĩa'),(3,'2023-11-10 09:29:33.077','2023-11-10 11:28:06.710',NULL,'Pepsi','Pepsi',10000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1699583373/go_food_delivery_dev/yiertay9ybmwk1aqaffp.jpg',3,97,3,6,'Lon'),(8,'2023-11-10 09:48:14.657','2023-11-10 09:48:14.657',NULL,'Coca','Coca',10000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1699584495/go_food_delivery_dev/zvc4wfgyqvwft1qjntkw.jpg',0,100,3,11,'Lon');
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
INSERT INTO `reset_passwords` VALUES (1,'2023-11-10 08:45:22.221','2023-11-10 08:45:22.221',NULL,2,'training.dso.xuanhoa@gmail.com');
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
INSERT INTO `roles` VALUES (1,'2023-11-10 08:25:39.847','2023-11-10 08:25:39.847',NULL,'Manager'),(2,'2023-11-10 08:25:49.660','2023-11-10 08:25:49.660',NULL,'Waiter'),(3,'2023-11-10 08:25:56.144','2023-11-10 08:25:56.144',NULL,'Chief');
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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tables`
--

LOCK TABLES `tables` WRITE;
/*!40000 ALTER TABLE `tables` DISABLE KEYS */;
INSERT INTO `tables` VALUES (1,'2023-11-10 09:32:27.647','2023-11-10 11:28:42.274',NULL,1,1,10,20000,1,1),(2,'2023-11-10 09:32:29.080','2023-11-10 09:32:29.080',NULL,1,1,10,20000,1,1),(3,'2023-11-10 09:32:29.631','2023-11-10 09:32:29.631',NULL,1,1,10,20000,1,1),(4,'2023-11-10 09:32:38.081','2023-11-10 09:32:38.081',NULL,0,1,20,30000,1,1),(5,'2023-11-10 09:32:42.475','2023-11-10 09:32:42.475',NULL,0,1,20,30000,1,2),(6,'2023-11-10 09:32:48.190','2023-11-10 09:32:48.190',NULL,0,1,20,30000,2,2),(7,'2023-11-10 09:32:52.035','2023-11-10 09:32:52.035',NULL,0,1,20,30000,2,1),(8,'2023-11-10 09:33:03.475','2023-11-10 09:33:03.475',NULL,0,1,15,20000,2,1),(9,'2023-11-10 09:33:03.995','2023-11-10 09:33:03.995',NULL,0,1,15,20000,2,1);
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

-- Dump completed on 2023-11-10 13:13:50

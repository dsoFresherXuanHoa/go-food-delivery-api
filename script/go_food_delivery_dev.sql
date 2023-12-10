-- MySQL dump 10.13  Distrib 8.0.34, for Win64 (x86_64)
--
-- Host: 172.104.181.251    Database: go_food_delivery_dev
-- ------------------------------------------------------
-- Server version	8.0.35-0ubuntu0.20.04.1

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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `accounts`
--

LOCK TABLES `accounts` WRITE;
/*!40000 ALTER TABLE `accounts` DISABLE KEYS */;
INSERT INTO `accounts` VALUES (1,'2023-11-26 08:32:16.690','2023-11-26 08:32:16.690',NULL,'$2a$05$zFnfNRsTDv.NoIABNO6kSO52RFV2j4BVZYkmxoUt0l8fP7BrxHefO','dso.intern.xuanhoa@gmail.com',9416,1,1),(2,'2023-11-26 08:32:52.254','2023-11-26 08:32:52.254',NULL,'$2a$05$5E3AoGWC4X5vCcy4oQuZ/eyMlu0pKFndsdIppo57f9HS3oDLl2pW.','nguyenvana@gmail.com',1475,2,2),(3,'2023-11-26 08:33:00.326','2023-11-26 08:33:00.326',NULL,'$2a$05$IW.3xU3jpTg1JPPeUjCWH.ItVgxuYIQ4.a3JRrmX21FMBTdRZ41Hy','nguyenvanb@gmail.com',6692,3,2),(4,'2023-11-26 08:33:10.586','2023-11-26 08:33:10.586',NULL,'$2a$05$UqPXmIR4Uqmek5AgDVO9gOlUAmrmKtKNVX.iEMoCCcH/PtcmzJczK','nguyenvanc@gmail.com',1885,4,2),(5,'2023-11-26 08:33:32.387','2023-11-26 08:33:32.387',NULL,'$2a$05$n7Bz5aYYft96Voa3F2VNgOWY6VjphsKswUMFF.wBi/Fem.5gRiWtu','nguyenthid@gmail.com',1453,5,2),(6,'2023-11-26 08:33:54.706','2023-11-26 08:33:54.706',NULL,'$2a$05$ENpZe1C20jazW.FhAx.5pexjnnk6yE3s7BKlFiVm4bXZQj5qQ4/4y','levana@gmail.com',1614,6,2),(7,'2023-11-26 08:34:02.157','2023-11-26 08:34:02.157',NULL,'$2a$05$E4k6S6y2BYXxO3vfYW3gx.jfK38D6Se32d575junTRn7W0KZSJ7IS','levanb@gmail.com',9324,7,2),(8,'2023-11-26 08:34:09.025','2023-11-26 08:34:09.025',NULL,'$2a$05$8gaIfYXOU/WhjaMCEF6VEelURCkHlzQcQyvcAC6kqVqTnJz1s.s36','levanc@gmail.com',4046,8,2),(9,'2023-11-26 08:34:25.747','2023-11-26 08:34:25.747',NULL,'$2a$05$ssH2a7QaThwly7w.mdTmQ.3FwxLwxXf7Rus6aIBygMdFHNHFbhspS','lethid@gmail.com',2226,9,2),(10,'2023-11-26 08:34:43.635','2023-11-26 08:34:43.635',NULL,'$2a$05$Lazx52ve87Djhhp2HNhr5OQjromlTIhW4O8O0MUbswfYHs5Qp1Qlu','huynhvana@gmail.com',2298,10,2),(11,'2023-11-26 08:35:09.090','2023-11-26 08:35:09.090',NULL,'$2a$05$1yGGRYRxM9oiQYsaGrvqb.s.MsM/jyynmu92LHOrLCKQiIQXl98gi','huynhvanb@gmail.com',2104,11,2),(12,'2023-11-26 08:38:02.450','2023-11-26 08:38:02.450',NULL,'$2a$05$cwFKPvwW.HnoICsGYP7JV.LVzm2CIYVZ2PE501OSndZuiVjsZV6G2','huynhvanc@gmail.com',5419,12,2),(13,'2023-11-26 08:38:29.868','2023-11-26 08:38:29.868',NULL,'$2a$05$/g3TEY4p.rRj3l43JBM3P.ys77/faZ8AxMc3M.1QFBgRcaxa9k1nG','huynhthid@gmail.com',3471,13,2),(14,'2023-11-28 08:21:12.163','2023-11-28 08:21:12.163',NULL,'$2a$05$/FbQmQ9IzL0nFRaNTjx.8uS9WzudyZc3QZ4Etvg5vjWiQXnp3apBS','vuhuyhoang@gmail.com',9384,14,3),(15,'2023-12-03 06:33:43.170','2023-12-03 06:33:43.170',NULL,'$2a$05$Lp0qoXkl6z3sO7ZURKEaGuD0q2AOIFWzctV4Lghyj4Px5rTSgCtL.','caovana@gmail.com',2957,15,2);
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
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `areas`
--

LOCK TABLES `areas` WRITE;
/*!40000 ALTER TABLE `areas` DISABLE KEYS */;
INSERT INTO `areas` VALUES (1,'2023-11-26 08:29:18.670','2023-11-26 08:29:18.670',NULL,'F1A1'),(2,'2023-11-26 08:29:24.600','2023-11-26 08:29:24.600',NULL,'F1A2'),(3,'2023-11-26 08:29:28.930','2023-11-26 08:29:28.930',NULL,'F1A3'),(4,'2023-11-26 08:29:32.799','2023-11-26 08:29:32.799',NULL,'F1A4'),(5,'2023-11-26 08:29:40.279','2023-11-26 08:29:40.279',NULL,'F2A1'),(6,'2023-11-26 08:29:45.270','2023-11-26 08:29:45.270',NULL,'F2A2'),(7,'2023-11-26 08:29:49.550','2023-11-26 08:29:49.550',NULL,'F2A3'),(8,'2023-11-26 08:29:53.990','2023-11-26 08:29:53.990',NULL,'F2A4'),(9,'2023-12-03 06:35:41.894','2023-12-03 06:35:41.894',NULL,'F3A1'),(10,'2023-12-03 06:35:49.334','2023-12-03 06:35:49.334',NULL,'F3A2'),(11,'2023-12-03 06:35:58.885','2023-12-03 06:35:58.885',NULL,'F3A3'),(12,'2023-12-03 06:36:05.850','2023-12-03 06:36:05.850',NULL,'F3A4');
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
) ENGINE=InnoDB AUTO_INCREMENT=105 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `bills`
--

LOCK TABLES `bills` WRITE;
/*!40000 ALTER TABLE `bills` DISABLE KEYS */;
INSERT INTO `bills` VALUES (1,'2023-11-26 10:19:32.094','2023-11-26 10:19:32.094',NULL,0,6,1,3,0),(2,'2023-11-26 10:19:32.097','2023-11-26 10:19:32.097',NULL,0,4,1,2,0),(3,'2023-11-26 10:19:32.101','2023-11-26 10:19:32.101',NULL,0,6,1,1,0),(4,'2023-11-26 10:19:41.250','2023-11-26 10:19:41.250',NULL,0,6,2,3,0),(5,'2023-11-26 10:19:41.255','2023-11-26 10:19:41.255',NULL,0,4,2,2,0),(6,'2023-11-26 10:19:41.257','2023-11-26 10:19:41.257',NULL,0,6,2,1,0),(7,'2023-11-26 10:20:09.348','2023-11-26 10:20:09.348',NULL,0,1,3,6,0),(8,'2023-11-26 10:20:09.352','2023-11-26 10:20:09.352',NULL,0,1,3,5,0),(9,'2023-11-26 10:20:09.356','2023-11-26 10:20:09.356',NULL,0,1,3,7,0),(10,'2023-11-26 10:20:32.762','2023-11-26 10:20:32.762',NULL,0,3,4,11,0),(11,'2023-11-26 10:20:32.766','2023-11-26 10:20:32.766',NULL,0,6,4,14,0),(12,'2023-11-26 10:20:32.770','2023-11-26 10:20:32.770',NULL,0,1,4,25,0),(13,'2023-11-26 10:21:28.654','2023-11-26 10:21:28.654',NULL,0,6,5,23,0),(14,'2023-11-26 10:21:28.657','2023-11-26 10:21:28.657',NULL,0,6,5,14,0),(15,'2023-11-26 10:21:28.661','2023-11-26 10:21:28.661',NULL,0,4,5,30,0),(16,'2023-11-26 10:21:28.663','2023-11-26 10:21:28.663',NULL,0,5,5,60,0),(17,'2023-11-28 07:53:59.396','2023-11-28 07:53:59.396',NULL,0,5,6,54,0),(18,'2023-11-28 07:53:59.402','2023-11-28 07:53:59.402',NULL,0,12,6,32,0),(19,'2023-11-28 07:53:59.404','2023-11-28 07:53:59.404',NULL,0,2,6,45,0),(20,'2023-11-28 07:54:09.925','2023-11-28 07:54:09.925',NULL,0,5,7,5,0),(21,'2023-11-28 07:54:09.929','2023-11-28 07:54:09.929',NULL,0,2,7,2,0),(22,'2023-11-28 07:54:09.932','2023-11-28 07:54:09.932',NULL,0,2,7,4,0),(23,'2023-11-28 07:54:11.005','2023-11-28 07:54:11.005',NULL,0,5,8,5,0),(24,'2023-11-28 07:54:11.009','2023-11-28 07:54:11.009',NULL,0,2,8,2,0),(25,'2023-11-28 07:54:11.012','2023-11-28 07:54:11.012',NULL,0,2,8,4,0),(26,'2023-11-28 07:54:17.603','2023-11-28 07:54:17.603',NULL,0,5,9,5,0),(27,'2023-11-28 07:54:17.606','2023-11-28 07:54:17.606',NULL,0,2,9,2,0),(28,'2023-11-28 07:54:17.610','2023-11-28 07:54:17.610',NULL,0,2,9,4,0),(29,'2023-11-28 07:54:38.972','2023-11-28 07:54:38.972',NULL,0,2,10,12,0),(30,'2023-11-28 07:54:38.976','2023-11-28 07:54:38.976',NULL,0,1,10,2,0),(31,'2023-11-28 07:54:38.980','2023-11-28 07:54:38.980',NULL,0,1,10,6,0),(32,'2023-11-28 07:56:34.787','2023-11-28 07:56:34.787',NULL,0,2,11,21,0),(33,'2023-11-28 07:56:34.791','2023-11-28 07:56:34.791',NULL,0,1,11,22,0),(34,'2023-11-28 07:56:34.795','2023-11-28 07:56:34.795',NULL,0,1,11,36,0),(35,'2023-11-28 07:56:38.308','2023-11-28 07:56:38.308',NULL,0,2,12,21,0),(36,'2023-11-28 07:56:38.311','2023-11-28 07:56:38.311',NULL,0,1,12,22,0),(37,'2023-11-28 07:56:38.315','2023-11-28 07:56:38.315',NULL,0,1,12,36,0),(38,'2023-11-28 07:56:40.233','2023-11-28 07:56:40.233',NULL,0,2,13,21,0),(39,'2023-11-28 07:56:40.237','2023-11-28 07:56:40.237',NULL,0,1,13,22,0),(40,'2023-11-28 07:56:40.241','2023-11-28 07:56:40.241',NULL,0,1,13,36,0),(41,'2023-11-28 07:56:40.981','2023-11-28 07:56:40.981',NULL,0,2,14,21,0),(42,'2023-11-28 07:56:40.984','2023-11-28 07:56:40.984',NULL,0,1,14,22,0),(43,'2023-11-28 07:56:40.987','2023-11-28 07:56:40.987',NULL,0,1,14,36,0),(44,'2023-11-28 07:56:57.093','2023-11-28 07:56:57.093',NULL,0,2,15,12,0),(45,'2023-11-28 07:56:57.096','2023-11-28 07:56:57.096',NULL,0,2,15,2,0),(46,'2023-11-28 07:56:57.099','2023-11-28 07:56:57.099',NULL,0,6,15,1,0),(47,'2023-11-28 07:57:12.874','2023-11-28 07:57:12.874',NULL,0,4,16,12,0),(48,'2023-11-28 07:57:12.877','2023-11-28 07:57:12.877',NULL,0,2,16,2,0),(49,'2023-11-28 07:57:12.880','2023-11-28 07:57:12.880',NULL,0,6,16,1,0),(50,'2023-11-28 07:58:17.294','2023-11-28 07:58:17.294',NULL,0,6,17,43,0),(51,'2023-11-28 07:58:17.297','2023-11-28 07:58:17.297',NULL,0,1,17,27,0),(52,'2023-11-28 07:58:17.299','2023-11-28 07:58:17.299',NULL,0,1,17,21,0),(53,'2023-11-28 07:58:23.346','2023-11-28 07:58:23.346',NULL,0,6,18,43,0),(54,'2023-11-28 07:58:23.350','2023-11-28 07:58:23.350',NULL,0,1,18,27,0),(55,'2023-11-28 07:58:23.353','2023-11-28 07:58:23.353',NULL,0,1,18,21,0),(56,'2023-11-28 07:58:36.123','2023-11-28 07:58:36.123',NULL,0,3,19,4,0),(57,'2023-11-28 07:58:36.128','2023-11-28 07:58:36.128',NULL,0,1,19,27,0),(58,'2023-11-28 07:58:36.131','2023-11-28 07:58:36.131',NULL,0,1,19,21,0),(59,'2023-11-28 07:58:57.711','2023-11-28 07:58:57.711',NULL,0,2,20,4,0),(60,'2023-11-28 07:58:57.714','2023-11-28 07:58:57.714',NULL,0,1,20,27,0),(61,'2023-11-28 07:58:57.716','2023-11-28 07:58:57.716',NULL,0,1,20,21,0),(62,'2023-11-28 07:59:04.651','2023-11-28 07:59:04.651',NULL,0,2,21,4,0),(63,'2023-11-28 07:59:04.653','2023-11-28 07:59:04.653',NULL,0,1,21,27,0),(64,'2023-11-28 07:59:04.656','2023-11-28 07:59:04.656',NULL,0,1,21,21,0),(65,'2023-11-28 07:59:10.424','2023-11-28 07:59:10.424',NULL,0,2,22,4,0),(66,'2023-11-28 07:59:10.428','2023-11-28 07:59:10.428',NULL,0,1,22,2,0),(67,'2023-11-28 07:59:10.431','2023-11-28 07:59:10.431',NULL,0,1,22,21,0),(68,'2023-11-28 08:05:36.138','2023-11-28 08:05:36.138',NULL,0,2,23,12,0),(69,'2023-11-28 08:05:36.141','2023-11-28 08:05:36.141',NULL,0,2,23,1,0),(70,'2023-11-28 08:05:36.144','2023-11-28 08:05:36.144',NULL,0,2,23,21,0),(71,'2023-11-28 08:05:47.121','2023-11-28 08:05:47.121',NULL,0,2,24,12,0),(72,'2023-11-28 08:05:47.125','2023-11-28 08:05:47.125',NULL,0,2,24,1,0),(73,'2023-11-28 08:05:47.129','2023-11-28 08:05:47.129',NULL,0,2,24,21,0),(74,'2023-11-28 08:06:04.041','2023-11-28 08:06:04.041',NULL,0,3,25,12,0),(75,'2023-11-28 08:06:04.045','2023-11-28 08:06:04.045',NULL,0,2,25,2,0),(76,'2023-11-28 08:06:04.050','2023-11-28 08:06:04.050',NULL,0,2,25,21,0),(77,'2023-11-28 08:07:39.175','2023-11-28 08:07:39.175',NULL,0,3,26,12,0),(78,'2023-11-28 08:07:39.178','2023-11-28 08:07:39.178',NULL,0,6,26,21,0),(79,'2023-11-28 08:07:39.181','2023-11-28 08:07:39.181',NULL,0,1,26,22,0),(80,'2023-11-28 08:07:52.283','2023-11-28 08:07:52.283',NULL,0,3,27,12,0),(81,'2023-11-28 08:07:52.286','2023-11-28 08:07:52.286',NULL,0,6,27,21,0),(82,'2023-11-28 08:08:07.985','2023-11-28 08:08:07.985',NULL,0,3,28,12,0),(83,'2023-11-28 08:08:07.988','2023-11-28 08:08:07.988',NULL,0,6,28,21,0),(84,'2023-11-28 08:08:07.992','2023-11-28 08:08:07.992',NULL,0,1,28,1,0),(85,'2023-11-28 08:08:07.995','2023-11-28 08:08:07.995',NULL,0,2,28,2,0),(86,'2023-11-28 08:08:08.004','2023-11-28 08:08:08.004',NULL,0,3,28,3,0),(87,'2023-11-28 08:08:23.502','2023-11-28 08:08:23.502',NULL,0,3,29,12,0),(88,'2023-11-28 08:08:23.506','2023-11-28 08:08:23.506',NULL,0,6,29,21,0),(89,'2023-11-28 08:08:23.509','2023-11-28 08:08:23.509',NULL,0,1,29,1,0),(90,'2023-11-28 08:08:23.511','2023-11-28 08:08:23.511',NULL,0,2,29,2,0),(91,'2023-11-28 08:08:23.514','2023-11-28 08:08:23.514',NULL,0,3,29,3,0),(92,'2023-11-28 08:11:23.773','2023-11-28 08:11:23.773',NULL,0,3,30,12,0),(93,'2023-11-28 08:11:23.777','2023-11-28 08:11:23.777',NULL,0,1,30,1,0),(94,'2023-11-28 08:11:23.780','2023-11-28 08:11:23.780',NULL,0,2,30,2,0),(95,'2023-11-28 08:11:23.784','2023-11-28 08:11:23.784',NULL,0,3,30,3,0),(96,'2023-11-28 08:11:46.472','2023-11-28 08:11:46.472',NULL,0,3,31,12,0),(97,'2023-11-28 08:11:46.476','2023-11-28 08:11:46.476',NULL,0,12,31,3,0),(98,'2023-11-28 08:11:46.479','2023-11-28 08:11:46.479',NULL,0,1,31,2,0),(99,'2023-11-28 08:11:46.482','2023-11-28 08:11:46.482',NULL,0,3,31,3,0),(100,'2023-11-28 08:11:53.909','2023-11-28 08:11:53.909',NULL,0,3,32,12,0),(101,'2023-11-28 08:11:53.912','2023-11-28 08:11:53.912',NULL,0,12,32,3,0),(102,'2023-11-28 08:11:53.916','2023-11-28 08:11:53.916',NULL,0,1,32,2,0),(103,'2023-11-28 08:11:53.920','2023-11-28 08:11:53.920',NULL,0,3,32,3,0),(104,'2023-11-28 08:16:46.148','2023-11-28 08:16:46.148',NULL,0,3,33,1,0);
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
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'2023-11-26 09:14:20.079','2023-11-26 09:14:20.079',NULL,'Loại viên','Loại viên!'),(2,'2023-11-26 09:14:38.179','2023-11-26 09:14:38.179',NULL,'Món ăn kinh điển','Món ăn kinh điển'),(3,'2023-11-26 09:14:49.392','2023-11-26 09:14:49.392',NULL,'Giải Khát','Giải Khát'),(4,'2023-11-26 09:15:00.442','2023-11-26 09:15:00.442',NULL,'Loại lẩu','Loại lẩu'),(5,'2023-11-26 09:15:13.638','2023-11-26 09:15:13.638',NULL,'Loại thịt','Loại thịt'),(6,'2023-11-26 09:15:27.079','2023-11-26 09:15:27.079',NULL,'Hải Sản','Hải Sản'),(7,'2023-11-26 09:15:38.868','2023-11-26 09:15:38.868',NULL,'Rau củ','Rau củ'),(8,'2023-11-26 09:15:50.919','2023-11-26 09:15:50.919',NULL,'Đồ Ăn Vặt','Đồ Ăn Vặt'),(9,'2023-11-26 09:16:10.328','2023-11-26 09:16:10.328',NULL,'Đồ uống có cồn','Đồ uống có cồn'),(10,'2023-11-26 09:57:32.626','2023-11-26 09:57:32.626',NULL,'Bánh Mì','Bánh Mì'),(11,'2023-11-26 09:57:51.015','2023-11-26 09:57:51.015',NULL,'Cơm','Cơm'),(12,'2023-11-26 09:58:00.956','2023-11-26 09:58:00.956',NULL,'Bún','Bún'),(13,'2023-11-26 09:58:12.336','2023-11-26 09:58:12.336',NULL,'Phở','Phở'),(14,'2023-11-26 09:58:26.283','2023-11-26 09:58:26.283',NULL,'Trái Cây','Trái Cây');
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
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discounts`
--

LOCK TABLES `discounts` WRITE;
/*!40000 ALTER TABLE `discounts` DISABLE KEYS */;
INSERT INTO `discounts` VALUES (1,'2023-11-26 09:25:22.000','2023-11-26 09:25:22.000',NULL,5,5),(2,'2023-11-26 09:25:50.571','2023-11-26 09:25:50.571',NULL,5,5),(3,'2023-11-26 09:26:16.831','2023-11-26 09:26:16.831',NULL,5,5),(4,'2023-11-26 09:26:35.925','2023-11-26 09:26:35.925',NULL,5,5),(5,'2023-11-26 09:26:49.935','2023-11-26 09:26:49.935',NULL,5,3),(6,'2023-11-26 09:28:07.686','2023-11-26 09:28:07.686',NULL,6,3),(7,'2023-11-26 09:28:25.553','2023-11-26 09:28:25.553',NULL,6,3),(8,'2023-11-26 09:28:57.589','2023-11-26 09:28:57.589',NULL,5,5),(9,'2023-11-26 09:29:38.178','2023-11-26 09:29:38.178',NULL,5,3),(10,'2023-11-26 09:30:43.017','2023-11-26 09:30:43.017',NULL,2,5),(11,'2023-11-26 09:31:15.518','2023-11-26 09:31:15.518',NULL,2,5),(12,'2023-11-26 09:31:39.163','2023-11-26 09:31:39.163',NULL,2,3),(13,'2023-11-26 09:32:31.256','2023-11-26 09:32:31.256',NULL,2,3),(14,'2023-11-26 09:33:11.577','2023-11-26 09:33:11.577',NULL,3,5),(15,'2023-11-26 09:33:32.633','2023-11-26 09:33:32.633',NULL,3,3),(16,'2023-11-26 09:34:13.291','2023-11-26 09:34:13.291',NULL,3,10),(17,'2023-11-26 09:35:10.757','2023-11-26 09:35:10.757',NULL,3,10),(18,'2023-11-26 09:35:34.972','2023-11-26 09:35:34.972',NULL,3,5),(19,'2023-11-26 09:36:03.375','2023-11-26 09:36:03.375',NULL,3,3),(20,'2023-11-26 09:36:33.153','2023-11-26 09:36:33.153',NULL,3,5),(21,'2023-11-26 09:37:30.266','2023-11-26 09:37:30.266',NULL,2,10),(22,'2023-11-26 09:38:00.942','2023-11-26 09:38:00.942',NULL,2,5),(23,'2023-11-26 09:38:34.495','2023-11-26 09:38:34.495',NULL,2,3),(24,'2023-11-26 09:39:02.221','2023-11-26 09:39:02.221',NULL,5,6),(25,'2023-11-26 09:39:36.542','2023-11-26 09:39:36.542',NULL,5,10),(26,'2023-11-26 09:40:20.682','2023-11-26 09:40:20.682',NULL,0,0),(27,'2023-11-26 09:41:01.862','2023-11-26 09:41:01.862',NULL,3,2),(28,'2023-11-26 09:42:33.572','2023-11-26 09:42:33.572',NULL,10,2),(29,'2023-11-26 09:43:03.096','2023-11-26 09:43:03.096',NULL,5,10),(30,'2023-11-26 09:43:34.059','2023-11-26 09:43:34.059',NULL,5,10),(31,'2023-11-26 09:44:09.522','2023-11-26 09:44:09.522',NULL,5,5),(32,'2023-11-26 09:44:42.768','2023-11-26 09:44:42.768',NULL,5,3),(33,'2023-11-26 09:45:21.455','2023-11-26 09:45:21.455',NULL,5,3),(34,'2023-11-26 09:45:56.520','2023-11-26 09:45:56.520',NULL,5,2),(35,'2023-11-26 09:46:26.961','2023-11-26 09:46:26.961',NULL,5,5),(36,'2023-11-26 10:00:19.913','2023-11-26 10:00:19.913',NULL,5,5),(37,'2023-11-26 10:00:50.369','2023-11-26 10:00:50.369',NULL,5,10),(38,'2023-11-26 10:01:23.702','2023-11-26 10:01:23.702',NULL,10,10),(39,'2023-11-26 10:02:07.381','2023-11-26 10:02:07.381',NULL,5,10),(40,'2023-11-26 10:02:36.882','2023-11-26 10:02:36.882',NULL,5,5),(41,'2023-11-26 10:03:19.775','2023-11-26 10:03:19.775',NULL,3,3),(42,'2023-11-26 10:03:48.153','2023-11-26 10:03:48.153',NULL,0,0),(43,'2023-11-26 10:04:22.460','2023-11-26 10:04:22.460',NULL,5,3),(44,'2023-11-26 10:05:12.293','2023-11-26 10:05:12.293',NULL,0,0),(45,'2023-11-26 10:05:35.785','2023-11-26 10:05:35.785',NULL,0,0),(46,'2023-11-26 10:06:11.000','2023-11-26 10:06:11.000',NULL,0,0),(47,'2023-11-26 10:06:31.601','2023-11-26 10:06:31.601',NULL,5,5),(48,'2023-11-26 10:06:49.851','2023-11-26 10:06:49.851',NULL,3,5),(49,'2023-11-26 10:07:27.295','2023-11-26 10:07:27.295',NULL,2,5),(50,'2023-11-26 10:07:58.493','2023-11-26 10:07:58.493',NULL,3,10),(51,'2023-11-26 10:08:48.881','2023-11-26 10:08:48.881',NULL,5,10),(52,'2023-11-26 10:09:25.274','2023-11-26 10:09:25.274',NULL,5,5),(53,'2023-11-26 10:09:58.757','2023-11-26 10:09:58.757',NULL,2,5),(54,'2023-11-26 10:10:34.028','2023-11-26 10:10:34.028',NULL,5,5),(55,'2023-11-26 10:12:02.183','2023-11-26 10:12:02.183',NULL,5,5),(56,'2023-11-26 10:12:38.312','2023-11-26 10:12:38.312',NULL,0,0),(57,'2023-11-26 10:13:00.989','2023-11-26 10:13:00.989',NULL,0,0),(58,'2023-11-26 10:13:16.560','2023-11-26 10:13:16.560',NULL,0,0),(59,'2023-11-26 10:13:37.514','2023-11-26 10:13:37.514',NULL,0,0),(60,'2023-11-26 10:14:03.497','2023-11-26 10:14:03.497',NULL,0,0);
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
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `employees`
--

LOCK TABLES `employees` WRITE;
/*!40000 ALTER TABLE `employees` DISABLE KEYS */;
INSERT INTO `employees` VALUES (1,'2023-11-26 08:32:16.684','2023-11-26 08:32:16.684',NULL,'Lê Xuân Hòa','0364015071',0),(2,'2023-11-26 08:32:52.249','2023-11-26 08:32:52.249',NULL,'Nguyen Van A','0364015071',0),(3,'2023-11-26 08:33:00.321','2023-11-26 08:33:00.321',NULL,'Nguyen Van B','0364015071',0),(4,'2023-11-26 08:33:10.580','2023-11-26 08:33:10.580',NULL,'Nguyen Van C','0364015071',0),(5,'2023-11-26 08:33:32.381','2023-11-26 08:33:32.381',NULL,'Nguyen Thi D','0364015071',1),(6,'2023-11-26 08:33:54.701','2023-11-26 08:33:54.701',NULL,'Le Van A','0364015071',0),(7,'2023-11-26 08:34:02.151','2023-11-26 08:34:02.151',NULL,'Le Van B','0364015071',0),(8,'2023-11-26 08:34:09.021','2023-11-26 08:34:09.021',NULL,'Le Van C','0364015071',0),(9,'2023-11-26 08:34:25.742','2023-11-26 08:34:25.742',NULL,'Le Thi D','0364015071',1),(10,'2023-11-26 08:34:43.631','2023-11-26 08:34:43.631',NULL,'Huynh Van A','0364015071',0),(11,'2023-11-26 08:35:09.085','2023-11-26 08:35:09.085',NULL,'Huynh Van B','0364015071',0),(12,'2023-11-26 08:38:02.446','2023-11-26 08:38:02.446',NULL,'Huynh Van C','0364015071',0),(13,'2023-11-26 08:38:29.863','2023-11-26 08:38:29.863',NULL,'Huynh Thi D ','0364015071',0),(14,'2023-11-28 08:21:12.157','2023-11-28 08:21:12.157',NULL,'Vũ Huy Hoàng','0364015071',0),(15,'2023-12-03 06:33:43.156','2023-12-03 06:33:43.156',NULL,'Cao Văn A','0364015071',0);
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
  KEY `fk_employees_orders` (`employee_id`),
  KEY `fk_tables_orders` (`table_id`),
  CONSTRAINT `fk_employees_orders` FOREIGN KEY (`employee_id`) REFERENCES `employees` (`id`),
  CONSTRAINT `fk_tables_orders` FOREIGN KEY (`table_id`) REFERENCES `tables` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,'2023-11-26 10:19:32.086','2023-11-28 08:22:36.407',NULL,'Không Cho Hành Lá',0,1,0,2,1,1,1421200,NULL),(2,'2023-11-26 10:19:41.243','2023-11-28 08:22:42.319',NULL,'Không Cho Hành Lá',0,1,0,2,1,0,1421200,NULL),(3,'2023-11-26 10:20:09.344','2023-11-28 08:23:01.276',NULL,'Không Cho Hành Lá',0,1,0,2,2,1,121000,NULL),(4,'2023-11-26 10:20:32.754','2023-11-28 08:23:12.051',NULL,'Không Cho Hành Lá',0,1,0,2,3,1,352825,NULL),(5,'2023-11-26 10:21:28.646','2023-11-28 08:23:16.865',NULL,'Ngon Quá Đi À',0,1,0,2,4,1,1074260,NULL),(6,'2023-11-28 07:53:59.386','2023-11-28 07:53:59.408',NULL,'Cho Nhiều Hành Phi',0,1,0,3,5,1,1633940,NULL),(7,'2023-11-28 07:54:09.918','2023-11-28 07:54:09.936',NULL,'Cho Nhiều Hành Phi',0,1,0,3,6,1,825000,NULL),(8,'2023-11-28 07:54:10.999','2023-11-28 07:54:11.014',NULL,'Cho Nhiều Hành Phi',0,1,0,3,6,1,825000,NULL),(9,'2023-11-28 07:54:17.596','2023-11-28 07:54:17.613',NULL,'Cho Nhiều Hành Phi',0,0,1,3,6,0,825000,NULL),(10,'2023-11-28 07:54:38.965','2023-11-28 07:54:38.985',NULL,'Cho Nhiều Hành Phi',0,1,0,3,7,0,176000,NULL),(11,'2023-11-28 07:56:34.782','2023-11-28 07:56:34.798',NULL,'Cho Nhiều Hành Phi',0,0,0,4,9,1,242000,NULL),(12,'2023-11-28 07:56:38.301','2023-11-28 07:56:38.323',NULL,'Cho Nhiều Hành Phi',0,1,0,4,9,1,242000,NULL),(13,'2023-11-28 07:56:40.228','2023-11-28 08:24:59.759',NULL,'Cho Nhiều Hành Phi',0,0,1,4,9,1,242000,'\"Không đủ nguyên liệu\"'),(14,'2023-11-28 07:56:40.976','2023-11-28 08:25:07.824',NULL,'Cho Nhiều Hành Phi',0,0,1,4,9,1,242000,'\"Không đủ nguyên liệu\"'),(15,'2023-11-28 07:56:57.088','2023-11-28 07:56:57.102',NULL,'Cho Nhiều Hành Phi',0,1,0,4,10,1,765600,NULL),(16,'2023-11-28 07:57:12.868','2023-11-28 07:57:12.885',NULL,'Cho Nhiều Hành Phi',0,0,1,4,11,1,817300,NULL),(17,'2023-11-28 07:58:17.289','2023-11-28 07:58:17.302',NULL,'Cho Nhiều Hành Phi',0,1,0,5,13,1,216040,NULL),(18,'2023-11-28 07:58:23.339','2023-11-28 07:58:23.356',NULL,'Cho Nhiều Hành Phi',0,1,0,5,13,0,216040,NULL),(19,'2023-11-28 07:58:36.116','2023-11-28 07:58:36.134',NULL,'Cho Nhiều Hành Phi',0,1,0,5,14,0,352000,NULL),(20,'2023-11-28 07:58:57.707','2023-11-28 07:58:57.720',NULL,'Cho Nhiều Hành Phi',0,0,1,5,15,1,264000,NULL),(21,'2023-11-28 07:59:04.646','2023-11-28 07:59:04.659',NULL,'Cho Nhiều Hành Phi',0,0,0,5,15,0,264000,NULL),(22,'2023-11-28 07:59:10.418','2023-11-28 07:59:10.435',NULL,'Cho Nhiều Hành Phi',0,1,0,5,15,0,346500,NULL),(23,'2023-11-28 08:05:36.132','2023-11-28 08:05:36.147',NULL,'Cho Ít Hành Phi',0,0,0,6,17,1,363000,NULL),(24,'2023-11-28 08:05:47.115','2023-11-28 08:05:47.134',NULL,'Cho Ít Hành Phi',0,1,0,6,18,0,363000,NULL),(25,'2023-11-28 08:06:04.034','2023-11-28 08:06:04.054',NULL,'Cho Ít Hành Phi',0,0,0,6,18,1,421025,NULL),(26,'2023-11-28 08:07:39.169','2023-11-28 08:07:39.184',NULL,'Cho Vừa Hành Phi',0,1,0,7,21,1,507925,NULL),(27,'2023-11-28 08:07:52.277','2023-11-28 08:07:52.290',NULL,'Cho Vừa Hành Phi',0,0,1,7,22,1,436425,NULL),(28,'2023-11-28 08:08:07.977','2023-11-28 08:08:08.007',NULL,'Cho Vừa Hành Phi',0,1,0,7,22,0,997425,NULL),(29,'2023-11-28 08:08:23.492','2023-11-28 08:08:23.517',NULL,'Cho Vừa Hành Phi',0,0,0,7,24,1,1019425,NULL),(30,'2023-11-28 08:11:23.765','2023-11-28 08:11:23.787',NULL,'Cho Vừa Vừa Hành Phi',0,0,0,8,25,1,641025,NULL),(31,'2023-11-28 08:11:46.464','2023-11-28 08:11:46.485',NULL,'Cho Vừa Vừa Hành Phi',0,1,0,8,27,1,1451725,NULL),(32,'2023-11-28 08:11:53.903','2023-11-28 08:11:53.924',NULL,'Cho Vừa Vừa Hành Phi',0,0,0,8,27,0,1451725,NULL),(33,'2023-11-28 08:16:46.143','2023-11-28 08:16:46.153',NULL,'Cho Vừa Vừa Hành Phi',0,0,1,9,29,1,286000,NULL);
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
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'2023-11-26 09:25:24.375','2023-11-28 08:16:46.150',NULL,'Lẩu Dầu Cay','Lẩu Dầu Cay',80000,'https://th.bing.com/th/id/OIP.W7jovBMzZr3cE6sEel-fkAHaE7?rs=1&pid=ImgDetMain',34,166,4,1,'Phần'),(2,'2023-11-26 09:25:52.450','2023-11-28 08:11:53.918',NULL,'Lẩu Thảo Dược','Lẩu Thảo Dược',95000,'https://th.bing.com/th/id/R.bf100e86a32a0c5b680640702c22df3e?rik=cnXxADtv1KZZ2A&pid=ImgRaw&r=0',30,170,4,2,'Phần'),(3,'2023-11-26 09:26:18.523','2023-11-28 08:11:53.922',NULL,'Lẩu Nấm','Lẩu Nấm',80000,'https://th.bing.com/th/id/OIP.ZhBDFzarleYd1_q-xaaLTAHaDk?rs=1&pid=ImgDetMain',51,149,4,3,'Phần'),(4,'2023-11-26 09:26:37.609','2023-11-28 07:59:10.426',NULL,'Lẩu Cà Chua','Lẩu Cà Chua',80000,'https://th.bing.com/th/id/R.95e4cbaa4d66ffeee824f0d608e73684?rik=DOn1AGWVvQu9IA&pid=ImgRaw&r=0',15,185,4,4,'Phần'),(5,'2023-11-26 09:26:51.729','2023-11-28 07:54:17.605',NULL,'Lẩu Thái','Lẩu Thái',80000,'https://th.bing.com/th/id/R.3dc54e8f5df2eaa53ca93758d080b2f0?rik=vygBSkUueMC71g&pid=ImgRaw&r=0',16,184,4,5,'Phần'),(6,'2023-11-26 09:28:09.373','2023-11-28 07:54:38.983',NULL,'Pepsi','Pepsi',15000,'https://th.bing.com/th/id/R.cfbef0028c9bbf4118512c896c6a22aa?rik=UsMKOcnz3t1FJw&pid=ImgRaw&r=0',2,598,3,6,'Lon'),(7,'2023-11-26 09:28:27.191','2023-11-26 10:20:09.358',NULL,'Coca','Coca',15000,'https://th.bing.com/th/id/OIP.tDh0_aD0IV6EoI8MZZHscAHaHa?rs=1&pid=ImgDetMain',1,599,3,7,'Lon'),(8,'2023-11-26 09:28:59.253','2023-11-26 09:28:59.253',NULL,'Bia Tiger','Bia Tiger',24000,'https://th.bing.com/th/id/R.7478626a15649bb8b355effc16190c68?rik=abXQuOvnUTbOog&pid=ImgRaw&r=0',0,600,9,8,'Lon'),(9,'2023-11-26 09:29:39.854','2023-11-26 09:29:39.854',NULL,'Nước ép Nhiệt Đới','Nước ép Nhiệt Đới',30000,'https://th.bing.com/th/id/R.7478626a15649bb8b355effc16190c68?rik=abXQuOvnUTbOog&pid=ImgRaw&r=0',0,300,3,9,'Chai'),(10,'2023-11-26 09:30:44.781','2023-11-26 09:30:44.781',NULL,'Nấm đùi gà','Nấm đùi gà',30000,'https://th.bing.com/th/id/OIP.srs_4BcTSyKEo59z6fTJrgHaHa?rs=1&pid=ImgDetMain',0,300,7,10,'Phần'),(11,'2023-11-26 09:31:17.186','2023-11-26 10:20:32.765',NULL,'Nấm Rơm','Nấm Rơm',35000,'https://th.bing.com/th/id/R.87d34129428cc01503749c2e6d234c6b?rik=WdjuqI%2flpVniyg&pid=ImgRaw&r=0',3,197,7,11,'Phần'),(12,'2023-11-26 09:31:40.862','2023-11-28 08:11:53.911',NULL,'Cải Ngọt','Cải Ngọt',25000,'https://th.bing.com/th/id/OIP.VX5uDPkGqtUXIU2pGKC2ywHaHa?rs=1&pid=ImgDetMain',36,164,7,12,'Phần'),(13,'2023-11-26 09:32:32.959','2023-11-26 09:32:32.959',NULL,'Cải Thìa','Cải Thìa',25000,'https://th.bing.com/th/id/OIP.Tb5ycKTOJKgN2VTr9LXrPAHaEL?rs=1&pid=ImgDetMain',0,200,7,13,'Phần'),(14,'2023-11-26 09:33:13.235','2023-11-26 10:21:28.660',NULL,'Cá Viên','Cá Viên',30000,'https://th.bing.com/th/id/R.728afac42cb8da23d9e99422a44ed6fb?rik=9GBvEi0PeW7Dxg&pid=ImgRaw&r=0',12,288,1,14,'Phần'),(15,'2023-11-26 09:33:34.243','2023-11-26 09:33:34.243',NULL,'Bò Viên','Bò Viên',30000,'https://th.bing.com/th/id/R.c4d9a705a80384829667056df3dea4f2?rik=CE4%2b5xEZJ3CQeQ&pid=ImgRaw&r=0',0,300,1,15,'Phần'),(16,'2023-11-26 09:34:14.985','2023-11-26 09:34:14.985',NULL,'Mực Sữa','Mực Sữa',70000,'https://th.bing.com/th/id/OIP.5yQu9BIxuol2-oXxhD1eLQHaHa?rs=1&pid=ImgDetMain',0,200,6,16,'Phần'),(17,'2023-11-26 09:35:12.434','2023-11-26 09:35:12.434',NULL,'Tôm Sú','Tôm Sú',90000,'https://th.bing.com/th/id/R.ea5875e0ff5e288d1e423ff788daa28c?rik=8ClxCE9Rqs8vkg&riu=http%3a%2f%2fwww.tomgionghoaphatcp.com.vn%2fuploads%2fpost%2fimg_1541476591742_o_1crjika3phas160smg11ad7nd7d.jpg&ehk=uj3Q1lURuBT1KuBPZMJgmf72GWH5IUYIAPg9FrccwMI%3d&risl=&pid=ImgRaw&r=0',0,200,6,17,'Phần'),(18,'2023-11-26 09:35:36.801','2023-11-26 09:35:36.801',NULL,'Cá Hồi','Cá Hồi',75000,'https://th.bing.com/th/id/OIP.hqm93NR354bg9AS3Ch1PpwHaHa?rs=1&pid=ImgDetMain',0,100,6,18,'Phần'),(19,'2023-11-26 09:36:04.898','2023-11-26 09:36:04.898',NULL,'Thịt Bò','Thịt Bò',120000,'https://th.bing.com/th/id/R.f12e4dd3c245d98996e6b7e88c081063?rik=kpJsO3N2plG1SA&riu=http%3a%2f%2fthucphamtoancau.com%2fupload%2ffiles%2fsanpham%2fthit-bo-uc%5b1%5d.jpg&ehk=H5oUlxWRcfq1O93B2j%2bCi1keZN2rMF4%2bKvcmGvfTa38%3d&risl=&pid=ImgRaw&r=0',0,150,5,19,'Phần'),(20,'2023-11-26 09:36:34.735','2023-11-26 09:36:34.735',NULL,'Ba Chỉ Heo','Ba Chỉ Heo',99000,'https://th.bing.com/th/id/OIP.Dko67zak3OZaO5ApLG0OLAHaHa?rs=1&pid=ImgDetMain',0,200,5,20,'Phần'),(21,'2023-11-26 09:37:31.823','2023-11-28 08:08:23.508',NULL,'Óc Heo','Óc Heo',60000,'https://yt.cdnxbvn.com/medias/uploads/68/68846-1.jpg',44,56,2,21,'Phần'),(22,'2023-11-26 09:38:02.625','2023-11-28 08:07:39.183',NULL,'Bao Tử Cá Basa','Bao Tử Cá Basa',65000,'https://th.bing.com/th/id/R.30129ecd2881e2ec091996af03e84995?rik=E1Gmain7cZ3Xmw&pid=ImgRaw&r=0',5,195,2,22,'Phần'),(23,'2023-11-26 09:38:36.179','2023-11-26 10:21:28.656',NULL,'Lá Sách','Lá Sách',80000,'https://i.pinimg.com/originals/88/41/3f/88413f8503ef047251866da3529f1cdd.jpg',6,144,2,23,'Phần'),(24,'2023-11-26 09:39:03.944','2023-11-26 09:39:03.944',NULL,'Xương To','Xương To',100000,'https://image.cooky.vn/recipe/g5/44818/s/cooky-recipe-cover-r44818.jpg',0,300,2,24,'Phần'),(25,'2023-11-26 09:39:38.166','2023-11-26 10:20:32.771',NULL,'Đậu Phụ Cá','Đậu Phụ Cá',50000,'https://th.bing.com/th/id/OIP.fQVHubj3rMvjHwmW38mhEQHaHa?w=736&h=736&rs=1&pid=ImgDetMain',1,199,2,25,'Phần'),(26,'2023-11-26 09:40:22.354','2023-11-26 09:40:22.354',NULL,'Cơm Trắng','Cơm Trắng',15000,'https://th.bing.com/th/id/OIP.Wp9Sc9zad0qd56VH1VXPHgHaGW?rs=1&pid=ImgDetMain',0,600,8,26,'Phần'),(27,'2023-11-26 09:41:03.561','2023-11-28 07:59:04.655',NULL,'Phở Chua Cay','Phở Chua Cay',20000,'https://th.bing.com/th/id/R.e6554b81fc6c9c1027122ba678f4ad9c?rik=HTIFeHELRMHZYA&pid=ImgRaw&r=0',5,295,8,27,'Phần'),(28,'2023-11-26 09:42:35.961','2023-11-26 09:42:35.961',NULL,'Màn Thầu Chiên','Màn Thầu Chiên',40000,'https://sgdirectory.com/wp-content/uploads/2022/10/Famous-Haidilao-Menu-Singapore.jpg',0,200,8,28,'Phần'),(29,'2023-11-26 09:43:04.728','2023-11-26 09:43:04.728',NULL,'Cơm Chiên','Cơm Chiên',20000,'https://th.bing.com/th/id/OIP.wDiRyTJaRlEVolLe5WZBZQHaHa?rs=1&pid=ImgDetMain',0,300,8,29,'Phần'),(30,'2023-11-26 09:43:35.726','2023-11-26 10:21:28.662',NULL,'Nem Tôm','Nem Tôm',30000,'https://th.bing.com/th/id/OIP.l-1uEZvyMXl6ZcvCEdV37wHaE7?rs=1&pid=ImgDetMain',4,196,8,30,'Phần'),(31,'2023-11-26 09:44:11.105','2023-11-26 09:44:11.105',NULL,'Thịt Bò Bông Tuyết Mỹ','Thịt Bò Bông Tuyết Mỹ',145000,'https://th.bing.com/th/id/OIP.dCdQzO1GOEq_rvKM9w6mSQHaGD?w=991&h=811&rs=1&pid=ImgDetMain',0,200,5,31,'Phần'),(32,'2023-11-26 09:44:44.465','2023-11-28 07:53:59.403',NULL,'Thịt Dê','Thịt Dê',110000,'https://th.bing.com/th/id/R.c4f1a5404d92650b1bdb9543c61ed466?rik=vN2EWPUbknjwxQ&riu=http%3a%2f%2fmedia.vietq.vn%2ffiles%2fcach-chon-thit-de.png&ehk=XeCMLfJ8DbaWEnhweMFPVPdayhLRWBxtVSgRS%2bqOr8g%3d&risl=&pid=ImgRaw&r=0',12,188,5,32,'Phần'),(33,'2023-11-26 09:45:23.142','2023-11-26 09:45:23.142',NULL,'Ba Chỉ Bò','Ba Chỉ Bò',115000,'https://th.bing.com/th/id/R.fb253fd561caf0dc6132e17fc1711410?rik=aV2B7oQtE7pmwA&pid=ImgRaw&r=0',0,200,5,33,'Phần'),(34,'2023-11-26 09:45:58.160','2023-11-26 09:45:58.160',NULL,'Viên Bò Phô Mai','Viên Bò Phô Mai',35000,'https://th.bing.com/th/id/R.1266ffae974a59b7f8315479aced58bc?rik=aKX8GzmUaCA%2fcw&pid=ImgRaw&r=0',0,200,1,34,'Phần'),(35,'2023-11-26 09:46:28.682','2023-11-26 09:46:28.682',NULL,'Viên Phật Sơn','Viên Phật Sơn',40000,'https://th.bing.com/th/id/R.31212934ac0a0823cd36aefbde22638a?rik=f8skooGMvsvQyw&pid=ImgRaw&r=0',0,200,1,35,'Phần'),(36,'2023-11-26 10:00:21.673','2023-11-28 07:56:40.989',NULL,'Cơm tấm sườn nướng','Cơm tấm sườn nướng',35000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700992821/go_food_delivery_dev/ogxh6me7jgsxv8yll4pw.jpg',4,96,11,36,'Phần'),(37,'2023-11-26 10:00:51.896','2023-11-26 10:00:51.896',NULL,'Cơm tấm chả cá','Cơm tấm chả cá',37000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700992851/go_food_delivery_dev/ysdysfgyjxpafx9xulrm.jpg',0,100,11,37,'Phần'),(38,'2023-11-26 10:01:24.812','2023-11-26 10:01:24.812',NULL,'Cơm tấm gà','Cơm tấm gà',38000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700992884/go_food_delivery_dev/wlfpdvz7dfylklzmsoua.jpg',0,100,11,38,'Phần'),(39,'2023-11-26 10:02:07.940','2023-11-26 10:02:07.940',NULL,'Cơm tấm tôm','Cơm tấm tôm',39000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700992927/go_food_delivery_dev/sg7wfqdnuckqpnukluno.jpg',0,100,11,39,'Phần'),(40,'2023-11-26 10:02:37.920','2023-11-26 10:02:37.920',NULL,'Cơm tấm thập cẩm','Cơm tấm thập cẩm',40000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700992957/go_food_delivery_dev/sbxoi91nesc8bc29x3e5.jpg',0,100,11,40,'Phần'),(41,'2023-11-26 10:03:21.187','2023-11-26 10:03:21.187',NULL,'Bánh mì thịt nướng','Bánh mì thịt nướng',25000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993000/go_food_delivery_dev/sgdnwrxlkqwzua845s7p.jpg',0,50,10,41,'Phần'),(42,'2023-11-26 10:03:49.571','2023-11-26 10:03:49.571',NULL,'Bánh mì ốp la','Bánh mì ốp la',15000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993029/go_food_delivery_dev/cxmtxz3muqd9o4s7wfgn.jpg',0,50,10,42,'Phần'),(43,'2023-11-26 10:04:23.531','2023-11-28 07:58:23.348',NULL,'Bánh mì chả lụa','Bánh mì chả lụa',20000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993063/go_food_delivery_dev/c4nqgzj5jset4hmjrr3h.jpg',12,38,10,43,'Phần'),(44,'2023-11-26 10:05:13.660','2023-11-26 10:05:13.660',NULL,'Bánh mì pate','Bánh mì pate',18000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993113/go_food_delivery_dev/efestluhnqhwbmtt9yqp.jpg',0,50,10,44,'Phần'),(45,'2023-11-26 10:05:36.961','2023-11-28 07:53:59.407',NULL,'Bánh mì chay','Bánh mì chay',15000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993136/go_food_delivery_dev/hjgutdpvyrjz7qgaf3n4.jpg',2,48,10,45,'Phần'),(46,'2023-11-26 10:06:12.082','2023-11-26 10:06:12.082',NULL,'Bún chả cá','Bún chả cá',30000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993171/go_food_delivery_dev/zjis7nuwoe3uzzymonip.jpg',0,100,12,46,'Phần'),(47,'2023-11-26 10:06:32.618','2023-11-26 10:06:32.618',NULL,'Bún bò Huế','Bún bò Huế',35000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993192/go_food_delivery_dev/yrovkfwqjn4h63debln2.jpg',0,100,12,47,'Phần'),(48,'2023-11-26 10:06:50.901','2023-11-26 10:06:50.901',NULL,'Bún đậu mắm tôm','Bún đậu mắm tôm',55000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993210/go_food_delivery_dev/cpaamdfndpf2mwrw5p6k.jpg',0,100,12,48,'Phần'),(49,'2023-11-26 10:07:28.312','2023-11-26 10:07:28.312',NULL,'Bún chả Hà Nội','Bún chả Hà Nội',30000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993247/go_food_delivery_dev/mhhr8ec3f8smwb3uk5q3.jpg',0,100,12,49,'Phần'),(50,'2023-11-26 10:07:59.619','2023-11-26 10:07:59.619',NULL,'Bún mọc','Bún mọc',25000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993279/go_food_delivery_dev/teoyz74axwyuhxfctcfy.jpg',0,100,12,50,'Phần'),(51,'2023-11-26 10:08:50.440','2023-11-26 10:08:50.440',NULL,'Phở bò','Phở bò',40000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993329/go_food_delivery_dev/b1oiu4au9kijeqcigxe1.jpg',0,100,13,51,'Phần'),(52,'2023-11-26 10:09:26.769','2023-11-26 10:09:26.769',NULL,'Phở gà','Phở gà',35000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993366/go_food_delivery_dev/qjtavlvbk1jzex02siq6.jpg',0,100,13,52,'Phần'),(53,'2023-11-26 10:09:59.769','2023-11-26 10:09:59.769',NULL,'Phở chay','Phở chay',30000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993399/go_food_delivery_dev/nyygcypx37lfzawihykq.jpg',0,100,13,53,'Phần'),(54,'2023-11-26 10:10:35.367','2023-11-28 07:53:59.401',NULL,'Phở trộn','Phở trộn',35000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993434/go_food_delivery_dev/rbdhm33wtnbsbxh9kehx.jpg',5,95,13,54,'Phần'),(55,'2023-11-26 10:12:03.259','2023-11-26 10:12:03.259',NULL,'Phở xào','Phở xào',40000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993522/go_food_delivery_dev/w8se9b4p6bn54jlql6xl.jpg',0,150,13,55,'Phần'),(56,'2023-11-26 10:12:39.902','2023-11-26 10:12:39.902',NULL,'Táo','Táo',25000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993559/go_food_delivery_dev/flbqkv6jomv9hnl3g0yh.jpg',0,150,14,56,'Phần'),(57,'2023-11-26 10:13:02.075','2023-11-26 10:13:02.075',NULL,'Cam','Cam',30000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993581/go_food_delivery_dev/vfimoitrmotgcuikonmw.jpg',0,150,14,57,'Phần'),(58,'2023-11-26 10:13:17.653','2023-11-26 10:13:17.653',NULL,'Chuối','Chuối',30000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993597/go_food_delivery_dev/rbffv0di82pp3gwfeidn.jpg',0,150,14,58,'Phần'),(59,'2023-11-26 10:13:39.183','2023-11-26 10:13:39.183',NULL,'Ổi','Ổi',35000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993618/go_food_delivery_dev/vqdt8fzpywr0ddyo9oht.jpg',0,150,14,59,'Phần'),(60,'2023-11-26 10:14:05.239','2023-11-26 10:21:28.665',NULL,'Dâu tây','Dâu tây',40000,'http://res.cloudinary.com/dxn6kjfnd/image/upload/v1700993644/go_food_delivery_dev/kuraniw2yqzteaeq0abq.jpg',5,145,14,60,'Phần');
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
INSERT INTO `roles` VALUES (1,'2023-11-26 08:30:50.372','2023-11-26 08:30:50.372',NULL,'Manager'),(2,'2023-11-26 08:30:57.180','2023-11-26 08:30:57.180',NULL,'Waiter'),(3,'2023-11-26 08:31:02.080','2023-11-26 08:31:02.080',NULL,'Chef');
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
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tables`
--

LOCK TABLES `tables` WRITE;
/*!40000 ALTER TABLE `tables` DISABLE KEYS */;
INSERT INTO `tables` VALUES (1,'2023-11-26 08:43:33.986','2023-11-26 10:19:41.261',NULL,1,0,10,0,2,1),(2,'2023-11-26 08:43:35.076','2023-11-26 10:20:09.360',NULL,1,0,10,0,2,1),(3,'2023-11-26 08:43:35.657','2023-11-26 10:20:32.773',NULL,1,0,10,0,2,1),(4,'2023-11-26 08:45:06.455','2023-11-26 10:21:28.667',NULL,0,0,15,20000,2,1),(5,'2023-11-26 08:45:39.399','2023-11-28 07:53:59.409',NULL,1,0,10,0,3,2),(6,'2023-11-26 08:45:39.995','2023-11-28 07:54:17.614',NULL,1,0,10,0,3,2),(7,'2023-11-26 08:45:40.497','2023-11-28 07:54:38.986',NULL,1,0,10,0,3,2),(8,'2023-11-26 08:45:51.257','2023-11-26 08:45:51.257',NULL,0,1,15,20000,3,2),(9,'2023-11-26 08:46:05.767','2023-11-28 07:56:40.992',NULL,1,0,10,0,4,3),(10,'2023-11-26 08:46:06.245','2023-11-28 07:56:57.104',NULL,1,0,10,0,4,3),(11,'2023-11-26 08:46:06.765','2023-11-28 07:57:12.887',NULL,1,0,10,0,4,3),(12,'2023-11-26 08:46:29.607','2023-11-26 08:46:29.607',NULL,0,1,15,20000,4,3),(13,'2023-11-26 08:46:47.677','2023-11-28 07:58:23.358',NULL,1,0,10,0,5,4),(14,'2023-11-26 08:46:48.276','2023-11-28 07:58:36.135',NULL,1,0,10,0,5,4),(15,'2023-11-26 08:46:48.829','2023-11-28 07:59:10.436',NULL,1,0,10,0,5,4),(16,'2023-11-26 08:47:00.517','2023-11-26 08:47:00.517',NULL,0,1,15,20000,5,4),(17,'2023-11-26 08:47:22.828','2023-11-28 08:05:36.148',NULL,1,0,10,0,6,5),(18,'2023-11-26 08:47:23.518','2023-11-28 08:06:04.055',NULL,1,0,10,0,6,5),(19,'2023-11-26 08:47:24.037','2023-11-26 08:47:24.037',NULL,1,1,10,0,6,5),(20,'2023-11-26 08:47:36.247','2023-11-26 08:47:36.247',NULL,0,1,15,20000,6,5),(21,'2023-11-26 08:47:50.897','2023-11-28 08:07:39.185',NULL,1,0,10,0,7,6),(22,'2023-11-26 08:47:51.317','2023-11-28 08:08:08.008',NULL,1,0,10,0,7,6),(23,'2023-11-26 08:47:51.767','2023-11-26 08:47:51.767',NULL,1,1,10,0,7,6),(24,'2023-11-26 08:48:02.333','2023-11-28 08:08:23.518',NULL,0,0,15,20000,7,6),(25,'2023-11-26 08:48:35.792','2023-11-28 08:11:23.789',NULL,1,0,10,0,8,7),(26,'2023-11-26 08:48:36.498','2023-11-26 08:48:36.498',NULL,1,1,10,0,8,7),(27,'2023-11-26 08:48:37.046','2023-11-28 08:11:53.926',NULL,1,0,10,0,8,7),(28,'2023-11-26 08:48:48.137','2023-11-26 08:48:48.137',NULL,0,1,15,20000,8,7),(29,'2023-11-26 08:49:17.498','2023-11-28 08:16:46.154',NULL,0,0,15,20000,9,8),(30,'2023-11-26 08:49:18.049','2023-11-26 08:49:18.049',NULL,0,1,15,20000,9,8),(31,'2023-11-26 08:49:18.641','2023-11-26 08:49:18.641',NULL,0,1,15,20000,9,8),(32,'2023-11-26 08:49:19.259','2023-11-26 08:49:19.259',NULL,0,1,15,20000,9,8),(33,'2023-12-03 06:38:38.857','2023-12-03 06:38:38.857',NULL,0,1,15,20000,15,9),(34,'2023-12-03 06:38:42.193','2023-12-03 06:38:42.193',NULL,0,1,15,20000,15,9),(35,'2023-12-03 06:38:43.292','2023-12-03 06:38:43.292',NULL,0,1,15,20000,15,9),(36,'2023-12-03 06:38:44.450','2023-12-03 06:38:44.450',NULL,0,1,15,20000,15,9),(37,'2023-12-03 06:38:55.381','2023-12-03 06:38:55.381',NULL,1,1,10,0,15,9),(38,'2023-12-03 06:38:56.298','2023-12-03 06:38:56.298',NULL,1,1,10,0,15,9),(39,'2023-12-03 06:38:56.994','2023-12-03 06:38:56.994',NULL,1,1,10,0,15,9),(40,'2023-12-03 06:38:57.825','2023-12-03 06:38:57.825',NULL,1,1,10,0,15,9),(41,'2023-12-03 06:38:58.516','2023-12-03 06:38:58.516',NULL,1,1,10,0,15,9);
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

-- Dump completed on 2023-12-03 13:40:14

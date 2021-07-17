-- MySQL dump 10.13  Distrib 8.0.23, for osx10.15 (x86_64)
--
-- Host: localhost    Database: lxd_warehouse
-- ------------------------------------------------------
-- Server version	8.0.23

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `tbl_distributors`
--

DROP TABLE IF EXISTS `tbl_distributors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tbl_distributors` (
  `id` int NOT NULL AUTO_INCREMENT,
  `distributor_name` varchar(50) NOT NULL,
  `distributor_address` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_distributors`
--

LOCK TABLES `tbl_distributors` WRITE;
/*!40000 ALTER TABLE `tbl_distributors` DISABLE KEYS */;
INSERT INTO `tbl_distributors` VALUES (1,'Mr. Robot Computer Shop','Vientiane, Laos');
/*!40000 ALTER TABLE `tbl_distributors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_export`
--

DROP TABLE IF EXISTS `tbl_export`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tbl_export` (
  `id` int NOT NULL AUTO_INCREMENT,
  `distributor_id` int NOT NULL,
  `user_id` int NOT NULL,
  `export_time` timestamp NOT NULL,
  `export_status` varchar(50) NOT NULL DEFAULT 'ORDERED',
  PRIMARY KEY (`id`),
  KEY `tbl_export_tbl_distributors_id_fk` (`distributor_id`),
  KEY `tbl_export_tbl_users_id_fk` (`user_id`),
  CONSTRAINT `tbl_export_tbl_distributors_id_fk` FOREIGN KEY (`distributor_id`) REFERENCES `tbl_distributors` (`id`),
  CONSTRAINT `tbl_export_tbl_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `tbl_users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_export`
--

LOCK TABLES `tbl_export` WRITE;
/*!40000 ALTER TABLE `tbl_export` DISABLE KEYS */;
INSERT INTO `tbl_export` VALUES (1,1,1,'2021-02-03 15:40:56','EXPORTED'),(2,1,1,'2021-02-03 15:46:51','EXPORTED');
/*!40000 ALTER TABLE `tbl_export` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_export_detail`
--

DROP TABLE IF EXISTS `tbl_export_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tbl_export_detail` (
  `id` int NOT NULL AUTO_INCREMENT,
  `export_id` int NOT NULL,
  `product_id` int NOT NULL,
  `product_qty` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `tbl_export_detail_tbl_export_id_fk` (`export_id`),
  KEY `tbl_export_detail_tbl_products_id_fk` (`product_id`),
  CONSTRAINT `tbl_export_detail_tbl_export_id_fk` FOREIGN KEY (`export_id`) REFERENCES `tbl_export` (`id`),
  CONSTRAINT `tbl_export_detail_tbl_products_id_fk` FOREIGN KEY (`product_id`) REFERENCES `tbl_products` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_export_detail`
--

LOCK TABLES `tbl_export_detail` WRITE;
/*!40000 ALTER TABLE `tbl_export_detail` DISABLE KEYS */;
INSERT INTO `tbl_export_detail` VALUES (1,1,1,5),(2,2,1,10);
/*!40000 ALTER TABLE `tbl_export_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_import`
--

DROP TABLE IF EXISTS `tbl_import`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tbl_import` (
  `id` int NOT NULL AUTO_INCREMENT,
  `supplier_id` int NOT NULL,
  `user_id` int NOT NULL,
  `import_time` timestamp NOT NULL,
  `import_status` varchar(20) NOT NULL DEFAULT 'ORDERED',
  PRIMARY KEY (`id`),
  KEY `tbl_import_tbl_suppliers_id_fk` (`supplier_id`),
  KEY `tbl_import_tbl_users_id_fk` (`user_id`),
  CONSTRAINT `tbl_import_tbl_suppliers_id_fk` FOREIGN KEY (`supplier_id`) REFERENCES `tbl_suppliers` (`id`),
  CONSTRAINT `tbl_import_tbl_users_id_fk` FOREIGN KEY (`user_id`) REFERENCES `tbl_users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_import`
--

LOCK TABLES `tbl_import` WRITE;
/*!40000 ALTER TABLE `tbl_import` DISABLE KEYS */;
INSERT INTO `tbl_import` VALUES (1,1,1,'2021-02-03 15:32:16','IMPORTED'),(2,1,1,'2021-02-03 15:46:17','IMPORTED');
/*!40000 ALTER TABLE `tbl_import` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_import_detail`
--

DROP TABLE IF EXISTS `tbl_import_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tbl_import_detail` (
  `id` int NOT NULL AUTO_INCREMENT,
  `import_id` int NOT NULL,
  `product_id` int NOT NULL,
  `product_qty` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `tbl_import_detail_tbl_import_id_fk` (`import_id`),
  KEY `tbl_import_detail_tbl_products_id_fk` (`product_id`),
  CONSTRAINT `tbl_import_detail_tbl_import_id_fk` FOREIGN KEY (`import_id`) REFERENCES `tbl_import` (`id`),
  CONSTRAINT `tbl_import_detail_tbl_products_id_fk` FOREIGN KEY (`product_id`) REFERENCES `tbl_products` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_import_detail`
--

LOCK TABLES `tbl_import_detail` WRITE;
/*!40000 ALTER TABLE `tbl_import_detail` DISABLE KEYS */;
INSERT INTO `tbl_import_detail` VALUES (1,1,1,10),(2,2,1,10);
/*!40000 ALTER TABLE `tbl_import_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_product_brand`
--

DROP TABLE IF EXISTS `tbl_product_brand`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tbl_product_brand` (
  `id` int NOT NULL AUTO_INCREMENT,
  `brand_name` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_product_brand`
--

LOCK TABLES `tbl_product_brand` WRITE;
/*!40000 ALTER TABLE `tbl_product_brand` DISABLE KEYS */;
INSERT INTO `tbl_product_brand` VALUES (1,'Apple');
/*!40000 ALTER TABLE `tbl_product_brand` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_product_type`
--

DROP TABLE IF EXISTS `tbl_product_type`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tbl_product_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `type_name` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_product_type`
--

LOCK TABLES `tbl_product_type` WRITE;
/*!40000 ALTER TABLE `tbl_product_type` DISABLE KEYS */;
INSERT INTO `tbl_product_type` VALUES (1,'Laptop Computer');
/*!40000 ALTER TABLE `tbl_product_type` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_products`
--

DROP TABLE IF EXISTS `tbl_products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tbl_products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_name` varchar(50) NOT NULL,
  `product_brand` int NOT NULL,
  `product_type` int NOT NULL,
  `product_description` varchar(255) NOT NULL,
  `product_price` decimal(10,2) NOT NULL,
  `product_qty` int NOT NULL DEFAULT '0',
  `product_image_url` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `tbl_products_tbl_product_brand_id_fk` (`product_brand`),
  KEY `tbl_products_tbl_product_type_id_fk` (`product_type`),
  CONSTRAINT `tbl_products_tbl_product_brand_id_fk` FOREIGN KEY (`product_brand`) REFERENCES `tbl_product_brand` (`id`),
  CONSTRAINT `tbl_products_tbl_product_type_id_fk` FOREIGN KEY (`product_type`) REFERENCES `tbl_product_type` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_products`
--

LOCK TABLES `tbl_products` WRITE;
/*!40000 ALTER TABLE `tbl_products` DISABLE KEYS */;
INSERT INTO `tbl_products` VALUES (1,'Apple MacBook Pro 13-inch',1,1,'Apple laptop powered by Apple Silicon',1299.00,5,'http://localhost:8110/upload/2021-02-01T10:33:30_mbp-spacegray-gallery2-202011_GEO_US.jpeg');
/*!40000 ALTER TABLE `tbl_products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_suppliers`
--

DROP TABLE IF EXISTS `tbl_suppliers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tbl_suppliers` (
  `id` int NOT NULL AUTO_INCREMENT,
  `supplier_name` varchar(50) NOT NULL,
  `supplier_address` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_suppliers`
--

LOCK TABLES `tbl_suppliers` WRITE;
/*!40000 ALTER TABLE `tbl_suppliers` DISABLE KEYS */;
INSERT INTO `tbl_suppliers` VALUES (1,'Apple Southeast Asia','Singapore');
/*!40000 ALTER TABLE `tbl_suppliers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_user_roles`
--

DROP TABLE IF EXISTS `tbl_user_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tbl_user_roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_name` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_user_roles`
--

LOCK TABLES `tbl_user_roles` WRITE;
/*!40000 ALTER TABLE `tbl_user_roles` DISABLE KEYS */;
INSERT INTO `tbl_user_roles` VALUES (1,'Admininstrator'),(2,'Product Manager'),(3,'Import-Export Staff');
/*!40000 ALTER TABLE `tbl_user_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tbl_users`
--

DROP TABLE IF EXISTS `tbl_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tbl_users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `user_firstname` varchar(100) NOT NULL,
  `user_lastname` varchar(100) NOT NULL,
  `user_email` varchar(100) NOT NULL,
  `user_password` varchar(60) NOT NULL,
  `user_role` int NOT NULL,
  `user_image_url` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `tbl_users_tbl_user_roles_id_fk` (`user_role`),
  CONSTRAINT `tbl_users_tbl_user_roles_id_fk` FOREIGN KEY (`user_role`) REFERENCES `tbl_user_roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tbl_users`
--

LOCK TABLES `tbl_users` WRITE;
/*!40000 ALTER TABLE `tbl_users` DISABLE KEYS */;
INSERT INTO `tbl_users` VALUES (1,'artydev','Philaphonh','Inthavongsa','atyinthavongsa@gmail.com','$2a$10$88hjibKIyodoVv5LZCfpQuWLOrYdXKXtvNSm7Nnlwh14vaR5q9fbW',1,'http://localhost:8110/upload/2021-02-01T10:38:08_shipit.jpg'),(2,'meeno','Soukkanya','XYKT','soukkanya@gmail.com','$2a$10$TiUrgHwkYe/cKgEmyw1hOuuY5ELToYgW1XiAI1YNbdwH7ta.NAioW',1,'http://localhost:8110/upload/2021-02-03T22:49:11_screaming-jake.jpg');
/*!40000 ALTER TABLE `tbl_users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-03-02 14:32:28

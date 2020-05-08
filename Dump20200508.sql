CREATE DATABASE  IF NOT EXISTS `gelibert` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `gelibert`;
-- MySQL dump 10.13  Distrib 8.0.20, for macos10.15 (x86_64)
--
-- Host: localhost    Database: gelibert
-- ------------------------------------------------------
-- Server version	8.0.18

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
-- Table structure for table `clients`
--

DROP TABLE IF EXISTS `clients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `clients` (
  `id` int(11) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `tel` varchar(45) DEFAULT NULL,
  `address` varchar(90) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `clients`
--

LOCK TABLES `clients` WRITE;
/*!40000 ALTER TABLE `clients` DISABLE KEYS */;
INSERT INTO `clients` VALUES (0,'Name_0','Tel_0','Address_0'),(1,'Name_1','Tel_1','Address_1'),(2,'Name_2','Tel_2','Address_2'),(3,'Name_3','Tel_3','Address_3'),(4,'Name_4','Tel_4','Address_4'),(5,'Name_5','Tel_5','Address_5'),(6,'Name_6','Tel_6','Address_6'),(7,'Name_7','Tel_7','Address_7'),(8,'Name_8','Tel_8','Address_8'),(9,'Name_9','Tel_9','Address_9'),(10,'Name_10','Tel_10','Address_10'),(11,'Name_11','Tel_11','Address_11'),(12,'Name_12','Tel_12','Address_12'),(13,'Name_13','Tel_13','Address_13'),(14,'Name_14','Tel_14','Address_14'),(15,'Name_15','Tel_15','Address_15'),(16,'Name_16','Tel_16','Address_16'),(17,'Name_17','Tel_17','Address_17'),(18,'Name_18','Tel_18','Address_18'),(19,'Name_19','Tel_19','Address_19');
/*!40000 ALTER TABLE `clients` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `consists_from`
--

DROP TABLE IF EXISTS `consists_from`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `consists_from` (
  `id` int(11) DEFAULT NULL,
  `product` varchar(45) DEFAULT NULL,
  `quantity` float DEFAULT NULL,
  `price` float DEFAULT NULL,
  `ext_info` varchar(100) NOT NULL DEFAULT ' '
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `consists_from`
--

LOCK TABLES `consists_from` WRITE;
/*!40000 ALTER TABLE `consists_from` DISABLE KEYS */;
INSERT INTO `consists_from` VALUES (0,'ProductFrom_0/0',1,0.2,' '),(0,'ProductFrom_0/1',2,0.4,' '),(0,'ProductFrom_0/2',3,0.6,' '),(0,'ProductFrom_0/3',4,0.8,' '),(1,'ProductFrom_1/0',1,0.2,' '),(1,'ProductFrom_1/1',2,0.4,' '),(1,'ProductFrom_1/2',3,0.6,' '),(1,'ProductFrom_1/3',4,0.8,' '),(2,'ProductFrom_2/0',1,0.2,' '),(2,'ProductFrom_2/1',2,0.4,' '),(2,'ProductFrom_2/2',3,0.6,' '),(2,'ProductFrom_2/3',4,0.8,' '),(3,'ProductFrom_3/0',1,0.2,' '),(3,'ProductFrom_3/1',2,0.4,' '),(3,'ProductFrom_3/2',3,0.6,' '),(3,'ProductFrom_3/3',4,0.8,' '),(4,'ProductFrom_4/0',1,0.2,' '),(4,'ProductFrom_4/1',2,0.4,' '),(4,'ProductFrom_4/2',3,0.6,' '),(4,'ProductFrom_4/3',4,0.8,' '),(5,'ProductFrom_5/0',1,0.2,' '),(5,'ProductFrom_5/1',2,0.4,' '),(5,'ProductFrom_5/2',3,0.6,' '),(5,'ProductFrom_5/3',4,0.8,' '),(6,'ProductFrom_6/0',1,0.2,' '),(6,'ProductFrom_6/1',2,0.4,' '),(6,'ProductFrom_6/2',3,0.6,' '),(6,'ProductFrom_6/3',4,0.8,' '),(7,'ProductFrom_7/0',1,0.2,' '),(7,'ProductFrom_7/1',2,0.4,' '),(7,'ProductFrom_7/2',3,0.6,' '),(7,'ProductFrom_7/3',4,0.8,' '),(8,'ProductFrom_8/0',1,0.2,' '),(8,'ProductFrom_8/1',2,0.4,' '),(8,'ProductFrom_8/2',3,0.6,' '),(8,'ProductFrom_8/3',4,0.8,' '),(9,'ProductFrom_9/0',1,0.2,' '),(9,'ProductFrom_9/1',2,0.4,' '),(9,'ProductFrom_9/2',3,0.6,' '),(9,'ProductFrom_9/3',4,0.8,' '),(10,'ProductFrom_10/0',1,0.2,' '),(10,'ProductFrom_10/1',2,0.4,' '),(10,'ProductFrom_10/2',3,0.6,' '),(10,'ProductFrom_10/3',4,0.8,' '),(11,'ProductFrom_11/0',1,0.2,' '),(11,'ProductFrom_11/1',2,0.4,' '),(11,'ProductFrom_11/2',3,0.6,' '),(11,'ProductFrom_11/3',4,0.8,' '),(12,'ProductFrom_12/0',1,0.2,' '),(12,'ProductFrom_12/1',2,0.4,' '),(12,'ProductFrom_12/2',3,0.6,' '),(12,'ProductFrom_12/3',4,0.8,' '),(13,'ProductFrom_13/0',1,0.2,' '),(13,'ProductFrom_13/1',2,0.4,' '),(13,'ProductFrom_13/2',3,0.6,' '),(13,'ProductFrom_13/3',4,0.8,' '),(14,'ProductFrom_14/0',1,0.2,' '),(14,'ProductFrom_14/1',2,0.4,' '),(14,'ProductFrom_14/2',3,0.6,' '),(14,'ProductFrom_14/3',4,0.8,' '),(15,'ProductFrom_15/0',1,0.2,' '),(15,'ProductFrom_15/1',2,0.4,' '),(15,'ProductFrom_15/2',3,0.6,' '),(15,'ProductFrom_15/3',4,0.8,' '),(16,'ProductFrom_16/0',1,0.2,' '),(16,'ProductFrom_16/1',2,0.4,' '),(16,'ProductFrom_16/2',3,0.6,' '),(16,'ProductFrom_16/3',4,0.8,' '),(17,'ProductFrom_17/0',1,0.2,' '),(17,'ProductFrom_17/1',2,0.4,' '),(17,'ProductFrom_17/2',3,0.6,' '),(17,'ProductFrom_17/3',4,0.8,' '),(18,'ProductFrom_18/0',1,0.2,' '),(18,'ProductFrom_18/1',2,0.4,' '),(18,'ProductFrom_18/2',3,0.6,' '),(18,'ProductFrom_18/3',4,0.8,' '),(19,'ProductFrom_19/0',1,0.2,' '),(19,'ProductFrom_19/1',2,0.4,' '),(19,'ProductFrom_19/2',3,0.6,' '),(19,'ProductFrom_19/3',4,0.8,' ');
/*!40000 ALTER TABLE `consists_from` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `consists_to`
--

DROP TABLE IF EXISTS `consists_to`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `consists_to` (
  `id` int(11) DEFAULT NULL,
  `product` varchar(45) DEFAULT NULL,
  `quantity` float DEFAULT NULL,
  `price` float DEFAULT NULL,
  `ext_info` varchar(100) NOT NULL DEFAULT ' '
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `consists_to`
--

LOCK TABLES `consists_to` WRITE;
/*!40000 ALTER TABLE `consists_to` DISABLE KEYS */;
INSERT INTO `consists_to` VALUES (0,'ProductTo_0/0',1,0.3,' '),(0,'ProductTo_0/1',2,0.6,' '),(0,'ProductTo_0/2',3,0.9,' '),(0,'ProductTo_0/3',4,1.2,' '),(1,'ProductTo_1/0',1,0.3,'Test'),(1,'ProductTo_1/1',2,0.6,'Test'),(1,'ProductTo_1/2',3,0.9,'Test'),(1,'ProductTo_1/3',4,1.2,'Test'),(2,'ProductTo_2/0',1,0.3,' '),(2,'ProductTo_2/1',2,0.6,' '),(2,'ProductTo_2/2',3,0.9,' '),(2,'ProductTo_2/3',4,1.2,' '),(3,'ProductTo_3/0',1,0.3,' '),(3,'ProductTo_3/1',2,0.6,' '),(3,'ProductTo_3/2',3,0.9,' '),(3,'ProductTo_3/3',4,1.2,' '),(4,'ProductTo_4/0',1,0.3,' '),(4,'ProductTo_4/1',2,0.6,' '),(4,'ProductTo_4/2',3,0.9,' '),(4,'ProductTo_4/3',4,1.2,' '),(5,'ProductTo_5/0',1,0.3,' '),(5,'ProductTo_5/1',2,0.6,' '),(5,'ProductTo_5/2',3,0.9,' '),(5,'ProductTo_5/3',4,1.2,' '),(6,'ProductTo_6/0',1,0.3,' '),(6,'ProductTo_6/1',2,0.6,' '),(6,'ProductTo_6/2',3,0.9,' '),(6,'ProductTo_6/3',4,1.2,' '),(7,'ProductTo_7/0',1,0.3,' '),(7,'ProductTo_7/1',2,0.6,' '),(7,'ProductTo_7/2',3,0.9,' '),(7,'ProductTo_7/3',4,1.2,' '),(8,'ProductTo_8/0',1,0.3,' '),(8,'ProductTo_8/1',2,0.6,' '),(8,'ProductTo_8/2',3,0.9,' '),(8,'ProductTo_8/3',4,1.2,' '),(9,'ProductTo_9/0',1,0.3,' '),(9,'ProductTo_9/1',2,0.6,' '),(9,'ProductTo_9/2',3,0.9,' '),(9,'ProductTo_9/3',4,1.2,' '),(10,'ProductTo_10/0',1,0.3,' '),(10,'ProductTo_10/1',2,0.6,' '),(10,'ProductTo_10/2',3,0.9,' '),(10,'ProductTo_10/3',4,1.2,' '),(11,'ProductTo_11/0',1,0.3,' '),(11,'ProductTo_11/1',2,0.6,' '),(11,'ProductTo_11/2',3,0.9,' '),(11,'ProductTo_11/3',4,1.2,' '),(12,'ProductTo_12/0',1,0.3,' '),(12,'ProductTo_12/1',2,0.6,' '),(12,'ProductTo_12/2',3,0.9,' '),(12,'ProductTo_12/3',4,1.2,' '),(13,'ProductTo_13/0',1,0.3,' '),(13,'ProductTo_13/1',2,0.6,' '),(13,'ProductTo_13/2',3,0.9,' '),(13,'ProductTo_13/3',4,1.2,' '),(14,'ProductTo_14/0',1,0.3,' '),(14,'ProductTo_14/1',2,0.6,' '),(14,'ProductTo_14/2',3,0.9,' '),(14,'ProductTo_14/3',4,1.2,' '),(15,'ProductTo_15/0',1,0.3,' '),(15,'ProductTo_15/1',2,0.6,' '),(15,'ProductTo_15/2',3,0.9,' '),(15,'ProductTo_15/3',4,1.2,' '),(16,'ProductTo_16/0',1,0.3,' '),(16,'ProductTo_16/1',2,0.6,' '),(16,'ProductTo_16/2',3,0.9,' '),(16,'ProductTo_16/3',4,1.2,' '),(17,'ProductTo_17/0',1,0.3,' '),(17,'ProductTo_17/1',2,0.6,' '),(17,'ProductTo_17/2',3,0.9,' '),(17,'ProductTo_17/3',4,1.2,' '),(18,'ProductTo_18/0',1,0.3,' '),(18,'ProductTo_18/1',2,0.6,' '),(18,'ProductTo_18/2',3,0.9,' '),(18,'ProductTo_18/3',4,1.2,' '),(19,'ProductTo_19/0',1,0.3,' '),(19,'ProductTo_19/1',2,0.6,' '),(19,'ProductTo_19/2',3,0.9,' '),(19,'ProductTo_19/3',4,1.2,' ');
/*!40000 ALTER TABLE `consists_to` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `couriers`
--

DROP TABLE IF EXISTS `couriers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `couriers` (
  `id` int(11) NOT NULL,
  `imei` bigint(11) NOT NULL,
  `tel` varchar(45) DEFAULT NULL,
  `name` varchar(45) DEFAULT NULL,
  `car_number` varchar(45) DEFAULT NULL,
  `timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `imei` (`imei`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `couriers`
--

LOCK TABLES `couriers` WRITE;
/*!40000 ALTER TABLE `couriers` DISABLE KEYS */;
INSERT INTO `couriers` VALUES (0,123456789012300,'Tel_0','Name_0','CarNumber_0','2020-03-04 14:39:42'),(1,352167058641169,'Tel_1','Name_1','CarNumber_1','2020-03-04 14:40:14'),(2,123456789012302,'Tel_2','Name_2','CarNumber_2','2020-03-04 14:39:42'),(3,123456789012303,'Tel_3','Name_3','CarNumber_3','2020-03-04 14:39:42'),(4,123456789012304,'Tel_4','Name_4','CarNumber_4','2020-03-04 14:39:42'),(5,123456789012305,'Tel_5','Name_5','CarNumber_5','2020-03-04 14:53:45'),(6,123456789012306,'Tel_6','Name_6','CarNumber_6','2020-03-04 14:39:42'),(7,123456789012307,'Tel_7','Name_7','CarNumber_7','2020-03-04 14:39:42'),(8,123456789012308,'Tel_8','Name_8','CarNumber_8','2020-03-04 14:39:42'),(9,123456789012309,'Tel_9','Name_9','CarNumber_9','2020-03-04 14:53:45'),(10,123456789012310,'Tel_10','Name_10','CarNumber_10','2020-03-04 14:39:42'),(11,123456789012311,'Tel_11','Name_11','CarNumber_11','2020-03-04 14:39:42'),(12,123456789012312,'Tel_12','Name_12','CarNumber_12','2020-03-05 07:46:11'),(13,123456789012313,'Tel_13','Name_13','CarNumber_13','2020-03-04 14:39:42'),(14,123456789012314,'Tel_14','Name_14','CarNumber_14','2020-03-04 14:39:42'),(15,123456789012315,'Tel_15','Name_15','CarNumber_15','2020-03-04 14:39:42'),(16,123456789012316,'Tel_16','Name_16','CarNumber_16','2020-03-04 14:39:42'),(17,123456789012317,'Tel_17','Name_17','CarNumber_17','2020-03-04 14:39:42'),(18,123456789012318,'Tel_18','Name_18','CarNumber_18','2020-03-04 14:39:42'),(19,123456789012319,'Tel_19','Name_19','CarNumber_19','2020-03-05 07:47:37');
/*!40000 ALTER TABLE `couriers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `geodata`
--

DROP TABLE IF EXISTS `geodata`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `geodata` (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `address` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `imei` bigint(11) NOT NULL,
  `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `courier_id` int(255) DEFAULT NULL,
  `longitude` float(12,0) NOT NULL DEFAULT '0',
  `latitude` float(12,0) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `geodata`
--

LOCK TABLES `geodata` WRITE;
/*!40000 ALTER TABLE `geodata` DISABLE KEYS */;
INSERT INTO `geodata` VALUES (1,'',0,'2020-04-30 07:52:13',0,0,0);
/*!40000 ALTER TABLE `geodata` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id` int(11) NOT NULL,
  `courier_id` int(11) DEFAULT NULL,
  `client_id` int(11) DEFAULT NULL,
  `payment_method` varchar(45) DEFAULT NULL,
  `order_cost` float DEFAULT NULL,
  `delivered` int(11) DEFAULT NULL,
  `delivery_delay` int(11) DEFAULT NULL,
  `date_start` datetime DEFAULT NULL,
  `date_finish` datetime DEFAULT '1000-01-01 00:00:00',
  `timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (0,0,0,'Cash',5,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(1,1,1,'Cash',15,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(2,2,2,'Cash',25,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(3,3,3,'Cash',35,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(4,4,4,'Cash',45,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(5,5,5,'Cash',55,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(6,1,6,'Cash',65,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:54:16'),(7,7,7,'Cash',75,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(8,8,8,'Cash',85,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(9,9,9,'Cash',95,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(10,1,10,'Cash',105,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:54:16'),(11,11,11,'Cash',115,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(12,12,12,'Cash',125,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(13,13,13,'Cash',135,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(14,14,14,'Cash',145,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(15,15,15,'Cash',155,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(16,16,16,'Cash',165,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(17,17,17,'Cash',175,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(18,18,18,'Cash',185,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42'),(19,19,19,'Cash',195,0,0,'2020-03-04 16:39:42','1000-01-01 00:00:00','2020-03-04 14:39:42');
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-05-08  9:54:13

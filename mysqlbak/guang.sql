-- MySQL dump 10.13  Distrib 5.7.19, for osx10.12 (x86_64)
--
-- Host: localhost    Database: guang
-- ------------------------------------------------------
-- Server version	5.7.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `moban`
--

DROP TABLE IF EXISTS `moban`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `moban` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `question_list` varchar(100) DEFAULT NULL,
  `place_list` varchar(100) DEFAULT NULL,
  `questiontype_list` varchar(10) DEFAULT NULL,
  `count_list` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `moban`
--

LOCK TABLES `moban` WRITE;
/*!40000 ALTER TABLE `moban` DISABLE KEYS */;
/*!40000 ALTER TABLE `moban` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `txlyd`
--

DROP TABLE IF EXISTS `txlyd`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `txlyd` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `starttime` varchar(100) DEFAULT NULL,
  `endtime` varchar(100) DEFAULT NULL,
  `question` varchar(100) DEFAULT NULL,
  `place` varchar(4) DEFAULT NULL,
  `questiontype` varchar(10) DEFAULT NULL,
  `count` varchar(3) DEFAULT NULL,
  `solvetime` varchar(100) DEFAULT NULL,
  `note` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=73 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `txlyd`
--

LOCK TABLES `txlyd` WRITE;
/*!40000 ALTER TABLE `txlyd` DISABLE KEYS */;
/*!40000 ALTER TABLE `txlyd` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `zhuce`
--

DROP TABLE IF EXISTS `zhuce`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `zhuce` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET latin1 DEFAULT NULL,
  `passwd` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `zhuce`
--

LOCK TABLES `zhuce` WRITE;
/*!40000 ALTER TABLE `zhuce` DISABLE KEYS */;
/*!40000 ALTER TABLE `zhuce` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-01-22 18:06:57

-- MySQL dump 10.13  Distrib 8.0.13, for macos10.14 (x86_64)
--
-- Host: localhost    Database: mylearning
-- ------------------------------------------------------
-- Server version	8.0.13

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `GLOBAL_INSIGHT_DATA`
--

DROP TABLE IF EXISTS `GLOBAL_INSIGHT_DATA`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `GLOBAL_INSIGHT_DATA` (
  `INSIGHT_KEYS` varchar(255) DEFAULT NULL,
  `INSIGHT_VALUES` varchar(255) DEFAULT NULL,
  `ID` int(100) NOT NULL AUTO_INCREMENT,
  `Country` varchar(255) DEFAULT NULL,
  `Environment` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `GLOBAL_INSIGHT_DATA`
--

LOCK TABLES `GLOBAL_INSIGHT_DATA` WRITE;
/*!40000 ALTER TABLE `GLOBAL_INSIGHT_DATA` DISABLE KEYS */;
/*!40000 ALTER TABLE `GLOBAL_INSIGHT_DATA` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `PROJECTS`
--

DROP TABLE IF EXISTS `PROJECTS`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `PROJECTS` (
  `Project_ID` int(255) NOT NULL AUTO_INCREMENT,
  `Project_name` varchar(255) DEFAULT NULL,
  `Project_desc` varchar(255) DEFAULT NULL,
  `Created_Date` datetime DEFAULT NULL,
  `Modified_Date` datetime DEFAULT NULL,
  `Project_status` enum('Active','Inactive') DEFAULT NULL,
  PRIMARY KEY (`Project_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PROJECTS`
--

LOCK TABLES `PROJECTS` WRITE;
/*!40000 ALTER TABLE `PROJECTS` DISABLE KEYS */;

/*!40000 ALTER TABLE `PROJECTS` ENABLE KEYS */;
UNLOCK TABLES;
--
-- WARNING: old server version. The following dump may be incomplete.
--
DELIMITER ;;
/*!50003 SET SESSION SQL_MODE="ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION" */;;
/*!50003 CREATE */ /*!50017 DEFINER=`root`@`localhost` */ /*!50003 TRIGGER `Insert_Project` BEFORE INSERT ON `PROJECTS` FOR EACH ROW begin
	set new.Created_Date=now(), new.Modified_Date=now();
end */;;
DELIMITER ;
--
-- WARNING: old server version. The following dump may be incomplete.
--
DELIMITER ;;
/*!50003 SET SESSION SQL_MODE="ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION" */;;
/*!50003 CREATE */ /*!50017 DEFINER=`root`@`localhost` */ /*!50003 TRIGGER `Update_Project` BEFORE UPDATE ON `PROJECTS` FOR EACH ROW begin
	set new.Modified_Date=now();
end */;;
DELIMITER ;

--
-- Table structure for table `TC_Design`
--

DROP TABLE IF EXISTS `TC_Design`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `TC_Design` (
  `TC_ID` int(11) NOT NULL,
  `STEP_ID` int(11) NOT NULL,
  `TEST_STEP_DESC` varchar(255) DEFAULT NULL,
  `KEYWORD` varchar(255) DEFAULT NULL,
  `LOCATOR_TYPE` varchar(255) DEFAULT NULL,
  `TARGET` varchar(255) DEFAULT NULL,
  `VALUE_S` varchar(255) DEFAULT NULL,
  `COMMENTS` varchar(255) DEFAULT NULL,
  `TC_NAME` varchar(255) DEFAULT NULL,
  `COUNTRY` varchar(255) DEFAULT NULL,
  `ENVIRONMENT` varchar(255) DEFAULT NULL,
  `CREATED_DATE` datetime DEFAULT NULL,
  `MODIFIED_DATE` datetime DEFAULT NULL,
  `STEP_STATUS` enum('Active','Inactive','InProgress','Depreciated') DEFAULT NULL,
  KEY `fk_tcid` (`TC_ID`),
  KEY `fk_tcname` (`TC_NAME`),
  CONSTRAINT `fk_tcid` FOREIGN KEY (`TC_ID`) REFERENCES `test_cases` (`tc_id`),
  CONSTRAINT `fk_tcname` FOREIGN KEY (`TC_NAME`) REFERENCES `test_cases` (`tc_name`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TC_Design`
--

LOCK TABLES `TC_Design` WRITE;
/*!40000 ALTER TABLE `TC_Design` DISABLE KEYS */;

/*!40000 ALTER TABLE `TC_Design` ENABLE KEYS */;
UNLOCK TABLES;
--
-- WARNING: old server version. The following dump may be incomplete.
--
DELIMITER ;;
/*!50003 SET SESSION SQL_MODE="ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION" */;;
/*!50003 CREATE */ /*!50017 DEFINER=`root`@`localhost` */ /*!50003 TRIGGER `TC_Design_On_Insert` BEFORE INSERT ON `TC_Design` FOR EACH ROW begin
	set new.CREATED_DATE = now(), new.MODIFIED_DATE= now();
end */;;
DELIMITER ;
--
-- WARNING: old server version. The following dump may be incomplete.
--
DELIMITER ;;
/*!50003 SET SESSION SQL_MODE="ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION" */;;
/*!50003 CREATE */ /*!50017 DEFINER=`root`@`localhost` */ /*!50003 TRIGGER `TC_Design_On_Update` BEFORE UPDATE ON `TC_Design` FOR EACH ROW begin
	set new.MODIFIED_DATE= now();
end */;;
DELIMITER ;

--
-- Table structure for table `test_cases`
--

DROP TABLE IF EXISTS `test_cases`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `test_cases` (
  `TC_ID` int(11) NOT NULL AUTO_INCREMENT,
  `TC_NAME` varchar(255) NOT NULL,
  `TC_DESC` varchar(255) DEFAULT NULL,
  `CREATED_DATE` datetime DEFAULT NULL,
  `MODIFIED_DATE` datetime DEFAULT NULL,
  `TC_STATUS` enum('Active','Inactive','InProgress','Depreciated') DEFAULT NULL,
  `Country` varchar(255) DEFAULT NULL,
  `Environment` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`TC_NAME`),
  KEY `TC_ID` (`TC_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_cases`
--

LOCK TABLES `test_cases` WRITE;
/*!40000 ALTER TABLE `test_cases` DISABLE KEYS */;

/*!40000 ALTER TABLE `test_cases` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 TRIGGER `Update_TCDate_On_Insert` BEFORE INSERT ON `test_cases` FOR EACH ROW Begin
		set new.CREATED_DATE=now(), new.MODIFIED_DATE=now();
	end */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_0900_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 TRIGGER `Update_TCDate_On_Update` BEFORE UPDATE ON `test_cases` FOR EACH ROW Begin
		set new.MODIFIED_DATE=now();
	end */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `USERS`
--

DROP TABLE IF EXISTS `USERS`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `USERS` (
  `ID` int(255) NOT NULL AUTO_INCREMENT,
  `USER_ID` varchar(255) DEFAULT NULL,
  `Username` varchar(255) DEFAULT NULL,
  `User_Password` varchar(255) DEFAULT NULL,
  `User_Status` enum('Active','Inactive') DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `USERS`
--

LOCK TABLES `USERS` WRITE;
/*!40000 ALTER TABLE `USERS` DISABLE KEYS */;
INSERT INTO `USERS` VALUES (1,'admin123','admin','admin','Active'),(2,'SnehaWilson','Sneha','Wilson','Inactive');
/*!40000 ALTER TABLE `USERS` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-03-15  6:09:40

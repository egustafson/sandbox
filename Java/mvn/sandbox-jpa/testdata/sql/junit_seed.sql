-- MySQL dump 10.11
--
-- Host: localhost    Database: junit_sandbox_jpa
-- ------------------------------------------------------
-- Server version	5.0.45-community-nt

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
-- Dumping data for table `bags`
--


--
-- Dumping data for table `boxes`
--

INSERT INTO `boxes` VALUES (1,'First Box');
INSERT INTO `boxes` VALUES (2,'Second Box');
INSERT INTO `boxes` VALUES (3,'Third Box');

--
-- Dumping data for table `extras`
--

INSERT INTO `extras` VALUES (1,1,'Box','Box 1 Extra 1','bogus');
INSERT INTO `extras` VALUES (2,1,'Box','Box 1 Extra 2','bogus too');
INSERT INTO `extras` VALUES (3,2,'Box','Box 2 Extra 1','bogus');
INSERT INTO `extras` VALUES (4,3,'Box','Box 3 Extra 1','bogus');
INSERT INTO `extras` VALUES (5,3,'Box','Box 3 Extra 2','bogus too');
INSERT INTO `extras` VALUES (6,1,'Bag','Bag 1 Extra 1','this is a bag value');
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2008-03-05 16:26:00

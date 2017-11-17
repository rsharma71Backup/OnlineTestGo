CREATE DATABASE  IF NOT EXISTS `onlinetestdb` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `onlinetestdb`;
-- MySQL dump 10.13  Distrib 5.7.12, for Win64 (x86_64)
--
-- Host: localhost    Database: onlinetestdb
-- ------------------------------------------------------
-- Server version	5.7.18-0ubuntu0.16.04.1

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
-- Table structure for table `Options`
--

DROP TABLE IF EXISTS `Options`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Options` (
  `id` bigint(10) NOT NULL AUTO_INCREMENT,
  `qid` bigint(10) DEFAULT NULL,
  `choices` text,
  `answers` text,
  PRIMARY KEY (`id`),
  KEY `qid_idx` (`qid`),
  CONSTRAINT `qid` FOREIGN KEY (`qid`) REFERENCES `questions` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=173 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Options`
--

LOCK TABLES `Options` WRITE;
/*!40000 ALTER TABLE `Options` DISABLE KEYS */;
INSERT INTO `Options` VALUES (1,1,'monitor','0'),(2,1,'printer','0'),(3,1,'both ','both'),(4,1,'none','0'),(5,2,'c','c'),(6,2,'objective c','0'),(7,2,'python','0'),(8,2,'java','0'),(9,3,'method overloading','0'),(10,3,'method overriding','method overriding'),(11,3,'abstract class','0'),(12,3,'interface','0'),(13,4,'hiding unnecessary information to outside world','0'),(14,4,'wraping up properties and method','wraping up properties and methods'),(15,4,'a way of delivering a class from another class ','0'),(16,4,'none of the above','0'),(17,5,'a method that loops through a list','0'),(18,5,'a method that calls itself','a method that calls itself'),(19,5,'a method that iterates over a block of code until a condition is solved','0'),(20,5,'a method that contains multiple defination','0'),(21,6,'linked list','0'),(22,6,'doubly linked list','0'),(23,6,'queue','0'),(24,6,'stack','stack'),(25,7,'performance testing','0'),(26,7,'unit testing','0'),(27,7,'big O notation','big O notation'),(28,7,'SLOC','0'),(29,8,'shows password of the user','0'),(30,8,'shows path of current directory','shows path of current directory'),(31,8,'shows the list of files in the directory','0'),(32,8,'shows the access defination for a diractory','0'),(33,9,'a process that is dead','0'),(34,9,'a process that is eating memory','0'),(35,9,'a long running background process','a long running background process'),(36,9,'a process that automatically starts up when the machine is started','0'),(37,10,'has to implement all method defination from the interface','has to implement all method defination from the interface'),(38,10,'has to implement at least one method from the interface','0'),(39,10,'has to implement none of the methods from interface','0'),(40,10,'has to overload the method from the interface','0'),(41,11,'abstract class','0'),(42,11,'parent class','0'),(43,11,'Final class','Final class'),(44,11,'None of above','0'),(45,12,'Yes','0'),(46,12,'No','No'),(47,12,'maybe','0'),(48,12,'w.q','0'),(49,13,'it is syntax','0'),(50,13,'Can store multiple values','Can store multiple values'),(51,13,'Both of above','0'),(52,13,'none of above','0'),(53,14,'start() method','0'),(54,14,'Suspend() method','0'),(55,14,'resume() method','resume() method'),(56,14,'yield() method','0'),(57,15,'java.util.HashSet','0'),(58,15,'java.util.LinkedList','0'),(59,15,'java.util.TreeMap','java.util.TreeMap'),(60,15,'java.util.SortedSet','0'),(61,16,'executeQuery()','0'),(62,16,'executeUpdate()','0'),(63,16,'getConnection()','getConnection()'),(64,16,'prepareCall()','0'),(65,17,'Multithreaded','Multithreaded'),(66,17,'Singlethreaded','0'),(67,17,'Both of above','0'),(68,17,'none of above','0'),(69,18,'java.util.Time','0'),(70,18,'java.sql.Time','0'),(71,18,'java.util.Date','java.util.Date'),(72,18,'None of above','0'),(73,19,'boxing','0'),(74,19,'wrapping','0'),(75,19,'instantiation','0'),(76,19,'autoboxing','autoboxing'),(77,20,'Destructors','0'),(78,20,'Object','0'),(79,20,'Variable','0'),(80,20,'Constructor','Constructor'),(81,21,'int','0'),(82,21,'char','0'),(83,21,'string','string'),(84,21,'double','0'),(85,22,'Queue','Queue'),(86,22,'Array','0'),(87,22,'Stack','0'),(88,22,'Maps','0'),(89,23,'System class','0'),(90,23,'InputStream class','0'),(91,23,'Object class','0'),(92,23,'PrintStream class','PrintStream class'),(93,24,'strictfp','0'),(94,24,'native','0'),(95,24,'transient','transient'),(96,24,'volatile','0'),(97,25,'Class','0'),(98,25,'Method','0'),(99,25,'Variable','0'),(100,25,'Interface','Interface'),(101,26,'Compilation error','0'),(102,26,'100','100'),(103,26,'101','0'),(104,26,'throws exception at runtime','throws exception at runtime'),(105,27,'-128 to 127','0'),(106,27,'-(215)�to (215) � 1','0'),(107,27,'0 to 32767','0'),(108,27,'0 to 65535','0 to 65535'),(109,28,'Serialization','Serialization'),(110,28,'File Filtering','0'),(111,28,'Garbage collection','0'),(112,28,'All of the mentioned','0'),(113,29,'clear()','0'),(114,29,'fflush()','0'),(115,29,'flush()','flush()'),(116,29,'close()','0'),(117,30,'max(Collection c)','0'),(118,30,'max(Collection c, Comparator comp)','0'),(119,30,'max(Comparator comp)','max(Comparator comp)'),(120,30,'max(List c)','0'),(121,31,'rand()','0'),(122,31,'randomize()','0'),(123,31,'shuffle()','shuffle()'),(124,31,'ambigous()','0'),(125,32,'EMPTY_SET','0'),(126,32,'EMPTY_LIST','0'),(127,32,'EMPTY_MAP','0'),(128,32,'All of the mentioned','All of the mentioned'),(129,33,'switch, if','switch, if'),(130,33,'if, switch','0'),(131,33,'if, break','0'),(132,33,'continue, if','0'),(133,34,'for(; ;)','0'),(134,34,'for(i=0 ; i<1; i--)','0'),(135,34,'for(i=0; ; i++)','0'),(136,34,'All of the above','All of the above'),(137,35,'super','0'),(138,35,'static','0'),(139,35,'final','final'),(140,35,'private','0'),(141,36,'flattered','0'),(142,36,'tormented','tormented'),(143,36,'made angry','0'),(144,36,'none of the above','0'),(145,37,'very busy','very busy'),(146,37,'not busy at all','0'),(147,37,'irritated','0'),(148,37,'all of the above','0'),(149,38,'written contract','written contract'),(150,38,'verbal promise','0'),(151,38,'dental plate','0'),(152,38,'partial invasion','0'),(153,39,'synopsis,reservations','0'),(154,39,'gist,ambiguity','gist,ambiguity'),(155,39,'magnitude,distortion','0'),(156,39,'totality,hedging','0'),(157,40,'diamond','0'),(158,40,'gold','gold'),(159,40,'nucleus','0'),(160,40,'atoms','0'),(161,41,'domesticity','0'),(162,41,'knowledge','0'),(163,41,'tractability','tractabilitytractability'),(164,41,'eulogy','0'),(165,42,'infallible,mythological','infallible,mythological'),(166,42,'dependable,ordinary','0'),(167,42,'reliable,inventive','0'),(168,42,'trustworthy,credulous','0'),(169,43,'have been agreed for','0'),(170,43,'has agreed to','has agreed to'),(171,43,'are agreeable to','0'),(172,43,'have agreed to','0');
/*!40000 ALTER TABLE `Options` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `admin`
--

DROP TABLE IF EXISTS `admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admin` (
  `id` bigint(10) NOT NULL,
  `userid` bigint(10) DEFAULT NULL,
  `fname` char(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `lname` char(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `type` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `score` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `userid_fk_idx` (`userid`),
  CONSTRAINT `userid_fk` FOREIGN KEY (`userid`) REFERENCES `registration` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admin`
--

LOCK TABLES `admin` WRITE;
/*!40000 ALTER TABLE `admin` DISABLE KEYS */;
INSERT INTO `admin` VALUES (1,1,'sweta','vahia','java',15),(2,1,'sweta','vahia','fundamentals',14),(3,1,'sweta','vahia','language',18);
/*!40000 ALTER TABLE `admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `answers`
--

DROP TABLE IF EXISTS `answers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `answers` (
  `id` bigint(10) NOT NULL AUTO_INCREMENT,
  `uid` bigint(10) DEFAULT NULL,
  `testtype` varchar(20) DEFAULT NULL,
  `score` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid_idx` (`uid`),
  KEY `uid_fk_idx` (`uid`),
  CONSTRAINT `uid_fk` FOREIGN KEY (`uid`) REFERENCES `registration` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `answers`
--

LOCK TABLES `answers` WRITE;
/*!40000 ALTER TABLE `answers` DISABLE KEYS */;
INSERT INTO `answers` VALUES (1,1,'java',15),(2,1,'programming',16),(3,1,'fundamentals',13),(4,180,'java',10),(5,180,'programming',10),(6,180,'fundamentals',10);
/*!40000 ALTER TABLE `answers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `attempts`
--

DROP TABLE IF EXISTS `attempts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `attempts` (
  `id` bigint(10) NOT NULL,
  `userid` bigint(10) DEFAULT NULL,
  `testtype` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `a_count` int(2) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid_idx` (`userid`),
  CONSTRAINT `user_fk` FOREIGN KEY (`userid`) REFERENCES `registration` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `attempts`
--

LOCK TABLES `attempts` WRITE;
/*!40000 ALTER TABLE `attempts` DISABLE KEYS */;
/*!40000 ALTER TABLE `attempts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `questions`
--

DROP TABLE IF EXISTS `questions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `questions` (
  `id` bigint(10) NOT NULL AUTO_INCREMENT,
  `question` text,
  `type` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `questions`
--

LOCK TABLES `questions` WRITE;
/*!40000 ALTER TABLE `questions` DISABLE KEYS */;
INSERT INTO `questions` VALUES (1,'what is a standard output','fundamental'),(2,'which of the following is not a oop language','fundamental'),(3,'which of the following is an example of runtime','fundamental'),(4,'what is encapsulation','fundamental'),(5,'what is a recursive method','fundamental'),(6,'which one of the following data structure has a lifo','fundamental'),(7,'which one of the following is used to represent the complexity of an algorithm','fundamental'),(8,'what does a pwd command does in unix environment','fundamental'),(9,'what is a daemon process','fundamental'),(10,'a class that implements an interface','fundamental'),(11,'Which class cannot be subclassed (or extended) in java?','java'),(12,'Can we declare abstract static method?','java'),(13,'Why we use array as a parameter of main method','java'),(14,'Suspend thread can be revived by using','java'),(15,'Which collection class associates values witch keys, and orders the keys according to their natural order','java'),(16,'Which statement is static and synchronized in JDBC API','java'),(17,'The JDBC-ODBC bridge is','java'),(18,'The class java.sql.Timestamp is associated with','java'),(19,'Converting a primitive type data into its corresponding wrapper class object instance is called','java'),(20,'If method have same name as class name and method don�t have any return type then it is known as','java'),(21,' Main method parameter has which type of data type','java'),(22,'Which among following classes is not part of Java\'s collection framework','java'),(23,'Which of the following class defines print() and println() method?','java'),(24,'Which of the following modifier is used to prevent a property from being serialized?','java'),(25,'Runnable is','java'),(26,'What will be the output of following code?','java'),(27,'What is the numerical range of a char?','java'),(28,'Which of these process occur automatically by java run time system?','java'),(29,'Which of these is a method of ObjectOutput interface used to finalize the output state so that any buffers are cleared?','java'),(30,' Which of these is an incorrect form of using method max() to obtain maximum element?','java'),(31,'Which of these methods can randomize all elements in a list?','java'),(32,'Which of these is static variable defined in Collections?','java'),(33,' In java, ............ can only test for equality, where as .......... can evaluate any type of the Boolean expression.','java'),(34,'Which of the following for loops will be an infinite loop?','java'),(35,'____________ method cannot be overridden.','java'),(36,'To be badgered is to be','language'),(37,'If you are inundated with work,you would be','language'),(38,'Identity the closest in meaning to indenture','language'),(39,'in to the limited space given him a headline writer must compress the ______ of the news and he must do it without______','language'),(40,'uranium:fissionable :: x:malleable,identify x','language'),(41,'If a person cannot be easily handled or dealt with,he will not be complimented for his','language'),(42,'Human memory is not_______,especially on ancient happenings that smack of the ______','language'),(43,'The cheif minister,as well as his advisers,have agreeed to adress the rally','language');
/*!40000 ALTER TABLE `questions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `registration`
--

DROP TABLE IF EXISTS `registration`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `registration` (
  `id` bigint(10) NOT NULL AUTO_INCREMENT,
  `fname` varchar(45) DEFAULT NULL,
  `lname` varchar(45) DEFAULT NULL,
  `phone` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `password` varchar(45) DEFAULT NULL,
  `usertype` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=235 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `registration`
--

LOCK TABLES `registration` WRITE;
/*!40000 ALTER TABLE `registration` DISABLE KEYS */;
INSERT INTO `registration` VALUES (1,'Sweta','Vahia','9898989898','test@t.com','123','admin'),(2,'rakesh','bharti','1212121212','test@1.com','123','user'),(3,'shradha','kharat','1231231231','test@2.com','123','user'),(4,'arun','hosamani','1234123412','test@3.com','123','user'),(5,'risabh','sharma','1234512345','test@4.com','123','user'),(175,'llam','sam','9729533391','rpqb@rpqb.com','123','user'),(176,'','','','rakesh.bharati@rapidqube.com','A@a12','user'),(177,'','','','s@u.com','Sweta1','user'),(178,'swetas','vahiav','9898984123','s@r.com','Sweta1','user'),(179,'Rakesh','Bhartti','9804608459','rakesh.bharati@rapidqube.com','A@a12','user'),(180,'lol','dsd','1234567890','test@6.com','123','user'),(181,'risabh','sharma','9029533390','rpqb@rpqb.com','123','user'),(193,'Rakesh','bharati','9681817926','','A@a12','user'),(194,'rakesh','Bharti','1234567893','s@s.com','S@s12','user'),(195,'risabh','sharma','9029533381','rpqb@rpqb.com','123',''),(196,'sunny','bharati','1239875464','rakesh.bharati@rapidqube.com','G@g12',''),(197,'Vikram','Viswanathan','7462826484','v@mailnesia.com','Test@123',''),(198,'risabh','sharma','8169226823','Risabh.sharma@gmail.com','Test123',''),(199,'risabh','sharma','12121212121','r@r.com','Test123',''),(200,'xyz','abc','9988899888','abc@xyz','Xyz@123',''),(201,'vinay','phenix','9666430037','vinayphenix@gmail.com','calyxN@143',''),(202,'pxyz','pabc','1234567899','pabc@a.com','Pxyz@123',''),(203,'Rahul','Desai','9930831907','rahul.desai@rapidqube.com','Rpqb@123',''),(204,'rahul','desai','2124542434','r@d.com','Test@123',''),(205,'piyusha','patel','123456789','piyusha@gmail.com','Piyu@1123',''),(206,'Suahs','Patil','8080910100','patilsuhas1004@gmail.com','India@143',''),(207,'jilani','shaik','9440697390','jilanishaik338@gmail.com','Jilani1234',''),(208,'Rahul','Desai','8930831907','rahul.desai@rapidqube.com','Rpqb@123',''),(209,'fghfdh','dfhdfh','2323432434','hh@b.gf','Abc@123',''),(210,'jhjkhg','fjkg','12334455666','ff@mm.jj','Abc@123',''),(211,'ram','jai','9098778909','ramjai@gmail.com','Ramjai@123',''),(212,'sai','ram','9989897760','ram@gmail.com','raM@123',''),(213,'sai','rama','9876567789','rama@gmail.com','Rama@123',''),(214,'sai','uma','9858675849','uma@gmail.com','Uma@123',''),(215,'sai','vinay','9848781828','sai1@gmail.com','Sai@123',''),(216,'vinay','kumar','9868483938','sai123@gmail.com','Sai@123',''),(217,'vinay','vyas','9857348393','sai1234@gmail.com','Sai@123',''),(218,'Sai','Vinay','9667493939','sai19@gmail.com','Sai@123',''),(219,'sai','vinay','9484728393','sai20@gmail.com','Sai@123',''),(220,'sai','vinay','4483403055','sai21@gmail.com','Sai@123',''),(221,'sai','vinay','9849349293','sai1997@gmail.com','Sai@123',''),(222,'sai','vinay','9828392838','sai2000@gmail.com','Sai@123',''),(223,'SAI','VINAY','9858748392','sai23@gmail.com','Sai@123',''),(224,'sai','vinay','9867867754','sai27@gmail.com','Sai@123',''),(225,'sai','vinay','9847473839','sai2001@gmail.com','Sai@123',''),(226,'sai','vinay','9847489384','sai919@gmail.com','Sai@123',''),(227,'sai','vinay','9585858940','sai9819@gmail.com','Sai@123',''),(228,'sai','vinay','9847493849','sai30@gmail.com','Sai@123',''),(229,'sai','vinay','9877657566','sai31@gmail.com','Sai@123',''),(230,'sai','vinay','9474848559','sai43@gmail.com','Sai@123',''),(231,'sai','vinay','9666458589','sai45@gmail.com','Sai@123',''),(232,'sai','vinay','8574848475','sai46@gmail.com','Sai@123',''),(233,'sai','vinay','9857483994','sai48@gmail.com','Sai@123',''),(234,'sweta','vahia','9892184841','rapid@y.com','Sweta123','');
/*!40000 ALTER TABLE `registration` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `token`
--

DROP TABLE IF EXISTS `token`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `token` (
  `id` bigint(10) NOT NULL AUTO_INCREMENT,
  `uid` bigint(10) DEFAULT NULL,
  `token` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `lastaccesstime` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `token_UNIQUE` (`token`),
  KEY `uid_idx` (`uid`),
  CONSTRAINT `uid` FOREIGN KEY (`uid`) REFERENCES `registration` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=1029 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `token`
--

LOCK TABLES `token` WRITE;
/*!40000 ALTER TABLE `token` DISABLE KEYS */;
INSERT INTO `token` VALUES (59,178,'IVdMzElHxoOZFYetMXRn','2017-03-24 15:42:39'),(62,177,'LJnPsdJYkYNFYdFwMGtf','2017-03-24 10:29:49'),(66,176,'VvtqGQLZVnSBbYhmBPfH','2017-03-24 11:10:03'),(103,3,'szYepYJYdlvDXOUybQcr','2017-03-25 07:20:24'),(114,194,'fVxUpZXkWPFPVGWZSllY','2017-03-25 10:43:09'),(127,179,'PMwUyoZMogmNGYGxQbpo','2017-03-25 10:50:11'),(146,197,'pPeHeXwNcdpOPfBZZjEC','2017-04-19 13:26:54'),(148,198,'EILLcvwZVHEMjncWkAHm','2017-04-20 04:26:59'),(155,195,'kcujUiuyvLfSLGupwAxa','2017-05-08 08:48:21'),(198,203,'cExwingnLWcjwjwbBHcZ','2017-05-08 13:39:25'),(330,199,'GBFtLhxGTbgnaDEmvkfS','2017-05-10 08:02:49'),(335,204,'UtMOFSprYlokUtlJIUnY','2017-05-10 14:38:47'),(825,206,'pLNgJkvGPUMtCYMkMBLu','2017-05-19 19:20:43'),(886,207,'XGypeRxaytxtOoczAwlO','2017-05-22 17:14:31'),(888,4,'NQLGxurRwEjpSYAWUGUh','2017-05-22 13:11:13'),(936,1,'ruVflKxPOreJmxwmdZrh','2017-05-24 05:01:42'),(949,201,'srwAkSrgaisdRnjZbGTz','2017-05-24 06:25:19'),(957,211,'QgEAwAamoQMPmeszKSrS','2017-05-24 07:46:42'),(959,212,'gZisXNdLQSPBwpfvVSSX','2017-05-24 07:56:36'),(960,213,'oEOklyztioSLiqhNuTQF','2017-05-24 08:42:51'),(961,214,'wXsXQCZdOvSvjchiWSRq','2017-05-24 08:48:12'),(964,215,'CyyOQhAYDIGzoeDHHiPR','2017-05-24 08:59:34'),(970,216,'hrraXMUaTyKErkAEfwkD','2017-05-24 09:44:05'),(971,217,'TFrhFrpWNbYzPlyQxOPM','2017-05-24 09:52:21'),(972,218,'YaZqvEoQNSSrdXZniIwW','2017-05-24 09:59:46'),(973,219,'pfSPoGeMgBuAeQLTCeSR','2017-05-24 10:04:14'),(975,220,'NzeoUJWQiTvAaPorlpzi','2017-05-24 10:13:12'),(981,221,'KfhoxzKpLzsrawVjYGJn','2017-05-24 10:25:25'),(984,222,'drDproESVlRcpMPMpxMh','2017-05-24 10:29:39'),(988,223,'UUjSTIgqqFLZxnmPYeFv','2017-05-24 10:39:38'),(990,224,'IAQvKlRkstnMCBYggKli','2017-05-24 10:59:24'),(993,225,'OgITXvqyhHxlRaCKjMoW','2017-05-24 11:11:13'),(994,226,'znyiphbRjlBcWtbcLxtD','2017-05-24 11:14:44'),(995,227,'psudfUpiyvOoqJDZNKyC','2017-05-24 11:18:20'),(997,228,'OnaXFIRUHpfgtaBOMHIR','2017-05-24 11:32:44'),(998,229,'rzfYSCiTQIuKrqWHacku','2017-05-24 11:54:22'),(999,230,'xtYEwcKNQfcXUHVYSlbc','2017-05-24 12:31:42'),(1000,231,'YENOgwerjHRpeooMXVIU','2017-05-24 12:45:00'),(1001,232,'QilDjhiqzRmbdAxgbDkX','2017-05-24 12:55:52'),(1002,233,'rhmQsomDfkXaigXACiQp','2017-05-24 13:01:59'),(1028,2,'jYoSKeKjPYcbxYLGmgnw','2017-05-26 10:39:02');
/*!40000 ALTER TABLE `token` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-05-26 17:08:34

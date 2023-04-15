-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 15, 2023 at 03:32 AM
-- Server version: 10.4.17-MariaDB
-- PHP Version: 7.3.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `skyshi`
--

-- --------------------------------------------------------

--
-- Table structure for table `activity_groups`
--

CREATE TABLE `activity_groups` (
  `id` int(11) NOT NULL,
  `title` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `activity_groups`
--

INSERT INTO `activity_groups` (`id`, `title`, `email`, `created_at`, `updated_at`) VALUES
(2, 'Testing', 'ramdanriawan4@gmail.com', '2023-04-14 23:46:32', '2023-04-15 01:17:58'),
(4, 'Title 3', 'ramdanriawan5@gmail.com', '2023-04-15 00:19:49', '2023-04-15 00:19:49'),
(6, 'Title 3', 'ramdanriawan6@gmail.com', '2023-04-15 00:26:23', '2023-04-15 00:26:23'),
(8, 'Title 3', 'ramdanriawan7@gmail.com', '2023-04-15 00:26:52', '2023-04-15 00:26:52'),
(10, 'Title 3', 'ramdanriawan8@gmail.com', '2023-04-15 00:36:25', '2023-04-15 00:36:25'),
(11, 'Test', 'test@mail.com', '2023-04-15 01:05:31', '2023-04-15 01:05:31');

-- --------------------------------------------------------

--
-- Table structure for table `todos`
--

CREATE TABLE `todos` (
  `id` int(11) NOT NULL,
  `activity_group_id` int(11) NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT 0,
  `priority` enum('very-high','medium','low') NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `todos`
--

INSERT INTO `todos` (`id`, `activity_group_id`, `is_active`, `priority`, `created_at`, `updated_at`) VALUES
(2, 2, 1, 'medium', '2023-04-14 23:49:21', '2023-04-15 01:29:41'),
(3, 2, 1, 'medium', '2023-04-15 01:02:14', '2023-04-15 01:02:14'),
(4, 2, 1, 'medium', '2023-04-15 01:02:20', '2023-04-15 01:02:20'),
(5, 2, 0, 'medium', '2023-04-15 01:03:28', '2023-04-15 01:03:28'),
(6, 2, 0, 'medium', '2023-04-15 01:03:29', '2023-04-15 01:03:29'),
(8, 2, 0, 'medium', '2023-04-15 01:03:46', '2023-04-15 01:03:46'),
(9, 2, 1, 'medium', '2023-04-15 01:25:58', '2023-04-15 01:25:58'),
(10, 2, 1, 'medium', '2023-04-15 01:26:11', '2023-04-15 01:26:11'),
(11, 2, 1, 'medium', '2023-04-15 01:26:13', '2023-04-15 01:26:13'),
(12, 2, 0, 'medium', '2023-04-15 01:28:48', '2023-04-15 01:28:48'),
(13, 2, 0, 'medium', '2023-04-15 01:28:55', '2023-04-15 01:28:55'),
(14, 2, 0, 'medium', '2023-04-15 01:28:56', '2023-04-15 01:28:56'),
(15, 2, 0, 'medium', '2023-04-15 01:28:57', '2023-04-15 01:28:57'),
(16, 2, 1, 'medium', '2023-04-15 01:29:02', '2023-04-15 01:29:02'),
(17, 2, 1, 'medium', '2023-04-15 01:29:05', '2023-04-15 01:29:05'),
(18, 2, 1, 'medium', '2023-04-15 01:29:25', '2023-04-15 01:29:25'),
(20, 2, 1, 'medium', '2023-04-15 01:30:08', '2023-04-15 01:30:08');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `activity_groups`
--
ALTER TABLE `activity_groups`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `todos`
--
ALTER TABLE `todos`
  ADD PRIMARY KEY (`id`),
  ADD KEY `activity_group_id` (`activity_group_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `activity_groups`
--
ALTER TABLE `activity_groups`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `todos`
--
ALTER TABLE `todos`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `todos`
--
ALTER TABLE `todos`
  ADD CONSTRAINT `todos_ibfk_1` FOREIGN KEY (`activity_group_id`) REFERENCES `activity_groups` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Oct 05, 2021 at 01:10 PM
-- Server version: 10.4.18-MariaDB
-- PHP Version: 8.0.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bansosman`
--

-- --------------------------------------------------------

--
-- Table structure for table `admins`
--

CREATE TABLE `admins` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` varchar(191) DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `role_id` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `admins`
--

INSERT INTO `admins` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `password`, `role_id`) VALUES
(1, '2021-10-04 15:51:15.005', '2021-10-04 15:51:15.005', NULL, 'admin1', '$2a$04$dlSg4GYdz9kHf/SA6PDaBuDMzi1O8ty2TBkc29FzQZg3OIQW94L7q', 1),
(2, '2021-10-05 07:00:01.315', '2021-10-05 07:00:01.315', NULL, 'admin12', '$2a$04$jIwE3s/WYeiPwIP.OMKu2u.n7ehn516wjU9UpWQVWsTl5E2A1Ilzy', 1);

-- --------------------------------------------------------

--
-- Table structure for table `apbns`
--

CREATE TABLE `apbns` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `dana_bansos` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `apbns`
--

INSERT INTO `apbns` (`id`, `created_at`, `updated_at`, `deleted_at`, `dana_bansos`) VALUES
(1, '2021-10-04 15:52:16.229', '2021-10-04 15:52:16.229', NULL, 1500000000),
(2, '2021-10-04 15:52:19.915', '2021-10-04 15:52:19.915', NULL, 1400000000),
(3, '2021-10-04 15:52:24.093', '2021-10-04 15:52:24.093', NULL, 1800000000),
(4, '2021-10-04 17:07:54.519', '2021-10-04 17:07:54.519', NULL, 1800000000),
(5, '2021-10-04 17:07:55.159', '2021-10-04 17:07:55.159', NULL, 1800000000),
(6, '2021-10-04 17:07:55.637', '2021-10-04 17:07:55.637', NULL, 1800000000),
(7, '2021-10-05 07:00:17.272', '2021-10-05 07:00:17.272', NULL, 1800000000),
(8, '2021-10-05 07:00:17.696', '2021-10-05 07:00:17.696', NULL, 1800000000),
(9, '2021-10-05 07:00:17.848', '2021-10-05 07:00:17.848', NULL, 1800000000),
(10, '2021-10-05 07:00:18.008', '2021-10-05 07:00:18.008', NULL, 1800000000),
(11, '2021-10-05 07:00:18.175', '2021-10-05 07:00:18.175', NULL, 1800000000),
(12, '2021-10-05 07:00:18.336', '2021-10-05 07:00:18.336', NULL, 1800000000),
(13, '2021-10-05 07:00:18.482', '2021-10-05 07:00:18.482', NULL, 1800000000);

-- --------------------------------------------------------

--
-- Table structure for table `daerahs`
--

CREATE TABLE `daerahs` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `apbn_id` bigint(20) UNSIGNED DEFAULT NULL,
  `provinsi` longtext DEFAULT NULL,
  `city` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `daerahs`
--

INSERT INTO `daerahs` (`id`, `created_at`, `updated_at`, `deleted_at`, `apbn_id`, `provinsi`, `city`) VALUES
(1, '2021-10-04 15:52:52.558', '2021-10-04 15:52:52.558', NULL, 2, 'Jawa Tengah', 'Kudus'),
(2, '2021-10-04 15:53:01.308', '2021-10-04 15:53:01.308', NULL, 1, 'Jawa Tengah', 'Semarang'),
(3, '2021-10-04 15:57:20.206', '2021-10-04 15:57:20.206', NULL, 1, 'jateng', 'Semarang'),
(4, '2021-10-04 17:08:02.220', '2021-10-04 17:08:02.220', NULL, 5, 'jateng', 'Semarang'),
(5, '2021-10-04 17:08:06.198', '2021-10-04 17:08:06.198', NULL, 4, 'jateng', 'Semarang'),
(6, '2021-10-04 17:10:35.069', '2021-10-04 17:10:35.069', NULL, 4, 'jateng', 'Kudus'),
(7, '2021-10-05 07:00:41.025', '2021-10-05 07:00:41.025', NULL, 4, 'Jakarta', 'Jakarta Selatan');

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` bigint(20) NOT NULL,
  `name` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `name`) VALUES
(1, 'Owner'),
(2, 'Admin');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nik` bigint(20) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `foto_rumah` longtext DEFAULT NULL,
  `gaji` bigint(20) DEFAULT NULL,
  `alamat` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `nik`, `name`, `email`, `password`, `foto_rumah`, `gaji`, `alamat`) VALUES
(1, '2021-10-04 15:51:12.829', '2021-10-04 15:51:12.829', NULL, 12345, 'ilhampras', 'ilhmprz.gmail', '$2a$04$WxZggjPP5rJnCFsNBCN/NOhZt/6MCcsx4kNCF..6LecisYIt59vRq', 'rumahku', 120000, 'alamatku'),
(3, '2021-10-04 17:33:13.992', '2021-10-04 17:33:13.992', NULL, 12345, 'ilhampr', 'ilhmprz.gmail', '$2a$04$I3xfCN62vWAd8GrKwqnpq.AKt97Ha.Q3QsWjOCPe50PoAstCaRRci', 'rumahku', 120000, 'alamatku'),
(5, '2021-10-05 06:59:56.016', '2021-10-05 06:59:56.016', NULL, 12345, 'ilhampraaaas', 'ilhmprz.gmail', '$2a$04$xeVfiikb/6IMF0DAdSuzmuqDc8CYk1s7FBUie3DgnnqwrF9hPlNpy', 'rumahku', 120000, 'alamatku');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admins`
--
ALTER TABLE `admins`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD KEY `idx_admins_deleted_at` (`deleted_at`),
  ADD KEY `fk_admins_role` (`role_id`);

--
-- Indexes for table `apbns`
--
ALTER TABLE `apbns`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_apbns_deleted_at` (`deleted_at`);

--
-- Indexes for table `daerahs`
--
ALTER TABLE `daerahs`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_daerahs_deleted_at` (`deleted_at`),
  ADD KEY `fk_daerahs_apbns` (`apbn_id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admins`
--
ALTER TABLE `admins`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `apbns`
--
ALTER TABLE `apbns`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- AUTO_INCREMENT for table `daerahs`
--
ALTER TABLE `daerahs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `admins`
--
ALTER TABLE `admins`
  ADD CONSTRAINT `fk_admins_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE NO ACTION;

--
-- Constraints for table `daerahs`
--
ALTER TABLE `daerahs`
  ADD CONSTRAINT `fk_daerahs_apbns` FOREIGN KEY (`apbn_id`) REFERENCES `apbns` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: db
-- Waktu pembuatan: 07 Jun 2020 pada 05.53
-- Versi server: 10.4.12-MariaDB-1:10.4.12+maria~bionic
-- Versi PHP: 7.4.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sistem_akademik`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `student_table`
--

CREATE TABLE `student_table` (
  `student_id` int(255) NOT NULL,
  `student_nim` varchar(255) NOT NULL,
  `student_name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data untuk tabel `student_table`
--

INSERT INTO `student_table` (`student_id`, `student_nim`, `student_name`) VALUES
(24, '1622009', 'Ridwan');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `student_table`
--
ALTER TABLE `student_table`
  ADD PRIMARY KEY (`student_id`),
  ADD UNIQUE KEY `student_nim` (`student_nim`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `student_table`
--
ALTER TABLE `student_table`
  MODIFY `student_id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

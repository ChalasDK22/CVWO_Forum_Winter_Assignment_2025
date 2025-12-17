-- phpMyAdmin SQL Dump
-- version 5.2.3
-- https://www.phpmyadmin.net/
--
-- Host: db
-- Generation Time: Dec 16, 2025 at 01:49 PM
-- Server version: 9.5.0
-- PHP Version: 8.3.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `forum-db`
--

-- --------------------------------------------------------

--
-- Table structure for table `COMMENTS`
--

CREATE TABLE `COMMENTS` (
  `topicID` int NOT NULL,
  `user_id` int NOT NULL,
  `post_date` date NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `post_id` int NOT NULL,
  `comment_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `COMMENTS`
--

INSERT INTO `COMMENTS` (`topicID`, `user_id`, `post_date`, `content`, `post_id`, `comment_id`) VALUES
(1, 2, '2024-06-02', 'Welcome to the forum!', 1, 1),
(1, 5, '2024-06-03', 'Nice introduction!', 1, 2),
(2, 3, '2024-06-04', 'AI is incredible lately.', 3, 3),
(2, 7, '2024-06-04', 'Totally agree about VR.', 4, 4),
(3, 8, '2024-06-06', 'I love RPG games too.', 5, 5),
(3, 1, '2024-06-06', 'Thanks for the recommendation!', 6, 6),
(4, 4, '2024-06-07', 'Japan is beautiful!', 7, 7),
(4, 9, '2024-06-08', 'Did you visit Osaka?', 7, 8),
(5, 10, '2024-06-09', 'That recipe looks great!', 9, 9),
(5, 6, '2024-06-10', 'Will try cooking this soon.', 9, 10),
(3, 3, '2024-06-11', 'What platform do you play on?', 5, 11),
(2, 2, '2024-06-11', 'Tech news is always exciting.', 3, 12),
(1, 7, '2024-06-12', 'Happy to have you here!', 2, 13),
(4, 8, '2024-06-13', 'Travel posts are inspiring.', 8, 14),
(5, 1, '2024-06-14', 'Cooking is fun!', 10, 15);

-- --------------------------------------------------------

--
-- Table structure for table `POSTS`
--

CREATE TABLE `POSTS` (
  `topicID` int NOT NULL,
  `user_id` int NOT NULL,
  `post_date` date NOT NULL,
  `title` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `post_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `POSTS`
--

INSERT INTO `POSTS` (`topicID`, `user_id`, `post_date`, `title`, `content`, `post_id`) VALUES
(1, 1, '2024-06-01', 'Hello everyone!', 'Glad to be here.', 1),
(1, 3, '2024-06-02', 'Nice to meet you all', 'Excited to join.', 2),
(2, 2, '2024-06-03', 'AI is growing fast', 'AI will shape the future.', 3),
(2, 7, '2024-06-04', 'Tech Opinion', 'VR and AR adoption is increasing.', 4),
(3, 3, '2024-06-05', 'Best game this year?', 'Looking for new game suggestions.', 5),
(3, 8, '2024-06-06', 'Game Thoughts', 'This year’service releases are amazing.', 6),
(4, 4, '2024-06-07', 'Trip to Japan', 'Kyoto temples were stunning.', 7),
(4, 6, '2024-06-08', 'Travel Help', 'Planning a trip soon, need ideas.', 8),
(5, 5, '2024-06-09', 'My Pasta Recipe', 'Sharing my best creamy pasta recipe.', 9),
(5, 9, '2024-06-10', 'Cooking Tips', 'Salt your pasta water!', 10);

-- --------------------------------------------------------

--
-- Table structure for table `TOPICS`
--

CREATE TABLE `TOPICS` (
  `topicID` int NOT NULL,
  `user_id` int NOT NULL,
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `TOPICS`
--

INSERT INTO `TOPICS` (`topicID`, `user_id`, `name`, `description`) VALUES
(6, 6, 'Music Lounge', 'Discuss your favorite songs, artists, and albums.'),
(7, 7, 'Fitness & Health', 'Share workouts, routines, and wellness tips.'),
(8, 8, 'Movies & TV Shows', 'Talk about films, series, and reviews.'),
(9, 9, 'Programming Help', 'Ask coding questions and share solutions.'),
(10, 10, 'School & Study Tips', 'Give advice on studying and exams.'),
(11, 1, 'Pets & Animals', 'Share cute pet photos and stories.'),
(12, 2, 'Sports Talk', 'Discuss football, basketball, and other sports.'),
(13, 3, 'Art & Creativity', 'Post drawings, designs, or creative projects.'),
(14, 4, 'Business & Finance', 'Talk about investments, savings, and career paths.'),
(15, 5, 'Random Chat', 'A place to talk about anything and everything.');

-- --------------------------------------------------------

--
-- Table structure for table `USERS`
--

CREATE TABLE `USERS` (
  `user_id` int NOT NULL,
  `username` varchar(20) NOT NULL,
  `password` varchar(40) NOT NULL,
  `registration_date` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `USERS`
--

INSERT INTO `USERS` (`user_id`, `username`, `password`, `registration_date`) VALUES
(1, 'john', 'password123', '2024-01-10'),
(2, 'alice', 'mypassword', '2024-02-15'),
(3, 'bob', 'securepass', '2024-03-01'),
(4, 'charlie', 'charliepass', '2024-03-20'),
(5, 'demo', 'test1234', '2024-04-05'),
(6, 'susan', 'susanpass', '2024-04-10'),
(7, 'michael', 'mike2024', '2024-05-01'),
(8, 'anna', 'anna_pass', '2024-05-12'),
(9, 'test_user1', 'abc12345', '2024-06-01'),
(10, 'devaccount', 'devpass', '2024-06-15');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `COMMENTS`
--
ALTER TABLE `COMMENTS`
  ADD PRIMARY KEY (`comment_id`);

--
-- Indexes for table `POSTS`
--
ALTER TABLE `POSTS`
  ADD PRIMARY KEY (`post_id`);

--
-- Indexes for table `TOPICS`
--
ALTER TABLE `TOPICS`
  ADD PRIMARY KEY (`topicID`);

--
-- Indexes for table `USERS`
--
ALTER TABLE `USERS`
  ADD PRIMARY KEY (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `COMMENTS`
--
ALTER TABLE `COMMENTS`
  MODIFY `comment_id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `POSTS`
--
ALTER TABLE `POSTS`
  MODIFY `post_id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `TOPICS`
--
ALTER TABLE `TOPICS`
  MODIFY `topicID` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `USERS`
--
ALTER TABLE `USERS`
  MODIFY `user_id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

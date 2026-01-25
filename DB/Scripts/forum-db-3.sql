SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

--
-- Database: `defaultdb`
--

-- --------------------------------------------------------
-- Table structure for table `USERS`
-- --------------------------------------------------------
CREATE TABLE `USERS` (
                         `user_id` int NOT NULL AUTO_INCREMENT,
                         `username` varchar(20) NOT NULL,
                         `password` varchar(255) NOT NULL,
                         `registration_date` date NOT NULL,
                         PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------
-- Table structure for table `TOPICS`
-- --------------------------------------------------------
CREATE TABLE `TOPICS` (
                          `topicID` int NOT NULL AUTO_INCREMENT,
                          `user_id` int NOT NULL,
                          `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                          `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                          PRIMARY KEY (`topicID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------
-- Table structure for table `POSTS`
-- --------------------------------------------------------
CREATE TABLE `POSTS` (
                         `topicID` int NOT NULL,
                         `user_id` int NOT NULL,
                         `post_date` date NOT NULL,
                         `title` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                         `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                         `post_id` int NOT NULL AUTO_INCREMENT,
                         PRIMARY KEY (`post_id`),
                         KEY `fk_posts_topics` (`topicID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------
-- Table structure for table `COMMENTS`
-- --------------------------------------------------------
CREATE TABLE `COMMENTS` (
                            `user_id` int NOT NULL,
                            `post_date` date NOT NULL,
                            `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
                            `post_id` int NOT NULL,
                            `comment_id` int NOT NULL AUTO_INCREMENT,
                            PRIMARY KEY (`comment_id`),
                            KEY `fk_comments_posts` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------
-- Constraints
-- --------------------------------------------------------
ALTER TABLE `POSTS`
    ADD CONSTRAINT `fk_posts_topics` FOREIGN KEY (`topicID`) REFERENCES `TOPICS` (`topicID`)
        ON DELETE RESTRICT ON UPDATE RESTRICT;

ALTER TABLE `COMMENTS`
    ADD CONSTRAINT `fk_comments_posts` FOREIGN KEY (`post_id`) REFERENCES `POSTS` (`post_id`)
        ON DELETE CASCADE ON UPDATE RESTRICT;

-- --------------------------------------------------------
-- Dumping data (keep exactly as your original)
-- --------------------------------------------------------

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

INSERT INTO `TOPICS` (`topicID`, `user_id`, `name`, `description`) VALUES
                                                                       (1, 1, 'Fitness & Health', 'Share workouts, routines, and wellness tips.'),
                                                                       (2, 2, 'Movies & TV Shows', 'Talk about films, series, and reviews.'),
                                                                       (3, 3, 'Programming Help', 'Ask coding questions and share solutions.'),
                                                                       (4, 4, 'Travel', 'Share travel stories and tips.'),
                                                                       (5, 5, 'Food & Cooking', 'Share recipes and cooking tips.'),
                                                                       (6, 6, 'Gaming', 'Discuss games, platforms, and releases.');

INSERT INTO `POSTS` (`topicID`, `user_id`, `post_date`, `title`, `content`, `post_id`) VALUES
                                                                                           (1, 1, '2024-06-01', 'Welcome to Fitness', 'Let’s talk about healthy habits.', 1),
                                                                                           (1, 2, '2024-06-02', 'Morning Workout', 'Best exercises to start your day.', 2),
                                                                                           (2, 3, '2024-06-03', 'Best Movie 2024', 'What is your favorite movie this year?', 3),
                                                                                           (2, 4, '2024-06-04', 'TV Series 추천', 'Any good series to binge-watch?', 4),
                                                                                           (3, 5, '2024-06-05', 'Go vs Java', 'Which language should I learn?', 5),
                                                                                           (3, 6, '2024-06-06', 'SQL Help', 'How do joins really work?', 6),
                                                                                           (4, 7, '2024-06-07', 'Japan Trip', 'Kyoto was absolutely beautiful.', 7),
                                                                                           (4, 8, '2024-06-08', 'Europe Travel', 'Backpacking tips for Europe.', 8),
                                                                                           (5, 9, '2024-06-09', 'Best Pasta', 'Creamy pasta recipe inside.', 9),
                                                                                           (5, 10, '2024-06-10', 'Baking Tips', 'How to bake fluffy bread.', 10),
                                                                                           (6, 11, '2024-06-11', 'Best RPGs', 'Top RPG games this year?', 11),
                                                                                           (6, 12, '2024-06-12', 'Console vs PC', 'Which do you prefer?', 12);

INSERT INTO `COMMENTS` (`user_id`, `post_date`, `content`, `post_id`, `comment_id`) VALUES
                                                                                        (2, '2024-06-01', 'Great topic!', 1, 1),
                                                                                        (3, '2024-06-01', 'Very helpful tips.', 1, 2),
                                                                                        (4, '2024-06-02', 'I love morning workouts.', 2, 3),
                                                                                        (5, '2024-06-03', 'Amazing year for movies.', 3, 4),
                                                                                        (6, '2024-06-03', 'So many good releases!', 3, 5),
                                                                                        (7, '2024-06-04', 'Any Netflix recommendations?', 4, 6),
                                                                                        (8, '2024-06-05', 'Go is very clean.', 5, 7),
                                                                                        (9, '2024-06-05', 'Java still has a big ecosystem.', 5, 8),
                                                                                        (10, '2024-06-06', 'SQL joins take time to master.', 6, 9),
                                                                                        (11, '2024-06-07', 'Kyoto is beautiful!', 7, 10),
                                                                                        (12, '2024-06-07', 'Japan is on my bucket list.', 7, 11),
                                                                                        (1, '2024-06-08', 'Europe backpacking is amazing.', 8, 12),
                                                                                        (2, '2024-06-09', 'This pasta looks delicious!', 9, 13),
                                                                                        (3, '2024-06-09', 'I will try this recipe.', 9, 14),
                                                                                        (4, '2024-06-10', 'Baking tips are always useful.', 10, 15),
                                                                                        (5, '2024-06-11', 'Elden Ring is incredible.', 11, 16),
                                                                                        (6, '2024-06-11', 'RPGs are my favorite genre.', 11, 17),
                                                                                        (7, '2024-06-12', 'PC gaming all the way!', 12, 18);

COMMIT;

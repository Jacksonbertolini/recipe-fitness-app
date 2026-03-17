-- =============================================================
-- FitMeals Database Setup
-- Run as MySQL root: mysql -u root -p < schema.sql
-- =============================================================

-- 1. Create database
CREATE DATABASE IF NOT EXISTS fitmeals CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 2. Create application user (change password before production)
CREATE USER IF NOT EXISTS 'fitmeals_user'@'localhost' IDENTIFIED BY 'fitmeals_password';
GRANT ALL PRIVILEGES ON fitmeals.* TO 'fitmeals_user'@'localhost';
FLUSH PRIVILEGES;

USE fitmeals;

-- =============================================================
-- Tables
-- =============================================================

CREATE TABLE IF NOT EXISTS users (
    id            INT AUTO_INCREMENT PRIMARY KEY,
    email         VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name          VARCHAR(100) NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email)
);

CREATE TABLE IF NOT EXISTS user_goals (
    id               INT AUTO_INCREMENT PRIMARY KEY,
    user_id          INT NOT NULL,
    goal_type        ENUM('weight_gain', 'weight_loss') NOT NULL,
    target_calories  INT NOT NULL,
    created_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE KEY unique_user_goal (user_id)
);

CREATE TABLE IF NOT EXISTS recipes (
    id                 INT AUTO_INCREMENT PRIMARY KEY,
    name               VARCHAR(255) NOT NULL,
    description        TEXT,
    goal_type          ENUM('weight_gain', 'weight_loss') NOT NULL,
    prep_time_minutes  INT,
    cook_time_minutes  INT,
    servings           INT NOT NULL,
    ingredients        JSON NOT NULL,
    instructions       JSON NOT NULL,
    image_url          VARCHAR(500),
    created_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_goal_type (goal_type),
    INDEX idx_name (name)
);

CREATE TABLE IF NOT EXISTS recipe_nutrition (
    id         INT AUTO_INCREMENT PRIMARY KEY,
    recipe_id  INT NOT NULL,
    calories   INT NOT NULL,
    protein_g  DECIMAL(5,1) NOT NULL,
    carbs_g    DECIMAL(5,1) NOT NULL,
    fats_g     DECIMAL(5,1) NOT NULL,
    fiber_g    DECIMAL(5,1) DEFAULT 0,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE,
    UNIQUE KEY unique_recipe_nutrition (recipe_id)
);

CREATE TABLE IF NOT EXISTS user_favorites (
    id         INT AUTO_INCREMENT PRIMARY KEY,
    user_id    INT NOT NULL,
    recipe_id  INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id)   REFERENCES users(id)   ON DELETE CASCADE,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE,
    UNIQUE KEY unique_favorite (user_id, recipe_id),
    INDEX idx_user_favorites (user_id)
);

# FitMeals - Database Schema

## Overview

This database schema supports the FitMeals MVP features:
- User authentication and goal setting
- Recipe browsing and filtering
- Favorites management
- Nutrition tracking

## Tables Summary

### users
- id (primary key)
- email (unique, indexed)
- password_hash (bcrypt)
- name
- created_at, updated_at

### user_goals
- id (primary key)
- user_id (foreign key → users, unique)
- goal_type (enum: 'weight_gain', 'weight_loss')
- target_calories (int)
- created_at, updated_at

### recipes
- id (primary key)
- name (varchar 255, indexed)
- **description (text)**
- goal_type (enum: 'weight_gain', 'weight_loss', indexed)
- **prep_time_minutes (int)**
- **cook_time_minutes (int)**
- **servings (int, not null)**
- ingredients (JSON array)
- instructions (JSON array)
- image_url (varchar 500)
- created_at, updated_at

### recipe_nutrition
- id (primary key)
- recipe_id (foreign key → recipes, unique)
- calories (int, not null)
- protein_g (decimal 5,1)
- carbs_g (decimal 5,1)
- fats_g (decimal 5,1)
- **fiber_g (decimal 5,1, default 0)**

### user_favorites (junction table)
- id (primary key)
- user_id (foreign key → users)
- recipe_id (foreign key → recipes)
- created_at
- Unique constraint on (user_id, recipe_id)

## Complete SQL Schema

```sql
-- Users table
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email)
);

-- User goals table (one goal per user)
CREATE TABLE user_goals (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    goal_type ENUM('weight_gain', 'weight_loss') NOT NULL,
    target_calories INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE KEY unique_user_goal (user_id)
);

-- Recipes table
CREATE TABLE recipes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    goal_type ENUM('weight_gain', 'weight_loss') NOT NULL,
    prep_time_minutes INT,
    cook_time_minutes INT,
    servings INT NOT NULL,
    ingredients JSON NOT NULL,
    instructions JSON NOT NULL,
    image_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_goal_type (goal_type),
    INDEX idx_name (name)
);

-- Recipe nutrition table (one row per recipe)
CREATE TABLE recipe_nutrition (
    id INT AUTO_INCREMENT PRIMARY KEY,
    recipe_id INT NOT NULL,
    calories INT NOT NULL,
    protein_g DECIMAL(5,1) NOT NULL,
    carbs_g DECIMAL(5,1) NOT NULL,
    fats_g DECIMAL(5,1) NOT NULL,
    fiber_g DECIMAL(5,1) DEFAULT 0,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE,
    UNIQUE KEY unique_recipe_nutrition (recipe_id)
);

-- User favorites (junction table)
CREATE TABLE user_favorites (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    recipe_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (recipe_id) REFERENCES recipes(id) ON DELETE CASCADE,
    UNIQUE KEY unique_favorite (user_id, recipe_id),
    INDEX idx_user_favorites (user_id)
);
```

## JSON Data Structures

### recipes.ingredients (JSON Array)
```json
[
  {
    "name": "chicken breast",
    "amount": "200",
    "unit": "g"
  },
  {
    "name": "olive oil",
    "amount": "1",
    "unit": "tbsp"
  }
]
```

### recipes.instructions (JSON Array)
```json
[
  "Preheat oven to 400°F (200°C)",
  "Season chicken with salt and pepper",
  "Bake for 25-30 minutes until cooked through",
  "Let rest for 5 minutes before serving"
]
```

## Design Decisions

**1. Ingredients & Instructions as JSON**
- Simpler than normalized tables for MVP
- Go's `json.Unmarshal` handles this easily
- Can normalize later if needed for advanced querying

**2. Separate nutrition table**
- Keeps recipes table cleaner
- Makes it easy to query recipes by calorie ranges
- One-to-one relationship with recipes

**3. Enum for goal_type**
- Enforces data integrity (only 'weight_gain' or 'weight_loss')
- Makes filtering simple and performant

**4. Indexes on commonly queried columns**
- `goal_type`, `name` for recipe searches
- `email` for login lookups
- `user_id` for favorites lookups

**5. CASCADE DELETE**
- When a user is deleted, their goals and favorites are automatically removed
- When a recipe is deleted, its nutrition data and favorites are automatically removed

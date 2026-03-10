# FitMeals - Module 4 Project

## Project Definition

**Name:** FitMeals

**Target Audience:** People on a weight journey (gaining or losing) who want goal-specific recipes and nutrition tracking

**Problem:** Hard to find goal-specific recipes and track nutrition without multiple apps

**Value Proposition:** One app for goal-specific recipes and simple tracking in one place

## Must-Have Features (MVP)

1. Browse recipes filtered by goal (Weight Gain or Weight Loss)
2. View recipe details (ingredients, instructions, nutrition facts)
3. Set personal goal (gain/loss) and target calories
4. Save favorite recipes (requires login/auth)
5. Search recipes by name

## Deferred Features (Not Building Yet)

- Recipe sharing
- Daily meal logging
- Meal planning calendar
- Shopping list generator
- Progress photos
- Weight tracking over time
- Custom recipe creation
- Recipe ratings/reviews

## Tech Stack

- Frontend: React + Vite
- Backend: Go (Gin framework)
- Database: MySQL
- Auth: JWT tokens
- Deployment: AWS Lightsail with nginx

## Database Tables (Updated)

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
- **description (text)** ← ADDED
- goal_type (enum: 'weight_gain', 'weight_loss', indexed)
- **prep_time_minutes (int)** ← ADDED
- **cook_time_minutes (int)** ← ADDED
- **servings (int, not null)** ← ADDED
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
- **fiber_g (decimal 5,1, default 0)** ← ADDED

### user_favorites (junction table)
- id (primary key)
- user_id (foreign key → users)
- recipe_id (foreign key → recipes)
- created_at
- Unique constraint on (user_id, recipe_id)

## Timeline

Week 1: Auth + recipe browsing
Week 2-3: Favorites + search
Week 4: Deployment + polish

## Claude's Critique - Accepted Changes

✅ Added recipe detail view
✅ Added search functionality
✅ Removed recipe sharing (too complex)
✅ Confirmed realistic scope for 3-4 weeks

## Pages & Navigation (Lesson 2)

### Page List

1. **Home (/)** - Entry point with goal selection buttons
2. **Recipe Browse (/recipes)** - Filtered recipe list + search
3. **Recipe Detail (/recipes/:id)** - Full recipe info, ingredients, instructions
4. **Login (/login)** - User authentication
5. **Register (/register)** - Create new account
6. **Profile (/profile)** - Set goal type and target calories
7. **Favorites (/favorites)** - View saved recipes (auth required)

### Navigation Flow

**Logged Out Navbar:**
- Logo/Home
- Browse Recipes
- Sign In

**Logged In Navbar:**
- Logo/Home
- Browse Recipes
- Favorites
- Profile
- Logout

### Key User Journeys

**Browse without login:**
Home → Recipe Browse → Recipe Detail → Try to Save → Login → Recipe Detail

**New user registration:**
Home → Login → Register → Profile (set goal) → Browse

**Logged-in user:**
Browse → Recipe Detail → Save to Favorites → View Favorites

### Shared Components Identified

- Navigation Header (appears on all pages)
- Recipe Card (used on Browse and Favorites pages)
- Footer (appears on all pages)

## API Endpoints (Lesson 3)

### Authentication (Public)
- **POST /api/auth/register** - Create new user account
  - Request: { name, email, password }
  - Response: { token, user }

- **POST /api/auth/login** - Authenticate user, return JWT
  - Request: { email, password }
  - Response: { token, user }

### Recipes (Public)
- **GET /api/recipes** - List recipes with filters
  - Query params: goal (weight_gain|weight_loss), search (name)
  - Response: Array of recipes with nutrition
  - If authenticated: includes is_favorited field

- **GET /api/recipes/:id** - Get single recipe with full details
  - Response: Recipe with ingredients, instructions, nutrition
  - If authenticated: includes is_favorited field

### User Profile (Protected - JWT Required)
- **GET /api/user/profile** - Get user info and goal
  - Response: { id, name, email, goal }

- **PUT /api/user/goal** - Set or update user's goal
  - Request: { goal_type, target_calories }
  - Response: Success message

### Favorites (Protected - JWT Required)
- **GET /api/favorites** - Get user's saved recipes
  - Response: Array of favorited recipes

- **POST /api/favorites/:recipeId** - Add recipe to favorites
  - Response: Success message

- **DELETE /api/favorites/:recipeId** - Remove from favorites
  - Response: Success message

### Total: 8 Endpoints
- 4 Public (no auth required)
- 4 Protected (JWT required)

### Authentication Flow
- Use JWT tokens in Authorization: Bearer <token> header
- Middleware verifies token and attaches user_id to request
- Password hashing with bcrypt

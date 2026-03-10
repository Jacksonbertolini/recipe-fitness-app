# FitMeals - Architecture

## Pages & Navigation

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

## API Endpoints

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

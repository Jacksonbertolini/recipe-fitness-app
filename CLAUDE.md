# FitMeals - AI Context Document

This file helps AI assistants understand the FitMeals project structure and context.

## What is FitMeals?

FitMeals is a web application that helps people on a weight journey (gaining or losing) find goal-specific recipes and track nutrition. Users can browse recipes filtered by their fitness goal, save favorites, and set personal calorie targets.

**Core Features:**
- Browse recipes by goal (Weight Gain / Weight Loss)
- View detailed recipe information (ingredients, instructions, nutrition)
- User authentication with JWT
- Save favorite recipes
- Search recipes by name
- Set personal fitness goals and calorie targets

## Project Structure

```
recipe-fitness-app/
├── frontend/                 # React + Vite application
│   ├── src/
│   │   ├── components/      # Reusable React components
│   │   ├── pages/           # Page components (Home, Browse, Detail, etc.)
│   │   ├── services/        # API service layer
│   │   ├── context/         # React context for auth state
│   │   └── App.jsx          # Main app component
│   ├── public/              # Static assets
│   └── package.json
│
├── backend/                  # Go (Gin framework) API server
│   ├── main.go              # Application entry point
│   ├── handlers/            # HTTP request handlers
│   ├── models/              # Database models
│   ├── middleware/          # Auth middleware (JWT verification)
│   ├── database/            # Database connection and migrations
│   └── go.mod
│
├── docs/                     # Project documentation
│   ├── project-proposal.md  # Project definition and scope
│   ├── architecture.md      # Pages, navigation, API endpoints
│   └── database-schema.md   # Complete database schema with SQL
│
└── CLAUDE.md                # This file - AI context
```

## Tech Stack

**Frontend:**
- React 18
- Vite (build tool)
- React Router (navigation)
- Axios (HTTP client)

**Backend:**
- Go 1.21+
- Gin (web framework)
- JWT for authentication (`github.com/golang-jwt/jwt`)
- bcrypt for password hashing (`golang.org/x/crypto/bcrypt`)
- MySQL driver (`github.com/go-sql-driver/mysql`)

**Database:**
- MySQL 8.0+

**Deployment:**
- AWS Lightsail
- nginx (reverse proxy)

## Development Setup

### Frontend (React + Vite)
```bash
cd frontend
npm install
npm run dev         # Runs on http://localhost:5173
```

### Backend (Go + Gin)
```bash
cd backend
go mod download
go run main.go      # Runs on http://localhost:8080
```

### Database
```bash
# Create database
mysql -u root -p
CREATE DATABASE fitmeals;

# Run migrations (from backend/)
go run database/migrate.go
```

## Key Architectural Decisions

### 1. JWT Authentication
- Tokens issued on login/register
- Stored in localStorage on frontend
- Sent in `Authorization: Bearer <token>` header
- Backend middleware verifies tokens and extracts user_id

### 2. Recipe Data Strategy
- **Manually pre-populated**: 20-30 recipes entered directly into MySQL
- No external API dependency (simplifies MVP)
- Recipe data includes: name, description, goal_type, ingredients (JSON), instructions (JSON), nutrition facts

### 3. JSON for Ingredients & Instructions
- Ingredients stored as JSON array: `[{"name": "chicken", "amount": "200", "unit": "g"}]`
- Instructions stored as JSON array: `["Step 1", "Step 2", ...]`
- **Why?** Simpler than normalized tables for MVP, easy to work with in Go/React
- Can normalize later if complex querying is needed

### 4. Public Recipe Browsing
- Recipe list and detail pages are publicly accessible (no auth required)
- Authentication only required for:
  - Saving favorites
  - Setting personal goals
  - Viewing saved recipes

### 5. Database Design
- `users` - Authentication and profile
- `user_goals` - One goal per user (goal_type, target_calories)
- `recipes` - Recipe metadata and instructions
- `recipe_nutrition` - Nutrition facts (1:1 with recipes)
- `user_favorites` - Many-to-many junction table

## API Overview

**8 total endpoints:**
- 4 public: Register, Login, List Recipes, Get Recipe
- 4 protected: Get Profile, Update Goal, List Favorites, Add/Remove Favorite

See `docs/architecture.md` for complete API documentation with request/response formats.

## Database Schema

All tables use AUTO_INCREMENT primary keys. Key relationships:
- `user_goals.user_id` → `users.id` (one-to-one)
- `recipe_nutrition.recipe_id` → `recipes.id` (one-to-one)
- `user_favorites` → junction table (many-to-many between users and recipes)

See `docs/database-schema.md` for complete SQL CREATE statements and JSON data structures.

## Development Timeline

- **Week 1:** Authentication + recipe browsing
- **Week 2-3:** Favorites + search functionality
- **Week 4:** Deployment + polish

## Important Notes for AI Assistants

- This is a college project with a 3-4 week timeline - keep solutions simple and focused
- Avoid over-engineering - no need for advanced features not in the MVP scope
- Recipe data is manually entered, not fetched from external APIs
- Focus on completing the 5 core features before adding extras
- Security: Always hash passwords, validate JWT tokens, prevent SQL injection

## Documentation References

- **Project scope and features:** `docs/project-proposal.md`
- **Pages, navigation, API design:** `docs/architecture.md`
- **Complete database schema:** `docs/database-schema.md`

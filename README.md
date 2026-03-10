# FitMeals

A goal-oriented recipe discovery and nutrition tracking web application for people on their fitness journey.

## Overview

FitMeals helps users find recipes tailored to their specific fitness goals—whether they're trying to gain or lose weight. Instead of juggling multiple apps for recipe discovery and nutrition tracking, FitMeals provides a streamlined experience in one place.

## Problem Statement

People on a weight journey struggle to find recipes that align with their specific goals (weight gain vs. weight loss) and often need multiple apps to discover recipes and track nutrition. FitMeals solves this by providing goal-specific recipe browsing and simple tracking in a single application.

## Features

This MVP includes five core features:

1. **Browse Recipes by Goal** - Filter recipes by weight gain or weight loss objectives
2. **View Recipe Details** - See complete recipe information including ingredients, instructions, and nutrition facts
3. **Set Personal Goals** - Define your fitness goal (gain/loss) and target daily calories
4. **Save Favorite Recipes** - Bookmark recipes for quick access later (requires authentication)
5. **Search Recipes** - Find recipes by name using the search function

## Tech Stack

**Frontend:**
- React 18
- Vite (build tool and dev server)
- React Router (client-side routing)
- Axios (HTTP requests)

**Backend:**
- Go 1.21+
- Gin Web Framework
- JWT authentication
- bcrypt password hashing

**Database:**
- MySQL 8.0+

**Deployment:**
- AWS Lightsail
- nginx (reverse proxy)

## Installation

### Prerequisites

- Node.js 18+ and npm
- Go 1.21+
- MySQL 8.0+
- Git

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/recipe-fitness-app.git
cd recipe-fitness-app
```

### 2. Database Setup

```bash
# Log into MySQL
mysql -u root -p

# Create database
CREATE DATABASE fitmeals;
exit;
```

### 3. Backend Setup

```bash
cd backend

# Install Go dependencies
go mod download

# Create .env file with your database credentials
cat > .env << EOF
DB_USER=root
DB_PASSWORD=yourpassword
DB_HOST=localhost
DB_PORT=3306
DB_NAME=fitmeals
JWT_SECRET=your-secret-key-change-this
PORT=8080
EOF

# Run database migrations
go run database/migrate.go

# (Optional) Seed database with sample recipes
go run database/seed.go
```

### 4. Frontend Setup

```bash
cd ../frontend

# Install npm dependencies
npm install

# Create .env file
cat > .env << EOF
VITE_API_URL=http://localhost:8080/api
EOF
```

## Running Locally

### Start the Backend Server

```bash
cd backend
go run main.go
# Server runs on http://localhost:8080
```

### Start the Frontend Development Server

```bash
cd frontend
npm run dev
# Application runs on http://localhost:5173
```

Open your browser and navigate to `http://localhost:5173`

## Deployment

The application is deployed to AWS Lightsail with the following configuration:

- **Frontend:** Built with `npm run build` and served via nginx
- **Backend:** Go binary running as a systemd service
- **Database:** MySQL instance on the same Lightsail server
- **Reverse Proxy:** nginx configured to proxy `/api/*` requests to the Go backend
- **Deployment Path:** Application accessible at `/recipe-app/`

## Project Structure

```
recipe-fitness-app/
├── frontend/              # React application
│   ├── src/
│   │   ├── components/   # Reusable UI components
│   │   ├── pages/        # Page components
│   │   ├── services/     # API service layer
│   │   └── context/      # React context (auth state)
│   └── package.json
│
├── backend/               # Go API server
│   ├── main.go           # Application entry point
│   ├── handlers/         # HTTP request handlers
│   ├── models/           # Data models
│   ├── middleware/       # Authentication middleware
│   └── database/         # DB connection and migrations
│
├── docs/                  # Project documentation
│   ├── project-proposal.md
│   ├── architecture.md
│   └── database-schema.md
│
├── README.md             # This file
└── CLAUDE.md             # AI assistant context
```

## API Endpoints

The backend exposes 8 RESTful API endpoints:

**Public Endpoints:**
- `POST /api/auth/register` - Create new account
- `POST /api/auth/login` - Authenticate user
- `GET /api/recipes` - List recipes (with optional filters)
- `GET /api/recipes/:id` - Get recipe details

**Protected Endpoints (JWT Required):**
- `GET /api/user/profile` - Get user profile and goal
- `PUT /api/user/goal` - Update user goal
- `GET /api/favorites` - Get user's favorite recipes
- `POST /api/favorites/:recipeId` - Add recipe to favorites
- `DELETE /api/favorites/:recipeId` - Remove favorite

See `docs/architecture.md` for detailed API documentation.

## Database Schema

The application uses 5 MySQL tables:
- `users` - User accounts and authentication
- `user_goals` - User fitness goals and calorie targets
- `recipes` - Recipe data with ingredients (JSON) and instructions (JSON)
- `recipe_nutrition` - Nutrition facts per recipe
- `user_favorites` - User-recipe favorites (junction table)

See `docs/database-schema.md` for complete SQL schema and design decisions.

## Development Timeline

- **Week 1:** User authentication + recipe browsing
- **Week 2-3:** Favorites functionality + search
- **Week 4:** Deployment and final polish

## Author

**[Your Name]**
Web Development - Module 4 Project
[Your College/University]
[Semester/Year]

## License

This project is created for educational purposes as part of a college web development course.

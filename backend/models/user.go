package models

import (
	"database/sql"
	"fitmeals/database"
	"time"
)

// ─── Structs ─────────────────────────────────────────────────────────────────

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // never serialised to JSON
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserGoal struct {
	GoalType       string `json:"goal_type"`
	TargetCalories int    `json:"target_calories"`
}

// ─── Queries ──────────────────────────────────────────────────────────────────

// GetUserByEmail is used during login.
func GetUserByEmail(email string) (*User, error) {
	u := &User{}
	row := database.DB.QueryRow(
		`SELECT id, email, password_hash, name, created_at FROM users WHERE email = ?`,
		email,
	)
	err := row.Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Name, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil // caller checks for nil → 401
	}
	return u, err
}

// CreateUser inserts a new user and returns the new ID.
func CreateUser(name, email, passwordHash string) (int, error) {
	res, err := database.DB.Exec(
		`INSERT INTO users (name, email, password_hash) VALUES (?, ?, ?)`,
		name, email, passwordHash,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return int(id), err
}

// GetUserByID returns a user row (without password_hash).
func GetUserByID(userID int) (*User, error) {
	u := &User{}
	row := database.DB.QueryRow(
		`SELECT id, email, name, created_at FROM users WHERE id = ?`,
		userID,
	)
	err := row.Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return u, err
}

// GetUserGoal returns the goal for a user, or nil if not set yet.
func GetUserGoal(userID int) (*UserGoal, error) {
	g := &UserGoal{}
	row := database.DB.QueryRow(
		`SELECT goal_type, target_calories FROM user_goals WHERE user_id = ?`,
		userID,
	)
	err := row.Scan(&g.GoalType, &g.TargetCalories)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return g, err
}

// UpsertUserGoal creates or replaces the user's goal.
func UpsertUserGoal(userID int, goalType string, targetCalories int) error {
	_, err := database.DB.Exec(
		`INSERT INTO user_goals (user_id, goal_type, target_calories)
		 VALUES (?, ?, ?)
		 ON DUPLICATE KEY UPDATE
		   goal_type       = VALUES(goal_type),
		   target_calories = VALUES(target_calories),
		   updated_at      = CURRENT_TIMESTAMP`,
		userID, goalType, targetCalories,
	)
	return err
}

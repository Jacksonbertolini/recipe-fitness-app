package models

import (
	"fitmeals/database"
)

// ListFavorites returns all favorited recipes (with nutrition) for a user.
func ListFavorites(userID int) ([]Recipe, error) {
	rows, err := database.DB.Query(`
		SELECT r.id, r.name, r.description, r.goal_type,
		       r.prep_time_minutes, r.cook_time_minutes, r.servings,
		       r.ingredients, r.instructions, COALESCE(r.image_url, ''),
		       r.created_at,
		       n.calories, n.protein_g, n.carbs_g, n.fats_g, n.fiber_g
		FROM user_favorites f
		JOIN recipes r         ON r.id = f.recipe_id
		LEFT JOIN recipe_nutrition n ON n.recipe_id = r.id
		WHERE f.user_id = ?
		ORDER BY f.created_at DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []Recipe
	for rows.Next() {
		r, n := Recipe{}, Nutrition{}
		err := rows.Scan(
			&r.ID, &r.Name, &r.Description, &r.GoalType,
			&r.PrepTimeMinutes, &r.CookTimeMinutes, &r.Servings,
			&r.Ingredients, &r.Instructions, &r.ImageURL, &r.CreatedAt,
			&n.Calories, &n.ProteinG, &n.CarbsG, &n.FatsG, &n.FiberG,
		)
		if err != nil {
			return nil, err
		}
		r.Nutrition = &n
		r.IsFavorited = true
		recipes = append(recipes, r)
	}
	return recipes, rows.Err()
}

// AddFavorite inserts a favorite row; silently ignores if already exists.
func AddFavorite(userID, recipeID int) error {
	_, err := database.DB.Exec(
		`INSERT IGNORE INTO user_favorites (user_id, recipe_id) VALUES (?, ?)`,
		userID, recipeID,
	)
	return err
}

// RemoveFavorite deletes the favorite row.
func RemoveFavorite(userID, recipeID int) error {
	_, err := database.DB.Exec(
		`DELETE FROM user_favorites WHERE user_id = ? AND recipe_id = ?`,
		userID, recipeID,
	)
	return err
}

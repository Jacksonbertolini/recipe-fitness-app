-- =============================================================
-- FitMeals Sample Recipes
-- Run after schema.sql: mysql -u fitmeals_user -pfitmeals_password fitmeals < seed.sql
-- =============================================================

USE fitmeals;

-- Clear existing seed data so this file is safe to re-run
DELETE FROM recipe_nutrition WHERE recipe_id <= 8;
DELETE FROM recipes WHERE id <= 8;

-- Reset auto-increment so IDs are predictable
ALTER TABLE recipes AUTO_INCREMENT = 1;

-- =============================================================
-- Recipes
-- =============================================================

INSERT INTO recipes (id, name, description, goal_type, prep_time_minutes, cook_time_minutes, servings, ingredients, instructions, image_url) VALUES

-- 1: weight_gain
(1, 'High-Protein Chicken & Rice Bowl',
 'A calorie-dense staple packed with lean protein and complex carbs. Great post-workout.',
 'weight_gain', 10, 25, 2,
 '[{"name":"chicken breast","amount":"300","unit":"g"},{"name":"brown rice","amount":"200","unit":"g"},{"name":"olive oil","amount":"2","unit":"tbsp"},{"name":"garlic powder","amount":"1","unit":"tsp"},{"name":"salt","amount":"0.5","unit":"tsp"},{"name":"black pepper","amount":"0.25","unit":"tsp"}]',
 '["Cook brown rice according to package instructions.","Season chicken with garlic powder, salt, and pepper.","Heat olive oil in a pan over medium-high heat.","Cook chicken 6-7 minutes per side until cooked through (165°F internal).","Slice chicken and serve over rice."]',
 NULL),

-- 2: weight_gain
(2, 'Peanut Butter Banana Mass Smoothie',
 'A quick high-calorie smoothie ideal for bulking. Blend it up between meals or after the gym.',
 'weight_gain', 5, 0, 1,
 '[{"name":"whole milk","amount":"350","unit":"ml"},{"name":"banana","amount":"2","unit":"whole"},{"name":"peanut butter","amount":"3","unit":"tbsp"},{"name":"Greek yogurt","amount":"150","unit":"g"},{"name":"honey","amount":"1","unit":"tbsp"},{"name":"rolled oats","amount":"40","unit":"g"}]',
 '["Add all ingredients to a blender.","Blend on high for 45-60 seconds until smooth.","Pour into a large glass and drink immediately."]',
 NULL),

-- 3: weight_gain
(3, 'Beef & Sweet Potato Power Bowl',
 'A hearty bowl of ground beef and sweet potato loaded with calories, iron, and complex carbs.',
 'weight_gain', 10, 20, 2,
 '[{"name":"lean ground beef (90%)","amount":"400","unit":"g"},{"name":"sweet potato","amount":"300","unit":"g"},{"name":"olive oil","amount":"1","unit":"tbsp"},{"name":"onion","amount":"1","unit":"medium"},{"name":"paprika","amount":"1","unit":"tsp"},{"name":"cumin","amount":"0.5","unit":"tsp"},{"name":"salt","amount":"0.5","unit":"tsp"}]',
 '["Peel and cube sweet potato into 2 cm pieces.","Toss sweet potato with half the olive oil and roast at 200°C (400°F) for 20 minutes.","While sweet potato roasts, dice onion and cook in remaining oil over medium heat for 3 minutes.","Add ground beef and break apart; cook until browned, about 8 minutes.","Season with paprika, cumin, and salt.","Serve beef mixture over roasted sweet potato."]',
 NULL),

-- 4: weight_gain
(4, 'Overnight Oats with Nuts & Fruit',
 'Prep the night before for an effortless high-calorie breakfast. Great for adding size.',
 'weight_gain', 10, 0, 1,
 '[{"name":"rolled oats","amount":"100","unit":"g"},{"name":"whole milk","amount":"250","unit":"ml"},{"name":"Greek yogurt","amount":"100","unit":"g"},{"name":"mixed nuts","amount":"40","unit":"g"},{"name":"banana","amount":"1","unit":"whole"},{"name":"honey","amount":"1","unit":"tbsp"},{"name":"chia seeds","amount":"1","unit":"tbsp"}]',
 '["Combine oats, milk, yogurt, chia seeds, and honey in a jar or bowl.","Stir well, cover, and refrigerate overnight (at least 6 hours).","In the morning, top with sliced banana and mixed nuts.","Eat cold straight from the jar."]',
 NULL),

-- 5: weight_loss
(5, 'Lemon Herb Baked Salmon',
 'A light, protein-rich dinner that is low in carbs and high in omega-3 fatty acids.',
 'weight_loss', 10, 15, 2,
 '[{"name":"salmon fillets","amount":"300","unit":"g"},{"name":"lemon","amount":"1","unit":"whole"},{"name":"garlic","amount":"2","unit":"cloves"},{"name":"fresh dill","amount":"1","unit":"tbsp"},{"name":"olive oil","amount":"1","unit":"tsp"},{"name":"salt","amount":"0.25","unit":"tsp"},{"name":"black pepper","amount":"0.25","unit":"tsp"}]',
 '["Preheat oven to 200°C (400°F).","Place salmon fillets on a lined baking sheet.","Mince garlic and mix with olive oil, dill, salt, and pepper.","Spread herb mixture over salmon and top with lemon slices.","Bake 12-15 minutes until salmon flakes easily with a fork."]',
 NULL),

-- 6: weight_loss
(6, 'Greek Chicken Salad',
 'A fresh, filling salad with lean grilled chicken. High protein and very low in calories.',
 'weight_loss', 15, 12, 1,
 '[{"name":"chicken breast","amount":"150","unit":"g"},{"name":"romaine lettuce","amount":"100","unit":"g"},{"name":"cucumber","amount":"0.5","unit":"whole"},{"name":"cherry tomatoes","amount":"100","unit":"g"},{"name":"red onion","amount":"0.25","unit":"whole"},{"name":"feta cheese","amount":"30","unit":"g"},{"name":"lemon juice","amount":"2","unit":"tbsp"},{"name":"olive oil","amount":"1","unit":"tsp"}]',
 '["Season chicken with salt and pepper and grill or pan-cook over medium-high heat for 6 minutes per side.","Let chicken rest 5 minutes, then slice.","Chop lettuce, dice cucumber, halve tomatoes, and thinly slice red onion.","Combine vegetables in a large bowl.","Whisk lemon juice and olive oil together and drizzle over salad.","Top with sliced chicken and crumbled feta."]',
 NULL),

-- 7: weight_loss
(7, 'Egg White Veggie Scramble',
 'A low-calorie, high-protein breakfast that keeps you full without the extra fat from yolks.',
 'weight_loss', 5, 8, 1,
 '[{"name":"egg whites","amount":"5","unit":"whole"},{"name":"spinach","amount":"60","unit":"g"},{"name":"mushrooms","amount":"80","unit":"g"},{"name":"red bell pepper","amount":"0.5","unit":"whole"},{"name":"olive oil spray","amount":"1","unit":"spray"},{"name":"salt","amount":"0.25","unit":"tsp"},{"name":"black pepper","amount":"0.25","unit":"tsp"}]',
 '["Slice mushrooms and dice bell pepper.","Spray a non-stick pan with olive oil and heat over medium.","Cook mushrooms and bell pepper for 3 minutes until softened.","Add spinach and cook 1 minute until wilted.","Pour in egg whites, season with salt and pepper.","Stir gently and cook until egg whites are just set, about 2-3 minutes."]',
 NULL),

-- 8: weight_loss
(8, 'Turkey Lettuce Wraps',
 'Light, crunchy wraps using lettuce instead of tortillas — big on flavour, low on calories.',
 'weight_loss', 10, 10, 2,
 '[{"name":"lean ground turkey","amount":"300","unit":"g"},{"name":"butter lettuce","amount":"8","unit":"leaves"},{"name":"garlic","amount":"2","unit":"cloves"},{"name":"ginger","amount":"1","unit":"tsp"},{"name":"low-sodium soy sauce","amount":"2","unit":"tbsp"},{"name":"rice vinegar","amount":"1","unit":"tbsp"},{"name":"green onions","amount":"2","unit":"stalks"},{"name":"sesame oil","amount":"0.5","unit":"tsp"}]',
 '["Mince garlic and grate ginger.","Cook ground turkey in a pan over medium-high heat, breaking it apart, until browned (about 7 minutes).","Add garlic and ginger; cook 1 minute.","Stir in soy sauce, rice vinegar, and sesame oil; cook 1 minute more.","Spoon turkey mixture into lettuce leaves.","Top with sliced green onions and serve immediately."]',
 NULL);

-- =============================================================
-- Nutrition (one row per recipe)
-- =============================================================

INSERT INTO recipe_nutrition (recipe_id, calories, protein_g, carbs_g, fats_g, fiber_g) VALUES
(1, 650, 55.0, 72.0, 14.0, 3.5),   -- Chicken & Rice Bowl
(2, 720, 28.0, 85.0, 30.0, 5.0),   -- PB Banana Smoothie
(3, 710, 52.0, 48.0, 30.0, 6.0),   -- Beef & Sweet Potato
(4, 680, 24.0, 88.0, 26.0, 8.0),   -- Overnight Oats
(5, 340, 38.0,  4.0, 19.0, 0.5),   -- Baked Salmon
(6, 310, 36.0, 11.0, 13.0, 2.5),   -- Greek Chicken Salad
(7, 160, 24.0,  8.0,  2.0, 2.0),   -- Egg White Scramble
(8, 280, 34.0, 10.0,  9.0, 1.5);   -- Turkey Lettuce Wraps

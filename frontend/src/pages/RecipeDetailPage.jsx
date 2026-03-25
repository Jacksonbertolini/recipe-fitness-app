import { useState, useEffect } from 'react'
import { useParams, Link, useNavigate } from 'react-router-dom'
import { recipesAPI } from '../services/api'

const s = {
  page: {
    maxWidth: '800px',
    margin: '0 auto',
    padding: '2rem 1.5rem',
  },
  back: {
    display: 'inline-flex',
    alignItems: 'center',
    gap: '0.4rem',
    color: '#94a3b8',
    fontSize: '0.875rem',
    marginBottom: '1.5rem',
    cursor: 'pointer',
  },
  badge: (goal) => ({
    display: 'inline-block',
    background: goal === 'weight_gain' ? '#166534' : '#1e3a5f',
    color: goal === 'weight_gain' ? '#86efac' : '#93c5fd',
    borderRadius: '9999px',
    padding: '0.25rem 0.85rem',
    fontSize: '0.8rem',
    fontWeight: '600',
    marginBottom: '1rem',
  }),
  title: {
    fontSize: '2rem',
    fontWeight: '700',
    marginBottom: '0.75rem',
    lineHeight: '1.2',
  },
  description: {
    color: '#94a3b8',
    lineHeight: '1.7',
    marginBottom: '2rem',
  },
  nutritionGrid: {
    display: 'grid',
    gridTemplateColumns: 'repeat(4, 1fr)',
    gap: '1rem',
    background: '#1e293b',
    border: '1px solid #334155',
    borderRadius: '12px',
    padding: '1.5rem',
    marginBottom: '2rem',
  },
  nutriItem: {
    textAlign: 'center',
  },
  nutriValue: {
    fontSize: '1.5rem',
    fontWeight: '700',
    color: '#22c55e',
  },
  nutriLabel: {
    fontSize: '0.75rem',
    color: '#64748b',
    textTransform: 'uppercase',
    letterSpacing: '0.05em',
    marginTop: '0.25rem',
  },
  section: {
    background: '#1e293b',
    border: '1px solid #334155',
    borderRadius: '12px',
    padding: '1.5rem',
    marginBottom: '1.5rem',
  },
  sectionTitle: {
    fontSize: '1.1rem',
    fontWeight: '600',
    marginBottom: '1rem',
    color: '#22c55e',
  },
  ingredientList: {
    listStyle: 'none',
    display: 'flex',
    flexDirection: 'column',
    gap: '0.5rem',
  },
  ingredientItem: {
    display: 'flex',
    gap: '0.5rem',
    alignItems: 'baseline',
    color: '#cbd5e1',
    fontSize: '0.95rem',
  },
  ingredientAmount: {
    color: '#94a3b8',
    fontSize: '0.875rem',
    whiteSpace: 'nowrap',
  },
  stepList: {
    listStyle: 'none',
    display: 'flex',
    flexDirection: 'column',
    gap: '1rem',
  },
  stepItem: {
    display: 'flex',
    gap: '1rem',
    alignItems: 'flex-start',
  },
  stepNumber: {
    background: '#22c55e',
    color: '#0f172a',
    borderRadius: '50%',
    width: '28px',
    height: '28px',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    fontWeight: '700',
    fontSize: '0.8rem',
    flexShrink: 0,
    marginTop: '1px',
  },
  stepText: {
    color: '#cbd5e1',
    lineHeight: '1.6',
  },
  loading: {
    color: '#64748b',
    textAlign: 'center',
    padding: '4rem',
  },
  error: {
    background: '#450a0a',
    border: '1px solid #7f1d1d',
    color: '#fca5a5',
    borderRadius: '8px',
    padding: '1rem',
  },
}

function parseJSON(val) {
  if (!val) return []
  if (Array.isArray(val)) return val
  try { return JSON.parse(val) } catch { return [] }
}

function goalLabel(goal) {
  return goal === 'weight_gain' ? 'Weight Gain' : 'Weight Loss'
}

export default function RecipeDetailPage() {
  const { id } = useParams()
  const navigate = useNavigate()
  const [recipe, setRecipe] = useState(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    recipesAPI.get(id)
      .then((res) => setRecipe(res.data.recipe || res.data))
      .catch(() => setError('Recipe not found.'))
      .finally(() => setLoading(false))
  }, [id])

  if (loading) return <div style={s.loading}>Loading recipe...</div>
  if (error) return <div style={{ ...s.page }}><div style={s.error}>{error}</div></div>
  if (!recipe) return null

  const ingredients = parseJSON(recipe.ingredients)
  const instructions = parseJSON(recipe.instructions)
  const nutrition = recipe.nutrition || {}

  return (
    <div style={s.page}>
      <div style={s.back} onClick={() => navigate(-1)}>← Back to recipes</div>

      <span style={s.badge(recipe.goal_type)}>{goalLabel(recipe.goal_type)}</span>
      <h1 style={s.title}>{recipe.name}</h1>
      <p style={s.description}>{recipe.description}</p>

      {(nutrition.calories || nutrition.protein_g) && (
        <div style={s.nutritionGrid}>
          <div style={s.nutriItem}>
            <div style={s.nutriValue}>{nutrition.calories ?? '—'}</div>
            <div style={s.nutriLabel}>Calories</div>
          </div>
          <div style={s.nutriItem}>
            <div style={s.nutriValue}>{nutrition.protein_g != null ? `${nutrition.protein_g}g` : '—'}</div>
            <div style={s.nutriLabel}>Protein</div>
          </div>
          <div style={s.nutriItem}>
            <div style={s.nutriValue}>{nutrition.carbs_g != null ? `${nutrition.carbs_g}g` : '—'}</div>
            <div style={s.nutriLabel}>Carbs</div>
          </div>
          <div style={s.nutriItem}>
            <div style={s.nutriValue}>{nutrition.fat_g != null ? `${nutrition.fat_g}g` : '—'}</div>
            <div style={s.nutriLabel}>Fat</div>
          </div>
        </div>
      )}

      {ingredients.length > 0 && (
        <div style={s.section}>
          <div style={s.sectionTitle}>Ingredients</div>
          <ul style={s.ingredientList}>
            {ingredients.map((ing, i) => (
              <li key={i} style={s.ingredientItem}>
                {typeof ing === 'object' ? (
                  <>
                    <span style={s.ingredientAmount}>{ing.amount} {ing.unit}</span>
                    <span>{ing.name}</span>
                  </>
                ) : (
                  <span>{ing}</span>
                )}
              </li>
            ))}
          </ul>
        </div>
      )}

      {instructions.length > 0 && (
        <div style={s.section}>
          <div style={s.sectionTitle}>Instructions</div>
          <ol style={s.stepList}>
            {instructions.map((step, i) => (
              <li key={i} style={s.stepItem}>
                <div style={s.stepNumber}>{i + 1}</div>
                <div style={s.stepText}>{typeof step === 'object' ? step.text || step.step || JSON.stringify(step) : step}</div>
              </li>
            ))}
          </ol>
        </div>
      )}
    </div>
  )
}

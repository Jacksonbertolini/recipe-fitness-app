import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { recipesAPI } from '../services/api'

const s = {
  page: {
    maxWidth: '1100px',
    margin: '0 auto',
    padding: '2rem 1.5rem',
  },
  header: {
    marginBottom: '2rem',
  },
  title: {
    fontSize: '2rem',
    fontWeight: '700',
    marginBottom: '0.5rem',
  },
  subtitle: {
    color: '#94a3b8',
  },
  filters: {
    display: 'flex',
    gap: '0.75rem',
    marginBottom: '2rem',
    flexWrap: 'wrap',
  },
  filterBtn: (active) => ({
    background: active ? '#22c55e' : '#1e293b',
    color: active ? '#0f172a' : '#94a3b8',
    border: `1px solid ${active ? '#22c55e' : '#334155'}`,
    borderRadius: '9999px',
    padding: '0.4rem 1.1rem',
    fontWeight: active ? '600' : '400',
    fontSize: '0.875rem',
    cursor: 'pointer',
    transition: 'all 0.15s',
  }),
  grid: {
    display: 'grid',
    gridTemplateColumns: 'repeat(auto-fill, minmax(280px, 1fr))',
    gap: '1.25rem',
  },
  card: {
    background: '#1e293b',
    border: '1px solid #334155',
    borderRadius: '12px',
    padding: '1.5rem',
    display: 'flex',
    flexDirection: 'column',
    gap: '0.75rem',
    transition: 'border-color 0.15s',
    textDecoration: 'none',
    color: 'inherit',
  },
  badge: (goal) => ({
    display: 'inline-block',
    background: goal === 'weight_gain' ? '#166534' : '#1e3a5f',
    color: goal === 'weight_gain' ? '#86efac' : '#93c5fd',
    borderRadius: '9999px',
    padding: '0.2rem 0.7rem',
    fontSize: '0.75rem',
    fontWeight: '600',
    alignSelf: 'flex-start',
  }),
  cardTitle: {
    fontSize: '1.1rem',
    fontWeight: '600',
  },
  cardDesc: {
    color: '#94a3b8',
    fontSize: '0.875rem',
    lineHeight: '1.5',
    display: '-webkit-box',
    WebkitLineClamp: 2,
    WebkitBoxOrient: 'vertical',
    overflow: 'hidden',
  },
  nutrition: {
    display: 'flex',
    gap: '1rem',
    marginTop: 'auto',
    paddingTop: '0.75rem',
    borderTop: '1px solid #334155',
  },
  nutriStat: {
    display: 'flex',
    flexDirection: 'column',
    gap: '0.15rem',
  },
  nutriLabel: {
    fontSize: '0.7rem',
    color: '#64748b',
    textTransform: 'uppercase',
    letterSpacing: '0.05em',
  },
  nutriValue: {
    fontSize: '0.9rem',
    fontWeight: '600',
    color: '#22c55e',
  },
  empty: {
    color: '#64748b',
    textAlign: 'center',
    padding: '4rem',
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
    marginBottom: '1rem',
  },
}

const FILTERS = [
  { label: 'All Recipes', value: '' },
  { label: 'Weight Gain', value: 'weight_gain' },
  { label: 'Weight Loss', value: 'weight_loss' },
]

function goalLabel(goal) {
  return goal === 'weight_gain' ? 'Weight Gain' : 'Weight Loss'
}

export default function RecipesPage() {
  const [recipes, setRecipes] = useState([])
  const [goalFilter, setGoalFilter] = useState('')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    setLoading(true)
    setError('')
    recipesAPI.list(goalFilter)
      .then((res) => setRecipes(res.data.recipes || res.data || []))
      .catch(() => setError('Failed to load recipes. Please try again.'))
      .finally(() => setLoading(false))
  }, [goalFilter])

  return (
    <div style={s.page}>
      <div style={s.header}>
        <h1 style={s.title}>Browse Recipes</h1>
        <p style={s.subtitle}>Goal-specific meals to fuel your fitness journey</p>
      </div>

      <div style={s.filters}>
        {FILTERS.map((f) => (
          <button
            key={f.value}
            style={s.filterBtn(goalFilter === f.value)}
            onClick={() => setGoalFilter(f.value)}
          >
            {f.label}
          </button>
        ))}
      </div>

      {error && <div style={s.error}>{error}</div>}

      {loading ? (
        <div style={s.loading}>Loading recipes...</div>
      ) : recipes.length === 0 ? (
        <div style={s.empty}>No recipes found.</div>
      ) : (
        <div style={s.grid}>
          {recipes.map((recipe) => (
            <Link key={recipe.id} to={`/recipes/${recipe.id}`} style={s.card}>
              <span style={s.badge(recipe.goal_type)}>{goalLabel(recipe.goal_type)}</span>
              <div style={s.cardTitle}>{recipe.name}</div>
              <div style={s.cardDesc}>{recipe.description}</div>
              {recipe.nutrition && (
                <div style={s.nutrition}>
                  <div style={s.nutriStat}>
                    <span style={s.nutriLabel}>Calories</span>
                    <span style={s.nutriValue}>{recipe.nutrition.calories}</span>
                  </div>
                  <div style={s.nutriStat}>
                    <span style={s.nutriLabel}>Protein</span>
                    <span style={s.nutriValue}>{recipe.nutrition.protein_g}g</span>
                  </div>
                  <div style={s.nutriStat}>
                    <span style={s.nutriLabel}>Carbs</span>
                    <span style={s.nutriValue}>{recipe.nutrition.carbs_g}g</span>
                  </div>
                  <div style={s.nutriStat}>
                    <span style={s.nutriLabel}>Fat</span>
                    <span style={s.nutriValue}>{recipe.nutrition.fat_g}g</span>
                  </div>
                </div>
              )}
            </Link>
          ))}
        </div>
      )}
    </div>
  )
}

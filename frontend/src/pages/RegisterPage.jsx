import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { authAPI } from '../services/api'
import { useAuth } from '../context/AuthContext'

const s = {
  page: {
    minHeight: 'calc(100vh - 60px)',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    padding: '2rem',
  },
  card: {
    background: '#1e293b',
    border: '1px solid #334155',
    borderRadius: '12px',
    padding: '2.5rem',
    width: '100%',
    maxWidth: '420px',
  },
  title: {
    fontSize: '1.75rem',
    fontWeight: '700',
    marginBottom: '0.5rem',
  },
  subtitle: {
    color: '#94a3b8',
    marginBottom: '2rem',
    fontSize: '0.9rem',
  },
  label: {
    display: 'block',
    color: '#94a3b8',
    fontSize: '0.875rem',
    marginBottom: '0.4rem',
  },
  input: {
    width: '100%',
    background: '#0f172a',
    border: '1px solid #334155',
    borderRadius: '8px',
    padding: '0.6rem 0.9rem',
    color: '#f1f5f9',
    fontSize: '0.95rem',
    outline: 'none',
    marginBottom: '1.25rem',
  },
  btn: {
    width: '100%',
    background: '#22c55e',
    color: '#0f172a',
    border: 'none',
    borderRadius: '8px',
    padding: '0.75rem',
    fontWeight: '700',
    fontSize: '1rem',
    cursor: 'pointer',
    marginTop: '0.5rem',
  },
  error: {
    background: '#450a0a',
    border: '1px solid #7f1d1d',
    color: '#fca5a5',
    borderRadius: '8px',
    padding: '0.75rem',
    fontSize: '0.875rem',
    marginBottom: '1rem',
  },
  footer: {
    marginTop: '1.5rem',
    textAlign: 'center',
    color: '#64748b',
    fontSize: '0.875rem',
  },
  footerLink: {
    color: '#22c55e',
    fontWeight: '600',
  },
}

export default function RegisterPage() {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)
  const { login } = useAuth()
  const navigate = useNavigate()

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError('')
    setLoading(true)
    try {
      const res = await authAPI.register(name, email, password)
      login(res.data.token)
      navigate('/')
    } catch (err) {
      setError(err.response?.data?.error || 'Registration failed. Please try again.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div style={s.page}>
      <div style={s.card}>
        <h1 style={s.title}>Create account</h1>
        <p style={s.subtitle}>Start your FitMeals journey today</p>
        {error && <div style={s.error}>{error}</div>}
        <form onSubmit={handleSubmit}>
          <label style={s.label}>Name</label>
          <input
            type="text"
            style={s.input}
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Your name"
            required
          />
          <label style={s.label}>Email</label>
          <input
            type="email"
            style={s.input}
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="you@example.com"
            required
          />
          <label style={s.label}>Password</label>
          <input
            type="password"
            style={s.input}
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="••••••••"
            required
            minLength={6}
          />
          <button type="submit" style={s.btn} disabled={loading}>
            {loading ? 'Creating account...' : 'Create Account'}
          </button>
        </form>
        <p style={s.footer}>
          Already have an account?{' '}
          <Link to="/login" style={s.footerLink}>Sign in</Link>
        </p>
      </div>
    </div>
  )
}

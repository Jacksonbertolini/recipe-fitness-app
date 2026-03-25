import { Link, useNavigate } from 'react-router-dom'
import { useAuth } from '../context/AuthContext'

const styles = {
  nav: {
    background: '#1e293b',
    borderBottom: '1px solid #334155',
    padding: '0 1.5rem',
    height: '60px',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  brand: {
    fontSize: '1.25rem',
    fontWeight: '700',
    color: '#22c55e',
  },
  links: {
    display: 'flex',
    gap: '1rem',
    alignItems: 'center',
  },
  link: {
    color: '#94a3b8',
    fontSize: '0.9rem',
    transition: 'color 0.15s',
  },
  btn: {
    background: '#22c55e',
    color: '#0f172a',
    border: 'none',
    borderRadius: '6px',
    padding: '0.4rem 1rem',
    fontWeight: '600',
    fontSize: '0.875rem',
    cursor: 'pointer',
  },
}

export default function Navbar() {
  const { isAuthenticated, logout } = useAuth()
  const navigate = useNavigate()

  const handleLogout = () => {
    logout()
    navigate('/login')
  }

  return (
    <nav style={styles.nav}>
      <Link to="/" style={styles.brand}>FitMeals</Link>
      <div style={styles.links}>
        <Link to="/" style={styles.link}>Recipes</Link>
        {isAuthenticated ? (
          <button style={styles.btn} onClick={handleLogout}>Logout</button>
        ) : (
          <>
            <Link to="/login" style={styles.link}>Login</Link>
            <Link to="/register">
              <button style={styles.btn}>Sign Up</button>
            </Link>
          </>
        )}
      </div>
    </nav>
  )
}

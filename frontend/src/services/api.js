import axios from 'axios'

const api = axios.create({
  baseURL: 'https://jackson-web-dev.duckdns.org/recipe-app/api',
})

api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

export const authAPI = {
  login: (email, password) => api.post('/auth/login', { email, password }),
  register: (name, email, password) => api.post('/auth/register', { name, email, password }),
}

export const recipesAPI = {
  list: (goalType) => api.get('/recipes', { params: goalType ? { goal: goalType } : {} }),
  get: (id) => api.get(`/recipes/${id}`),
}

export default api

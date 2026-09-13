/**
 * Senior Data Scientist.: Dr. Eddy Giusepe Chirinos Isidro
 *
 * main.jsx
 * ========
 * React entry point. Loads index.css, mounts App.jsx into #root,
 * and is the script referenced by index.html.
 */
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <App />
  </StrictMode>,
)

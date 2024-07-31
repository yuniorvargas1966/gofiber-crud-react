
import {BrowserRouter, Routes, Route} from "react-router-dom"
import { TalleresPage } from './pages/TalleresPage'
import { Navigation } from './components/Navigation'
import { TallerFormPage } from './pages/TallerFormPage'

function App() {
  return (
    <BrowserRouter>
    <Navigation />
      <Routes>  
      <Route path="/" element={<TalleresPage />} />
      <Route path="/crear" element={<TallerFormPage />} />
      </Routes>
      
    </BrowserRouter>
  
  )
}
export default App


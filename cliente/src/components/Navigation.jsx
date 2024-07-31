import { Link } from 'react-router-dom';

export function Navigation() {
    return (
        <div> 
            <Link to="/"><h1>Taller</h1></Link>
            <Link to="/crear">Crear Servicio</Link>
        </div>
    )
}

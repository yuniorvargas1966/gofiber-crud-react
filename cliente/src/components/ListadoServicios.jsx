/* eslint-disable react/jsx-key */
import { useEffect, useState} from 'react'
import { getAllServicios } from '../api/Servicios.api';

export function ListadoServicios() {
   
    const [servicios, setServicios] = useState([]);

    useEffect(() => {
        async function loadServicios() {
           const res = await getAllServicios()
           setServicios(res.data)
        }
        loadServicios()
    }, []);
       
    return  ( 
        <div>
            {servicios.map(servicio => (
                <div key={servicio.id}>
                    <h4>{servicio.data}</h4>       
                    </div>
            ))}
        </div>
    )
}

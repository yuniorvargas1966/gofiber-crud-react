import axios from "axios";

export const getAllServicios = () => {
    return axios.get("http://localhost:8080")
}
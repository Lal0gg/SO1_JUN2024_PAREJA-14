import axios from 'axios'


// api para la conexion con el backend

const instance = axios.create({
    baseURL: 'http://127.0.0.1:8080/',
});



export const GetInfoRam = async () => {
    const res = await instance.get('ram')
    console.log("Soy la respuesta de la api", res.data)
    return res
}

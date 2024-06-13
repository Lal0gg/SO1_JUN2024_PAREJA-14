import axios from 'axios'


// api para la conexion con el backend

const instance = axios.create({
    baseURL: 'http://127.0.0.1:8080/',
});



export const GetInfoRam = async () => {
    const res = await instance.get('ram')
    console.log("Soy la respuesta Obteniendo RamInfo", res.data)
    return res
}


export const GetInfoCpu = async () => {
    const res = await instance.get('cpu')
    console.log("Soy la respuesta Obteniendo CpuInfo", res.data)
    return res
}


export const CreateProcess = async () => {
    const res = await instance.post('/create-process')
    console.log("Soy la respuesta Creando Proceso", res.data)
    return res
}

export const KillProcess = async (pid) => {
    const res = await instance.post('kill-process', { pid })
    console.log("Soy la respuesta Eliminando Proceso", res.data)
    return res
}
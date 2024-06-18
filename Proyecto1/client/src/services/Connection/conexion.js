export const GetInfoRam = async () => {
    try {
        const response = await fetch(`/sopes1/ram`);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data = await response.json();
        console.log("Soy la respuesta Obteniendo RamInfo", data);
        return data;
    } catch (error) {
        console.error("Error al obtener información de RAM:", error);
        throw error;
    }
}

export const GetInfoCpu = async () => {
    try {
        const response = await fetch(`/sopes1/cpu`);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data = await response.json();
        console.log("Soy la respuesta Obteniendo CpuInfo", data);
        return data;
    } catch (error) {
        console.error("Error al obtener información de CPU:", error);
        throw error;
    }
}

export const CreateProcess = async () => {
    try {
        const response = await fetch(`/sopes1/create-process`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
        });
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data = await response.json();
        console.log("Soy la respuesta Creando Proceso", data);
        return data;
    } catch (error) {
        console.error("Error al crear proceso:", error);
        throw error;
    }
}

export const KillProcess = async (pid) => {
    try {
        const response = await fetch(`/sopes1/kill-process`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ pid })
        });
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data = await response.json();
        console.log("Soy la respuesta Eliminando Proceso", data);
        return data;
    } catch (error) {
        console.error("Error al eliminar proceso:", error);
        throw error;
    }
}
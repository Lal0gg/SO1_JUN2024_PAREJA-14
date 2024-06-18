# Universidad de San Carlos de Guatemala
# Ingenieria en Ciencias y Sistemas
# Ingeniero Sergio Arnaldo Méndez Aguilar
# Auxiliar Luis Leonel Aguilar Sánchez
# Sección A

# Proyecto 1
| Carnet | Nombre |
| ------ | -------  |
| 201900647 |Eduardo Josué González Cifuentes|
| 201902301 |Piter Angel Esaú Valiente de León|

## Objetivo del Proyecto
Implementar una plataforma integral de monitoreo de recursos del sistema y gestión de procesos en tiempo real, utilizando tecnologías y lenguajes de programación modernos, para proporcionar una interfaz amigable y eficiente que permita a los usuarios obtener y administrar información clave sobre el rendimiento del sistema y los procesos en ejecución, desplegada en un entorno de máquina virtual sin interfaz gráfica.

# Componentes Utilizados
El proyecto utiliza una combinación de tecnologías modernas y robustas para crear un sistema de monitoreo eficiente y de alto rendimiento, desplegado en un entorno virtualizado. Incluye el uso de módulos del kernel para obtener datos del sistema, contenedores para gestión y despliegue, programación asíncrona para eficiencia, y una interfaz web para facilitar la interacción del usuario.

Máquina Virtual:

    Sistema Operativo: Ubuntu Server 22.04
    Hipervisor: KVM (Kernel-based Virtual Machine)

Módulos del Kernel de Linux:

    Módulo de RAM:
        Archivo en /proc/ram_so1_jun2024
        Librería: <linux/mm.h>
    Módulo de CPU:
        Archivo en /proc/cpu_so1_1s2024
        Librería: <linux/sched.h>


Contenedores:

    Plataforma de Contenedores: Docker
    Base de Datos: MongoDB (con persistencia mediante Volumen de Docker)
    Gestión de Contenedores: Docker Compose
    Repositorio de Imágenes: Docker Hub

Frontend:

    Framework Web: React con Vite
    Características:
        Gráfica en tiempo real del uso de RAM (obtenida del módulo del kernel)
        Gráfica en tiempo real del uso de CPU (obtenida mediante mpstat)
        Tabla de procesos con detalles de procesos y subprocesos
        Botones para crear y eliminar procesos sleep infinity

Backend:

    API:
        Lenguaje: Golang
        Funciones:
            Llamadas a los módulos en /proc
            Almacenamiento de datos en MongoDB
            Envío de datos para gráficos en tiempo real
            Gestión de procesos (sleep infinity)

Base de Datos:

    Tipo: NoSQL
    Motor: MongoDB
    Despliegue: Docker container con persistencia

Pruebas:

    Pruebas de Stress: Usando el módulo de Linux para verificar el funcionamiento bajo carga


# Comando utilizados
Estos comandos proporcionan una guía básica para la instanciación y configuración de los componentes necesarios para la plataforma de monitoreo de procesos en un entorno Linux.

Máquina Virtual:

    Instalar KVM:

    bash

sudo apt update
sudo apt install qemu-kvm libvirt-daemon-system libvirt-clients bridge-utils virt-manager

Crear y gestionar la VM:

bash

    virt-manager

Módulos del Kernel de Linux:

    Crear módulo de RAM:

    bash

    cd /usr/src
    mkdir ram_module
    cd ram_module
    # Crear y editar el archivo ram_so1_jun2024.c
    # Compilar el módulo
    make -C /lib/modules/$(uname -r)/build M=$PWD modules
    # Cargar el módulo
    sudo insmod ram_so1_jun2024.ko

Crear módulo de CPU:

bash

    cd /usr/src
    mkdir cpu_module
    cd cpu_module
    # Crear y editar el archivo cpu_so1_1s2024.c
    # Compilar el módulo
    make -C /lib/modules/$(uname -r)/build M=$PWD modules
    # Cargar el módulo
    sudo insmod cpu_so1_1s2024.ko



Contenedores:

    Instalar Docker:

    bash

    sudo apt update
    sudo apt install docker.io
    sudo systemctl start docker
    sudo systemctl enable docker

Instalar Docker Compose:

    bash

    sudo apt install docker-compose

Frontend (React con Vite):

    Instalar Node.js y npm:

    bash

    curl -sL https://deb.nodesource.com/setup_16.x | sudo -E bash -
    sudo apt install -y nodejs

Crear proyecto React con Vite:

    bash

    npm create vite@latest my-react-app --template react
    cd my-react-app
    npm install
    npm run dev

Backend (API en Golang):

    Instalar Golang:

    bash

    sudo apt update
    sudo apt install golang

    Crear proyecto Golang:

    bash

    mkdir go_api
    cd go_api
    go mod init go_api
    # Crear y editar archivos fuente
    go build

Base de Datos (MongoDB en Docker):

    Ejecutar MongoDB en Docker:

    bash

    docker pull mongo
    docker run -d -p 27017:27017 --name mongodb -v mongo_data:/data/db mongo

Gestión de Contenedores con Docker Compose:

    Crear archivo docker-compose.yml:

    yaml

    version: '3'
    services:
    mongodb:
        image: mongo
        ports:
        - "27017:27017"
        volumes:
        - mongo_data:/data/db
    api:
        build: ./path_to_go_api
        ports:
        - "8080:8080"
        depends_on:
        - mongodb
    frontend:
        build: ./path_to_react_app
        ports:
        - "3000:3000"
    volumes:
    mongo_data:

Ejecutar Docker Compose:

bash

    docker-compose up

Pruebas de Stress:

    Instalar y usar stress-ng:

    bash

    sudo apt update
    sudo apt install stress-ng
    stress-ng --cpu 4 --io 2 --vm 2 --vm-bytes 128M --timeout 60s
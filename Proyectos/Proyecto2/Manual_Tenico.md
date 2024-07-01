# 🎓 Universidad de San Carlos de Guatemala
## 💻 Ingeniería en Ciencias y Sistemas
## 👨‍🏫 Ing. Sergio Arnaldo Méndez Aguilar
## 👨‍🏫 Aux. Daniel Velásquez
## 🏫 Sección A

# 📂 Proyecto 2

| 🎓 Carnet | 📛 Nombre |
| --------- | --------- |
| 201900647 | Eduardo Josué González Cifuentes |
| 201902301 | Piter Angel Esaú Valiente de León |

# Manual Técnico
## 📚 Contenido
1. [🎯 Objetivo](#-objetivo-del-proyecto)
2. [🚀 Arquitectura del Proyecto](#-arquitectura-del-proyecto)
3. [📃 Requerimientos](#)
    - [📍Tecnologias](#-tecnologías)
    - [📍Herramientas](#️-herramientas)
    - [🛠️Componentes Utilizados](#️-componentes-utilizados)

# 🎯 Objetivo del Proyecto

Diseñar y desarrollar una arquitectura de sistema distribuida genérica y escalable que pueda procesar y mostrar tweets sobre el clima de diferentes partes del mundo. Además, se busca medir y monitorear el consumo de energía y las emisiones de CO2 de las implementaciones, promoviendo la sostenibilidad ambiental.

# 👷🏻 Arquitectura del Proyecto
![alt text](images/image-7.png)

# 📍 Tecnologías

#### Estas son las tecnologías y herramientas utilizadas en el proyecto:


- **Docker:** 26.1.4
- **Cuenta en DockerHub**
- **Git:** 2.34.1
- **Go:** 1.18.1 
- **Ubuntu:** 22.04 LTS
- **MongoDB:** 7.0.11 
- **Locust**
- **Strimzi-kafka**
- **Kubernetes (kubectl)**
- **gRPC**
- **protobuf**
- **Helm**
- **Cuenta Google Cloud**

# 🛠️ Herramientas
- **Visual Studio Code:** 1.90.1
- **Postman**
- **GitKraken:** 10.0.2
- **MongoDB Compass:** 1.43.0
- **Navegador Web**
- **virtualenv**
- **Redis insight**

# 🛠️ Componentes Utilizados
El proyecto utiliza una combinación de tecnologías modernas y robustas para crear un sistema eficiente y de alto rendimiento. Incluye el uso de Kafka para mensajería, gRPC para comunicación, y Kubernetes para orquestación de contenedores.


### 🐳 Contenedores:

- **Plataforma de Contenedores**: Docker
- **Gestión de Contenedores**: Kubernetes (kubectl)
- **Repositorio de Imágenes**: Docker Hub

### 💬 Mensajería:

- **Kafka**: Desplegado usando Strimzi en Kubernetes

### 🌐 Comunicación:

- **gRPC**: Para comunicación eficiente entre servicios
- **protobuf**: Para definir y compilar archivos proto

### ⚙️ Orquestación:

- **Kubernetes**: Para desplegar y gestionar contenedores
- **Helm**: Para gestionar paquetes de Kubernetes

### 🧪 Pruebas:

- **Locust**: Para pruebas de carga



# 🔧 Comandos Utilizados

Estos comandos proporcionan una guía básica para la instalación, configuración y despliegue de los componentes necesarios en el entorno de desarrollo y producción.


### 📦 Instalación de Strimzi

```bash
kubectl create namespace kafka
kubectl create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka
kubectl get pod -n kafka --watch
```

### 🚀 Deploy Kafka usando Strimzi
```
kubectl apply -f https://strimzi.io/examples/latest/kafka/kafka-persistent-single.yaml -n kafka 

```

### 📡 Iniciar gRPC


```
sudo apt install protobuf-compiler
sudo apt-get install golang-goprotobuf-dev

```

### 📚 Dependencias Librerías para empezar con gRPC


```
go get github.com/gofiber/fiber/v2
go get google.golang.org/grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### 🛠️ Compilar archivo proto
```
protoc --go_out=. --go-grpc_out=. client.proto

```


### 🐍 Instalar Locust
```
sudo apt install python3-locust
```


### 🌐 Instalar y crear un entorno virtual para Locust
```
pip3 install virtualenv
virtualenv env1
source env1/bin/activate
```

### 📊 Usar Locust
```
locust -f traffic.py
```

### ☁️ Crear Cluster
```
gcloud init
# seguir las instrucciones del asistente
gcloud config set compute/zone us-central1-a
gcloud container clusters create sopes-proyecto --num-nodes 2 --machine-type n1-standard-2 --zone us-central1-a
```

### 🐋 Crear Contenedor en K8S
```
docker build -t gcr.io/id_proyecto/nombre_imagen .
docker push gcr.io/id_proyecto/nombre_imagen
kubectl create namespace nombre_namespace
kubectl create deployment nombre_deployment --image=gcr.io/id_proyecto/nombre_imagen -n=nombre_namespace
kubectl expose deployment nombre_deployment --type=LoadBalancer --port puerto -n=nombre_namespace
```


### 📋 Ver logs de los pods y contenedores
```
kubectl logs [-f] grpc-producer-745788bbbd-w8vhn [--container grpc-server] -n so1jun2024
```

### 🐳 Creación de los Dockerfile y despliegue
```
#Crear imagen:
docker build -t lalogg/p2_so1_p14_redis_rust:11.0.0 .

#Subirla a Docker Hub:
docker push lalogg/p2_so1_p14_redis_rust:11.0.0
```

### 📦 Despliegue de archivos YAML en K8S 
```
kubectl apply -f <file.yml>
```

### 🌐 Servicio de NGINX
```
#Crear namespace:
kubectl create ns nginx-ingress

#gregar repositorio a nuestro namespace:
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx

#Actualizar repositorio:
helm repo update

#Instalar NGINX Ingress:
helm install nginx-ingress ingress-nginx/ingress-nginx -n nginx-ingress
```



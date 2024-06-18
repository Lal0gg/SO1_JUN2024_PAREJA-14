# Desde que lenguaje o framework se va a construir la imagen
FROM node:20-alpine as builder

#Creamos un directorio de trabajo
WORKDIR /client

#Copiar los archivos necesarios para la aplicacion
COPY package.json ./
COPY package-lock.json ./

# Ejecutar el comando para obtener las librerias necesarias
RUN npm install

# Copiar el resto de los archivos
COPY . ./

# Exponer el puerto 80
#EXPOSE 5173

# Ejecutar el comando para iniciar la aplicacion
#CMD ["npm", "run","dev"]

# Construir la aplicacion
FROM nginx:1.21-alpine as final

# Copiar los archivos de la aplicacion
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf

# Copiar los archivos de la aplicacion
COPY  --from=builder /client/dist /usr/share/nginx/html

# Exponer el puerto 80
EXPOSE 80

# Iniciar el servidor
CMD ["nginx", "-g", "daemon off;"]

# Description: Archivo de configuracion para la creacion de la imagen de docker
FROM golang:alpine AS builder

# Creamos un directorio de trabajo
WORKDIR /go/src/app

# Copiar los archivos necesarios para la aplicacion
COPY . ./

# Ejecutar el comando para obtener las librerias necesarias
RUN go get -d -v
RUN go build -o /go/bin/app

# Variables de entorno
# ENV DB_HOST="localhost"
# ENV DB_PORT="27017"
# ENV DB_NAME="DB"

# Exponer el puerto 3000
#EXPOSE 3000

# Ejecutar el comando para iniciar la aplicacion
#CMD ["go", "run", "main.go"]

# Construir la aplicacion
FROM alpine:latest

# Instalar las librerias necesarias
RUN apk --no-cache add ca-certificates sysstat

# Copiar la aplicacion
WORKDIR /app
COPY --from=builder /go/bin/app /app
COPY .env /app/.env

# Ejecutar la aplicacion
ENTRYPOINT ["/app/app"]

EXPOSE 3000


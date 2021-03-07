# Descripción

Proyecto api de contacto

## Obtener todos los paquetes

go get -d -u github.com/henryinfanteg/heyadivinen_backgo/api-contacto/...

## Crear Modulo

export GO111MODULE=auto
export GO111MODULE=on

go mod init github.com/henryinfanteg/heyadivinen_backgo/api-contacto

## Docker

docker build -t api-contacto .

docker build --no-cache -t api-contacto .

docker run -it api-contacto bash

### Docker Linux

docker run -it --name api-contacto -p 3001:3001 api-contacto

### Docker Windows

winpty docker run -it --name api-contacto -p 3001:3001 api-contacto

### Docker con enviroment

#### En Linux

docker run -it --name api-contacto -p 3001:3001 -v ${MAXADIVINA_CONFIG_PATH}:/config -e MAXADIVINA_CONFIG_PATH=/config api-contacto
docker run -it --name api-contacto -p 3001:3001 -v /home/ubuntu/maxadivina/_CONFIG_FILES:/config -e MAXADIVINA_CONFIG_PATH=/config api-contacto

#### En Windows

docker run -it --name api-contacto -p 3001:3001 -v "%MAXADIVINA_CONFIG_PATH%":/config -e MAXADIVINA_CONFIG_PATH=/config api-contacto
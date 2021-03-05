# Descripción

Proyecto api de palabras

## Obtener todos los paquetes

go get -d -u gitlab.com/adivinagame/backend/maxadivinabackend/api-palabras/...

## Crear Modulo

export GO111MODULE=auto
export GO111MODULE=on

go mod init gitlab.com/adivinagame/backend/maxadivinabackend/api-palabras

## Docker

docker build -t api-palabras .

docker build --no-cache -t api-palabras .

docker run -it api-palabras bash

### Docker Linux

docker run -it --name api-palabras -p 3000:3000 api-palabras

### Docker Windows

winpty docker run -it --name api-palabras -p 3000:3000 api-palabras

### Docker con enviroment

#### En Linux

docker run --net=host -it --name api-palabras -p 3000:3000 -v ${MAXADIVINA_CONFIG_PATH}:/config -e MAXADIVINA_CONFIG_PATH=/config api-palabras
docker run --net=host -it --name api-palabras -p 3000:3000 -v /home/ec2-user/adivina-con-max/_CONFIG_FILES:/config -e MAXADIVINA_CONFIG_PATH=/config api-palabras

#### En Windows

docker run -it --name api-palabras -p 3000:3000 -v "%MAXADIVINA_CONFIG_PATH%":/config -e MAXADIVINA_CONFIG_PATH=/config api-palabras

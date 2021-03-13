# Descripción

Proyecto api de palabras

## Obtener todos los paquetes

go get -d -u github.com/henryinfanteg/heyadivinen_backgo/api-words/...

## Crear Modulo

export GO111MODULE=auto
export GO111MODULE=on

go mod init github.com/henryinfanteg/heyadivinen_backgo/api-words

## Docker

docker build -t api-words .

docker build --no-cache -t api-words .

docker run -it api-words bash

### Docker Linux

docker run -it --name api-words -p 3000:3000 api-words

### Docker Windows

winpty docker run -it --name api-words -p 3000:3000 api-words

### Docker con enviroment

#### En Linux

docker run --net=host -it --name api-words -p 3000:3000 -v ${MAXADIVINA_CONFIG_PATH}:/config -e MAXADIVINA_CONFIG_PATH=/config api-words
docker run --net=host -it --name api-words -p 3000:3000 -v /home/ec2-user/adivina-con-max/_CONFIG_FILES:/config -e MAXADIVINA_CONFIG_PATH=/config api-words

#### En Windows

docker run -it --name api-words -p 3000:3000 -v "%MAXADIVINA_CONFIG_PATH%":/config -e MAXADIVINA_CONFIG_PATH=/config api-words

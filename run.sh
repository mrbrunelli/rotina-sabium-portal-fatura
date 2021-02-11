#!/bin/sh

echo "---> Gerando build..."

go build main.go

echo "---> Build gerado com sucesso!"

chmod +x main*

echo "---> Inicializando servidor http..."

sleep 2s

echo -e "---> $ curl http://localhost:3000/ \n"

./main*



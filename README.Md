# fri-go

## lancer le projet

lancer redis-server

go build cmd/pubsub/pubsub.go  
go build cmd/server/server.go

./pubsub  
./server

cd front  
npm install  
gatsby develop

## Structure des fichiers

- cmd : launchers de l'application, importe des méthodes de internal

- config : fichiers de config

- data_lake : dossier contenant les fichiers csv constituant un data lake

- front : projet Gatsby contenant le front-end

- internal : code métier

- types : structures de données

# fri-go

## Contributeurs

- Adrien Mazet
- Charlotte Leclère
- Isuru Haupe
- Clémentine Biais

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

## Documentation de l'API

[Swagger fri-go](https://app.swaggerhub.com/apis-docs/f393/weather-reports_api/1.0.0?fbclid=IwAR12l7smon1b4lH5xQ1gKTxNN0P-Pr_qe3pP8dx83P5kgp21OFviog7Dw_c)

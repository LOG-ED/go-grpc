### Docente: Federico Minzoni  
- fminzoni@enter.eu  
- https://github.com/f-minzoni  

Sono un programmatore e smanettone della prima ora. Appassionato da sempre di Git, MongoDB, Rich Internet Application, IOT, Cloud Computing, da diversi mesi ho aggiunto alla lista, Docker e le architetture a Microservizi. Due temi che, insieme, stanno rivoluzionando l'intero processo di sviluppo e rilascio delle applicazioni sul Cloud. Ambito in cui lavoro da 3 anni, in Enter, occupandomi della piattaforma [Enter Cloud Suite](http://www.entercloudsuite.com).

### Avvertenze per lo svolgimento e l'utilizzo di questo corso

Questa tipologia di corso, definita e organizzata da [LOG.ED](https://loged.it) è strutturata in modo tale che tutti i materiali e le esercitazioni siano pubblici e accessibili liberamente. Si prega di rispettare le condizioni di utilizzo e/o licenza di qualsiasi codice qui presente. Per ogni algoritmo o codice che si desidera riutilizzare, si richiede di indicare nei crediti la fonte originale, con un commento in linea.

### Usage
 
1. Run the following command to start the grpc services (server and http proxy):   
   `$ docker-compose up -d proxy`   

2. Test the grpc server running the related client, with the following command:   
   `$ docker-compose run client`  

3. Test the grpc proxy running the following HTTP request:   
    `curl -v -X POST http://localhost:9000/v1/task/run -d '{"method":"GET"}' `       

4. Use the **docker-compose.dev.yml** file instead of the defaut one if you want to make changes in a dev environment:   
    `$ docker-compose -f docker-compose.dev.yml run srv bash`  

5. If changes are made to the Protocol Buffer file use the following command (from the proto dir) to regenerate the service        stub:  
   `protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. task.proto`    
Use the following command to regenerate the stub for the grpc proxy:   
   `protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. task.proto`
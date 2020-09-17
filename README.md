# Overview
1. There are four gRPC microservices: ```productpage```, ```ratings```, ```reviews``` and ```details```.
2. One version each of ```productpage```, ```ratings``` and ```details``` microservice is built using docker and pushed to docker registry.
3. Three versions of ```reviews``` microservice is build using docker and pushed to docker registry.
4. ```productpage``` microservice calls ```details``` microservice to get ```Price``` and ```Genre``` of the book.
5. ```productpage``` microservice calls ```reviews``` microservice to get ```Review``` and ```Rating``` of the book.
6. ```reviews v1``` microservice does not call ```ratings``` microservice. It returns 0 rating.
7. ```reviews v2``` microservice calls ```ratings``` microservice. It returns rating from 1 to 5.
8. ```reviews v3``` microservice calls ```ratings``` microservice. It returns rating from 6 to 10.

# SETUP
1. There are four gRPC microservices: ```productpage```, ```ratings```, ```reviews``` and ```details```.
2. Go to each microservice and run build.sh to create docker build image and push them in docker registry.
3. For ```reviews``` service, build three docker images with versions v1, v2 and v3
4. Create local minikube cluster using ```istio-kubernetes/scripts/setup.sh```
5. Run sample client in ```productpage/sample-client/sampleClient.go``` to send request to product page through istio-ingressgateway
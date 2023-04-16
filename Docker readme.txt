Docker readme
1. create .dockerigonre, that contains:
-Not to include the package folder(s) in docker container
-<package folder name/> ie handlers/
-<executable file name> ie *.exe

2. Create testing docker file:
- also start the mini cube  cluster at the same time at terminal. ie
$ minikube status
- configure docker CLi to use the docker engine running inside the mini queue.
- then just used docker PS it is going to get all containers information in mini cube cluster.
- So the "from" will use Golang and use the latest as it tag: ie 
FROM golang:latest as builder
- Next add working directoty, ie
WORKDIR /app
- Download all the dependencies, first one is go.mod and go.sum inside the current (working) directory, ie
COPY go.mod go.sum ./
- then run go mod download all the dependencies used in the GO project. ie
RUN go mod download
- copy the rest of the source code in current working directory using copy. ie
COPY . .
- Build the go project and save it as main in the current working directory. ie
RUN GO build -o main .
- Run command to expose 80. ie
EXPOSE 80.
- Run the final executable at entry point. ie
ENTRYPOINT [ "./main" ]
- Finally, the test docker file would be looking like:
FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GO build -o main .
EXPOSE 80
ENTRYPOINT [ "./main" ]
- using DOCKER command for full DOCKET container image build and the tag stick to the latest. ie
docker build -t go.app-normal:latest .
- Once the docker container image it ready, then run the docker container using docker run. ie
docker run -d -p 80:80 --name web app go-app-normal:latest
ps: web is the name of the container
- Kick start Go web application so it start listening to request. ie
docker logs -f web
- Get the IP of the minikube. ie
minikube ip
- we have one problem with the go build, the size problem!
- Now we remove docker container by rm command ie
docker rm -f web
- check the Go build application file size. ie
docker image ls
- this bring us to multistage builds to further reduce the Go executable size. 


3. Create production docker file, which is multistage bills: 
- check the google container distroless image. ie
# Start by building the application.
FROM golang:1.17-bullseye as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/app /
CMD [ "/app" ]

- Now it is time to distroless the docker image. ie
# Name it as a builder stage
FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GO build -o main .
# Do not put the entry point and expose port 80 here
# EXPOSE 80
# ENTRYPOINT [ "./main" ]

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian11
# using COPY command from state: builder and the location inside the builder. PS: use "." to copy everything into the current stage
COPY --from=builder /app/main . /
# Expose port: 80 here, run CMD instead ENTRYPOINT
EXPOSE 80
CMD [ "/main" ]

- Showtime, ps: ms=multiseries
docker build -t go-app-ms:latest .
- check the docker image. ie docker image ls
docker run -d -p 80:80 --name web app go-app-ms:latest
docker logs -f web

- create docker compose file named docker-compose.yml
PS: take a reference from the web for quick template at docs.docker.com/compose/

﻿
docker-compose.yml
version: "3.9" # optional since v1.27.0
services:
	web:
		build:
			context: .
			dockerfile: Dockerfile
		image: go-app-ms:latest
		ports:
			- "80:80"
		restart: always
		networks: 
			- web
	networks:
		web: 
----------------------------------------The following are not needed
		volumes:
			-.:/code
			- logvolume01:/var/log
		links:
			- redis
	redis:
		image: redis
volumes:
	logvolume01: {}
---------------------------------------The following are not needed

- build the docker container image, ie
docker-compose build
- once it is build and tagged, run the docker container, as a result networking is created and it will put the container inside that network, ie
docker-compose up -d
- run docker PS to get the container service status in docker-compose, ie
docker-compose ps
- bring down the docker complete stage for kubernetes deployment, ie
docker-compose down
- create deployment.yml file in kubernetes folder, also create the service.yml file. get the deployment in kubernetes for reference. ie 
﻿deployment.yml:
apiVersion: apps/v1 
kind: Deployment 
metadata:
	name: web-deployment 
	labels:
		app: web
spec:
	replicas: 3
	selector: 
		matchLabels:
			app: web
	template: 
		metadata: 
			Labels:
				app: web
		spec:
			containers:
			- name: go-web-app
			  image: go-web-ms:latest
			  imagePullPolicy: IfNotPresent
			  ports:
			  - containerPort: 80

- then prepare for NodePort service. ie
service.yml

﻿apiVersion: v1
kind: Service 
metadata:
	name: web-service
spec:
	type: NodePort
	selector:
		app: web
	ports:
		#By default and for convenience, the 'targetPort is set to the same value as the 'port' field.
		- port: 80
		  targetPort: 80
		  #Optional field
		  #By default and for convenience, the Kubernetes control plane will allocate a port from a range (default: 30000-32767) nodePort: 30007

- apply the deployment in complete kubernetes folder
kubectl apply -f kubernetes/
- list out all the kubernetes services, ie
kubectl get,svc
- delete all the created kubernetes services. ie
kubectl delete deploy --all

Load balancing:
- create a new port named curl running on nginx, ie
kubectl run -it curl --rm --image=nginx:alpine sh
PS: remove image whenever exit the curl port image
- the intention is call the internally a service which is web service
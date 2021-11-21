FROM golang:alpine

WORKDIR /app

#renombrar
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping
#COPY ./AlgoritmoGenetico.go ./api-covid.go
#port
EXPOSE 8000
# run api-svc001.go
CMD [ "/docker-gs-ping" ]
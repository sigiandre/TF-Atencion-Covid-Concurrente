FROM golang:alpine
#renombrar
COPY ./AlgoritmoGenetico.go ./api-covid.go
#port
EXPOSE 8000
# run api-svc001.go
CMD ["go", "run", "api-covid.go"]
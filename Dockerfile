# stiahne package podľa aktuálnej architektúry dockeru (windows/linux/...)
FROM golang:1.18 AS build
WORKDIR /app
COPY models ./models
COPY api ./api
COPY routers ./routers
COPY pkg ./pkg
COPY main.go ./

# zabezpečí statické linkovanie knižníc do výsledného programu
ENV CGO_ENABLED=0

# inicializujeme nas modul definovany v go.mod
RUN go mod init ASOS
# module requirements and sums
RUN go mod tidy 

# výsledkom príkazu je vykonávateľný súbor asos-dojo
RUN go build -a -o ./asos-dojo .

# Multi-stage builds 
FROM scratch
# nasledujúci riadok treba odkomentovať v prípade produkčnej verzie
#ENV GIN_MODE=release
COPY --from=build /app/asos-dojo ./
EXPOSE 9999/tcp
ENTRYPOINT ["./asos-dojo"]
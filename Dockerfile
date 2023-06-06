FROM golang:alpine3.18 as buildStage
WORKDIR /restaurant-mgmt
COPY . .
RUN go mod download
RUN go build -o ./res-app main.go

FROM alpine:3.16.2 as runtimeStage
WORKDIR /app
COPY --from=buildStage /restaurant-mgmt/res-app .
ENTRYPOINT ["/app/res-app"]
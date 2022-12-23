FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /api-search

ENV ZINC_SEARCH_HOST="http://localhost:4080/api/"
ENV ZINC_SEARCH_HOST_ES="http://localhost:4080/es/"
ENV ZINC_SEARCH_INDEX="mail"
ENV ZINC_SEARCH_USER="admin"
ENV ZINC_SEARCH_PASSWORD="Complexpass#123"
ENV BATCH_SIZE=100
ENV PORT=3001


CMD [ "/api-search" ]
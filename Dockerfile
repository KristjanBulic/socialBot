FROM golang:1.17-alpine as build

WORKDIR /app

COPY . .

RUN go build -o main .
RUN ls

FROM golang:1.17-alpine

COPY --from=build /app/main socialBot
ENTRYPOINT ./socialBot
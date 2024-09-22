FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./ 

RUN go mod download

COPY . .

RUN go build -o auth-service ./main.go

EXPOSE 50051

CMD [ "./auth-service" ]

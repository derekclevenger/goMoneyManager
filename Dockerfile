FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get github.com/gorilla/mux
RUN go get github.com/jinzhu/gorm
RUN go get 	github.com/jinzhu/gorm/dialects/postgres
RUN go get github.com/joho/godotenv
RUN go build -o main .
CMD ["/app/main"]



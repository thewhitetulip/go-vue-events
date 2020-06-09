FROM golang

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/gorilla/mux
RUN go build -o main .
CMD ["/app/main"]
COPY . .

EXPOSE 8080
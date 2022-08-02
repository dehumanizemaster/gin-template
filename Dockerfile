FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go get github.com/githubnemo/CompileDaemon

#EXPOSE 8000

RUN go build main.go
#CMD ["go", "run", "main.go"]
#ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
CMD ["./main"]


#Docker Command

#docker build -t streaming_api .
#docker run -i -t -d -p 8080:8080 streaming_api
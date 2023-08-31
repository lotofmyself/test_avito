FROM debian

RUN apt-get update && apt-get install -y curl && mkdir /app
RUN curl -O https://dl.google.com/go/go1.17.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.17.linux-amd64.tar.gz
RUN rm go1.17.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/go
EXPOSE 6379
WORKDIR /app
COPY ./app/ .
RUN go get -u github.com/gorilla/mux && go get -u github.com/go-redis/redis/v8 && go get -u golang.org/x/net/context
RUN go mod download 
RUN go build -o app .

CMD ["./app"]

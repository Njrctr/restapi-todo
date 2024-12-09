FROM golang:1.22.2


RUN go version
ENV GOPATH=/

COPY ./ ./

#install psql
RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
RUN echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ bionic main" > /etc/apt/sources.list.d/migrate.list
RUN apt-get update
RUN apt-get install -y migrate
RUN apt-get -y install postgresql-client

#make executable
RUN chmod +x wait-for-postgres.sh


RUN go mod download
RUN go build -o restapi-todo ./cmd/main.go

CMD [ "./restapi-todo" ]
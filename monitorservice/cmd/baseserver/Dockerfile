
FROM golang:1.18 as build

LABEL Version="1.0.0"
LABEL Author="monitorservice"


WORKDIR /usr/local/go/src/order

COPY . /usr/local/go/src/order

RUN go mod download 

WORKDIR /usr/local/go/src/order/cmd/baseserver
# You must use parameter settings, otherwise you cannot run on the alpin image
# like GOOS=linux GOARCH=amd64 CGO_ENABLED=0
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o monitorservice main.go


# Move the file of completed to another docker image
FROM alpine

WORKDIR /monitorservice
RUN chmod 777 /monitorservice

COPY --from=build /usr/local/go/src/order/cmd/baseserver/monitorservice /monitorservice
COPY --from=build /usr/local/go/src/order/cmd/baseserver/config.yml /monitorservice
COPY --from=build /usr/local/go/src/order/migrate/sql /monitorservice/sql

RUN mv /monitorservice/config.yml /monitorservice/config.yml

EXPOSE 17782
CMD [ "./monitorservice" ]

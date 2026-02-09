
FROM golang:1.22.3 as build

LABEL Version="1.0.0"
LABEL Author="dcc"


WORKDIR /usr/local/go/src/dcc

COPY . /usr/local/go/src/dcc

RUN go mod download 

WORKDIR /usr/local/go/src/dcc/cmd/baseserver
# You must use parameter settings, otherwise you cannot run on the alpin image
# like GOOS=linux GOARCH=amd64 CGO_ENABLED=0
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o chatservice main.go


# Move the file of completed to another docker image
FROM alpine

WORKDIR /chatservice
RUN chmod 777 /chatservice

ARG web_mode=debug

COPY --from=build /usr/local/go/src/dcc/cmd/baseserver/chatservice /chatservice
COPY --from=build /usr/local/go/src/dcc/cmd/baseserver/config-${web_mode}.yml /chatservice

RUN mv /chatservice/config-${web_mode}.yml /chatservice/config.yml

EXPOSE 8896
CMD [ "./chatservice" ]

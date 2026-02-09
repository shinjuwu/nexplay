FROM golang:1.22.3 as build_backend

LABEL Version="1.0.0"
LABEL Author="dcc"

WORKDIR /usr/local/go/src/definition

COPY /definition /usr/local/go/src/definition

WORKDIR /usr/local/go/src/dcc

COPY /backend /usr/local/go/src/dcc

RUN chmod +x /usr/local/go/src/dcc/swag
RUN ./swag init
# You must use parameter settings, otherwise you cannot run on the alpin image
# like GOOS=linux GOARCH=amd64 CGO_ENABLED=0

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o dcc_backend main.go


# Move the file of completed to another docker image
FROM alpine

WORKDIR /backend
RUN chmod 777 /backend

ARG web_mode=debug

COPY --from=build_backend /usr/local/go/src/dcc/dcc_backend /backend
COPY --from=build_backend /usr/local/go/src/dcc/config-${web_mode}.yml /backend
COPY --from=build_backend /usr/local/go/src/dcc/docs /backend/docs
COPY --from=build_backend /usr/local/go/src/dcc/db/migrations /backend/db/migrations


RUN mv /backend/config-${web_mode}.yml /backend/config.yml

ENV web_mode=${web_mode}

# EXPOSE 9986
# CMD [ "./dcc_backend" ]

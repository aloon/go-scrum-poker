# Dockerfile.production

FROM golang:1.20 as builder

ENV APP_HOME /go/src/go-scrum-poker

WORKDIR "$APP_HOME"
COPY ./ .

RUN go mod download
RUN go mod verify
RUN go build -o go-scrum-poker -buildvcs=false

FROM golang:1.20

ENV APP_HOME /go/src/go-scrum-poker
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY templates/ templates/
COPY static/ static/
COPY resources/ resources/
COPY --from=builder "$APP_HOME"/go-scrum-poker $APP_HOME 

EXPOSE 8080
CMD ["./go-scrum-poker"]
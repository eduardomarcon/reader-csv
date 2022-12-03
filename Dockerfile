FROM golang:1.19.3 as download

WORKDIR /go/src/reader

COPY go.mod .
COPY go.sum .

RUN go mod download


FROM golang:1.19.3 as build

WORKDIR /go/src/reader

COPY . .

COPY --from=download /go/pkg/mod/ /go/pkg/mod/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o reader cmd/reader/main.go


FROM scratch as image

COPY --from=build /go/src/reader/reader .
COPY --from=build /go/src/reader/.env .
COPY --from=build /go/src/reader/users.csv .

EXPOSE 8080

CMD ["/reader"]
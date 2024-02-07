FROM golang:1.22 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app ./cmd

FROM scratch AS exec
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
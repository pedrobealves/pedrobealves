FROM golang:1.21.6 as build
WORKDIR /app
COPY . .
RUN go mod download && \
  go mod verify
RUN go build -o autometedReadme ./cmd/automated-readme

FROM scratch

COPY --from=build /app .

CMD ["./autometedReadme"]


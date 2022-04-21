FROM golang:1.17-alpine AS BACKEND

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o backend

# -------------------------------------------------------
FROM alpine

COPY --from=BACKEND /app/backend /usr/bin/backend

RUN chmod +x /usr/bin/backend

CMD ["/usr/bin/backend"]

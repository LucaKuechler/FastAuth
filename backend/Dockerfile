FROM golang:1.17-alpine AS BACKEND

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o backend

# -------------------------------------------------------
FROM alpine

COPY --from=BACKEND /app/backend /usr/bin/backend

ENV JWT_SECRET_KEY "GITHUB_YOUR_SECRET_KEY"
ENV DB_USER "GITHUB_DB_USER"
ENV DB_PASSWORD "GITHUB_DB_PASSWORD"
ENV DB_HOST "GITHUB_DB_HOST"
ENV DB_NAME "GITHUB_DB_NAME"

RUN chmod +x /usr/bin/backend

CMD ["/usr/bin/backend"]

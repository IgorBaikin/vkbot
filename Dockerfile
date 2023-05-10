FROM golang:1.19

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

EXPOSE 80/tcp

RUN CGO_ENABLED=0 GOOS=linux go build -o /vk-bot-golang .

CMD ["/vk-bot-golang"]
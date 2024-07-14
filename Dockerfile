FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .
CMD ["sh", "-c", "./main -telegram-token=$TELEGRAM_TOKEN -vk-token=$VK_TOKEN"]
EXPOSE 8080

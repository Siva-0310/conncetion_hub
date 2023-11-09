FROM golang:latest as builder
WORKDIR /app
COPY . /app
RUN rm -f app/auth_service/Dockerfile
RUN rm -f .gitignore
RUN rm -f compose.yaml
RUN rm -f db.sql 
RUN rm -f README.md 
RUN rm -rf auth_protos
RUN rm -rf authorization
RUN rm -rf user_service
COPY go.mod /app
COPY go.sum /app
RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

FROM alpine:latest
WORKDIR /
ENV MY_SQL_USER=user 
ENV MY_SQL_PASS=password
COPY --from=builder /server /server
EXPOSE 8000
CMD [ "/server" ]
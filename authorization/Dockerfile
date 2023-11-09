FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN rm -f .gitignore
RUN rm -f db.sql
RUN rm -f README.md
RUN rm -f compose.yaml
RUN rm -rf auth_service
RUN rm -rf user_service
RUN rm -f app/authorization/Dockerfile
RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

FROM alpine:latest
WORKDIR /
COPY --from=builder /server /server
EXPOSE 9080
CMD [ "/server" ]
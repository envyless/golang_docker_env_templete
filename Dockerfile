# FROM golang:1.15-alpine

# COPY . .
 
# EXPOSE 9999

# CMD ["go", "run", "/app/main.go"]

FROM golang:1.15-alpine

RUN mkdir /hello
COPY ./app/main.go /hello
CMD ["go", "run", "/hello/main.go"]
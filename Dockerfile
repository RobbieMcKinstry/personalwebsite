FROM golang:alpine AS build-env
ADD . /src/github.com/RobbieMcKinstry/personal-website
ENV GOPATH=/
RUN cd /src/github.com/RobbieMcKinstry/personal-website && go build -o /bin/server

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /bin/server /app/
COPY ./static /app/static
EXPOSE 8080
ENTRYPOINT ["./server"]

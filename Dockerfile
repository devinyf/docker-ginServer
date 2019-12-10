FROM golang
LABEL maintainer="dyfsquall"
WORKDIR /app
COPY ./go.mod ./
COPY ./go.sum ./
RUN go get -u github.com/gin-gonic/gin
COPY ./ ./
EXPOSE 8888
CMD ["go","run","main.go"]

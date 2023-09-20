FROM golang:1.19-alpine

EXPOSE 8080

#WORKDIR /usr/local/go/src/groove-gallery
WORKDIR /opt/groove-gallery

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
ADD controllers ./controllers
ADD models ./models
RUN go mod download


RUN go build -o /groove-gallery

CMD [ "/groove-gallery" ]
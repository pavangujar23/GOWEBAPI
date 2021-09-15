FROM golang:latest

RUN mkdir /build
WORKDIR /build

COPY go.* ./


#RUN go mod download

RUN export GO111MODULE=on

RUN go get github.com/pavangujar23/GOWEBAPI

#sShttps://github.com/pavangujar23/GOWEBAPI.gitSS
RUN cd /build && git clone https://github.com/pavangujar23/GOWEBAPI.git

RUN cd /build/GOWEBAPI && go build
EXPOSE 8080

ENTRYPOINT [ "/build/GOWEBAPI"]
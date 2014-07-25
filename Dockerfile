FROM google/golang:1.3

MAINTAINER "Sam Jantz<sjantz0@gmail.com>"

ADD greg.go /app/greg.go
WORKDIR "/app/"

RUN go get github.com/theTechnoWeenie/greg/server
RUN go build greg.go && chmod +x greg

CMD /app/greg

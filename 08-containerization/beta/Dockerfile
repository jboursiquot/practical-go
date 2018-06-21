FROM golang:latest 
RUN mkdir /beta
ADD . /beta/ 
WORKDIR /beta
EXPOSE 8082
RUN go build -o main . 
CMD ["/beta/main"]
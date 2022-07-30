FROM golang:latest

RUN apt-get update
RUN apt-get -y install python3
RUN apt-get -y install python3-setuptools
RUN apt-get -y install python3-pip

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN pip install -r requirements.txt
RUN go build -o main .
CMD ["/app/main"]

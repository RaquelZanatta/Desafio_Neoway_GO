FROM golang

RUN mkdir /app
WORKDIR /app

ADD main.go /app/
ADD base_teste[802].txt /app/
COPY . /app/
RUN go mod init go-neoway
RUN go get github.com/lib/pq
RUN go mod tidy



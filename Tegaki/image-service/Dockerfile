FROM golang:1.11 

RUN mkdir -p /app && mkdir -p /app/imgs
WORKDIR /app

# copy the content 
COPY . .

# install dependencies
RUN export GO111MODULES=on
RUN go build

# execute
CMD ["./image-service"]

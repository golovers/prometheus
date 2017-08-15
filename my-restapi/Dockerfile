FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
COPY my-restapi.bin /app 
CMD ["/app/my-restapi.bin"]


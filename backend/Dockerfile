FROM golang:1.18
WORKDIR /app
EXPOSE 8001
ADD backend .
RUN go install ./
ENTRYPOINT backend

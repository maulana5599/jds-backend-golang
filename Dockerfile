from golang as builder

WORKDIR /app
ADD . .
RUN go build -o /usr/local/bin/jds-backend-app

EXPOSE 3001
CMD ["/usr/local/bin/jds-backend-app"]

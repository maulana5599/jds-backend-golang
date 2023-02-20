from golang as builder

WORKDIR /app
ADD . .
RUN go build -o /usr/local/bin/jds-backend-spp

EXPOSE 3000
CMD ["/usr/local/bin/jds-backend-app"]

FROM golang:1.17.0-alpine3.14
RUN apk update && apk add gcc git 
WORKDIR /xxx
COPY . .
RUN go build ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /xxx/app .
EXPOSE 8080
CMD ["./app"]  
ENV DATABASE_URL=postgres://ednlnmoi:huVNUQvl1bR9puqxPOrGkJ-pPiVQ23hY@otto.db.elephantsql.com:5432/ednlnmoi
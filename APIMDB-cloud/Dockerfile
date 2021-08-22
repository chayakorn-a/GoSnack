FROM golang:1.11.2-alpine3.8
WORKDIR /xxx
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
EXPOSE 8080
CMD ["./app"]
ENV DATABASE_URL=postgres://ednlnmoi:huVNUQvl1bR9puqxPOrGkJ-pPiVQ23hY@otto.db.elephantsql.com:5432/ednlnmoi
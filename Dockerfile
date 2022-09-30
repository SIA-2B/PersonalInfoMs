FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /personalInfoMs

EXPOSE 3000

CMD [ "/personalInfoMs" ]
# docker build -t personal_info_ms .
# docker run -it -d -p 3000:3000 -v $PWD:/app --name personalInfoMs personal_info_ms
# docker start 
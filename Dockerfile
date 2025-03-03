FROM golang:1.17-alpine

ENV GOPROXY https://goproxy.cn
RUN apk add git
RUN  go get github.com/google/go-github/v69@v69.2.0 \
    && go get "github.com/olekukonko/tablewriter"

WORKDIR /

COPY . /

ENTRYPOINT ["go", "run", "/github.go"]

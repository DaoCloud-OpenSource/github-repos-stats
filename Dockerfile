FROM golang:1.17-alpine

ENV GOPROXY https://goproxy.cn
RUN apk add git
RUN  go get github.com/google/go-github/v42@v42.0.0 \
    && go get "github.com/olekukonko/tablewriter"

WORKDIR /

COPY . /

ENTRYPOINT ["go", "run", "/github.go"]
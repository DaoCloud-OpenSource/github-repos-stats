FROM golang:1.17-alpine

# ENV GOPROXY https://goproxy.cn
RUN apk add git
RUN  go get github.com/google/go-github/v72@v72.0.0 \
    && go get "github.com/olekukonko/tablewriter"@v0.0.5

WORKDIR /

COPY . /

ENTRYPOINT ["go", "run", "/github.go"]

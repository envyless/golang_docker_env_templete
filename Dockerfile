FROM golang:1.15-alpine


RUN apk update && apk add alpine-sdk git bash build-base


ENV GOTOOLSTOBUILD \
        github.com/FiloSottile/vendorcheck \
        github.com/GoASTScanner/gas \
        github.com/alecthomas/gocyclo \
        github.com/alecthomas/gometalinter \
        github.com/client9/misspell \
        golang.org/x/lint/golint \
        github.com/gordonklaus/ineffassign \
        github.com/jgautheron/goconst \
        github.com/golang/dep/cmd/dep \
        github.com/kisielk/errcheck \
        github.com/mdempsky/unconvert \
        github.com/mibk/dupl \
        mvdan.cc/interfacer \
        mvdan.cc/unparam \
        github.com/opennota/check/cmd/aligncheck \
        github.com/opennota/check/cmd/structcheck \
        github.com/opennota/check/cmd/varcheck \
        github.com/stretchr/testify \
        github.com/stripe/safesql \
        github.com/tsenart/deadcode \
        github.com/walle/lll \
        golang.org/x/tools/cmd/goimports \
        honnef.co/go/tools/cmd/staticcheck \
        github.com/client9/misspell/cmd/misspell


WORKDIR /go/src/app
COPY . .

#RUN go get -d -v ./...
#RUN go install -v ./...

CMD ["app"]
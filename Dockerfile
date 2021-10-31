FROM golang:1.17.2

COPY . "/go/src/github.com/rosfran/interview-accountapi"
WORKDIR "/go/src/github.com/rosfran/interview-accountapi"

COPY . .

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN ls -lath /go/src/github.com/rosfran/interview-accountapi
COPY *.go ./

RUN go get -d -v ./...
#RUN go get github.com/rosfran/interview-accountapi/

RUN go install -v ./...
RUN go install -v github.com/rosfran/interview-accountapi/cmd


# RUN set -ex; \
#     apk update; \
#     apk add --no-cache git 
# WORKDIR /go/src/app/interview-accountapi

#RUN go build 

#RUN GO111MODULE=on  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build /go/src/github.com/rosfran/interview-accountapi/account/account.go
#RUN GO111MODULE=on  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build /go/src/github.com/rosfran/interview-accountapi/account/account_rest.go
#RUN GO111MODULE=on  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build /go/src/github.com/rosfran/interview-accountapi/account/accounttest.go


#CMD CGO_ENABLED=0 go test ./...

WORKDIR "/go/src/github.com/rosfran/interview-accountapi/account"

EXPOSE 8089

ENTRYPOINT ["go", "test", "-v", "./..."]





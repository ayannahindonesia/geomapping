FROM golang:alpine  AS build-env

ARG APPNAME="asira_geomapping"
ARG ENV="staging"

#RUN adduser -D -g '' golang
#USER root

ADD . $GOPATH/src/"${APPNAME}"
WORKDIR $GOPATH/src/"${APPNAME}"

RUN apk add --update git gcc libc-dev;
RUN apk --no-cache add curl
#  tzdata wget gcc libc-dev make openssl py-pip;
RUN go get -u github.com/golang/dep/cmd/dep

RUN cd $GOPATH/src/"${APPNAME}"
RUN cp deploy/conf.yaml config.yaml
RUN dep ensure -v
RUN go build -v -o "${APPNAME}-res"

RUN ls -alh $GOPATH/src/
RUN ls -alh $GOPATH/src/"${APPNAME}"
RUN ls -alh $GOPATH/src/"${APPNAME}"/vendor
RUN pwd

FROM alpine

WORKDIR /go/src/
COPY --from=build-env /go/src/asira_geomapping/asira_lender-res /go/src/asira_geomapping
COPY --from=build-env /go/src/asira_geomapping/deploy/conf.yaml /go/src/config.yaml
COPY --from=build-env /go/src/asira_geomapping/permissions.yaml /go/src/permissions.yaml

RUN pwd
#ENTRYPOINT /app/asira_lender-res
CMD ["/go/src/asira_geomapping","run"]

EXPOSE 8000

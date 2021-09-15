FROM golang:1.16.6 AS buildgo

WORKDIR /work
COPY . .
RUN env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build .

# This image name is documented below with the com.ymeadows.baseimage label
# If you update the line below, update the label to match.
FROM bash

COPY --from=buildgo /work/versiontool /versiontool

LABEL com.ymeadows.baseimage=scratch
LABEL org.opencontainers.image.source=https://github.com/ymeadows/versiontool
LABEL org.opencontainers.image.authors=devops

ENTRYPOINT ["/versiontool"]

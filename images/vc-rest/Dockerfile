#
#  Go-lang-Vcs-NEXTClan 
#

ARG GO_VER
ARG ALPINE_VER

FROM golang:${GO_VER}-alpine${ALPINE_VER} as golang
RUN apk add --no-cache \
	gcc \
	musl-dev \
	git \
	libtool \
	bash \
	make;
ADD . src/github.com/trustbloc/vcs
WORKDIR src/github.com/trustbloc/vcs
ENV EXECUTABLES go git

FROM golang as vcs
RUN make vc-rest


FROM alpine:${ALPINE_VER} as base
LABEL org.opencontainers.image.source https://github.com/trustbloc/vcs

COPY --from=vcs /go/src/github.com/trustbloc/vcs/.build/bin/vc-rest /usr/local/bin

ENTRYPOINT ["vc-rest"]

FROM alpine:3.14.2

LABEL maintainer="allysson.v1@gmail.com"

WORKDIR /songs

RUN set -ex; \
	apk update --no-cache && \
	apk add --no-cache git ffmpeg python3 py3-pip && \
	pip3 install git+https://github.com/flyingrub/scdl

ENTRYPOINT ["scdl", "-l"]
CMD []

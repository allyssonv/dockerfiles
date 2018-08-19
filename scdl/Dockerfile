FROM alpine:3.7

LABEL maintainer="allysson.v1@gmail.com"

COPY scdl /scdl

WORKDIR /scdl

RUN set -ex; \
	apk update --no-cache && \
	apk add --no-cache python3 && \
	addgroup -g 1000 -S jarvis && \
	adduser -u 1000 -S jarvis -G jarvis && \
	python3 setup.py install

USER jarvis

ENTRYPOINT ["scdl", "-l"]
CMD []
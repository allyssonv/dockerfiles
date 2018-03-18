FROM alpine:3.7

LABEL maintainer="allysson.v1@gmail.com"

COPY scdl /scdl

WORKDIR /scdl

RUN set -ex; \
	apk update --no-cache && \
	apk add --no-cache python3; \
	python3 setup.py install

ENTRYPOINT ["scdl", "-l"]
CMD []


######################RUN AS######################

# Pass as parameter the LINK ADDRESS of the song

# docker run --rm -it -v /home/user:/scdl scdl:flyingrub [LINK ADDRESS]
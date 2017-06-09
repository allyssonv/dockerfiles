# docker run --rm --name utorrent -v /host/folder:/utorrent -it -p 8080:8080 utorrent:debian
# after, head over to http://localhost:8080/gui
# user: admin
# password:

FROM debian:wheezy-slim

COPY . /

RUN set -ex pipefail; apt-get update && apt-get install -y ca-certificates \
	--no-install-recommends \
	&& rm -rf /var/lib/apt/lists/*

WORKDIR utorrent

CMD ["./utserver"]
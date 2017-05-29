# execute that way
#
# docker run --rm -it -v /your/host/folder:/spotdl/Music spotify-dl:debian

FROM debian:wheezy-slim

RUN set -eux pipefail; apt-get update && apt-get install --no-install-recommends \
    git \
    libav-tools \
    ca-certificates \
    python-pip -y && \
    git clone https://github.com/Ritiek/Spotify-Downloader.git spotdl && \
    cd spotdl && \
    pip install -U -r requirements.txt && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /spotdl

CMD ["python", "spotdl.py"]
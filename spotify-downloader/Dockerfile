FROM ubuntu:16.04

COPY spotdl spotdl

WORKDIR spotdl

RUN set -eux pipefail; apt-get update && apt-get install \
    libav-tools \
    python3-pip -yq && \
    pip3 install -U -r requirements.txt && \
    rm -rf /var/lib/apt/lists/*

CMD ["python3", "spotdl.py"]
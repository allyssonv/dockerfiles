# docker run --rm --name nes \
# --device /dev/snd \
# --device /dev/input \
# --group-add audio \
# -e DISPLAY \
# -v /tmp/.X11-unix:/tmp/.X11-unix \
# -v /home/user/zsnes/roms:/roms \
# -v /home/user/zsnes/files:/root/.zsnes \
# zsnes:small


FROM debian:wheezy-slim

RUN dpkg --add-architecture i386 && \
    apt-get update && \
    apt-get install --no-install-recommends -y zsnes && \
    rm -rf /var/lib/apt/lists/*

CMD ["/usr/bin/zsnes"]
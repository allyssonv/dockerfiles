FROM python:3.7-alpine3.7
RUN apk --no-cache update && apk add --no-cache --virtual .build-deps \
	gcc \
	musl-dev \
	linux-headers && \
	pip install psutil glances bottle && \
	apk del .build-deps

EXPOSE 61208 61209

CMD ["glances"]

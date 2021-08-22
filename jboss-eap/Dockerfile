FROM allysson/openjdk:11-alpine3.14.1

COPY . /opt/jboss

ENV EAP_HOME /opt/jboss/jboss-eap-7.0

# install dumb-init for container shutdown (CTRL-C) purposes
RUN apk --no-cache update && \
	apk --no-cache add openssl && \
	${EAP_HOME}/bin/add-user.sh admin admin --silent && \
	wget -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.0/dumb-init_1.2.0_amd64 && \
	chmod +x /usr/local/bin/dumb-init

EXPOSE 8080 9990

CMD ["/usr/local/bin/dumb-init", "/opt/jboss/jboss-eap-7.0/bin/standalone.sh", "-b", "0.0.0.0", "-bmanagement", "0.0.0.0"]
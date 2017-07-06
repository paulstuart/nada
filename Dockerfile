FROM scratch

MAINTAINER Paul Stuart <pauleyphonic@gmail.com>

COPY nada /

ENTRYPOINT ["/nada"]

EXPOSE 8080


FROM alpine
MAINTAINER Alex Gaesser (agaesser@gmail.com)

ENV TERM linux
RUN apk --no-cache add apache2-utils
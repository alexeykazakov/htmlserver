FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

LABEL maintainer "Devtools <devtools@redhat.com>"
LABEL author "Devtools <devtools@redhat.com>"

ENV HTML_SERVICE=/usr/local/bin/htmlserver \
    USER_UID=1001 \
    USER_NAME=htmlserver \
    LANG=en_US.utf8

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

# install backsoon-service binary
COPY build/_output/bin/htmlserver ${HTML_SERVICE}

ENTRYPOINT ["/usr/local/bin/htmlserver"]

EXPOSE 8090

USER ${USER_UID}

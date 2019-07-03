FROM debian:stretch

WORKDIR /home/homo

COPY . /home/homo/

# Golang env
ENV GOLANG_VERSION 1.12.6
ENV GOLANG_TAR_BALL go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/$GOLANG_TAR_BALL
ENV GOLANG_DOWNLOAD_SHA256 dbcf71a3c1ea53b8d54ef1b48c85a39a6c9a935d01fc8291ff2b92028e59913c
ENV GOPATH /home/homo/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Install system dependence
RUN \
    apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates git wget tar && \
    apt-get install -y --no-install-recommends gcc automake autoconf libtool build-essential && \
    apt-get install -y --no-install-recommends bison swig python-dev libpulse-dev portaudio19-dev

# Install PocketSphinx
RUN make deps

# Install Golang
RUN wget $GOLANG_DOWNLOAD_URL && \
    echo "$GOLANG_DOWNLOAD_SHA256  $GOLANG_TAR_BALL" | sha256sum -c - && \
    tar -C /usr/local -xzf $GOLANG_TAR_BALL && \
    rm $GOLANG_TAR_BALL

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN go env

RUN go version

# Build homo webview
RUN \
    cd homo && \
    make gen && \
    make webview

# Replace 1000 with your user / group id
#RUN export uid=1000 gid=1000 && \
#    mkdir -p /home/homo && \
#    echo "homo:x:${uid}:${gid}:homo,,,:/home/homo:/bin/bash" >> /etc/passwd && \
#    echo "homo:x:${uid}:" >> /etc/group && \
#    echo "homo ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/homo && \
#    chmod 0440 /etc/sudoers.d/homo && \
#    chown ${uid}:${gid} -R /home/homo

USER homo
ENV HOME /home/homo
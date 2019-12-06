FROM golang:1.13.4-buster
MAINTAINER Charo Nuguid <me@thegeekettespeaketh.com>


# Used for debian-based OS
RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
 && rm -rf /var/lib/apt/lists/*

# RUN wget https://github.com/gohugoio/hugo/releases/download/v0.59.0/hugo_extended_0.59.0_Linux-64bit.deb \
# 	&& dpkg -i hugo_extended_0.59.0_Linux-64bit.deb \
# 	&& rm hugo_extended_0.59.0_Linux-64bit.deb

# Used for Alpine
# RUN apk add --no-cache \
#         git


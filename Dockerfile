FROM golang:1.10.2

RUN apt-get update && apt-get install -y --no-install-recommends \
  curl \
  git \
  && rm -rf /var/lib/apt/lists/*

# Download and install the latest release of dep
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && \
    chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/github.com/twelvelabs/wordcount
WORKDIR /go/src/github.com/twelvelabs/wordcount

# copies the Gopkg.toml and Gopkg.lock to WORKDIR
COPY Gopkg.toml Gopkg.lock ./
# install the dependencies without checking for go code
RUN dep ensure -vendor-only

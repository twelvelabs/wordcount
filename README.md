## Overview

A simple word counting service I was asked to put together as a [programming exercise](./docs/requirements.md).

## Prerequisites

- [Docker](https://docs.docker.com/install/)
- [Docker Compose](https://docs.docker.com/compose/install/) (already installed w/ Docker for Mac)

## Install

```bash
# Clone the app
git clone git@github.com:twelvelabs/wordcount.git
cd ./wordcount
docker-compose build
```

## Running

```bash
docker-compose up
open http://0.0.0.0
```

## Tests

```bash
docker-compose run --rm app go test
```

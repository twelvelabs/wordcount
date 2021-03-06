## Overview

A simple word counting service I was asked to put together as a [programming exercise](./docs/requirements.md).

## Prerequisites

- [Docker](https://docs.docker.com/install/)
- [Docker Compose](https://docs.docker.com/compose/install/) (already installed w/ Docker for Mac)

## Install

```bash
git clone git@github.com:twelvelabs/wordcount.git
cd ./wordcount
# Configure ansible-vault password (ping @twelvelabs for it).
echo $VAULT_PASS > ./ansible/vault_pass.txt
# Build the app image
docker-compose build
# Decrypt app secrets into ./home (which will be mounted into the app container)
docker-compose run --rm ansible ansible-playbook /ansible/setup.yml
```

## Running

```bash
docker-compose up
open https://0.0.0.0
```

## Unit tests

```bash
docker-compose run --rm app go test
```

## Deploy

```bash
# Build the app into ./ansible/files/wordcount
docker-compose run --rm app bin/build
# Deploy to the remote server
docker-compose run --rm ansible ansible-playbook /ansible/deploy.yml
```

## Integration tests

This assumes that you have both [HTTPie](https://httpie.org) and [jq](https://stedolan.github.io/jq/) installed.

```bash
# Should return a 401 response
echo "Hey ho, let's go. Hey ho, let's go." | http --verify=no POST https://192.241.204.44/wordcount

# So, get an auth token...
TOKEN_JSON=$(http --verify=no POST https://192.241.204.44/token name="YOURNAME" password="YOURPASS")
TOKEN=$(jq -r '.token' <<< "$TOKEN_JSON")

# Then retry authenticated
echo "Hey ho, let's go. Hey ho, let's go." | http --verify=no POST https://192.241.204.44/wordcount "Authorization: Bearer $TOKEN"

# Now w/ a 3M text file :)
http --verify=no POST https://192.241.204.44/wordcount "Authorization: Bearer $TOKEN" < ./fixtures/war-and-peace.txt
```

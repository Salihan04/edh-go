# edh-go

[![Build Status](https://travis-ci.com/Salihan04/edh-go.svg?token=8qjDfnsQn45zaysJX5gL&branch=master)](https://travis-ci.com/Salihan04/edh-go)

Go module to obtain data from the Enterprise Data Hub

## Developer Guide

### Folder Structure

This repo follows the folder structure specified in [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

### EDH API Specs

[v1.1.0](https://public.cloud.myinfo.gov.sg/edh/edh-tuo-specs.html)

### Keys and Certs

* You should have generated a private and public key-pair
* Pass the public key to EDH team
* Put the private key into the **ssl** folder. This will be used to sign the request and decrypt the response
* EDH team should give you their public cert
  * Put this into the **ssl** folder. This will be used to verify the response

### Contributing Your Code

If you would like to contribute to this repo, please open an issue, fork the repo, implement your code and tests and create a PR

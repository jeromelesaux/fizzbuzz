# Fizzbuzz Golang implementation 

## Introduction 
This project is a litte fizzbuzz (https://en.wikipedia.org/wiki/Fizz_buzz)server implementation. 
You can use as standalone app or as a docker image.

## Requirements 
You will need at least  a docker installation on your computer (https://www.docker.com/get-started).
If you want to install the standalone application, you will need a golang installation (https://golang.org)

## Configuration
You can manage the listen port of the HTTP server. 
In the standalone app before runnnig, export as environnement variable : `export PORT="8080"`. 
In the docker version, modify in the head of the makefile file, the variable HTTP_PORT to the desired PORT. And run by `make docker-run`.

## Installation 
To get help on command line, just type `make help` in a terminal.

To build the standalone app type `make build`. 

To build the docker image type `make docker`. 
And to run the docker container just type `make docker-run`.

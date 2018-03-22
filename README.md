# goratp

[![godoc](https://godoc.org/github.com/kahlys/goratp?status.svg)](https://godoc.org/github.com/kahlys/goratp) 
[![go report](https://goreportcard.com/badge/github.com/kahlys/goratp)](https://goreportcard.com/report/github.com/kahlys/goratp)

Golang wrapper arround the [RATP rest-api by pierre grimaud](https://api-ratp.pierre-grimaud.fr/v3/documentation). 

## Installation

With a correctly configured [Go toolchain](https://golang.org/doc/install):

```
go get -u github.com/kahlys/proxy/cmd/tcpproxy
```

## Package

It gives you real informations and time schedules for any given RERs, Metros, Tramways, Bus or Noctiliens station in real time on the RATP network.

## Tool

```
$ goratp -line=12 -station=saint+lazare -transport=metros
```
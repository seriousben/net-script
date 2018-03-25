# netscript communication language

[![GoDoc](https://godoc.org/github.com/seriousben/net-script?status.svg)](https://godoc.org/github.com/seriousben/net-script)
[![CircleCI](https://circleci.com/gh/seriousben/net-script/tree/master.svg?style=shield)](https://circleci.com/gh/seriousben/net-script/tree/master)
[![Go Reportcard](https://goreportcard.com/badge/github.com/seriousben/net-script)](https://goreportcard.com/report/github.com/seriousben/net-script)
[![codecov](https://codecov.io/gh/seriousben/net-script/branch/master/graph/badge.svg)](https://codecov.io/gh/seriousben/net-script)

Write runnable programs against remote servers.

## Usage

```
Usage:
  netscript [flags]
  netscript [command]

Available Commands:
  help        Help about any command
  lint        Lint a netscript program
  run         Run a netscript program

Flags:
  -h, --help   help for netscript

```

## Examples

### httpbin.org example

```
$ cat example/httpbin.ns
# Get IP
GET https://httpbin.org/ip

# Get headers
GET https://httpbin.org/headers
X-Custom-Header: Custom value

# Post data
POST https://httpbin.org/post

DATA

# Post JSON
POST https://httpbin.org/post

{"json": true}

# Patch DATA
PATCH https://httpbin.org/patch

DATA

# Put
PUT https://httpbin.org/put

# Delete
DELETE https://httpbin.org/delete

# Get image
GET https://httpbin.org/image/png

```

```
$ netscript run http.ns
# Get IP
GET https://httpbin.org/ip
\\  HTTP/1.1 200 OK
\\ Content-Length: 34
\\ Access-Control-Allow-Credentials: true
\\ Access-Control-Allow-Origin: *
\\ Connection: keep-alive
\\ Content-Type: application/json
\\ Date: Sun, 25 Mar 2018 15:32:12 GMT
\\ Server: meinheld/0.6.1
\\ Via: 1.1 vegur
\\ X-Powered-By: Flask
\\ X-Processed-Time: 0

{
  "origin": "192.222.149.100"
}

# Post data
POST https://httpbin.org/post
\\  HTTP/1.1 200 OK
\\ Content-Length: 257
\\ Access-Control-Allow-Credentials: true
\\ Access-Control-Allow-Origin: *
\\ Connection: keep-alive
\\ Content-Type: application/json
\\ Date: Sun, 25 Mar 2018 15:32:12 GMT
\\ Server: meinheld/0.6.1
\\ Via: 1.1 vegur
\\ X-Powered-By: Flask
\\ X-Processed-Time: 0

{
  "args": {},
  "data": "DATA",
  "files": {},
  "form": {},
  "headers": {
    "Connection": "close",
    "Content-Length": "4",
    "Host": "httpbin.org"
  },
  "json": null,
  "origin": "192.222.149.100",
  "url": "https://httpbin.org/post"
}

# Post JSON
POST https://httpbin.org/post
\\  HTTP/1.1 200 OK
\\ Content-Length: 288
\\ Access-Control-Allow-Credentials: true
\\ Access-Control-Allow-Origin: *
\\ Connection: keep-alive
\\ Content-Type: application/json
\\ Date: Sun, 25 Mar 2018 15:32:11 GMT
\\ Server: meinheld/0.6.1
\\ Via: 1.1 vegur
\\ X-Powered-By: Flask
\\ X-Processed-Time: 0

{
  "args": {},
  "data": "{\"json\": true}",
  "files": {},
  "form": {},
  "headers": {
    "Connection": "close",
    "Content-Length": "14",
    "Host": "httpbin.org"
  },
  "json": {
    "json": true
  },
  "origin": "192.222.149.100",
  "url": "https://httpbin.org/post"
}

# Patch DATA
PATCH https://httpbin.org/patch
\\  HTTP/1.1 200 OK
\\ Content-Length: 258
\\ Access-Control-Allow-Credentials: true
\\ Access-Control-Allow-Origin: *
\\ Connection: keep-alive
\\ Content-Type: application/json
\\ Date: Sun, 25 Mar 2018 15:32:11 GMT
\\ Server: meinheld/0.6.1
\\ Via: 1.1 vegur
\\ X-Powered-By: Flask
\\ X-Processed-Time: 0

{
  "args": {},
  "data": "DATA",
  "files": {},
  "form": {},
  "headers": {
    "Connection": "close",
    "Content-Length": "4",
    "Host": "httpbin.org"
  },
  "json": null,
  "origin": "192.222.149.100",
  "url": "https://httpbin.org/patch"
}

# Put
PUT https://httpbin.org/put
\\  HTTP/1.1 200 OK
\\ Content-Length: 296
\\ Access-Control-Allow-Credentials: true
\\ Access-Control-Allow-Origin: *
\\ Connection: keep-alive
\\ Content-Type: application/json
\\ Date: Sun, 25 Mar 2018 15:32:12 GMT
\\ Server: meinheld/0.6.1
\\ Via: 1.1 vegur
\\ X-Powered-By: Flask
\\ X-Processed-Time: 0

{
  "args": {},
  "data": "# Delete\nDELETE https://httpbin.org/delete",
  "files": {},
  "form": {},
  "headers": {
    "Connection": "close",
    "Content-Length": "42",
    "Host": "httpbin.org"
  },
  "json": null,
  "origin": "192.222.149.100",
  "url": "https://httpbin.org/put"
}

# Get image
GET https://httpbin.org/image/png
\\  HTTP/1.1 200 OK
\\ Content-Length: 8090
\\ Access-Control-Allow-Credentials: true
\\ Access-Control-Allow-Origin: *
\\ Connection: keep-alive
\\ Content-Type: image/png
\\ Date: Sun, 25 Mar 2018 15:32:12 GMT
\\ Server: meinheld/0.6.1
\\ Via: 1.1 vegur
\\ X-Powered-By: Flask
\\ X-Processed-Time: 0

<No Text Representation (type: image/png)>

```

### Simple Websocket script

```
$ cat ws.ns
WS wss://echo.websocket.org 

message
$ netscript run ws.ns
message
```

## Vision

```
#
# Ideas taken from https://github.com/pashky/restclient.el 
#

# Assign jwt from js script
#   Calls `function ns_entry(/*no args*/) { return JSON.stringify({ jwt: "val" }) }`
:jwt |= generate_jwt.js | .jwt 

# Websocket connection
WS http://example.com
User-Agent: net script
Authorization: :jwt

ws message 1
ws message 2
ws message 3
ws message 4
ws message 5
ws message 6
ws message 7

# Pipe data to websocket.js
#   Calls `function ns_entry(allSentMessages, allReceivedMessages) { return JSON.stringify({ key: "val", key2: "val2" }) }`
:wsdata |= websocket.js

# POST with json body
POST http://example.com 
User-Agent: net script
Content-Type: application/json

{
  "example": "{{ :wsdata | .key2 }}"
}

# GET with a query param
GET http://example.com
? key "escaped query param value"
```

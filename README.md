# netscript communication language

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

### Simple HTTP script

```
$ cat http.ns
GET http://example.com
$ netscript run http.ns
status: 200
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

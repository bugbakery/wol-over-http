# Wake-on-LAN over HTTP

A very simple and very small application (and docker image) to trigger Wake-on-LAN packets via HTTP.

The size of the docker image is about 10 MB.

## Trigger WoL packet

There is only one endpoint `POST /wake/{mac_addr}`, while
`mac_addr` can be in the form of `AA:BB:CC:DD:EE:FF` or `AA-BB-CC-DD-EE-FF`.

```
curl -X POST http://localhost:8080/wake/AA:BB:CC:DD:EE:DD:FF
```

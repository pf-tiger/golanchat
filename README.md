# golanchat
2023-1-31
CLI LAN chat tool made with golang. Both client and server applications are available.

## Server usage:
Just run with the following command:
```shell
go run ./main.go
```

## Client usage:
Put the server's IPv4 address or the hostname for the first argument.
```shell
go run ./main.go <server IPv4 address or hostname>
```

## Roadmaps: 
- create secure connection with encryption
- create logging/storing feature
- flexible debug options:
  - choice for displaying datetime
  - converting address to hostname/associated username using AD

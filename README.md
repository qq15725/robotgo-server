# RobotgoServer

[![GitHub release](https://img.shields.io/github/release/qq15725/robotgo-server.svg)](https://github.com/qq15725/robotgo-server/releases/latest)

## Features

- Remotely invoke the Server machine Robotgo method via WebSocket
- Monitor mouse selection text notification WebSocket

## 🦄 Usage

Download the binary file for the corresponding operating system from the [release](https://github.com/qq15725/robotgo-server/releases) page

```shell
# MaxOS/Linux
sudo mv robotgo_server_xx_xxx /usr/local/bin/robotgo_server
chmod +x /usr/local/bin/robotgo_server
robotgo_server -p 8080

# Windows
# double click robotgo_server_xx_xxx_xx_xxx.exe
```

WebSocket client test

```shell
# pnpm add -g wscat
wscat -c ws://localhost:8080

# Call robotgo.Move(100, 100)
> {"jsonrpc":"2.0","method":"Move","params":[100,100]}

# Mouse select text notify
< {"jsonrpc":"2.0","method":"onSelectText","params":{"text":"# onSelectText","x":368,"y":730}}
```

## All callable methods

See the [rpc-map.go](./rpc-map.go) and [go-vgo/robotgo](https://github.com/go-vgo/robotgo)



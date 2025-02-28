# RobotgoServer

[![GitHub release](https://img.shields.io/github/release/qq15725/robotgo-server.svg)](https://github.com/qq15725/robotgo-server/releases/latest)

## 🦄 Usage

Download the binary file for the corresponding operating system from the [release](https://github.com/qq15725/robotgo-server/releases) page

```shell
# MaxOS/Linux
sudo mv robotgo_server_xx_xxx /usr/local/bin/robotgo_server;
chmod +x /usr/local/bin/robotgo_server;
robotgo_server -p 8080;

# Windows
# double click robotgo_server_xx_xxx_xx_xxx.exe
```

WebSocket client test

```shell
# pnpm add -g wscat
wscat -c ws://localhost:8080

> {"jsonrpc":"2.0","method":"Move","params":[100,100]}
```

## All callable methods

See the [rpc-map.go](./rpc-map.go) with [go-vgo/robotgo](https://github.com/go-vgo/robotgo)



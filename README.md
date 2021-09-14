# Weber

[![GitMoji](https://img.shields.io/badge/Gitmoji-%F0%9F%8E%A8%20-FFDD67.svg)](https://gitmoji.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
![Lines Of Code](https://img.shields.io/tokei/lines/github.com/UltiRequiem/weber?color=blue&label=Total%20Lines)
![CodeQL](https://github.com/UltiRequiem/weber/workflows/CodeQL/badge.svg)
![Build](https://github.com/UltiRequiem/weber/workflows/Build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/weber)](https://goreportcard.com/report/github.com/UltiRequiem/yamlfmt)

Asynchronous fetch unlimited times a website.

## Installation

```bash
go install github.com/UltiRequiem/weber@latest
```

Make sure your `$PATH` includes `$GOPATH/bin` so this command can be used anywhere.

You can also use the binaries from
[releases](https://github.com/UltiRequiem/yamlfmt/releases).

## Usage

Basic Usage:

```bash
weber -u http://localhost:3000
```

This will silently fetch _1_ time `http://localhost:3000`. You can use `--url` instead of `-u`.

If you want to log what is happening:

```bash
weber -u http://localhost:3000 -l
```

You can use the `-l` or `--log` flag.

To fetch more than one time:

```bash
weber -u http://localhost:3000 -t 99999 -l
```

Is also valid to use `--times` instead of `-t`.

If you want to do too large a number of fetchs,
by default the requests will be split into small chunks of 100.

Eg: `weber -u http: // localhost: 3000 -t 1000`

You told `Weber` to make 1000 fetches, but is like that the server is gonna ban you
if you make all the fetches at the same time. So `Weber` divides the 1000 requests in
10 parts of 100 fetches.

You can modify the max value of the chunks (default: 100) by:

```bash
weber -u http://localhost:3000 -t 99999 -m 200 -l
```

Instead of `-m` you can also use `maxChunkValue`.

#### License

This project is licensed under the [MIT license](./LICENSE.md).

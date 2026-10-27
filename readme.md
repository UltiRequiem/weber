# Weber

Asynchronous fetch a website.

## Installation

```sh
go install github.com/UltiRequiem/weber@latest
```

You can also use the binaries from
[releases](https://github.com/UltiRequiem/yamlfmt/releases).

## Usage

### Basic Usage

```sh
weber -u http://localhost:3000
```

This will silently fetch _1_ time `http://localhost:3000`. You can use `--url`
instead of `-u`.

### `-l` / `--log`

This Option will Log all what is happening, and the parameters that you passed.

```sh
weber -u http://localhost:3000 -l
```

### `-t` / `--times`

```sh
weber -u http://localhost:3000 -t 99999 -l
```

If you want to do too large a number of fetch, by default the requests will be
split into small chunks of 100.

Example: `weber -u http://localhost:3000 -t 1000`

You told `Weber` to make 1000 fetches, but is like that the server is going to
ban you if you make all the fetches at the same time. So `Weber` divides the
1000 requests in 10 parts of 100 fetches.

You can modify the default value of 100 by using
[`--maxChunkValue`](#-m----maxchunkvalue).

### `-m` / `--maxChunkValue`

If you see that with 100 requests per chunk you are not having any problem, you
can increase the limit.

```sh
weber -u http://localhost:3000 -t 99999 -m 200 -l
```

## Support

Open an Issue, I will check it a soon as possible 👀

If you want to hurry me up a bit
[send me a tweet](https://x.com/UltiRequiem) 😆

Don't forget to star the repo ⭐

## Authors

[Eliaz Bobadilla](https://ultirequiem.com) - Creator and Maintainer 💪

See also the full list of
[contributors](https://github.com/UltiRequiem/REPO/contributors) who
participated in this project ✨

## Licence

Licensed under the MIT License 📄

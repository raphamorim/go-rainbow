# go-rainbow [![Build Status](https://img.shields.io/travis/fatih/color.svg?style=flat-square)](https://travis-ci.org/raphamorim/go-rainbow)

Inspired by [chalk.js](https://github.com/chalk/chalk)

go-rainbow allows you to use `hex` colors or rainbow mask on a String.

![Sample](assets/sample.png)

## Install

```bash
go get github.com/raphamorim/go-rainbow
```

## Modifiers

- `reset`
- `bold`
- `dim`
- `italic (Not widely supported)`
- `underline`
- `inverse`
- `hidden`
- `strikethrough (Not widely supported)`

## Colors

### Usage

```go
Rainbow.Blue("String")
```

### List

- `black`
- `red`
- `green`
- `yellow`
- `blue` (On Windows the bright version is used since normal blue is illegible)
- `magenta`
- `cyan`
- `white`
- `gray` ("bright black")
- `redBright`
- `greenBright`
- `yellowBright`
- `blueBright`
- `magentaBright`
- `cyanBright`
- `whiteBright`

## Background Colors

```go
Rainbow.bgBlue("String")
```

### List

- `bgBlack`
- `bgRed`
- `bgGreen`
- `bgYellow`
- `bgBlue`
- `bgMagenta`
- `bgCyan`
- `bgWhite`
- `bgBlackBright`
- `bgRedBright`
- `bgGreenBright`
- `bgYellowBright`
- `bgBlueBright`
- `bgMagentaBright`
- `bgCyanBright`
- `bgWhiteBright`


## License

The MIT License (MIT) - see [`LICENSE.md`](https://github.com/raphamorim/go-rainbow/blob/master/LICENSE.md) for more details

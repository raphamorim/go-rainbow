# go-rainbow [![Build Status](https://img.shields.io/travis/fatih/color.svg?style=flat-square)](https://travis-ci.org/raphamorim/go-rainbow)

Extremely inspired by [chalk.js](https://github.com/chalk/chalk)

`go-rainbow` allows you to: 

- Iterate modifiers and colors on a String
- Use HEX as you wish.
- Use RGB as you wish.
- Bring magical Rainbow on a String.

![Sample](assets/sample.png)

## Install

```bash
go get github.com/raphamorim/go-rainbow
```

## Iterable Samples

```go
abc := Bold(BgYellow(Blue("Bold+BgYellow+Blue")))
xyz := Dim(Underscore("Dim+Underscore"))
```

## Modifiers

### Usage

```go
Rainbow.Bold("String in Bold")
```

- `Bold`
- `Dim`
- `Italic (Not widely supported)`
- `Underline`
- `Reverse`
- `Hidden`
- `Strikethrough (Not widely supported)`

## Colors

### Usage

```go
Rainbow.Blue("String in Blue")
```

### List

- `Black`
- `Red`
- `Green`
- `Yellow`
- `Blue`
- `Magenta`
- `Cyan`
- `White`

#### Soon (send a PR to help us!):

- `Gray`
- `RedBright`
- `GreenBright`
- `YellowBright`
- `BlueBright`
- `MagentaBright`
- `CyanBright`
- `WhiteBright`

## Background Colors

```go
Rainbow.BgBlue("String in Blue Background")
```

### List

- `BgBlack`
- `BgRed`
- `BgGreen`
- `BgYellow`
- `BgBlue`
- `BgMagenta`
- `BgCyan`
- `BgWhite`

#### Soon (send a PR to help us!):

- `BgBlackBright`
- `BgRedBright`
- `BgGreenBright`
- `BgYellowBright`
- `BgBlueBright`
- `BgMagentaBright`
- `BgCyanBright`
- `BgWhiteBright`


## Hex (working on it)

```go
Rainbow.Hex('#FF8800')
```

## RGB (working on it)

```go
Rainbow.RGB(100, 110, 100)
```

## License

The MIT License (MIT) - see [`LICENSE.md`](https://github.com/raphamorim/go-rainbow/blob/master/LICENSE.md) for more details

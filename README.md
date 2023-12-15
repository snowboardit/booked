# Reserved

_Reserved_ is a CLI built with Go that determines reserved words in programming/database languages or your favorite stack.

## Installation

> ⚠️ Assumes you have already setup your Go environment

**Download**

```sh
$ go get github.com/snowboardit/reserved
```

**Install**

```sh
$ go install
```

## Usage

```sh
# Basic usage
$ reserved [options] [words...]

# More usage
$ reserved -h
$ reserved --help
```

## Data

> ⚠️ **This list is by no means current or complete!**

To get started, I asked GPT-4 to produce a list of all the programming and database languages, omitting those that are obscure and rarely used. GPT-4 isn't perfect and neither am I, so the data requires validation before being considered a reliable source.

The data file is located here: `pkg/data/data.json`

Please feel free to contribute to the data to include missing information, or new languages and words. Learn more in the [contributing](#contributing) section below.

## Contributing

Pull requests are welcome. For major changes, like adding a language or new property, please open an issue first
to discuss what you would like to add/change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)

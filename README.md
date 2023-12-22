![reserved](https://github.com/snowboardit/reserved/assets/23141894/95aff914-eb3c-432a-9ce3-25d74695c067)

# Reserved

_Reserved_ is a CLI built with Go that determines reserved words in programming/database languages or your favorite stack.

Go [here](https://github.com/snowboardit/reserved/issues/3) for a list of currently supported languages.

## Installation

> 💡 Assumes you have Go installed (written using `v1.21.3`)

```sh
go install github.com/snowboardit/reserved@latest
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

To get started, I asked GPT-4 to produce a list of all the programming and database languages, omitting those that are obscure and rarely used. GPT-4 isn't perfect and neither am I, so the data requires validation before being considered a reliable source. Validated languages are being tracked [here](https://github.com/snowboardit/reserved/issues/4).

The data file is located [here](https://github.com/snowboardit/reserved/edit/master/pkg/data/data.json): `pkg/data/data.json`

Please feel free to contribute to the data to include missing information, or new languages and words. Learn more in the [contributing](#contributing) section below.

### Relational Database

For you data curious developers, I've created a script to convert and load the `data.json` file above into a PostgreSQL database to play with.

1. Run `scripts/db/reset` to setup the database
2. Then run the conversion script `scripts/convert/json-to-sql.go` to load the database

**🏆 Bonus:**
Execute `db/sql/top-10.sql` to list the top 10 words across all languages.

## Contributing

Pull requests are welcome. For major changes, like adding a language or new property, please open an issue first
to discuss what you would like to add/change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)

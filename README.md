# invok

A simple command line tool to Invoke Go lambda functions locally

## Install

From source (you need to have go 1.11 or above installed)

```bash
go get github.com/lmammino/invok
go install github.com/lmammino/invok
```

## Usage

Once you have a lambda running locally on a given port:

```bash
cat someJSONevent.json | invok --host "localhost:1234"
```

Make sure to replace the event file and the host accordingly.

If you just want to test if the lambda is up and running:

```bash
invok --host "localhost:1234" --ping
```

If this command exits with status 0 than the lambda is up and running and reachable

## Contributing

Everyone is very welcome to contribute to this project.
You can contribute just by submitting bugs or suggesting improvements by
[opening an issue on GitHub](https://github.com/lmammino/invok/issues).


## License

Licensed under [MIT License](LICENSE). © Luciano Mammino.

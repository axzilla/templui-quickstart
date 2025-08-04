# templUI Quickstart

Get started with templUI components for Go + templ.

## Quick Start

```bash
git clone https://github.com/templui/templui-quickstart.git my-app
cd my-app
make dev
```

Your app is now running at [http://localhost:7331](http://localhost:7331)

## Commands

```bash
make dev          # Start development server
make templ        # Watch templ files
make server       # Run server only
```

## Production

```bash
docker build -t my-app .
docker run -p 8090:8090 my-app
```

## Documentation

Visit [templui.io/docs](https://templui.io/docs) for component documentation.

## License

MIT


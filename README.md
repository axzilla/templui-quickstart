# templUI Quickstart

Get started with templUI, an enterprise-ready UI component library for Go and templ. This template provides a pre-configured setup for building professional web applications with templUI components.

## Installation

For installation instructions, visit our [documentation](https://templui.io/docs/how-to-use#requirements).

## Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/axzilla/templui-quickstart.git
   cd templui-quickstart
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

3. **Configure Tailwind**
   Since we're using templUI as a package, you need to configure Tailwind to process its components:

   ```bash
   make prepare-input-css
   ```

## Development

Start the development server with hot reload:

```bash
make dev
```

Your application will be running at [http://localhost:7331](http://localhost:7331)

## Deployment

This template includes a production-ready Dockerfile for easy deployment:

```bash
# Build the image
docker build -t templui-app .

# Run the container
docker run -p 8090:8090 templui-app
```

Your application will be available at `http://localhost:8090`

## Contributing

Issues and pull requests are welcome! Please read our [contributing guidelines](https://github.com/axzilla/templui/blob/main/CONTRIBUTING.md) before submitting a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

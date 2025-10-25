# wasmbuild

A command-line tool for building and serving Go WebAssembly (WASM) applications with live reload support.

## Features

- **Build WASM applications** from Go source code
- **Development server** with live reload
- **Dependency tracking** with automatic recompilation
- **Asset management** for static files

## Installation

```bash
go install github.com/djthorpe/go-wasmbuild/cmd/wasmbuild@latest
```

Or build from source (places into a 'build' directory):

```bash
git clone https://github.com/djthorpe/go-wasmbuild.git
cd go-wasmbuild
make wasmbuild
```

## Usage

Use `wasmbuild --help` to see available commands and flags. In your project directory,  you will need to include a configuration file (by default `wasmbuild.yaml`) or use the ``--config`` flag to specify a different file. See below
for more details on the configuration file.

### Build Command

Compile a WASM application and copy all necessary files to the output directory.

```bash
wasmbuild build [PATH] [flags]
```

**Arguments:**

- `PATH` - Source path to WASM application, optional (default: current directory)

**Flags:**

- `-o, --output PATH` - Output directory, optional. If not specified, a temporary directory will be created and used.
- `-v, --verbose` - Enable verbose output, optional.
- `--config FILE` - Path to configuration YAML file, optional. By default: `wasmbuild.yaml` which is located in the source path directory.
- `--go PATH` - Optional, path to go tool (default: `go`)
- `--go-flags="FLAGS"` - Optional, additional flags to pass to `go build`

**Example:**

```bash
# Build current directory, output to temporary dir
wasmbuild build

# Build specific path with custom output
wasmbuild build ./cmd/wasm/helloworld -o ./dist

# Build with verbose output
wasmbuild build -v
```

### Serve Command

Start a development server with optional live reload support.

```bash
wasmbuild serve [PATH] [flags]
```

**Arguments:**

- `PATH` - Source path to WASM application, optional (default: current directory)

**Flags:**

- `-w, --watch` - Watch for changes in dependencies and trigger automatic recompilation
- `--listen ADDRESS` - Address to listen on, optional (default: `localhost:9090`)
- `-v, --verbose` - Enable verbose output, optional
- `--config FILE` - Path to configuration YAML file, optional. By default: `wasmbuild.yaml` which is located in the source path directory.
- `--go PATH` - Optional, path to go tool (default: `go`)
- `--go-flags="FLAGS"` - Optional, additional flags to pass to `go build`

**Example:**

```bash
# Serve with live reload
wasmbuild serve -w

# Serve on custom address, which allows for non-local access
wasmbuild serve --listen 0.0.0.0:8080

# Serve specific application with watch
wasmbuild serve ./cmd/wasm/bootstrap-app -w
```

#### Dep Command

Display dependency information for a WASM application.

```bash
wasmbuild dep [PATH] [flags]
```

**Arguments:**

- `PATH` - Source path to WASM application (default: current directory)

**Flags:**

- `-w, --watch` - Watch for changes in dependencies
- `-v, --verbose` - Enable verbose output

**Example:**

```bash
# Show dependencies
wasmbuild dep

# Watch dependencies
wasmbuild dep -w
```

### Configuration File

Create a `wasmbuild.yaml` file in your project root:

```yaml
# Custom variables for HTML template
vars:
  Title: "My WASM App"
  Header: "<style>body { font-family: sans-serif; }</style>"
  Footer: "<script>console.log('App loaded');</script>"
  SecretToken: "${SECRET_TOKEN_FROM_ENV_VAR}"

# Optional: Static assets to copy to output directory
assets:
  - assets/css
  - assets/images
```

**Template Variables:**

- `Title` - HTML page title (defaults to directory name)
- `Header` - Custom HTML injected into `<head>`
- `Footer` - Custom HTML injected at the end of the `</body>`

Other variables can be included as needed. Each variable can expand environment variables using `${VAR_NAME}` syntax.

## Development Workflow

### Basic Workflow

1. Create your WASM application:

   ```go
   // main.go
   package main
   
   import "fmt"
   
   func main() {
       fmt.Println("Hello, WASM!")
       select {} // Keep program running
   }
   ```

2. Start the development server:

   ```bash
   wasmbuild serve -w
   ```

3. Open your browser to `http://localhost:9090`

4. Edit your source files - the browser will automatically reload

In order to create a production build, use the `wasmbuild build` command to compile and package your application.

## License

See LICENSE file in the repository root.

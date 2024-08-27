# Hathora CLI for CI

This repo houses the CLI that can be used to deploy your game server builds to Hathora Cloud in your CI/CD pipelines.

## Docs

For documentation on how to use this CLI, check out our [docs.](https://hathora.dev/docs/guides/ci-cd)

## Development

### Run the CLI locally

To run the CLI locally, execute the following:

```sh
go run hathora/main.go --help
```

### Running tests

To run tests:

```sh
make test
```

### Building binaries

To build a CLI binary, you can run the following:

```sh
make build
```

The target OS and architecture can be specified by setting the `OS` and `ARCH` environment variables, respectively. For example, to build a binary for the `linux` OS and the `amd64` architecture, you can do the following:

```sh
TARGETOS=linux TARGETARCH=amd64 make build
```

> [!NOTE]
> When building on windows you'll want the built binaries to include the `.exe` suffix. To achieve this, you can use the `BINARY_SUFFIX`
> variable, e.g.:
>
> ```sh
> TARGETOS=windows TARGETARCH=amd64 BINARY_SUFFIX=.exe make build
> ```

The binary will be available in the `bin` directory.

```sh
./bin/hathora-* --help
```

### Rebuilding binaries

To quickly rebuild a CLI binary, you can use:

```sh
make clean && make build
```

### Regenerating the SDK

The SDK can be regenerated based on the OpenAPI spec hosted at `https://hathora.dev/swagger.json`. To do this, run the following command from the root of the project. You must have the speakeasy CLI installed.

```sh
make sdk-clean
rm internal/sdk/go.mod
```

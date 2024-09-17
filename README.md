# Hathora CLI for CI

This repo houses the CLI that can be used to deploy your game server builds to Hathora Cloud in your CI/CD pipelines.

## Hathora CLI Upgrade Notes

This covers the relevant notes for users upgrading from previous CLI versions. We aim to cover compatibility changes that you must be aware of.

For more details about changes on each release please refer to the [Official Release Notes](https://github.com/hathora/ci/releases).

### Hathora CLI v1.0.0

Hathora CLI v1.0.0 now uses Hathora AppsV2, ProcessesV3, BuildsV3 and DeploymentsV3 APIs. Using these APIs for a given application implicitly migrates it to Global Builds.

**Nothing in your CLI command needs to be changed to migrate**. However,  usage of old SDKs should be updated to the latest versions (AppsV2, ProcessesV3, BuildsV3 and DeploymentsV3). This should be done before upgrading your continuous deployment to use Global Builds (CLI v1.0.0).

When you create build with CLI v1.0.0, the following will be impacted on old APIs:

- listing builds - new builds will not appear

When you create a deployment with CLI v1.0.0, the following will be impacted on old APIs:

- listing applications - deployment id and build id returned will be invalid
- listing deployments - new deployments will not appear
- listing processes - deployment id returned will be invalid

Notes:

- *appId* being passed into `hathora build` will now be ignored - it can safely be removed

### Hathora CLI v0.x.x

Hathora CLI v0.x.x is no longer compatible with builds and application deployments created via the Hathora UI Console.

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

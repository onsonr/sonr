### ⚡️ Installing Binaries

> [!NOTE]
>
> ```bash
> wget https://github.com/sonrhq/sonr/releases/latest/download/install.sh
> ```

Binaries for Linux and Darwin (amd64 and arm64) are available below.

### 🔨 Build from source

If you prefer to build from source, you can use the following commands:

```bash
git clone https://github.com/sonrhq/sonr
cd core && git checkout v{{ .Version }}
make install
# Then run sonrd
sonrd version
# v{{ .Version }}
```

### 🐳 Run with Docker

As an alternative to installing and running sonrd on your system, you may run sonrd in a Docker container.
The following Docker images are available in our registry:

| Image Name                               | Base                         | Description                       | Registry Source                                                                   |
| ---------------------------------------- | ---------------------------- | --------------------------------- | --------------------------------------------------------------------------------- |
| `sonrhq/sonrd:{{ .Version }}`            | `distroless/static-debian11` | Default image based on Distroless | [ghcr.io/sonrhq/sonrd](https://ghcr.io/sonrhq/sonrd:latest)                       |
| `sonrhq/sonrd:{{ .Version }}-standalone` | `distroless/static-debian11` | Standalone testing node image     | [ghcr.io/sonrhq/sonrd-standalone](https://ghcr.io/sonrhq/sonrd:latest-standalone) |
| `sonrhq/faucet:{{ .Version }}`           | `distroless/static-debian11` | Development environment faucet    | [ghcr.io/sonr-io/sonr-faucet](https://ghcr.io/sonr-io/sonr-faucet:latest)         |

Example run:

```bash
docker run sonrhq/sonrd:{{ .Version }} version
# v{{ .Version }}
```

All the images support `amd64` architectures.

# Installation

## Using Homebrew (Mac and linux)

```bash
brew install kishaningithub/tap/terralens
```

## Using docker

```bash
alias terralens="docker run -i ghcr.io/kishaningithub/terralens:latest"

terraform show -json | terralens
```

## Others

Head over to the [releases page](https://github.com/kishaningithub/terralens/releases) and download a binary for your
platform.
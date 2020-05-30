# drone-convert-starlark

A conversion plugin that provides optional support for generating pipeline configuration files via Starlark scripting. _Please note this project requires Drone server version 1.4 or higher._

Issue Tracker:
https://github.com/yckao/drone-convert-advanced/issues

## Installation

Create a shared secret:

```text
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

Download and run the plugin:

```text
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --env=BIND=:3000 \
  --env=DRONE_GITHUB_TOKEN=d6a9d0e91cf587e84b41278be630909352f1a2ff\
  --restart=always \
  --name=convert-advanced yckao/drone-convert-advanced
```

## Configuration

### Minimal Settings

```text
DRONE_SECRET

# One of scm settings list below.
# Like DRONE_GITHUB_TOKEN=d6a9d0e91cf587e84b41278be630909352f1a2ff
```

### Avaliable Environments

```text
DRONE_LOGS_TRACE
DRONE_LOGS_COLOR
DRONE_LOGS_PRETTY
DRONE_LOGS_TEXT

DRONE_BIND
DRONE_DEBUG
DRONE_SECRET

DRONE_BITBUCKET_USERNAME
DRONE_BITBUCKET_APP_PASSWORD
DRONE_BITBUCKET_DEBUG
DRONE_BITBUCKET_SKIP_VERIFY

DRONE_GITEA_SERVER
DRONE_GITEA_TOKEN
DRONE_GITEA_DEBUG
DRONE_GITEA_SKIP_VERIFY

DRONE_GITHUB_SERVER
DRONE_GITHUB_API_SERVER
DRONE_GITHUB_TOKEN
DRONE_GITHUB_USER_RATELIMIT
DRONE_GITHUB_DEBUG
DRONE_GITHUB_SKIP_VERIFY

DRONE_GITLAB_SERVER
DRONE_GITLAB_TOKEN
DRONE_GITLAB_DEBUG
DRONE_GITLAB_SKIP_VERIFY

DRONE_GOGS_SERVER
DRONE_GOGS_TOKEN
DRONE_GOGS_DEBUG
DRONE_GOGS_SKIP_VERIFY
```

## Loads/Imports

Starlark/Bazel has support for [loading](https://docs.bazel.build/versions/master/build-ref.html#load) extensions (modules). This is useful for cases where you'd like to share re-usable logic. drone-convert-advanced currently supports the ability to load starlark file in executing repo.

For example:

```python
# Relative load.
load("//:steps_extension.star", "example_step")

# Absolute load.
load("//subpackage:pipelines.star", "example_pipeline")

def main(ctx):
    return example_pipeline("sample", steps = example_step())
```

In this case. The first `load` imports an extension named `steps_extension.star` and extracts the `example_step` symbol for use in our Drone pipeline. The second example drills into the `subpackage` directory to load an extension called `pipelines.star`, then extracts a symbol named `example_pipeline`.

## Testing

Currently we can _NOT_ use `drone plugins convert` to test this plugin because load from repo will failed.

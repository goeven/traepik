# Traepik - Traefik Proper Plugins Package

## Disclaimer: This is an EXEPRIMENTAL implementation and was not tested with production workloads.

## Introduction

Traepik (træpik) adds support for compiled plugins to the traefik command. This could also include private plugins, as the control, but also the effort of compiling `traefik`, is shifted towards you. For known caveats see the [Caveats section](#caveats).

The benefits over the current Traefik plugins approach are:
- The plugins Go code is compiled as part of building the actual `traefik` command
- The plugins don't have to be distributed to Traefik Pilot.


## How it works

This module contains a small re-write of the Traefik command's [`main` package](https://github.com/traefik/traefik/blob/v2.3.2/cmd/traefik/traefik.go), transforming it into an importable Go package. This implementation also adds support to pass Traefik plugins directly to the command's constructor, making it possible to have custom plugins that don't have to be distributed through Traefik Pilot.

_NOTE: Traefik's plugins system still works, but plugins given to `cmd.New` will take precedence over the ones configured with Traefik Pilot if you use the same name._

### Usage

If you want to dive into some full code examples, take a look in the `examples` directory.

#### Run the Traefik Command

```go
package main

import (
	"github.com/goeven/traepik/v2/pkg/cmd"
	"github.com/goeven/traepik/v2/pkg/plugins"
)

func main() {
    myPlugins := map[string]plugins.PluginBuilder{ /* ... */ }

	traefikCmd, err := cmd.New(myPlugins)
	if err != nil { /* ... */ }

	if err := traefikCmd.Execute(); err != nil { /* ... */ }
}
```

#### Use a Custom Plugin Definition

A plugin represents an entry in the map passed to `cmd.New`, the key being the plugin's name – used to load the plugin's configuration from the dynamic configuration – and the value is a `plugins.PluginBuilder`, described below.

##### Logging Plugin Example

Code:
```go
map[string]plugins.PluginBuilder{
    "loggingPlugin": plugins.PluginBuilderFunc(func(config map[string]interface{}, middlewareName string) (cmd.PluginConstructor, error) {
        msg, ok := config["log-message"]
        if !ok {
            msg = "hello from the logging plugin"
        }
        return func(ctx context.Context, next http.Handler) (http.Handler, error) {
            return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                log.Printf("traepik: %s", msg)
                next.ServeHTTP(w, r)
            }), nil
        }, nil
    }),
}
```

Dynamic Config (with `file` provider):
```yaml
middlewares:
    exampleMiddleware:
        plugin:
            loggingPlugin:
                log-message: hi-ya

```

## Caveats

There are some caveats, including, but not limited to:

- Responsibility for building the package containing the `traefik` command shifts to the person also implementing the custom plugins (basically, you).
- The `replace` directives from this repo have to also be added to your `go.mod` file.
- The implementation relies on the particularities of how plugins are implemented right now in Traefik, and that could change anytime.

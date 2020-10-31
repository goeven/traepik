package plugins

import (
	"fmt"

	"github.com/traefik/traefik/v2/pkg/plugins"
)

// Constructor is a type alias for the Constructor type in github.com/traefik/traefik/v2/pkg/plugins.
// Useful for not importing traefik directly into a project.
type Constructor = plugins.Constructor

// PluginBuilder defines a Build method for a plugin.
type PluginBuilder interface {
	Build(config map[string]interface{}, middlewareName string) (Constructor, error)
}

// PluginBuilderFunc is a type adapter helper, similar to http.HandlerFunc. It is used to
// ease provinding a function as a PluginBuilder.
type PluginBuilderFunc func(config map[string]interface{}, middlewareName string) (Constructor, error)

// Build implements PluginBuilder for the PluginBuilderFunc type.
func (b PluginBuilderFunc) Build(config map[string]interface{}, middlewareName string) (Constructor, error) {
	return b(config, middlewareName)
}

// PluginMapBuilder is middleware.PluginsBuilder backed by a map
type PluginMapBuilder map[string]PluginBuilder

// Build implements the github.com/traefik/traefik/v2/pkg/server/middleware.PluginsBuilder interface.
func (b PluginMapBuilder) Build(pName string, config map[string]interface{}, middlewareName string) (Constructor, error) {
	p, ok := b[pName]
	if !ok {
		return nil, fmt.Errorf("plugin: unknown plugin type: %s", pName)
	}
	return p.Build(config, middlewareName)
}

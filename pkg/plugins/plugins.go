package plugins

import (
	"github.com/traefik/traefik/v2/pkg/plugins"
	"github.com/traefik/traefik/v2/pkg/server/middleware"
)

// Constructor is a type alias for the Constructor type in github.com/traefik/traefik/v2/pkg/plugins.
// Useful for not importing traefik directly into a project.
type Constructor = plugins.Constructor

// PluginBuilder defines a Build method for a single plugin.
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

// BuilderWithBasePlugins wraps a middleware.PluginsBuilder, providing a base set of plugins.
// It implements the middleware.PluginsBuilder interface.
type BuilderWithBasePlugins struct {
	BasePlugins map[string]PluginBuilder
	Builder     middleware.PluginsBuilder
}

// Build tries to find a plugin in basePlugins, based on the passed pName. If one is found, its Build method is called. If not,
// it fallbacks to using the wrapped middleware.PluginsBuilder.
func (b BuilderWithBasePlugins) Build(pName string, config map[string]interface{}, middlewareName string) (Constructor, error) {
	p, ok := b.BasePlugins[pName]
	if !ok {
		return b.Builder.Build(pName, config, middlewareName)
	}
	return p.Build(config, middlewareName)
}

// The MIT License (MIT)

// Copyright (c) 2016-2020 Containous SAS; 2020 Traefik Labs

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/traefik/traefik/v2/pkg/config/static"
	"github.com/traefik/traefik/v2/pkg/plugins"
)

const outputDir = "./plugins-storage/"

func initPlugins(staticCfg *static.Configuration) (*plugins.Client, map[string]plugins.Descriptor, *plugins.DevPlugin, error) {
	if !isPilotEnabled(staticCfg) || !hasPlugins(staticCfg) {
		return nil, map[string]plugins.Descriptor{}, nil, nil
	}

	opts := plugins.ClientOptions{
		Output: outputDir,
		Token:  staticCfg.Pilot.Token,
	}

	client, err := plugins.NewClient(opts)
	if err != nil {
		return nil, nil, nil, err
	}

	err = plugins.Setup(client, staticCfg.Experimental.Plugins, staticCfg.Experimental.DevPlugin)
	if err != nil {
		return nil, nil, nil, err
	}

	return client, staticCfg.Experimental.Plugins, staticCfg.Experimental.DevPlugin, nil
}

func isPilotEnabled(staticCfg *static.Configuration) bool {
	return staticCfg.Pilot != nil && staticCfg.Pilot.Token != ""
}

func hasPlugins(staticCfg *static.Configuration) bool {
	return staticCfg.Experimental != nil &&
		(len(staticCfg.Experimental.Plugins) > 0 || staticCfg.Experimental.DevPlugin != nil)
}

package main

import (
	"context"

	"github.com/urfave/cli"
)

type ConfigSource interface {
	List(ctx context.Context) ([]string, error)
}

const metadataConfigSourceKey = "configSource"

func appendConfigSource(c *cli.Context, src ConfigSource) {
	m, ok := c.App.Metadata[metadataConfigSourceKey].([]ConfigSource)
	if !ok {
		m = []ConfigSource{}
	}
	c.App.Metadata[metadataConfigSourceKey] = append(m, src)
}

func configSourcesFromContext(c *cli.Context) []ConfigSource {
	if src, ok := c.App.Metadata[metadataConfigSourceKey].([]ConfigSource); ok {
		return src
	}
	return nil
}

package main

import "context"

type StringSliceConfigSource struct {
	Config []string
}

func NewStringSliceConfigSource(config []string) *StringSliceConfigSource {
	src := &StringSliceConfigSource{
		Config: make([]string, len(config)),
	}
	copy(src.Config, config)
	return src
}

func (src *StringSliceConfigSource) List(ctx context.Context) ([]string, error) {
	return src.Config, nil
}

package converter

import (
	"context"

	"github.com/yckao/drone-convert-advanced/core"
)

func Combine(services ...core.ConvertService) core.ConvertService {
	return &combined{services}
}

type combined struct {
	sources []core.ConvertService
}

func (c *combined) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	for _, source := range c.sources {
		config, err := source.Convert(ctx, req)
		if err != nil {
			return nil, err
		}
		if config == nil {
			continue
		}
		if config.Data == "" {
			continue
		}
		return config, nil
	}
	return req.Config, nil
}

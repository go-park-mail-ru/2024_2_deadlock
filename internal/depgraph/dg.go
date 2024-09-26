package depgraph

import (
	"go.uber.org/zap"
)

type DepGraph struct {
	logger *dgEntity[*zap.SugaredLogger]
}

func NewDepGraph() *DepGraph {
	return &DepGraph{
		logger: &dgEntity[*zap.SugaredLogger]{},
	}
}

func (d *DepGraph) GetLogger() (*zap.SugaredLogger, error) {
	return d.logger.get(func() (*zap.SugaredLogger, error) {
		logger := zap.Must(zap.NewDevelopment()).Sugar()
		return logger, nil
	})
}

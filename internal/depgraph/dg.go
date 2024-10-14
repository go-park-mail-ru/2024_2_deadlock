package depgraph

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type DepGraph struct {
	logger   *dgEntity[*zap.SugaredLogger]
	validate *dgEntity[*validator.Validate]
}

func NewDepGraph() *DepGraph {
	return &DepGraph{
		logger:   &dgEntity[*zap.SugaredLogger]{},
		validate: &dgEntity[*validator.Validate]{},
	}
}

func (d *DepGraph) GetLogger() (*zap.SugaredLogger, error) {
	return d.logger.get(func() (*zap.SugaredLogger, error) {
		logger := zap.Must(zap.NewDevelopment()).Sugar()
		return logger, nil
	})
}

func (d *DepGraph) GetValidator() (*validator.Validate, error) {
	return d.validate.get(func() (*validator.Validate, error) {
		validate := validator.New(validator.WithRequiredStructEnabled())
		return validate, nil
	})
}

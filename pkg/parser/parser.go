package parser

import (
	"context"
	"github.com/scagogogo/sca-base-module-components/pkg/models"
)

// Parser 表示一个解析依赖的Parser
type Parser[EcosystemDependencyInput, ProjectEcosystem, ModuleEcosystem, ComponentEcosystem, ComponentDependencyEcosystem any] interface {

	// GetName parser的名字，即使是同一个Ecosystem也可能会有多个parser，会根据条件选择parser
	GetName() string

	// Init 初始化
	Init(ctx context.Context) error

	// Parse 解析依赖
	Parse(ctx context.Context, input EcosystemDependencyInput) (*models.Project[ProjectEcosystem, ModuleEcosystem, ComponentEcosystem, ComponentDependencyEcosystem], error)

	// Close 关闭解析器
	Close(ctx context.Context) error
}

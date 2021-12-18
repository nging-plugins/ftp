package ftpmanager

import (
	"gitee.com/hkcom/nging-core/application/library/config/extend"
	"github.com/admpub/nging/v4/application/library/config/cmder"
	"github.com/admpub/nging/v4/application/library/module"

	"github.com/nging-plugins/ftpmanager/pkg/handler"
	pluginCmder "github.com/nging-plugins/ftpmanager/pkg/library/cmder"
	"github.com/nging-plugins/ftpmanager/pkg/library/setup"
)

const ID = `ftp`

var Module = module.Module{
	Startup: `ftpserver`,
	Extend: map[string]extend.Initer{
		`ftpserver`: pluginCmder.Initer,
	},
	Cmder: map[string]cmder.Cmder{
		`ftpserver`: pluginCmder.New(),
	},
	TemplatePath: map[string]string{
		ID: `ftpmanager/template/backend`,
	},
	AssetsPath:    []string{},
	SQLCollection: setup.RegisterSQL,
	Navigate:      RegisterNavigate,
	Route:         handler.RegisterRoute,
	DBSchemaVer:   0.0000,
}

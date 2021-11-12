package pipelines

import (
	"github.com/kubesphere/kubekey/pkg/certs"
	"github.com/kubesphere/kubekey/pkg/common"
	"github.com/kubesphere/kubekey/pkg/core/module"
	"github.com/kubesphere/kubekey/pkg/core/pipeline"
)

func CheckCertsPipeline(runtime *common.KubeRuntime) error {
	m := []module.Module{
		&certs.CheckCertsModule{},
		&certs.PrintClusterCertsModule{},
	}

	p := pipeline.Pipeline{
		Name:    "CheckCertsPipeline",
		Modules: m,
		Runtime: runtime,
	}
	if err := p.Start(); err != nil {
		return err
	}
	return nil
}

func CheckCerts(args common.Argument) error {
	var loaderType string
	if args.FilePath != "" {
		loaderType = common.File
	} else {
		loaderType = common.AllInOne
	}

	runtime, err := common.NewKubeRuntime(loaderType, args)
	if err != nil {
		return err
	}

	if err := CheckCertsPipeline(runtime); err != nil {
		return err
	}
	return nil
}
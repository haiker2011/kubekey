package templates

import (
	kubekeyapiv1alpha2 "github.com/kubesphere/kubekey/apis/kubekey/v1alpha2"
	"github.com/kubesphere/kubekey/pkg/core/util"
	"github.com/lithammer/dedent"
	"text/template"
)

// Cluster defines the template of cluster configuration file default.
var Cluster = template.Must(template.New("Cluster").Parse(
	dedent.Dedent(`
apiVersion: kubekey.kubesphere.io/v1alpha2
kind: Cluster
metadata:
  name: {{ .Options.Name }}
spec:
  hosts:
  - {name: node1, address: 172.16.0.2, internalAddress: 172.16.0.2, user: ubuntu, password: Qcloud@123}
  - {name: node2, address: 172.16.0.3, internalAddress: 172.16.0.3, user: ubuntu, password: Qcloud@123}
  roleGroups:
    etcd:
    - node1
    master: 
    - node1
    worker:
    - node1
    - node2
  controlPlaneEndpoint:
    ##Internal loadbalancer for apiservers 
    #internalLoadbalancer: haproxy

    domain: lb.kubesphere.local
    address: ""
    port: 6443
  kubernetes:
    version: {{ .Options.KubeVersion }}
    clusterName: cluster.local
  network:
    plugin: calico
    kubePodsCIDR: 10.233.64.0/18
    kubeServiceCIDR: 10.233.0.0/18
  registry:
    registryMirrors: []
    insecureRegistries: []
  addons: []

{{ if .Options.KubeSphereEnabled }}
{{ .Options.KubeSphereConfigMap }}
{{ end }}
    `)))

// Options defines the parameters of cluster configuration.
type Options struct {
	Name                string
	KubeVersion         string
	KubeSphereEnabled   bool
	KubeSphereConfigMap string
}

// GenerateCluster is used to generate cluster configuration content.
func GenerateCluster(opt *Options) (string, error) {
	return util.Render(Cluster, util.Data{
		"KubeVersion": kubekeyapiv1alpha2.DefaultKubeVersion,
		"Options":     opt,
	})
}
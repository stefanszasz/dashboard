package auth

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const KubeDashboardSecEnvKey = "KUBE_DASHBOARD_SEC_NAMESPACE"

func ParseKubeConfigContent(inFile string) {
	_, found := os.LookupEnv(KubeDashboardSecEnvKey)
	if found {
		return
	}

	cont, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Panic("Cannot find file " + inFile)
	}
	kubeConfig := new(KubeConfig)
	if err := yaml.Unmarshal(cont, kubeConfig); err != nil {
		log.Panic("Cannot parse yaml")
	}

	for _, context := range kubeConfig.Contexts {
		if context.Name == kubeConfig.CurrentContext {
			if context.Context.Namespace != "" {
				SecuredNamespace = context.Context.Namespace
			}
		}
	}
}

func InitSecurityNamespace() {
	envNs, found := os.LookupEnv(KubeDashboardSecEnvKey)
	if found {
		SecuredNamespace = envNs
	}
}

package config

import (
	"github.com/cloudnative-pg/cnpg-i-machinery/pkg/pluginhelper/common"
	"github.com/go-viper/mapstructure/v2"
	"k8s.io/apimachinery/pkg/types"
)

// Configuration represents the plugin configuration parameters
type Configuration struct {
	// The name of the PodSpecTemplate object to be used for the
	// PostgreSQL instance
	InstanceTemplateName string `mapstructure:"instanceTemplateName"`

	// The namespace of the PodSpecTemplate object to be used for
	// the PostgreSQL instance
	InstanceTemplateNamespace string `mapstructure:"instanceTemplateNamespace"`
}

// FromParameters builds a plugin configuration from the configuration parameters
func FromParameters(
	helper *common.Plugin,
) *Configuration {
	var result Configuration
	mapstructure.Decode(helper.Parameters, &result)
	return &result
}

// ObjectKey transforms this configuration into the object key
// ot the relative API object
func (c Configuration) InstanceTemplateObjectKey() types.NamespacedName {
	return types.NamespacedName{
		Name:      c.InstanceTemplateName,
		Namespace: c.InstanceTemplateNamespace,
	}
}

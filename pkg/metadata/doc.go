// Package metadata contains the metadata of this plugin
package metadata

import "github.com/cloudnative-pg/cnpg-i/pkg/identity"

// PluginName is the name of the plugin
const PluginName = "cnpg-i-podspec.cloudnative-pg.io"

// Data is the metadata of this plugin
var Data = identity.GetPluginMetadataResponse{
	Name:          PluginName,
	Version:       "0.0.1",
	DisplayName:   "PodSpec reconciler",
	ProjectUrl:    "https://github.com/leonardoce/cnpg-i-podspec",
	RepositoryUrl: "https://github.com/leonardoce/cnpg-i-podspec",
	License:       "Apache 2.0",
	LicenseUrl:    "https://github.com/leonardoce/cnpg-i-podspec/LICENSE",
	Maturity:      "alpha",
}

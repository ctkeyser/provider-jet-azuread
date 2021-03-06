package group

import "github.com/crossplane/terrajet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_group", func(r *config.Resource) {
		r.ShortGroup = "group"
	})
}

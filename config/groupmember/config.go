package groupmember

import "github.com/crossplane/terrajet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("azuread_group_member", func(r *config.Resource) {
		r.ShortGroup = "group_member"

		r.References["group"] = config.Reference{
			Type: "github.com/ctkeyser/provider-jet-azuread/apis/repository/v1alpha1.Group",
		}
	})
}

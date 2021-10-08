package secrets

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"get_secret_oci": resourceGetSecret(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}

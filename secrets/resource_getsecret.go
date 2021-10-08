package secrets

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/oci-go-sdk/v48/common"
	"github.com/oracle/oci-go-sdk/v48/example/helpers"
	"github.com/oracle/oci-go-sdk/v48/secrets"
)

func resourceGetSecret() *schema.Resource {
	return &schema.Resource{
		Create: resourceGetSecretCreate,
		Read:   resourceGetSecretRead,
		Update: resourceGetSecretUpdate,
		Delete: resourceGetSecretDelete,
		Schema: map[string]*schema.Schema{
			"vaultid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"secretname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"secret_content": {
				Type:     schema.TypeString,
				Computed: true,
				//Sensitive: true,
			},
		},
	}
}

func resourceGetSecretCreate(d *schema.ResourceData, m interface{}) error {

	vaultid := d.Get("vaultid").(string)
	secretname := d.Get("secretname").(string)
	//ensure that args and scrript is not null
	d.SetId(vaultid)
	d.SetId(secretname)
	//define new secret managment
	client, err := secrets.NewSecretsClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// Create a request and dependent object(s).

	req := secrets.GetSecretBundleByNameRequest{VaultId: common.String(vaultid),
		SecretName: common.String(secretname),
	}

	// Send the request using the service client
	resp, err := client.GetSecretBundleByName(context.Background(), req)
	helpers.FatalIfError(err)
	output, err := json.Marshal(resp.SecretBundle.SecretBundleContent)
	// Retrieve value from the response.
	x := map[string]string{}

	json.Unmarshal([]byte(output), &x)
	fmt.Println()
	data, err := base64.StdEncoding.DecodeString(x["content"])
	if err != nil {
		return err
	}
	d.Set("secret_content", string(data))
	return resourceGetSecretRead(d, m)
}

func resourceGetSecretRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceGetSecretUpdate(d *schema.ResourceData, m interface{}) error {

	return resourceGetSecretRead(d, m)
}

func resourceGetSecretDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

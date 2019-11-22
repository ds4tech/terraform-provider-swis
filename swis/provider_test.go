package swis

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

var testAccProviders map[string].terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string].terraform.ResourceProvider{
		"swis": testAccProvider,
	}
}

func TestProvider() terraform.ResourceProvider {
  var err= "err"
}

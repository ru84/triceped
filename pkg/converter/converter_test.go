package converter_test

import (
	"testing"

	"github.com/ru84/triceped/pkg/converter"
	"github.com/stretchr/testify/require"
)

func TestTerraformToBicep(t *testing.T) {
	expr := `resource "azurerm_resource_group" "main" {
  name     = var.resource_group_name
  location = var.location
}`

	expected := `// Generated Bicep Template

resource main 'Microsoft.azurerm_resource_group@latest' = {
  name: null
  location: null
}

`

	out, err := converter.TerraformToBicep(expr)
	require.NoError(t, err)
	require.Equal(t, expected, out)
}

func TestTerraformToBicepWithInvalidInput(t *testing.T) {
	expr := `invalid {`

	_, err := converter.TerraformToBicep(expr)
	require.Error(t, err)
	require.Contains(t, err.Error(), "parsing Terraform input")
}

package converter

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
)

// TerraformToBicep converts a Terraform configuration string to a Bicep configuration string.
func TerraformToBicep(inputTF string) (string, error) {
	// Initialize the HCL parser
	parser := hclparse.NewParser()

	// Parse the Terraform input as HCL
	file, diags := parser.ParseHCL([]byte(inputTF), "")
	if diags.HasErrors() {
		return "", fmt.Errorf("parsing Terraform input: %s", diags.Error())
	}

	// Generate the Bicep equivalent
	var bicepBuffer bytes.Buffer
	bicepBuffer.WriteString("// Generated Bicep Template\n\n")

	// Ensure the file body is not nil
	if file.Body == nil {
		return "", errors.New("parsed file body is nil")
	}

	// Define the schema for the resource block attributes
	resourceAttrSchema := &hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{Name: "name"},
			{Name: "location"},
			// TODO: Add other commonly used attributes
		},
	}

	// Define the schema for parsing blocks
	blockSchema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{Type: "resource", LabelNames: []string{"type", "name"}},
			// TODO: Add other block types ("variable", "output", etc.)
		},
	}

	// Extract the body content
	content, _, diags := file.Body.PartialContent(blockSchema)
	if diags.HasErrors() {
		return "", fmt.Errorf("parsing Terraform body: %s", diags.Error())
	}

	for _, block := range content.Blocks {
		switch block.Type {
		case "resource":
			if len(block.Labels) >= 2 {
				resourceType := block.Labels[0]
				resourceName := block.Labels[1]
				bicepBuffer.WriteString(fmt.Sprintf("resource %s 'Microsoft.%s@latest' = {\n", resourceName, resourceType))
				attrs, _, diags := block.Body.PartialContent(resourceAttrSchema)
				if diags.HasErrors() {
					return "", fmt.Errorf("parsing resource %s: %s", resourceName, diags.Error())
				}
				for attrName, attrValue := range attrs.Attributes {
					bicepBuffer.WriteString(fmt.Sprintf("  %s: %s\n", attrName, ctyToBicep(attrValue.Expr)))
				}
				bicepBuffer.WriteString("}\n\n")
			}
		default:
			// TODO: Handle other block types
		}
	}

	return bicepBuffer.String(), nil
}

// Helper function to convert Terraform cty values to Bicep equivalents
func ctyToBicep(expr hcl.Expression) string {
	val, diags := expr.Value(nil) // No variables, evaluating static value
	if diags.HasErrors() {
		return "null" // Fallback in case of errors
	}
	switch val.Type() {
	case cty.String:
		return fmt.Sprintf("'%s'", val.AsString())
	case cty.Number:
		return val.AsBigFloat().String()
	case cty.Bool:
		return fmt.Sprintf("%t", val.True())
	default:
		return "null" // Handle complex or unsupported types
	}
}

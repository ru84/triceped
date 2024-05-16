package converter

import (
	"fmt"
)

func TerraformToBicep(inputTF string) (string, error) {
	content := fmt.Sprintf("%s\n", inputTF)
	return content, nil
}

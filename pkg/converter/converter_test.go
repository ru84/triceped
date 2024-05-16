package converter_test

import (
	"testing"

	"github.com/ru84/triceped/pkg/converter"
	"github.com/stretchr/testify/require"
)

func TestTerraformToBicep(t *testing.T) {
	out, err := converter.TerraformToBicep("")
	require.NoError(t, err)
	require.Equal(t, "\n", out)
}

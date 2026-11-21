package names_test

import (
	"testing"

	"github.com/HolyShitDutch/kashanka/internal/names"
	"github.com/stretchr/testify/require"
)

func TestGenerateName(t *testing.T) {
	actual := names.GenerateName()
	require.NotEmpty(t, actual)
}

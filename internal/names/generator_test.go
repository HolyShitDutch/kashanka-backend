package names_test

import (
	"fmt"
	"testing"

	"github.com/moroz/kashanka/internal/names"
	"github.com/stretchr/testify/require"
)

func TestGenerateName(t *testing.T) {
	actual := names.GenerateName()
	fmt.Println(actual)
	require.NotEmpty(t, actual)
}

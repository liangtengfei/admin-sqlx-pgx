package captcha

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateCaptcha(t *testing.T) {
	res, err := GenerateCaptcha("math")
	require.NoError(t, err)
	require.NotEmpty(t, res)

	t.Log(res.Body)
	answer := store.Get(res.Id, false)
	t.Log(answer)
	err = VerifyCaptcha(res.Id, answer)
	require.NoError(t, err)
}

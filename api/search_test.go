package api_test

import (
	"testing"

	"github.com/derad6709org/vksdk/v2/api"
)

func TestVK_SearchGetHints(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.SearchGetHints(api.Params{
		"q":             "VK API",
		"limit":         20,
		"search_global": true,
	})
	noError(t, err)
}

package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSlackQuery(t *testing.T) {
	http, err := slackQuery{Text: "https://www.example.com/path"}.URL()
	require.Nil(t, err)
	assert.Equal(t, "www.example.com", http.Host)
	assert.Equal(t, "/path", http.Path)

	noHTTP, err := slackQuery{Text: "www.example.com/path"}.URL()
	require.Nil(t, err)
	assert.Equal(t, "www.example.com", noHTTP.Host)
	assert.Equal(t, "/path", noHTTP.Path)
}

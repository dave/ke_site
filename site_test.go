package ke_site

import (
	"testing"

	_ "kego.io/gallery/data/types"
	"kego.io/kerr/assert"
)

func TestOpen(t *testing.T) {
	assert.Equal(t, "https://www.brophy.uk/me.jpg", Site.Images["dave"].GetUrl())
	assert.Equal(t, "Bonjour le monde.", Site.Title.Localize([]string{"fr"}))
}

package ke_site

import (
	"github.com/davelondon/ke_common/images"
	"github.com/davelondon/ke_common/words"
	"kego.io/json"
	"kego.io/system"
	"reflect"
)

// Automatically created basic rule for gallery
type Gallery_rule struct {
	*system.Base
	*system.RuleBase
}

// This represents a gallery - it's just a list of images
type Gallery struct {
	*system.Base
	Images map[string]images.Image
	Title  words.Localizer
}

func init() {
	json.RegisterType("github.com/davelondon/ke_site", "@gallery", reflect.TypeOf(&Gallery_rule{}))
	json.RegisterType("github.com/davelondon/ke_site", "gallery", reflect.TypeOf(&Gallery{}))
}

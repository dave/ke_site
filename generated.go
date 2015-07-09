package ke_site

import (
	images "github.com/davelondon/ke_common/images"
	words "github.com/davelondon/ke_common/words"
	json "kego.io/json"
	system "kego.io/system"
	reflect "reflect"
)

// This represents a gallery - it's just a list of images
type Gallery struct {
	*system.Base
	Images map[string]images.Image
	Title  words.Localizer
}

// Automatically created basic rule for gallery
type Gallery_rule struct {
	*system.Base
	*system.RuleBase
}

func init() {
	json.RegisterType("github.com/davelondon/ke_site", "gallery", reflect.TypeOf(&Gallery{}))
	json.RegisterType("github.com/davelondon/ke_site", "@gallery", reflect.TypeOf(&Gallery_rule{}))
}

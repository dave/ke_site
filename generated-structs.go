package ke_site

import (
	"reflect"

	"github.com/davelondon/ke_common/images"
	"github.com/davelondon/ke_common/words"
	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for gallery1
type Gallery1_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for gallery1a
type Gallery1a_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for gallery2
type Gallery2_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for gallery2a
type Gallery2a_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for gallery2b
type Gallery2b_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for gallery3
type Gallery3_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for gallery3a
type Gallery3a_rule struct {
	*system.Base
	*system.RuleBase
}

// This represents a gallery - it's just a list of images
type Gallery1 struct {
	*system.Base
	Title system.String
}

// This represents a gallery - it's just a list of images
type Gallery1a struct {
	*system.Base
	Title system.String
}

// This represents a gallery - it's just a list of images
type Gallery2 struct {
	*system.Base
	Images map[string]*images.Photo
	Title  system.String
}

// This represents a gallery - it's just a list of images
type Gallery2a struct {
	*system.Base
	Images map[string]*images.Photo
	Title  system.String
}

// This represents a gallery - it's just a list of images
type Gallery2b struct {
	*system.Base
	Images map[string]*images.Photo
	Title  system.String
}

// This represents a gallery - it's just a list of images
type Gallery3 struct {
	*system.Base
	Images map[string]images.Image
	Title  words.Localizer
}

// This represents a gallery - it's just a list of images
type Gallery3a struct {
	*system.Base
	Images map[string]images.Image
	Title  words.Localizer
}

func init() {
	json.RegisterType("github.com/davelondon/ke_site", "@gallery1", reflect.TypeOf(&Gallery1_rule{}), 0x1e7dc2fe69100697)
	json.RegisterType("github.com/davelondon/ke_site", "@gallery1a", reflect.TypeOf(&Gallery1a_rule{}), 0x71f3e78aef58e8d3)
	json.RegisterType("github.com/davelondon/ke_site", "@gallery2", reflect.TypeOf(&Gallery2_rule{}), 0x34c477fbf667c293)
	json.RegisterType("github.com/davelondon/ke_site", "@gallery2a", reflect.TypeOf(&Gallery2a_rule{}), 0x15a72f5e15ac5b15)
	json.RegisterType("github.com/davelondon/ke_site", "@gallery2b", reflect.TypeOf(&Gallery2b_rule{}), 0x7a55f651c3e84b3e)
	json.RegisterType("github.com/davelondon/ke_site", "@gallery3", reflect.TypeOf(&Gallery3_rule{}), 0x5dbf985bda22bd2c)
	json.RegisterType("github.com/davelondon/ke_site", "@gallery3a", reflect.TypeOf(&Gallery3a_rule{}), 0x52157a49e1a4a1d6)
	json.RegisterType("github.com/davelondon/ke_site", "gallery1", reflect.TypeOf(&Gallery1{}), 0x1e7dc2fe69100697)
	json.RegisterType("github.com/davelondon/ke_site", "gallery1a", reflect.TypeOf(&Gallery1a{}), 0x71f3e78aef58e8d3)
	json.RegisterType("github.com/davelondon/ke_site", "gallery2", reflect.TypeOf(&Gallery2{}), 0x34c477fbf667c293)
	json.RegisterType("github.com/davelondon/ke_site", "gallery2a", reflect.TypeOf(&Gallery2a{}), 0x15a72f5e15ac5b15)
	json.RegisterType("github.com/davelondon/ke_site", "gallery2b", reflect.TypeOf(&Gallery2b{}), 0x7a55f651c3e84b3e)
	json.RegisterType("github.com/davelondon/ke_site", "gallery3", reflect.TypeOf(&Gallery3{}), 0x5dbf985bda22bd2c)
	json.RegisterType("github.com/davelondon/ke_site", "gallery3a", reflect.TypeOf(&Gallery3a{}), 0x52157a49e1a4a1d6)
}

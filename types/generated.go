package types

import (
	images "github.com/davelondon/ke_common/images"
	_ "github.com/davelondon/ke_common/images/types"
	_ "github.com/davelondon/ke_common/units/types"
	words "github.com/davelondon/ke_common/words"
	_ "github.com/davelondon/ke_common/words/types"
	system "kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr8728906240 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery", Exists: true}, Description: "This represents a gallery - it's just a list of images", Rules: []system.Rule(nil)}
	ptr8728909952 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Force https on all photos in the gallery by directly accessing the value of the protocol field", Rules: []system.Rule(nil)}
	ptr8728885312 := &system.RuleBase{Selector: "{images:photo}>.protocol"}
	ptr8728725920 := &system.String_rule{Base: ptr8728909952, RuleBase: ptr8728885312, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{Value: "https", Exists: true}, Pattern: system.String{}, Format: system.String{}}
	ptr8728910080 := &system.Base{Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@image", Exists: true}, Id: system.Reference{}, Description: "Force https on all images in the gallery using a custom validator", Rules: []system.Rule(nil)}
	ptr8728885760 := &system.RuleBase{Selector: "{images:image}"}
	ptr8728885696 := &images.Image_rule{Base: ptr8728910080, RuleBase: ptr8728885760, Secure: system.Bool{Value: true, Exists: true}}
	ptr8728909824 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Rules: []system.Rule{ptr8728725920, ptr8728885696}}
	ptr8728909696 := &system.Base{Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@image", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728884832 := &images.Image_rule{Base: ptr8728909696, RuleBase: (*system.RuleBase)(nil), Secure: system.Bool{}}
	ptr8728751232 := &system.Map_rule{Base: ptr8728909824, RuleBase: (*system.RuleBase)(nil), Items: ptr8728884832, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728910208 := &system.Base{Type: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "@localizer", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729211632 := &words.Localizer_rule{Base: ptr8728910208, RuleBase: (*system.RuleBase)(nil)}
	ptr8728977520 := &system.Type{Base: ptr8728906240, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8728751232, "title": ptr8729211632}, Rule: (*system.Type)(nil)}
	ptr8728910336 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "@gallery", Exists: true}, Description: "Automatically created basic rule for gallery", Rules: []system.Rule(nil)}
	ptr8728978192 := &system.Type{Base: ptr8728910336, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	system.RegisterType("github.com/davelondon/ke_site", "gallery", ptr8728977520)
	system.RegisterType("github.com/davelondon/ke_site", "@gallery", ptr8728978192)
}

package types

import (
	"github.com/davelondon/ke_common/images"
	_ "github.com/davelondon/ke_common/images/types"
	_ "github.com/davelondon/ke_common/units/types"
	"github.com/davelondon/ke_common/words"
	_ "github.com/davelondon/ke_common/words/types"
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Base{Description: "Automatically created basic rule for gallery", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "@gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Base: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Base{Description: "Force https on all photos in the gallery by directly accessing the value of the protocol field", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr4 := &system.RuleBase{Selector: "{images:photo}>.protocol"}
	ptr5 := &system.String_rule{Base: ptr3, RuleBase: ptr4, Default: system.String{}, Enum: []string(nil), Equal: system.String{Value: "https", Exists: true}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr6 := &system.Base{Description: "Force https on all images in the gallery using a custom validator", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@image", Exists: true}}
	ptr7 := &system.RuleBase{Selector: "{images:image}"}
	ptr8 := &images.Image_rule{Base: ptr6, RuleBase: ptr7, Secure: system.Bool{Value: true, Exists: true}}
	ptr9 := &system.Base{Id: system.Reference{}, Rules: []system.Rule{ptr5, ptr8}, Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr10 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@image", Exists: true}}
	ptr11 := &images.Image_rule{Base: ptr10, RuleBase: (*system.RuleBase)(nil), Secure: system.Bool{}}
	ptr12 := &system.Map_rule{Base: ptr9, RuleBase: (*system.RuleBase)(nil), Items: ptr11, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr13 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "@localizer", Exists: true}}
	ptr14 := &words.Localizer_rule{Base: ptr13, RuleBase: (*system.RuleBase)(nil)}
	ptr15 := &system.Type{Base: ptr2, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr12, "title": ptr14}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.RegisterType("github.com/davelondon/ke_site", "@gallery", ptr1, 0x1891c27cbbdefbc5)
	system.RegisterType("github.com/davelondon/ke_site", "gallery", ptr15, 0x1891c27cbbdefbc5)
}

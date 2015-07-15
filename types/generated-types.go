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
	ptr0 := &system.Base{Description: "Automatically created basic rule for gallery1", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "@gallery1", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Base: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Base{Description: "Automatically created basic rule for gallery1a", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "@gallery1a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Type{Base: ptr2, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr4 := &system.Base{Description: "Automatically created basic rule for gallery2", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "@gallery2", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr5 := &system.Type{Base: ptr4, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr6 := &system.Base{Description: "Automatically created basic rule for gallery2a", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "@gallery2a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr7 := &system.Type{Base: ptr6, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr8 := &system.Base{Description: "Automatically created basic rule for gallery2b", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "@gallery2b", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr9 := &system.Type{Base: ptr8, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr10 := &system.Base{Description: "Automatically created basic rule for gallery3", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "@gallery3", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr11 := &system.Type{Base: ptr10, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr12 := &system.Base{Description: "Automatically created basic rule for gallery3a", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "@gallery3a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr13 := &system.Type{Base: ptr12, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr14 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery1", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr16 := &system.String_rule{Base: ptr15, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr17 := &system.Type{Base: ptr14, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"title": ptr16}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr18 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery1a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr19 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr20 := &system.String_rule{Base: ptr19, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 20, Exists: true}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr21 := &system.Type{Base: ptr18, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"title": ptr20}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr22 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery2", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr23 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr24 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@photo", Exists: true}}
	ptr25 := &images.Photo_rule{Base: ptr24, RuleBase: (*system.RuleBase)(nil)}
	ptr26 := &system.Map_rule{Base: ptr23, RuleBase: (*system.RuleBase)(nil), Items: ptr25, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr27 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr28 := &system.String_rule{Base: ptr27, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr29 := &system.Type{Base: ptr22, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr26, "title": ptr28}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr30 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery2a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr31 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr32 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr33 := &system.RuleBase{Selector: ".protocol"}
	ptr34 := &system.String_rule{Base: ptr32, RuleBase: ptr33, Default: system.String{}, Enum: []string(nil), Equal: system.String{Value: "https", Exists: true}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr35 := &system.Base{Id: system.Reference{}, Rules: []system.Rule{ptr34}, Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@photo", Exists: true}}
	ptr36 := &images.Photo_rule{Base: ptr35, RuleBase: (*system.RuleBase)(nil)}
	ptr37 := &system.Map_rule{Base: ptr31, RuleBase: (*system.RuleBase)(nil), Items: ptr36, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr38 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr39 := &system.String_rule{Base: ptr38, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr40 := &system.Type{Base: ptr30, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr37, "title": ptr39}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr41 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr42 := &system.RuleBase{Selector: "{images:photo} {system:int}.width"}
	ptr43 := &system.Int_rule{Base: ptr41, RuleBase: ptr42, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 800, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr44 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery2b", Exists: true}, Rules: []system.Rule{ptr43}, Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr45 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr46 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@photo", Exists: true}}
	ptr47 := &images.Photo_rule{Base: ptr46, RuleBase: (*system.RuleBase)(nil)}
	ptr48 := &system.Map_rule{Base: ptr45, RuleBase: (*system.RuleBase)(nil), Items: ptr47, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr49 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr50 := &system.String_rule{Base: ptr49, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr51 := &system.Type{Base: ptr44, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr48, "title": ptr50}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr52 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery3", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr53 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr54 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@image", Exists: true}}
	ptr55 := &images.Image_rule{Base: ptr54, RuleBase: (*system.RuleBase)(nil), Secure: system.Bool{}}
	ptr56 := &system.Map_rule{Base: ptr53, RuleBase: (*system.RuleBase)(nil), Items: ptr55, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr57 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "@localizer", Exists: true}}
	ptr58 := &words.Localizer_rule{Base: ptr57, RuleBase: (*system.RuleBase)(nil)}
	ptr59 := &system.Type{Base: ptr52, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr56, "title": ptr58}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr60 := &system.Base{Description: "This represents a gallery - it's just a list of images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery3a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr61 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr62 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "@image", Exists: true}}
	ptr63 := &images.Image_rule{Base: ptr62, RuleBase: (*system.RuleBase)(nil), Secure: system.Bool{Value: true, Exists: true}}
	ptr64 := &system.Map_rule{Base: ptr61, RuleBase: (*system.RuleBase)(nil), Items: ptr63, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr65 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "@localizer", Exists: true}}
	ptr66 := &words.Localizer_rule{Base: ptr65, RuleBase: (*system.RuleBase)(nil)}
	ptr67 := &system.Type{Base: ptr60, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"images": ptr64, "title": ptr66}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.RegisterType("github.com/davelondon/ke_site", "@gallery1", ptr1, 0x1e7dc2fe69100697)
	system.RegisterType("github.com/davelondon/ke_site", "@gallery1a", ptr3, 0x71f3e78aef58e8d3)
	system.RegisterType("github.com/davelondon/ke_site", "@gallery2", ptr5, 0x34c477fbf667c293)
	system.RegisterType("github.com/davelondon/ke_site", "@gallery2a", ptr7, 0x15a72f5e15ac5b15)
	system.RegisterType("github.com/davelondon/ke_site", "@gallery2b", ptr9, 0x7a55f651c3e84b3e)
	system.RegisterType("github.com/davelondon/ke_site", "@gallery3", ptr11, 0x5dbf985bda22bd2c)
	system.RegisterType("github.com/davelondon/ke_site", "@gallery3a", ptr13, 0x52157a49e1a4a1d6)
	system.RegisterType("github.com/davelondon/ke_site", "gallery1", ptr17, 0x1e7dc2fe69100697)
	system.RegisterType("github.com/davelondon/ke_site", "gallery1a", ptr21, 0x71f3e78aef58e8d3)
	system.RegisterType("github.com/davelondon/ke_site", "gallery2", ptr29, 0x34c477fbf667c293)
	system.RegisterType("github.com/davelondon/ke_site", "gallery2a", ptr40, 0x15a72f5e15ac5b15)
	system.RegisterType("github.com/davelondon/ke_site", "gallery2b", ptr51, 0x7a55f651c3e84b3e)
	system.RegisterType("github.com/davelondon/ke_site", "gallery3", ptr59, 0x5dbf985bda22bd2c)
	system.RegisterType("github.com/davelondon/ke_site", "gallery3a", ptr67, 0x52157a49e1a4a1d6)
}

package ke_site

import (
	"github.com/davelondon/ke_common/images"
	"github.com/davelondon/ke_common/units"
	"github.com/davelondon/ke_common/words"
	"kego.io/system"
)

var ptr0 = &system.Base{Description: "Rule to enforce this object always has a .uk server", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
var ptr1 = &system.RuleBase{Selector: "{images:photo}>.server"}
var ptr2 = &system.String_rule{Base: ptr0, RuleBase: ptr1, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{Value: "\\.uk$", Exists: true}}
var ptr3 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "site", Exists: true}, Rules: []system.Rule{ptr2}, Type: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery", Exists: true}}
var ptr4 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "photo", Exists: true}}
var ptr5 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/units", Name: "rectangle", Exists: true}}
var ptr6 = &units.Rectangle{Base: ptr5, Height: system.Number{Value: 600, Exists: true}, Width: system.Number{Value: 800, Exists: true}}
var ptr7 = &images.Photo{Base: ptr4, Path: system.String{Value: "/me.jpg", Exists: true}, Protocol: system.String{Value: "https", Exists: true}, Server: system.String{Value: "www.brophy.uk", Exists: true}, Size: ptr6}
var ptr8 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "translation", Exists: true}}
var ptr9 = &words.Translation{Base: ptr8, English: system.String{Value: "Hello, world.", Exists: true}, Translations: map[string]system.String{"es": system.String{Value: "Hola mundo.", Exists: true}, "fr": system.String{Value: "Bonjour le monde.", Exists: true}}}
var ptr10 = &Gallery{Base: ptr3, Images: map[string]images.Image{"dave": ptr7}, Title: ptr9}
var Site = ptr10

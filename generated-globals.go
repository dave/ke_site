package ke_site

import (
	"github.com/davelondon/ke_common/images"
	"github.com/davelondon/ke_common/units"
	"github.com/davelondon/ke_common/words"
	"kego.io/system"
)

var ptr0 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "site1", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery1", Exists: true}}
var ptr1 = &Gallery1{Base: ptr0, Title: system.String{Value: "Hello, world.", Exists: true}}
var ptr2 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "site1a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery1a", Exists: true}}
var ptr3 = &Gallery1a{Base: ptr2, Title: system.String{Value: "Hello, world.", Exists: true}}
var ptr4 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "site2", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery2", Exists: true}}
var ptr5 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "photo", Exists: true}}
var ptr6 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/units", Name: "rectangle", Exists: true}}
var ptr7 = &units.Rectangle{Base: ptr6, Height: system.Int{Value: 600, Exists: true}, Width: system.Int{Value: 800, Exists: true}}
var ptr8 = &images.Photo{Base: ptr5, Path: system.String{Value: "/dave.jpg", Exists: true}, Protocol: system.String{Value: "https", Exists: true}, Server: system.String{Value: "www.brophy.co.uk", Exists: true}, Size: ptr7}
var ptr9 = &Gallery2{Base: ptr4, Images: map[string]*images.Photo{"dave": ptr8}, Title: system.String{Value: "Hello, world.", Exists: true}}
var ptr10 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "site2a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery2a", Exists: true}}
var ptr11 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "photo", Exists: true}}
var ptr12 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/units", Name: "rectangle", Exists: true}}
var ptr13 = &units.Rectangle{Base: ptr12, Height: system.Int{Value: 600, Exists: true}, Width: system.Int{Value: 800, Exists: true}}
var ptr14 = &images.Photo{Base: ptr11, Path: system.String{Value: "/dave.jpg", Exists: true}, Protocol: system.String{Value: "https", Exists: true}, Server: system.String{Value: "www.brophy.co.uk", Exists: true}, Size: ptr13}
var ptr15 = &Gallery2a{Base: ptr10, Images: map[string]*images.Photo{"dave": ptr14}, Title: system.String{Value: "Hello, world.", Exists: true}}
var ptr16 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "site2b", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery2b", Exists: true}}
var ptr17 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "photo", Exists: true}}
var ptr18 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/units", Name: "rectangle", Exists: true}}
var ptr19 = &units.Rectangle{Base: ptr18, Height: system.Int{Value: 600, Exists: true}, Width: system.Int{Value: 800, Exists: true}}
var ptr20 = &images.Photo{Base: ptr17, Path: system.String{Value: "/dave.jpg", Exists: true}, Protocol: system.String{Value: "https", Exists: true}, Server: system.String{Value: "www.brophy.co.uk", Exists: true}, Size: ptr19}
var ptr21 = &Gallery2b{Base: ptr16, Images: map[string]*images.Photo{"dave": ptr20}, Title: system.String{Value: "Hello, world.", Exists: true}}
var ptr22 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "site3", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery3", Exists: true}}
var ptr23 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "photo", Exists: true}}
var ptr24 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/units", Name: "rectangle", Exists: true}}
var ptr25 = &units.Rectangle{Base: ptr24, Height: system.Int{Value: 600, Exists: true}, Width: system.Int{Value: 800, Exists: true}}
var ptr26 = &images.Photo{Base: ptr23, Path: system.String{Value: "/dave.jpg", Exists: true}, Protocol: system.String{Value: "https", Exists: true}, Server: system.String{Value: "www.brophy.co.uk", Exists: true}, Size: ptr25}
var ptr27 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "icon", Exists: true}}
var ptr28 = &images.Icon{Base: ptr27, Url: system.String{Value: "https://www.foo.com/james.png", Exists: true}}
var ptr29 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "translation", Exists: true}}
var ptr30 = &words.Translation{Base: ptr29, English: system.String{Value: "Hello, world.", Exists: true}, Translations: map[string]system.String{"es": system.String{Value: "Hola mundo.", Exists: true}, "fr": system.String{Value: "Bonjour le monde.", Exists: true}}}
var ptr31 = &Gallery3{Base: ptr22, Images: map[string]images.Image{"dave": ptr26, "james": ptr28}, Title: ptr30}
var ptr32 = &system.Base{Description: "This is an instance of the gallery type containing two nested images", Id: system.Reference{Package: "github.com/davelondon/ke_site", Name: "site3a", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_site", Name: "gallery3a", Exists: true}}
var ptr33 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "photo", Exists: true}}
var ptr34 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/units", Name: "rectangle", Exists: true}}
var ptr35 = &units.Rectangle{Base: ptr34, Height: system.Int{Value: 600, Exists: true}, Width: system.Int{Value: 800, Exists: true}}
var ptr36 = &images.Photo{Base: ptr33, Path: system.String{Value: "/dave.jpg", Exists: true}, Protocol: system.String{Value: "https", Exists: true}, Server: system.String{Value: "www.brophy.co.uk", Exists: true}, Size: ptr35}
var ptr37 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/images", Name: "icon", Exists: true}}
var ptr38 = &images.Icon{Base: ptr37, Url: system.String{Value: "https://www.foo.com/james.png", Exists: true}}
var ptr39 = &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "github.com/davelondon/ke_common/words", Name: "translation", Exists: true}}
var ptr40 = &words.Translation{Base: ptr39, English: system.String{Value: "Hello, world.", Exists: true}, Translations: map[string]system.String{"es": system.String{Value: "Hola mundo.", Exists: true}, "fr": system.String{Value: "Bonjour le monde.", Exists: true}}}
var ptr41 = &Gallery3a{Base: ptr32, Images: map[string]images.Image{"dave": ptr36, "james": ptr38}, Title: ptr40}
var Site1 = ptr1
var Site1a = ptr3
var Site2 = ptr9
var Site2a = ptr15
var Site2b = ptr21
var Site3 = ptr31
var Site3a = ptr41

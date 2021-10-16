package widgets

import (
    // "github.com/tyagdit/toolie/state"
    "github.com/tyagdit/toolie/tools"

    "github.com/gcla/gowid"
    // "github.com/gcla/gowid/widgets/button"
)


type Tool struct {
    Label string
    Callback func(app gowid.IApp, w gowid.IWidget)
}


var ToolMap = map[string]Tool {
    "to_b64": {"To Base64", tools.ToBase64CallBack},
    "from_b64": {"From Base64", tools.FromBase64CallBack},
    "to_hex": {"To Hex", tools.ToHexCallback},
    "from_hex": {"To Hex", tools.FromHexCallback},
    "sha1": {"SHA-1", tools.SHA1Callback},
    "sha256": {"SHA-256", tools.SHA256Callback},
    "sha512": {"SHA-512", tools.SHA512Callback},
    "md5": {"MD5", tools.MD5Callback},
    "url_en": {"Encode URL", tools.URLEncodeCallback},
    "uel_de": {"Decode URL", tools.URLDecodeCallback},
    "pretty_json": {"Beautify JSON", tools.JSONBeautifyCallback},
    "mini_json": {"Minify JSON", tools.JSONMinifyCallback},
}

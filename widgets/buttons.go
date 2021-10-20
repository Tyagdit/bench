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

var ToolSlice1 = []string {
    "to_b64",
    "from_b64",
    "to_hex",
    "from_hex",
    "sha1",
    "sha256",
    "sha512",
    "md5",
    "url_en",
    "url_de",
    "pretty_json",
    "mini_json",
}

var ToolSlice2 = []Tool {
    {"To Base64", tools.ToBase64CallBack},
    {"From Base64", tools.FromBase64CallBack},
    {"To Hex", tools.ToHexCallback},
    {"From Hex", tools.FromHexCallback},
    {"SHA-1", tools.SHA1Callback},
    {"SHA-256", tools.SHA256Callback},
    {"SHA-512", tools.SHA512Callback},
    {"MD5", tools.MD5Callback},
    {"Encode URL", tools.URLEncodeCallback},
    {"Decode URL", tools.URLDecodeCallback},
    {"Beautify JSON", tools.JSONBeautifyCallback},
    {"Minify JSON", tools.JSONMinifyCallback},
}

var ToolMap = make(map[string]Tool, len(ToolSlice1))

package tools

import (
    "encoding/hex"

    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"

    "github.com/tyagdit/toolie/state"
    "github.com/gcla/gowid"
)

func SHA1Callback(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "sha1"
    hash := sha1.New()
    hash.Write([]byte(state.Input.Text()))
    bs := hash.Sum(nil)
    txt := hex.EncodeToString(bs)
    state.Output.SetText(txt, app)
}

func SHA256Callback(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "sha256"
    hash := sha256.New()
    hash.Write([]byte(state.Input.Text()))
    bs := hash.Sum(nil)
    txt := hex.EncodeToString(bs)
    state.Output.SetText(txt, app)
}

func SHA512Callback(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "sha512"
    hash := sha512.New()
    hash.Write([]byte(state.Input.Text()))
    bs := hash.Sum(nil)
    txt := hex.EncodeToString(bs)
    state.Output.SetText(txt, app)
}

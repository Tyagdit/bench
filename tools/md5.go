package tools

import (
    "crypto/md5"
    "encoding/hex"
    "github.com/tyagdit/toolie/state"
    "github.com/gcla/gowid"
)

func MD5Callback(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "md5"
    hash := md5.New()
    hash.Write([]byte(state.Input.Text()))
    bs := hash.Sum(nil)
    txt := hex.EncodeToString(bs)
    state.Output.SetText(txt, app)
}

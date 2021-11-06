package tools

import (
    b64 "encoding/base64"
    "github.com/tyagdit/bench/state"
    "github.com/gcla/gowid"
)

func ToBase64CallBack(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "to_b64"
    txt := b64.StdEncoding.EncodeToString([]byte(state.Input.Text()))
    state.Output.SetText(txt, app)
}

func FromBase64CallBack(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "from_b64"
    txt, _ := b64.StdEncoding.DecodeString(state.Input.Text())
    state.Output.SetText(string(txt), app)
}

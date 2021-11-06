package tools

import (
    "encoding/hex"
    "github.com/tyagdit/bench/state"
    "github.com/gcla/gowid"
)

func ToHexCallback(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "to_hex"
    txt := hex.EncodeToString([]byte(state.Input.Text()))
    state.Output.SetText(txt, app)
}

func FromHexCallback(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "from_hex"
    txt, _ := hex.DecodeString(state.Input.Text())
    state.Output.SetText(string(txt), app)
}

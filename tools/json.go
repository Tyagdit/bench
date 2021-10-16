package tools

import (
    "github.com/tidwall/pretty"
    "github.com/tyagdit/toolie/state"
    "github.com/gcla/gowid"
)

func JSONBeautifyCallback(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "pretty_json"
    txt := pretty.Pretty([]byte(state.Input.Text()))
    state.Output.SetText(string(txt), app)
}

func JSONMinifyCallback(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "mini_json"
    txt := pretty.Ugly([]byte(state.Input.Text()))
    state.Output.SetText(string(txt), app)
}

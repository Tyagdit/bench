package tools

import (
    "net/url"
    "github.com/tyagdit/toolie/state"
    "github.com/gcla/gowid"
)

func URLEncodeCallback(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "url_en"
    txt := url.QueryEscape(state.Input.Text())
    state.Output.SetText(txt, app)
}

func URLDecodeCallback(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "url_de"
    txt, _ := url.QueryUnescape(state.Input.Text())
    state.Output.SetText(txt, app)
}

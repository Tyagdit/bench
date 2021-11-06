package tools

import (
    "fmt"
    "net/url"
    "github.com/tyagdit/bench/state"
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

func URLParse(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "url_parse"
    parsedUrl, _ := url.Parse(state.Input.Text())
    txt := ""
    if parsedUrl.Scheme != "" {txt += fmt.Sprintf("Scheme: %s\n", parsedUrl.Scheme)}
    if parsedUrl.User.Username() != "" {txt += fmt.Sprintf("User: %s\n", parsedUrl.User.Username())}
    if parsedUrl.Hostname() != "" {txt += fmt.Sprintf("Host: %s\n", parsedUrl.Hostname())}
    if parsedUrl.Port() != "" {txt += fmt.Sprintf("Port: %s\n", parsedUrl.Port())}
    if parsedUrl.Path != "" {txt += fmt.Sprintf("Path: %s\n", parsedUrl.Path)}
    if parsedUrl.Fragment != "" {txt += fmt.Sprintf("Fragment: %s\n", parsedUrl.Fragment)}

    q, _ := url.ParseQuery(parsedUrl.RawQuery)
    if len(q) != 0 {
        txt += "Query: \n"
        for k, v := range q {
            txt += fmt.Sprintf("  %q: ", k)
            if len(v) == 0 {
                txt += "\n"
            } else {
                for _, i := range v {
                    txt += fmt.Sprintf("%q, ", i)
                }
            }
            txt += "\n"
        }
    }
    state.Output.SetText(txt, app)
}

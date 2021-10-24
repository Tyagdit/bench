package tools

import (
    "encoding/json"
    "github.com/tidwall/pretty"
    "strings"

    "github.com/tyagdit/toolie/state"
    "github.com/gcla/gowid"
)

func CSVToJSON(app gowid.IApp, w gowid.IWidget) {
    state.SelectedTool = "csv_2_json"

    csv_str := state.Input.Text()
    rows := strings.Split(csv_str, "\n")
    rows = rows[:len(rows)-1]
    json_slice := make([][]string, len(rows))

    for i, row := range rows {
        json_slice[i] = strings.Split(row, ",")
    }

    json, err := json.Marshal(json_slice)
    if err != nil {
        state.Output.SetText(err.Error(), app)
    } else {
        txt := pretty.Pretty(json)
        state.Output.SetText(string(txt), app)
    }
}

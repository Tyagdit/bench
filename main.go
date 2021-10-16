package main

import (
    // "fmt"

    "github.com/tyagdit/toolie/state"
    "github.com/tyagdit/toolie/widgets"
    // "github.com/tyagdit/toolie/tools"

    "github.com/gcla/gowid"
    "github.com/gcla/gowid/examples"
    "github.com/gcla/gowid/widgets/pile"
    "github.com/gcla/gowid/widgets/framed"
    "github.com/gcla/gowid/widgets/text"
    "github.com/gcla/gowid/widgets/columns"
    "github.com/gcla/gowid/widgets/list"
    "github.com/gcla/gowid/widgets/button"

    "github.com/sirupsen/logrus"
)

type feature struct {
    label string
    featureFunc func(string) string
}

var featureList map[string]feature

func createBtn(selection string, f feature) *button.Widget {
    btn := button.NewDecorated(
        text.New(f.label),
        button.Decoration{Left: "", Right: ""},
    )
    btn.OnClick(gowid.WidgetCallback{
        Name: "cb",
        WidgetChangedFunction: func(app gowid.IApp, w gowid.IWidget) {
            state.SelectedTool = selection
            txt := f.featureFunc(state.Input.Text())
            state.Output.SetText(txt, app)
        },
    })
    return btn
}


func main() {

    framedInput := framed.New(
        state.Input,
        framed.Options{
            Frame: framed.UnicodeFrame,
            Title: "Input",
        },
    )

    framedOutput := framed.New(
        state.Output, 
        framed.Options{
            Frame: framed.UnicodeFrame,
            Title: "Output",
        },
    )

    ioPile := pile.New([]gowid.IContainerWidget{
        &gowid.ContainerWidget{IWidget: framedInput, D: gowid.RenderWithRatio{R: 0.5}},
        &gowid.ContainerWidget{IWidget: framedOutput, D: gowid.RenderWithRatio{R: 0.49}},
    })

    state.Input.OnTextSet(gowid.WidgetCallback{
        Name: "cb",
        WidgetChangedFunction: func(app gowid.IApp, w gowid.IWidget) {
            if state.SelectedTool == "" {
                state.Output.SetText(state.Input.Text(), app)
            } else {
                widgets.ToolMap[state.SelectedTool].Callback(app, w)
            }
            // txt := featureList[state.SelectedTool].featureFunc(state.Input.Text())
            // state.Output.SetText(txt, app)
            // switch state.SelectedTool {
            // case "":
            //     state.Output.SetText(state.Input.Text(), app)
            // case "to_b64":
            //     txt := b64.StdEncoding.EncodeToString([]byte(state.Input.Text()))
            //     state.Output.SetText(txt, app)
            // case "from_b64":
            //     txt, _ := b64.StdEncoding.DecodeString(state.Input.Text())
            //     state.Output.SetText(string(txt), app)
            // case "to_hex":
            //     txt := hex.EncodeToString([]byte(state.Input.Text()))
            //     state.Output.SetText(txt, app)
            // case "from_hex":
            //     txt, _ := hex.DecodeString(state.Input.Text())
            //     state.Output.SetText(string(txt), app)
            // case "sha1":
            //     hash := sha1.New()
            //     hash.Write([]byte(state.Input.Text()))
            //     bs := hash.Sum(nil)
            //     txt := hex.EncodeToString(bs)
            //     state.Output.SetText(txt, app)
            // case "sha256":
            //     hash := sha256.New()
            //     hash.Write([]byte(state.Input.Text()))
            //     bs := hash.Sum(nil)
            //     txt := hex.EncodeToString(bs)
            //     state.Output.SetText(txt, app)
            // case "sha512":
            //     hash := sha512.New()
            //     hash.Write([]byte(state.Input.Text()))
            //     bs := hash.Sum(nil)
            //     txt := hex.EncodeToString(bs)
            //     state.Output.SetText(txt, app)
            // case "md5":
            //     hash := md5.New()
            //     hash.Write([]byte(state.Input.Text()))
            //     bs := hash.Sum(nil)
            //     txt := hex.EncodeToString(bs)
            //     state.Output.SetText(txt, app)
            // case "url_en":
            //     txt := url.QueryEscape(state.Input.Text())
            //     state.Output.SetText(txt, app)
            // case "url_de":
            //     txt, _ := url.QueryUnescape(state.Input.Text())
            //     state.Output.SetText(txt, app)
            // case "pretty_json":
            //     txt := pretty.Pretty([]byte(state.Input.Text()))
            //     state.Output.SetText(string(txt), app)
            // case "mini_json":
            //     txt := pretty.Ugly([]byte(state.Input.Text()))
            //     state.Output.SetText(string(txt), app)
            // }
        },
    })

    // Buttons
    
    // btn1 := button.NewDecorated(
    //     text.New("To Base64"),
    //     button.Decoration{Left: "", Right: ""},
    // )
    // btn1.OnClick(gowid.WidgetCallback{
    //     Name: "cb",
    //     WidgetChangedFunction: tools.ToBase64CallBack,
    // })
    
    // btn2 := button.NewDecorated(
    //     text.New("From Base64"),
    //     button.Decoration{Left: "", Right: ""},
    // )
    // btn2.OnClick(gowid.WidgetCallback{
    //     Name: "cb",
    //     WidgetChangedFunction: tools.FromBase64CallBack,
    // })
    
    // btn3 := button.NewDecorated(
    //     text.New("To Hex"),
    //     button.Decoration{Left: "", Right: ""},
    // )
    // btn3.OnClick(gowid.WidgetCallback{
    //     Name: "cb",
    //     WidgetChangedFunction: func(app gowid.IApp, w gowid.IWidget) {
    //         state.SelectedTool = "to_hex"
    //         txt := hex.EncodeToString([]byte(state.Input.Text()))
    //         state.Output.SetText(txt, app)
    //     },
    // })
    
    // btn4 := button.NewDecorated(
    //     text.New("From Hex"),
    //     button.Decoration{Left: "", Right: ""},
    // )
    // btn4.OnClick(gowid.WidgetCallback{
    //     Name: "cb",
    //     WidgetChangedFunction : func(app gowid.IApp, w gowid.IWidget) {
    //         state.SelectedTool = "from_hex"
    //         txt, _ := hex.DecodeString(state.Input.Text())
    //         state.Output.SetText(string(txt), app)
    //     },
    // })
    
    // btn5 := button.NewDecorated(
    //     text.New("SHA-1"),
    //     button.Decoration{Left: "", Right: ""},
    // )
    // btn5.OnClick(gowid.WidgetCallback{
    //     Name: "cb",
    //     WidgetChangedFunction: func(app gowid.IApp, w gowid.IWidget) {
    //         state.SelectedTool = "sha1"
    //         hash := sha1.New()
    //         hash.Write([]byte(state.Input.Text()))
    //         bs := hash.Sum(nil)
    //         txt := hex.EncodeToString(bs)
    //         state.Output.SetText(txt, app)
    //     },
    // })

    // btn6 := button.NewDecorated(
    //     text.New("SHA-256"),
    //     button.Decoration{Left: "", Right: ""},
    // )
    // btn6.OnClick(gowid.WidgetCallback{
    //     Name: "cb",
    //     WidgetChangedFunction: func(app gowid.IApp, w gowid.IWidget) {
    //         state.SelectedTool = "sha256"
    //         hash := sha256.New()
    //         hash.Write([]byte(state.Input.Text()))
    //         bs := hash.Sum(nil)
    //         txt := hex.EncodeToString(bs)
    //         state.Output.SetText(txt, app)
    //     },
    // })

    // btn7 := button.NewDecorated(
    //     text.New("SHA-512"),
    //     button.Decoration{Left: "", Right: ""},
    // )
    // btn7.OnClick(gowid.WidgetCallback{
    //     Name: "cb",
    //     WidgetChangedFunction: func(app gowid.IApp, w gowid.IWidget) {
    //         state.SelectedTool = "sha512"
    //         hash := sha512.New()
    //         hash.Write([]byte(state.Input.Text()))
    //         bs := hash.Sum(nil)
    //         txt := hex.EncodeToString(bs)
    //         state.Output.SetText(txt, app)
    //     },
    // })

    // btn8 := button.NewDecorated(
    //     text.New("MD5"),
    //     button.Decoration{Left: "", Right: ""},
    // )
    // btn8.OnClick(gowid.WidgetCallback{
    //     Name: "cb",
    //     WidgetChangedFunction: func(app gowid.IApp, w gowid.IWidget) {
    //         state.SelectedTool = "md5"
    //         hash := md5.New()
    //         hash.Write([]byte(state.Input.Text()))
    //         bs := hash.Sum(nil)
    //         txt := hex.EncodeToString(bs)
    //         state.Output.SetText(txt, app)
    //     },
    // })

    // btn9 := button.NewDecorated(
    //     text.New("URL Encode"),
    //     button.Decoration{Left: "", Right: ""},
    // )
    // btn9.OnClick(gowid.WidgetCallback{
    //     Name: "cb",
    //     WidgetChangedFunction: func(app gowid.IApp, w gowid.IWidget) {
    //         state.SelectedTool = "url_en"
    //         txt := url.QueryEscape(state.Input.Text())
    //         state.Output.SetText(txt, app)
    //     },
    // })

    // btn10 := button.NewDecorated(
    //     text.New("URL Decode"),
    //     button.Decoration{Left: "", Right: ""},
    // )
    // btn10.OnClick(gowid.WidgetCallback{
    //     Name: "cb",
    //     WidgetChangedFunction: func(app gowid.IApp, w gowid.IWidget) {
    //         state.SelectedTool = "url_de"
    //         txt, _ := url.QueryUnescape(state.Input.Text())
    //         state.Output.SetText(txt, app)
    //     },
    // })

    // btn11 := button.NewDecorated(
    //     text.New("Beautify JSON"),
    //     button.Decoration{Left: "", Right: ""},
    // )
    // btn11.OnClick(gowid.WidgetCallback{
    //     Name: "cb",
    //     WidgetChangedFunction: func(app gowid.IApp, w gowid.IWidget) {
    //         state.SelectedTool = "pretty_json"
    //         txt := pretty.Pretty([]byte(state.Input.Text()))
    //         state.Output.SetText(string(txt), app)
    //     },
    // })

    // btn12 := button.NewDecorated(
    //     text.New("Minify JSON"),
    //     button.Decoration{Left: "", Right: ""},
    // )
    // btn12.OnClick(gowid.WidgetCallback{
    //     Name: "cb",
    //     WidgetChangedFunction: func(app gowid.IApp, w gowid.IWidget) {
    //         state.SelectedTool = "mini_json"
    //         txt := pretty.Ugly([]byte(state.Input.Text()))
    //         state.Output.SetText(string(txt), app)
    //     },
    // })

    var ToolButtonList = make([]gowid.IWidget, state.ToolButtonCount)
    for _, v := range widgets.ToolMap {
        btn := button.NewDecorated(
            text.New(v.Label),
            button.Decoration{Left: "", Right: ""},
        )
        btn.OnClick(gowid.WidgetCallback{
            Name: "cb",
            WidgetChangedFunction: v.Callback,
        })

        ToolButtonList = append(ToolButtonList, btn)
    }
    opsList := list.New(list.NewSimpleListWalker(ToolButtonList))
    framedOpsList := framed.New(
        opsList,
        framed.Options{
            Frame: framed.UnicodeFrame,
            Title: "Operations",
        },
    )

    cols := columns.New([]gowid.IContainerWidget{
        &gowid.ContainerWidget{IWidget: framedOpsList, D: gowid.RenderWithRatio{R: 0.2}},
        &gowid.ContainerWidget{IWidget: ioPile, D: gowid.RenderWithRatio{R: 0.8}},
    })

    palette := gowid.Palette{
        "body":  gowid.MakePaletteEntry(gowid.ColorBlack, gowid.ColorCyan),
        "fbody": gowid.MakePaletteEntry(gowid.ColorWhite, gowid.ColorBlack),
        "div": gowid.MakePaletteEntry(gowid.ColorNone, gowid.MakeRGBColor("#a06")),
    }

    app, err := gowid.NewApp(gowid.AppArgs{
        View:    cols,
        Palette: &palette,
        Log: logrus.StandardLogger(),
        // Log: logrus.Logger{
        //         Out: os.Stderr,
        //         Formatter: new(logrus.TextFormatter),
        //         Hooks: make(logrus.LevelHooks),
        //         Level: logrus.DebugLevel,
        //     },
    })
    examples.ExitOnErr(err)

    app.SimpleMainLoop()
}

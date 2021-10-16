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
        },
    })

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
    })
    examples.ExitOnErr(err)

    app.SimpleMainLoop()
}

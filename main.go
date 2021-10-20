package main

import (
    // "fmt"

    "github.com/tyagdit/toolie/state"
    "github.com/tyagdit/toolie/widgets"
    // "github.com/tyagdit/toolie/tools"

    "github.com/atotto/clipboard"

    "github.com/gcla/gowid"
    "github.com/gcla/gowid/examples"
    "github.com/gcla/gowid/widgets/button"
    "github.com/gcla/gowid/widgets/columns"
    "github.com/gcla/gowid/widgets/framed"
    "github.com/gcla/gowid/widgets/list"
    "github.com/gcla/gowid/widgets/overlay"
    "github.com/gcla/gowid/widgets/pile"
    "github.com/gcla/gowid/widgets/text"
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

    copyButton := button.NewDecorated(
        text.New(
            " Copy Output ",
            text.Options{Align: gowid.HAlignMiddle{}},
        ),
        button.Decoration{Left: "", Right: ""},
    )
    copyButton.OnClick(gowid.WidgetCallback{
        Name: "cb",
        WidgetChangedFunction: func(app gowid.IApp, w gowid.IWidget) {
            clipboard.WriteAll(state.Output.Content().String())
        },
    })

    overlayedFramedOutput := overlay.New(
        copyButton, framedOutput,
        gowid.VAlignTop{}, gowid.RenderWithUnits{U: 1},
        gowid.HAlignRight{}, gowid.RenderWithUnits{U: 17},
    )

    ioPile := pile.New([]gowid.IContainerWidget{
        &gowid.ContainerWidget{IWidget: framedInput, D: gowid.RenderWithRatio{R: 0.5}},
        &gowid.ContainerWidget{IWidget: overlayedFramedOutput, D: gowid.RenderWithRatio{R: 0.49}},
    })

    // To sort by the order in slice
    for i, tool := range widgets.ToolSlice1 {
        widgets.ToolMap[tool] = widgets.ToolSlice2[i]
    }

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

    var ToolButtonList = make([]gowid.IWidget, len(widgets.ToolSlice1))
    for _, v := range widgets.ToolSlice1 {
        btn := button.NewDecorated(
            text.New(widgets.ToolMap[v].Label),
            button.Decoration{Left: "", Right: ""},
        )
        btn.OnClick(gowid.WidgetCallback{
            Name: "cb",
            WidgetChangedFunction: widgets.ToolMap[v].Callback,
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
    })
    examples.ExitOnErr(err)

    app.SimpleMainLoop()
}

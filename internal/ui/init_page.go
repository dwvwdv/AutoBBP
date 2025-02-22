package ui

import (
    "github.com/rivo/tview"
)

func CreateInitPage(app *App) *tview.Flex {
    initPage := tview.NewFlex().SetDirection(tview.FlexRow)
    initPage.SetBorder(true).SetTitle("AutoBBP - Bug Bounty Program Hunter")
    
    list := createMainMenu(app)
    
    // 垂直置中
    initPage.AddItem(nil, 0, 1, false)
    initPage.AddItem(list, 10, 0, true)
    initPage.AddItem(nil, 0, 1, false)
    
    // 水平置中
    horizontalFlex := tview.NewFlex()
    horizontalFlex.AddItem(nil, 0, 1, false)
    horizontalFlex.AddItem(initPage, 20, 0, true)
    horizontalFlex.AddItem(nil, 0, 1, false)
    
    return horizontalFlex
}

func createMainMenu(app *App) *tview.List {
    list := tview.NewList().
        AddItem("New Project", "Create a new bug bounty project", 'n', func() {
            ShowNewProjectPage(app)
        }).
        AddItem("Import", "Import existing project", 'i', nil).
        AddItem("Export", "Export current project", 'e', nil).
        AddItem("Quit", "Press to exit", 'q', func() {
            app.Stop()
        })
    
    list.SetBorder(true)
    list.SetTitle("Menu")
    
    return list
}


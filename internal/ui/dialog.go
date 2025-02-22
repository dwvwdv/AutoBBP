package ui

import (
    "github.com/rivo/tview"
)

func ShowConfirmDialog(app *App) {
    modal := tview.NewModal().
        SetText("Do you want to quit without saving?").
        AddButtons([]string{"Yes", "No"}).
        SetDoneFunc(func(buttonIndex int, buttonLabel string) {
            if buttonLabel == "Yes" {
                app.Pages.SwitchToPage("init")
            }
            app.Pages.RemovePage("confirm_dialog")
        })
    
    flex := tview.NewFlex().
        AddItem(nil, 0, 1, false).
        AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
            AddItem(nil, 0, 1, false).
            AddItem(modal, 7, 1, true).
            AddItem(nil, 0, 1, false), 40, 1, true).
        AddItem(nil, 0, 1, false)
    
    app.Pages.AddPage("confirm_dialog", flex, true, true)
}

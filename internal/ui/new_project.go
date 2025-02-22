package ui

import (
    "github.com/rivo/tview"
)

func ShowNewProjectPage(app *App) {
    mainFlex := tview.NewFlex().SetDirection(tview.FlexRow)
    form := createProjectForm(app)
    
    // 垂直置中
    mainFlex.AddItem(nil, 0, 1, false)
    mainFlex.AddItem(form, 0, 2, true)
    mainFlex.AddItem(nil, 0, 1, false)
    
    // 水平置中
    horizontalFlex := tview.NewFlex()
    horizontalFlex.AddItem(nil, 0, 1, false)
    horizontalFlex.AddItem(mainFlex, 100, 0, true)
    horizontalFlex.AddItem(nil, 0, 1, false)
    
    app.Pages.AddPage("new_project", horizontalFlex, true, false)
    app.Pages.SwitchToPage("new_project")
}

func createProjectForm(app *App) *tview.Form {
    form := tview.NewForm()
    form.SetBorder(true)
    form.SetTitle("New Project")
    
    // 添加輸入字段
    form.AddInputField("Company Name", "", 50, nil, nil)
    form.AddTextArea("Terms", "", 50, 10, 0, nil)
    form.AddTextArea("Scope", "", 50, 10, 0, nil)
    form.AddTextArea("Valid Vulnerabilities", "", 50, 5, 0, nil)
    form.AddTextArea("Invalid Vulnerabilities", "", 50, 5, 0, nil)
    
    // 添加按鈕
    form.AddButton("Save", func() {
        // TODO: 保存項目數據
        app.Pages.SwitchToPage("init")
    })
    form.AddButton("Cancel", func() {
        ShowConfirmDialog(app)
    })
    
    // 設置樣式
    form.SetButtonsAlign(tview.AlignCenter)
    form.SetFieldTextColor(tcell.ColorWhite)
    form.SetButtonTextColor(tcell.ColorBlack)
    form.SetButtonBackgroundColor(tcell.ColorWhite)
    
    return form
}


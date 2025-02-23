package ui

import (
    "AutoBBP/internal/models"
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"
)

func ShowNewProjectPage(app *App) {
    mainFlex := tview.NewFlex().SetDirection(tview.FlexRow)
    form := createProjectForm(app)
    
    // 垂直置中
    mainFlex.AddItem(nil, 0, 1, false)
    mainFlex.AddItem(form, 0, 8, true)
    mainFlex.AddItem(nil, 0, 1, false)
    
    // 水平置中
    horizontalFlex := tview.NewFlex()
    horizontalFlex.AddItem(nil, 0, 1, false)
    horizontalFlex.AddItem(mainFlex, 100, 1, true)
    horizontalFlex.AddItem(nil, 0, 1, false)
    
    app.Pages.AddPage("new_project", horizontalFlex, true, true)
    app.Pages.SwitchToPage("new_project")
}

func createProjectForm(app *App) *tview.Form {
    form := tview.NewForm()
    form.SetBorder(true)
    form.SetTitle("New Project")
    
    // 創建一個新的 Project 實例
    project := &models.Project{}
    
    // 添加輸入字段並綁定數據
    form.AddInputField("Company Name", "", 50, nil, func(text string) {
        project.CompanyName = text
    })
    
    form.AddTextArea("Terms", "", 50, 10, 0, func(text string) {
        project.Terms = text
    })
    
    form.AddTextArea("Scope", "", 50, 10, 0, func(text string) {
        project.Scope = text
    })
    
    form.AddTextArea("Valid Vulnerabilities", "", 50, 5, 0, func(text string) {
        project.ValidVulns = text
    })
    
    form.AddTextArea("Invalid Vulnerabilities", "", 50, 5, 0, func(text string) {
        project.InvalidVulns = text
    })
    
    // 添加按鈕
    form.AddButton("Save", func() {
        if validateProject(project) {
            saveProject(app, project)
            app.Pages.SwitchToPage("init")
        } else {
            showErrorDialog(app, "Please fill in all required fields")
        }
    })
    
    form.AddButton("Cancel", func() {
        ShowConfirmDialog(app)
    })
    
    // 設置樣式
    form.SetButtonsAlign(tview.AlignCenter)
    form.SetFieldTextColor(tcell.ColorWhite)
    form.SetButtonTextColor(tcell.ColorBlack)
    form.SetButtonBackgroundColor(tcell.ColorWhite)
    
    // 添加快捷鍵
    form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Key() == tcell.KeyEsc {
            ShowConfirmDialog(app)
            return nil
        }
        return event
    })
    
    return form
}

// 驗證項目數據
func validateProject(p *models.Project) bool {
    return p.CompanyName != "" &&
           p.Terms != "" &&
           p.Scope != "" &&
           p.ValidVulns != ""
}

// 保存項目
func saveProject(app *App, project *models.Project) {
    // TODO: 實現項目保存邏輯
    // 1. 保存到文件或數據庫
    // 2. 更新應用程序狀態
}

// 顯示錯誤對話框
func showErrorDialog(app *App, message string) {
    modal := tview.NewModal().
        SetText(message).
        AddButtons([]string{"OK"}).
        SetDoneFunc(func(buttonIndex int, buttonLabel string) {
            app.Pages.RemovePage("error_dialog")
        })
    
    flex := tview.NewFlex().
        AddItem(nil, 0, 1, false).
        AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
            AddItem(nil, 0, 1, false).
            AddItem(modal, 7, 1, true).
            AddItem(nil, 0, 1, false), 40, 1, true).
        AddItem(nil, 0, 1, false)
    
    app.Pages.AddPage("error_dialog", flex, true, true)
}


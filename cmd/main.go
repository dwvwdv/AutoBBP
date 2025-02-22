package main

import (
    "log"
    "AutoBBP/internal/ui"
)

func main() {
    app := ui.NewApp()
    
    // 添加初始頁面
    initPage := ui.CreateInitPage(app)
    app.Pages.AddPage("init", initPage, true, true)
    
    if err := app.SetRoot(app.Pages, true).EnableMouse(true).Run(); err != nil {
        log.Fatal(err)
    }
}

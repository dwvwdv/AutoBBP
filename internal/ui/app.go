package ui

import (
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"
)

type App struct {
    *tview.Application
    Pages *tview.Pages
}

func NewApp() *App {
    app := &App{
        Application: tview.NewApplication(),
        Pages:      tview.NewPages(),
    }
    
    // 設置全域按鍵
    app.SetInputCapture(app.globalInputHandler)
    
    return app
}

func (a *App) globalInputHandler(event *tcell.EventKey) *tcell.EventKey {
    // vim-style 移動
    switch event.Rune() {
    case 'h':
        return tcell.NewEventKey(tcell.KeyLeft, 0, tcell.ModNone)
    case 'j':
        return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
    case 'k':
        return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
    case 'l':
        return tcell.NewEventKey(tcell.KeyRight, 0, tcell.ModNone)
    }
    return event
}


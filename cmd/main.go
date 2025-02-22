package main

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

func main() {
	app := tview.NewApplication()
	
	// 創建主視圖
	pages := tview.NewPages()
	
	// 初始化頁面
	initPage := tview.NewFlex().SetDirection(tview.FlexRow)
	initPage.SetBorder(true).SetTitle("AutoBBP - Bug Bounty Program Hunter")
	
	// 創建選項列表
	list := tview.NewList().
		AddItem("New Project", "Create a new bug bounty project", 'n', func() {
			showNewProjectPage(app, pages) // 添加新項目頁面的處理函數
		}).
		AddItem("Import", "Import existing project", 'i', nil).
		AddItem("Export", "Export current project", 'e', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	
	// 設置列表樣式
	list.SetBorder(true)
	list.SetTitle("Menu")
	
	// 添加上下空白區域來實現垂直置中
	initPage.AddItem(nil, 0, 1, false)
	initPage.AddItem(list, 10, 0, true)
	initPage.AddItem(nil, 0, 1, false)
	
	// 創建水平佈局來實現水平置中
	horizontalFlex := tview.NewFlex()
	horizontalFlex.AddItem(nil, 0, 1, false)
	horizontalFlex.AddItem(initPage, 50, 0, true)
	horizontalFlex.AddItem(nil, 0, 1, false)
	
	// 添加初始頁面
	
	pages.AddPage("init", horizontalFlex, true, true)
	// 設置全域按鍵處理
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
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
		case 'G':
			// 移動到最後一項
			list.SetCurrentItem(list.GetItemCount() - 1)
			return nil
		case 'g':
			// 移動到第一項
			list.SetCurrentItem(0)
			return nil
		}
		return event
	})
	
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

// 新增 showNewProjectPage 函數
func showNewProjectPage(app *tview.Application, pages *tview.Pages) {
	// 創建主要佈局
	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	
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
		pages.SwitchToPage("init")
	})
	form.AddButton("Cancel", func() {
		showConfirmDialog(pages)
	})
	
	// 設置表單樣式
	form.SetButtonsAlign(tview.AlignCenter)
	form.SetFieldTextColor(tcell.ColorWhite)
	form.SetButtonTextColor(tcell.ColorBlack)
	form.SetButtonBackgroundColor(tcell.ColorWhite)
	
	// 垂直置中
	mainFlex.AddItem(nil, 0, 1, false)
	mainFlex.AddItem(form, 0, 2, true)
	mainFlex.AddItem(nil, 0, 1, false)
	
	// 水平置中
	horizontalFlex := tview.NewFlex()
	horizontalFlex.AddItem(nil, 0, 1, false)
	horizontalFlex.AddItem(mainFlex, 100, 0, true)
	horizontalFlex.AddItem(nil, 0, 1, false)
	
	// 添加新項目頁面
	pages.AddPage("new_project", horizontalFlex, true, false)
	pages.SwitchToPage("new_project")
	
	// 設置表單的按鍵處理
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			showConfirmDialog(pages)
			return nil
		}
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
	})
}

// 新增確認對話框函數
func showConfirmDialog(pages *tview.Pages) {
	modal := tview.NewModal().
		SetText("Do you want to quit without saving?").
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				pages.SwitchToPage("init")
			}
			pages.RemovePage("confirm_dialog")
		})
	
	// 創建一個半透明的背景
	flex := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(modal, 7, 1, true).
			AddItem(nil, 0, 1, false), 40, 1, true).
		AddItem(nil, 0, 1, false)
		
	// 添加確認對話框頁面
	pages.AddPage("confirm_dialog", flex, true, true)
}

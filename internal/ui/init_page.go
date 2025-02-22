package ui

import (
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

func CreateInitPage(app *App) *tview.Flex {
	initPage := tview.NewFlex().SetDirection(tview.FlexRow)
	initPage.SetBorder(true).SetTitle("AutoBBP - Bug Bounty Program Hunter")

	list := createMainMenu(app)

	// 垂直置中
	initPage.AddItem(nil, 0, 1, false)
	initPage.AddItem(list, 0, 1, true)
	initPage.AddItem(nil, 0, 1, false)

	// 水平置中
	horizontalFlex := tview.NewFlex()
	horizontalFlex.AddItem(nil, 0, 1, false)
	horizontalFlex.AddItem(initPage, 50, 0, true)
	horizontalFlex.AddItem(nil, 0, 1, false)

	return horizontalFlex
}

func createMainMenu(app *App) *tview.List {
    list := tview.NewList().
        AddItem("New Project", "Create a new bug bounty project", 'n', func() {
            ShowNewProjectPage(app)
        }).
        AddItem("", "", 0, nil).
        AddItem("Asset Collection", "Manage target assets", 'a', func() {
            ShowAssetPage(app)
        }).
        AddItem("", "", 0, nil).
        AddItem("Import", "Import existing project", 'i', nil).
        AddItem("", "", 0, nil).
        AddItem("Export", "Export current project", 'e', nil).
        AddItem("", "", 0, nil).
        AddItem("Quit", "Press to exit", 'q', func() {
            app.Stop()
        })

    list.SetBorder(true)
    list.SetTitle("Menu")
    
    // 設置自定義的輸入處理
    list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        currentItem := list.GetCurrentItem()
        maxItems := list.GetItemCount()

        // 處理移動
        var moveUp bool
        switch {
        case event.Key() == tcell.KeyUp || event.Rune() == 'k':
            moveUp = true
        case event.Key() == tcell.KeyDown || event.Rune() == 'j':
            moveUp = false
        default:
            return event
        }

        // 計算新位置
        newItem := currentItem
        if moveUp {
            // 向上移動到前一個非空項目
            for newItem > 0 {
                newItem--
                mainText, _ := list.GetItemText(newItem)  // 只使用第一個返回值
                if mainText != "" {
                    list.SetCurrentItem(newItem)
                    break
                }
            }
        } else {
            // 向下移動到下一個非空項目
            for newItem < maxItems-1 {
                newItem++
                mainText, _ := list.GetItemText(newItem)  // 只使用第一個返回值
                if mainText != "" {
                    list.SetCurrentItem(newItem)
                    break
                }
            }
        }
        return nil
    })

    return list
}

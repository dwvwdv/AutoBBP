package ui

import (
    "AutoBBP/internal/models"
    "AutoBBP/internal/config"
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

type AssetPage struct {
	app       *App
	mainFlex  *tview.Flex
	assetList *tview.List
	assets    []*models.Asset
	deletedAssets [][]*models.Asset
	config        *config.Config
}

// 添加 ShowError 方法到 App 結構體
func (app *App) ShowError(message string) {
    modal := tview.NewModal().
        SetText(message).
        AddButtons([]string{"OK"}).
        SetDoneFunc(func(buttonIndex int, buttonLabel string) {
            app.Pages.RemovePage("error-modal")
        })
    app.Pages.AddPage("error-modal", modal, true, true)
}

func ShowAssetPage(app *App) {
    // 加載配置
    cfg, err := config.LoadConfig()
    if err != nil {
        app.ShowError("無法加載配置文件: " + err.Error())
        return
    }

    page := &AssetPage{
        app:           app,
        assets:        make([]*models.Asset, 0),
        deletedAssets: make([][]*models.Asset, 0),
        config:        cfg,
    }
    page.Setup()
}

func (p *AssetPage) Setup() {
	// 創建主佈局
	p.mainFlex = tview.NewFlex().SetDirection(tview.FlexRow)

	// 創建頂部按鈕欄
	buttonBar := tview.NewFlex().SetDirection(tview.FlexColumn)
	addButton := tview.NewButton("Add Asset").SetSelectedFunc(p.showAddAssetForm)
	deleteButton := tview.NewButton("Delete").SetSelectedFunc(p.deleteSelectedAsset)
	exportButton := tview.NewButton("Export").SetSelectedFunc(p.exportAssets)

	// 設置按鈕焦點樣式
	addButton.SetBackgroundColorActivated(tcell.ColorDarkBlue)
	deleteButton.SetBackgroundColorActivated(tcell.ColorDarkBlue)
	exportButton.SetBackgroundColorActivated(tcell.ColorDarkBlue)

	buttonBar.AddItem(addButton, 10, 0, true)
	buttonBar.AddItem(deleteButton, 10, 0, true)
	buttonBar.AddItem(exportButton, 10, 0, true)
	buttonBar.AddItem(nil, 0, 1, false) // 填充空間

	// 創建資產列表
	p.assetList = tview.NewList().
		SetHighlightFullLine(true).
		SetSelectedBackgroundColor(tcell.ColorDarkBlue)
	p.assetList.SetBorder(true).SetTitle("Assets")

    // 添加公司選擇功能
    p.mainFlex.AddItem(tview.NewButton("Select Company").
        SetSelectedFunc(func() {
            // 顯示公司選擇對話框
            companies := p.getCompanyList()
            modal := tview.NewModal().
                SetText("選擇公司").
                AddButtons(companies).
                SetDoneFunc(func(buttonIndex int, buttonLabel string) {
                    if buttonIndex >= 0 {
                        if err := p.LoadCompanyAssets(buttonLabel); err != nil {
                            p.app.ShowError("無法加載資產: " + err.Error())
                        }
                    }
                    p.app.Pages.RemovePage("company-modal")
                })
            p.app.Pages.AddPage("company-modal", modal, true, true)
        }), 0, 1, false)

	// 組裝主佈局
	p.mainFlex.AddItem(buttonBar, 3, 0, true)  // 修改為 true 使按鈕欄可以獲得焦點
	p.mainFlex.AddItem(p.assetList, 0, 1, true)

	// 當前焦點的按鈕索引
	currentButtonIndex := 0
	buttons := []*tview.Button{addButton, deleteButton, exportButton}

	// 添加快捷鍵
	p.mainFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if p.app.GetFocus() == p.assetList {
			action := p.listEvent(event) 
			if action != nil{
				return p.listEvent(event)
			}
		}
		switch event.Rune() {
		case 'h':
			if currentButtonIndex > 0 {
				currentButtonIndex--
				p.app.SetFocus(buttons[currentButtonIndex])
			}
			return nil
		case 'l':
			if currentButtonIndex < len(buttons)-1 {
				currentButtonIndex++
				p.app.SetFocus(buttons[currentButtonIndex])
			}
			return nil
		case 'j':
			p.app.SetFocus(p.assetList)
			return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
		case 'k':
			p.app.SetFocus(p.assetList)
			return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
		}

		switch event.Key() {
		case tcell.KeyEsc:
			p.app.Pages.SwitchToPage("init")
			return nil
		case tcell.KeyTab:
			// Tab 鍵在按鈕之間循環
			currentButtonIndex = (currentButtonIndex + 1) % len(buttons)
			p.app.SetFocus(buttons[currentButtonIndex])
			return nil
		}
		return event
	})

	// 設置初始焦點
	p.app.SetFocus(buttons[0])

	// 將頁面添加到應用程序
	p.app.Pages.AddPage("asset", p.mainFlex, true, true)
	p.app.Pages.SwitchToPage("asset")
}

// 獲取可用的公司列表
func (p *AssetPage) getCompanyList() []string {
    files, err := os.ReadDir(p.config.AssetPath)
    if err != nil {
        return []string{}
    }

    var companies []string
    for _, file := range files {
        if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
            companies = append(companies, strings.TrimSuffix(file.Name(), ".json"))
        }
    }
    return companies
}

func (p *AssetPage) showAddAssetForm() {
	form := tview.NewForm()
	asset := models.NewAsset()

	form.AddInputField("URL", "", 40, nil, func(text string) {
		asset.URL = text
	})

	form.AddDropDown("Type", []string{"web", "api", "mobile", "other"}, 0, func(option string, index int) {
		asset.Type = option
	})

	form.AddTextArea("Description", "", 50, 4, 0, func(text string) {
		asset.Notes = text
	})

	form.AddTextArea("Notes", "", 50, 4, 0, func(text string) {
		asset.Notes = text
	})

	form.AddButton("Save", func() {
		if p.validateAsset(asset) {
			p.addAsset(asset)
			p.app.Pages.RemovePage("add_asset")
		}
	})

	form.AddButton("Cancel", func() {
		p.app.Pages.RemovePage("add_asset")
	})

	// 設置表單樣式
	form.SetBorder(true).SetTitle("Add New Asset")
	form.SetButtonsAlign(tview.AlignCenter)

	// 創建模態對話框
	flex := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(
		tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(form, 20, 1, true).
			AddItem(nil, 0, 1, false),
		60, 1, true,
		).
		AddItem(nil, 0, 1, false)

	p.app.Pages.AddPage("add_asset", flex, true, true)
}

func (p *AssetPage) validateAsset(asset *models.Asset) bool {
	return asset.URL != "" && asset.Type != ""
}

func (p *AssetPage) addAsset(asset *models.Asset) {
	p.assets = append(p.assets, asset)
	p.refreshAssetList()
}

func (p *AssetPage) deleteSelectedAsset() {
	if len(p.assets) == 0 {
		return
	}

	index := p.assetList.GetCurrentItem()
	if index >= 0 && index < len(p.assets) {
		// 顯示確認對話框
		modal := tview.NewModal().
			SetText("Are you sure you want to delete this asset?").
			AddButtons([]string{"Yes", "No"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Yes" {
					p.assets = append(p.assets[:index], p.assets[index+1:]...)
					p.refreshAssetList()
				}
				p.app.Pages.RemovePage("confirm_delete")
			})

		p.app.Pages.AddPage("confirm_delete", modal, true, true)
	}
}

func (p *AssetPage) LoadCompanyAssets(companyName string) error {
    assets, err := models.LoadAssets(p.config.AssetPath, companyName)
    if err != nil {
        return err
    }

    p.assets = assets
    p.refreshAssetList()
    return nil
}

func (p *AssetPage) refreshAssetList() {
	p.assetList.Clear()
	for _, asset := range p.assets {
		p.assetList.AddItem(asset.URL, asset.Description, 0, nil)
        p.assetList.AddItem(
            fmt.Sprintf("%s (%s)", asset.Name, asset.IP),
            fmt.Sprintf("Ports: %v", asset.Ports),
            0,
            nil,
        )
	}
}

func (p *AssetPage) exportAssets() {
	// TODO: 實現資產導出功能
}

func (p *AssetPage) listEvent(event *tcell.EventKey) *tcell.EventKey {
	if event.Rune() == 'd' {
		currentItem := p.assetList.GetCurrentItem()
		if currentItem >= 0 {
			currentState := make([]*models.Asset, len(p.assets))
			copy(currentState, p.assets)
			p.deletedAssets = append(p.deletedAssets, currentState)

			// 執行刪除
			p.assets = append(p.assets[:currentItem], p.assets[currentItem+1:]...)
			p.refreshAssetList()
		}
	}
	if event.Rune() == 'u' {
		if len(p.deletedAssets) > 0 {
			lastIndex := len(p.deletedAssets) - 1
			p.assets = make([]*models.Asset, len(p.deletedAssets[lastIndex]))
			copy(p.assets, p.deletedAssets[lastIndex])

			p.deletedAssets = p.deletedAssets[:lastIndex]

			p.refreshAssetList()
		}
	}

	switch event.Rune() {
	case 'j':
		return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
	case 'k':
		return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
	}

	return nil
}

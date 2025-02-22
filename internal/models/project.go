
package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	
	// 創建主視圖
	pages := tview.NewPages()
	
	// 初始化頁面
	initPage := tview.NewFlex()
	initPage.SetBorder(true).SetTitle("AutoBBP - Bug Bounty Program Hunter")
	
	// 創建選項列表
	list := tview.NewList().
		AddItem("New Project", "Create a new bug bounty project", 'n', nil).
		AddItem("Import", "Import existing project", 'i', nil).
		AddItem("Export", "Export current project", 'e', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	
	initPage.AddItem(list, 0, 1, true)
	
	// 添加初始頁面
	pages.AddPage("init", initPage, true, true)
	
	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
package models

// Project 代表一個漏洞獎勵計劃項目
type Project struct {
	CompanyName    string   `json:"company_name"`
	Terms          string   `json:"terms"`
	Scope          []string `json:"scope"`
	ValidVulns     []string `json:"valid_vulns"`
	InvalidVulns   []string `json:"invalid_vulns"`
	CollectedAssets []Asset  `json:"collected_assets"`
	Commands       []Command `json:"commands"`
}

// Asset 代表收集到的資產
type Asset struct {
	Type     string `json:"type"`
	Value    string `json:"value"`
	Source   string `json:"source"`
	DateTime string `json:"date_time"`
}

// Command 代表執行的命令
type Command struct {
	Command  string `json:"command"`
	Output   string `json:"output"`
	DateTime string `json:"date_time"`
}

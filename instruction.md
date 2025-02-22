# AutoBBP
當前目錄為 ~/code/Golang/AutoBBP，AutoBBP該為項目名稱，妳需要協助我完成這個項目

### Bug Bounty Program Hounter
開發一個Bug Bounty AI工具，呈現形式為TUI，shortcut採用vim motion，主要功能為根據公司提供的條款規範、範圍、有效漏洞等資訊，自動生成命令並執行對應的資產掃描，並嘗試進一步進行攻擊
開發語言採用`go 1.23`，`tview`，整個項目只需要一個go.mod，並調用`grok 2` API
項目為本地開發，目前module 無須加上 github.com/dwvwdv


- 初始化頁面
    - 新項目
    - 導入
    - 導出

- 新項目頁面
    - 輸入公司 
    - 條款規範 
    - 測試範圍
    - 有效漏洞或不接受漏洞

- 資產收集頁面 分為3個panel
    - 詢問問題，該面板可用`tab`進行`提示詞模板`快速選取
    - AI回覆，並將生成命令記錄起來，當用戶按下[允許]快捷鍵(y)時，進行命令調用
    - 數據紀錄，當前蒐集資產思路、待做事項

- 導出當前工作紀錄
    - 保存為json
    - 可在初始化頁面進行導入


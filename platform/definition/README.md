# definition

This project is a constant library for other dcc backend project.

## How to use in other project
***

1. Install as a git submodule
```
git submodule add https://gitlab.int.dayongtek.com/dcc/dev/backend/definition.git definition
```

2. 錯誤碼範圍定義說明

* 0 為成功。表示系統操作成功完成。
* 0 ~ 99 為後台使用。表示系統出現後台操作相關的錯誤，例如數據庫連接失敗、文件系統錯誤等。
* 600 ~ 699 為內部溝通使用。表示系統在內部溝通過程中出現的錯誤，例如 API 請求失敗、網絡連接中斷等。
* 800 ~ 899 為 chatservice 使用。
* 1001 ~ 1999 為外部串接API使用錯誤碼。

1. api 代碼範圍定義說明
* 100100 ~ 100199 為遊戲SERVER串接使用
* 100200 ~ - 為後台使用
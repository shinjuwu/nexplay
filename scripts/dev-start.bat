@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

:: ============================================================
::  NextPlay 本地開發環境一鍵啟動腳本 (Windows)
::  用法: scripts\dev-start.bat [選項]
::
::  選項:
::    --all        啟動所有服務 (含 mock-gameserver, chatservice, orderservice)
::    --mock       額外啟動 mock-gameserver
::    --chat       額外啟動 chatservice
::    --order      額外啟動 orderservice
::    --agent      啟動 Agent 前端 (預設為 Manager)
::    --no-front   不啟動前端開發伺服器
::    --help       顯示說明
:: ============================================================

set "ROOT_DIR=%~dp0.."
set "START_MOCK=0"
set "START_CHAT=0"
set "START_ORDER=0"
set "START_FRONT=1"
set "FRONT_MODE=manager"

:: 解析參數
:parse_args
if "%~1"=="" goto :start
if /i "%~1"=="--all" (
    set "START_MOCK=1"
    set "START_CHAT=1"
    set "START_ORDER=1"
    shift & goto :parse_args
)
if /i "%~1"=="--mock" ( set "START_MOCK=1" & shift & goto :parse_args )
if /i "%~1"=="--chat" ( set "START_CHAT=1" & shift & goto :parse_args )
if /i "%~1"=="--order" ( set "START_ORDER=1" & shift & goto :parse_args )
if /i "%~1"=="--agent" ( set "FRONT_MODE=agent" & shift & goto :parse_args )
if /i "%~1"=="--no-front" ( set "START_FRONT=0" & shift & goto :parse_args )
if /i "%~1"=="--help" ( goto :show_help )
echo [ERROR] 未知選項: %~1
goto :show_help

:show_help
echo.
echo  NextPlay 本地開發環境一鍵啟動腳本
echo  ──────────────────────────────────
echo  用法: scripts\dev-start.bat [選項]
echo.
echo  選項:
echo    --all        啟動所有服務
echo    --mock       啟動 mock-gameserver
echo    --chat       啟動 chatservice
echo    --order      啟動 orderservice
echo    --agent      啟動 Agent 前端 (預設 Manager)
echo    --no-front   不啟動前端開發伺服器
echo    --help       顯示此說明
echo.
echo  範例:
echo    scripts\dev-start.bat              啟動基礎設施 + Manager 前端
echo    scripts\dev-start.bat --agent      啟動基礎設施 + Agent 前端
echo    scripts\dev-start.bat --all        啟動所有服務 + Manager 前端
echo    scripts\dev-start.bat --no-front   只啟動後端服務
echo.
exit /b 0

:: ============================================================
:start
echo.
echo  ============================================
echo    NextPlay 本地開發環境啟動中...
echo  ============================================
echo.

:: ------------------------------------------------------------
:: 1. 前置檢查
:: ------------------------------------------------------------
echo [1/7] 前置檢查...

:: 檢查 Docker
docker info >nul 2>&1
if errorlevel 1 (
    echo [ERROR] Docker 未運行，請先啟動 Docker Desktop。
    exit /b 1
)
echo       Docker 正常運行。

:: 檢查 Git Submodule
if not exist "%ROOT_DIR%\platform\definition\error_code_constant.go" (
    echo [ERROR] Git submodule 未初始化，platform\definition\ 目錄缺少檔案。
    echo         請先執行: git submodule update --init --recursive
    exit /b 1
)
echo       Git submodule 已初始化。

:: 檢查 Node.js
where node >nul 2>&1
if errorlevel 1 (
    if "%START_FRONT%"=="1" (
        echo [ERROR] 未偵測到 Node.js，前端開發伺服器無法啟動。
        echo         請安裝 Node.js 16+ : https://nodejs.org/
        exit /b 1
    )
)
echo       Node.js 已安裝。

:: ------------------------------------------------------------
:: 2. 建立 Docker Volumes (若不存在)
:: ------------------------------------------------------------
echo [2/7] 檢查 Docker Volumes...

for %%V in (dcc-pgsql-volumes dcc-redis-volumes dcc-pgadmin-volumes) do (
    docker volume inspect %%V >nul 2>&1
    if errorlevel 1 (
        echo       建立 volume: %%V
        docker volume create %%V >nul
    ) else (
        echo       %%V 已存在
    )
)

:: ------------------------------------------------------------
:: 3. 啟動基礎設施 (Platform + PostgreSQL + Redis + pgAdmin)
:: ------------------------------------------------------------
echo [3/7] 啟動基礎設施 (Platform + PostgreSQL + Redis + pgAdmin)...
docker-compose -f "%ROOT_DIR%\platform\deployments\docker-compose.yml" up -d
if errorlevel 1 (
    echo [ERROR] 基礎設施啟動失敗。
    exit /b 1
)

:: ------------------------------------------------------------
:: 4. 啟動選用服務
:: ------------------------------------------------------------
echo [4/7] 啟動選用服務...

if "%START_MOCK%"=="1" (
    echo       啟動 Mock Game Server...
    docker-compose -f "%ROOT_DIR%\docker-compose.dev.yml" up -d
)

if "%START_CHAT%"=="1" (
    echo       啟動 ChatService...
    docker-compose -f "%ROOT_DIR%\chatservice\deployments\docker-compose.yml" up -d
)

if "%START_ORDER%"=="1" (
    echo       啟動 OrderService...
    docker-compose -f "%ROOT_DIR%\orderservice\deployments\docker-compose.yml" up -d
)

if "%START_MOCK%%START_CHAT%%START_ORDER%"=="000" (
    echo       無額外選用服務 (使用 --all 可啟動全部)
)

:: ------------------------------------------------------------
:: 5. 等待後端就緒
:: ------------------------------------------------------------
echo [5/7] 等待 Platform Backend 就緒...
set "RETRY=0"
:wait_backend
if %RETRY% geq 30 (
    echo [WARN] 等待逾時，後端可能尚未完全啟動。繼續執行...
    goto :start_frontend
)
curl -s -o nul -w "%%{http_code}" http://127.0.0.1:9986 >nul 2>&1
if errorlevel 1 (
    set /a RETRY+=1
    timeout /t 2 /nobreak >nul
    goto :wait_backend
)
echo       Platform Backend 已就緒。

:: ------------------------------------------------------------
:: 6. 啟動前端開發伺服器
:: ------------------------------------------------------------
:start_frontend
if "%START_FRONT%"=="0" (
    echo [6/7] 跳過前端開發伺服器 (--no-front)
    goto :done
)

echo [6/7] 安裝前端依賴並啟動 %FRONT_MODE% 開發伺服器...

cd /d "%ROOT_DIR%\dcc_front\frontend"
if not exist "node_modules" (
    echo       執行 npm install...
    call npm install
)

echo.
echo  ============================================
echo    啟動 %FRONT_MODE% 前端開發伺服器
echo    http://localhost:9985
echo  ============================================
echo.

call npm run %FRONT_MODE%-dev

:: npm run 結束 = 使用者按了 Ctrl+C
goto :done

:: ============================================================
:done
echo.
echo  ============================================
echo    開發環境已啟動完成!
echo  ============================================
echo.
echo  服務列表:
echo    前端:     http://localhost:9985
echo    後端 API: http://localhost:9986
echo    pgAdmin:  http://localhost:5050
echo.
if "%START_FRONT%"=="0" ( pause )
exit /b 0

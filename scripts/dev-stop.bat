@echo off
chcp 65001 >nul

:: ============================================================
::  NextPlay 本地開發環境一鍵停止腳本 (Windows)
:: ============================================================

set "ROOT_DIR=%~dp0.."

echo.
echo  ============================================
echo    NextPlay 本地開發環境停止中...
echo  ============================================
echo.

echo [1/4] 停止 OrderService...
docker-compose -f "%ROOT_DIR%\orderservice\deployments\docker-compose.yml" down 2>nul

echo [2/4] 停止 ChatService...
docker-compose -f "%ROOT_DIR%\chatservice\deployments\docker-compose.yml" down 2>nul

echo [3/4] 停止 Mock Game Server...
docker-compose -f "%ROOT_DIR%\docker-compose.dev.yml" down 2>nul

echo [4/4] 停止基礎設施 (Platform + PostgreSQL + Redis + pgAdmin)...
docker-compose -f "%ROOT_DIR%\platform\deployments\docker-compose.yml" down

echo.
echo  所有服務已停止。
echo.
pause

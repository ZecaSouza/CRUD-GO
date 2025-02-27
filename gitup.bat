@'
@echo off
setlocal enabledelayedexpansion

:: Verifica branch
if "%~1"=="" (
    echo Erro: Nome da branch nao informado.
    echo Uso: gitup.bat ^<nome-branch^> [mensagem]
    exit /b 1
)

:: Configurações
set BRANCH=%~1
set MENSAGEM=%~2

:: Solicita mensagem se vazia
if "%MENSAGEM%"=="" (
    set /p MENSAGEM="Digite a mensagem do commit: "
)

:: Comandos Git
git fetch
git rev-parse --verify origin/!BRANCH! >nul 2>&1
if errorlevel 1 (
    echo Criando nova branch: !BRANCH!...
    git checkout -b !BRANCH!
) else (
    git checkout !BRANCH!
)

git add -A
git commit -m "!MENSAGEM!"
git push origin !BRANCH!
git pull origin !BRANCH!

echo ✔ Push realizado para branch !BRANCH! com sucesso!
endlocal
'@ | Out-File -Encoding ASCII gitup.bat
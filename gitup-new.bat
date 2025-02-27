@echo off
:: Script para commit e push simplificado no Windows
:: Uso: gitup.bat <nome-branch> [mensagem]

setlocal enabledelayedexpansion

:: Verifica se o nome da branch foi informado
if "%1"=="" (
    echo Erro: Nome da branch n??o informado.
    echo Uso: gitup.bat ^<nome-branch^> [mensagem]
    exit /b 1
)

set BRANCH=%1
set MENSAGEM=%2

:: Verifica se a mensagem foi informada
if "%MENSAGEM%"=="" (
    set /p MENSAGEM="Digite a mensagem do commit: "
)

:: Atualiza a lista de branches remotas
git fetch

:: Verifica se a branch existe
git rev-parse --verify origin/%BRANCH% >nul 2>&1
if %errorlevel% neq 0 (
    echo Criando nova branch: %BRANCH%...
    git checkout -b %BRANCH%
) else (
    git checkout %BRANCH%
)

:: Adiciona todas as altera????es
git add -A

:: Executa o commit
git commit -m "%MENSAGEM%"

:: Faz o push para a branch
git push origin %BRANCH%

:: Atualiza a branch local
git pull origin %BRANCH%

echo ?????? Push realizado para branch %BRANCH% com sucesso!
endlocal

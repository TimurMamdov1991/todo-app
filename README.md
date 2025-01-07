# Todo-app

Для подключению к бд надо выполнить ряд условий 
1) Установить Docker : https://www.docker.com/get-started/
2) Войти в консоль и запустить команду: docker pull postgres 
3) Запустить команду : docker run --name=todo-db -e POSTGRES_PASSWORD=qwerty -p 60062:5432 -d --rm postgres
4) Получить список контейнеров командой: docker ps
5) Зайти в оболочку PowerShell:
 и установить scoop: 1)Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
                     2)Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression
                     3)C помощью scoop установить Migrate: scoop install migrate
6) запустить команду внутри проекта в CMD: migrate -path ./schema -database "postgres://postgres:qwerty@localhost:60062/postgres?sslmode=disable" up
7) Если хотите дропнуть бд: migrate -path ./schema -database "postgres://postgres:qwerty@localhost:60062/postgres?sslmode=disable" down
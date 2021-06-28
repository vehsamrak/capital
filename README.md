# Capital
Расчет капитала

### Структура проекта
Используется [project-layout](https://github.com/golang-standards/project-layout).

## Dev-окружение
Запуск контейнеров через docker-compose: `docker-compose -f build/package/docker-compose.yml up -d`.

## Зависимости
Для обновления зависимостей используется `go mod tidy && go mod vendor`.

## Миграции
Миграции запускаются через bash-скрипт Shmig в отдельном докер-контейнере. 
Миграции хранятся в директории `tools\migrations`.

Создание новой миграции: `tools/shmig.sh -t postgresql -d capital -m tools/migrations create migration-name`.

## TODO
### Backlog
- ..

### In progress

### Done
- ..

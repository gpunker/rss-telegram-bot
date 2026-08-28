# Миграции
## Создать миграцию
`migrate create -ext sql -dir db/migrations -seq {migration_name}`

где {migration_name} название вашей миграции

Далее в появившихся UP и DOWN файлах необходимо написать, на SQL, изменения в схеме БД

## Выполнить миграцию
`migrate -path db/migrations -database "mysql://user:pass@tcp(db:3306)/rss_feed_bot" up [N]`

где N - количество шагов(не обязательно)

## Откатить миграцию
`migrate -path db/migrations -database "mysql://user:pass@tcp(db:3306)/rss_feed_bot" down [N]`

где N - количество шагов(не обязательно)
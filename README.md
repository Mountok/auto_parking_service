# Документация к сервису

Тема: Парковочная стоянка

### Требования

Go версия: 1.21.6

База данных: PostgreSQL 13

### Важно настроить перед запуском

Файл миграции БД: ./migrations/migration.sql (выполнить в СУБД для миграции бд в вашу локальную БД)

Настройки конфига: ./internals/cfg/config.go (изменить переменые окружения под свои)



# Запуск:

- Используйте `go build` или `go run cmd/main.go` - расширения должны установить автоматически и запустит сервис

- Если первый вариант не помог используй вендор `go mod vendor` - это загрузит все расщирения в локальную директорию в проект и запустите снова.

> Если что то не работает это возможно из-за ОС. 
> 
> Чел, удали Windows!!!!

# Роутеры:

### Для автомобилей (cars):

#### GET:

- `/cars/list` - вернет список всех машин и их владельцев.

- `/cars/list?brand=bmw` - вернет список всех машин c брендом авто `bmw`.

- `cars/list?brand=bmw&colour=red` - вернет список всех машин c брендом авто `bmw` где цвет `red`.

- `/cars/list?brand=Nissan&colour=blue&license_plate=LB081G` - вернет список всех машин c бредом `nissan` с цветом  `blue` и номерами `LB081G`.

- `/cars/find/2` - вернет  ответ машину с id = 2.

#### POST:

- `/cars/create` - создает запись в бд. В тело запроса передаеться `json` в след формате. 
  
  ```json
    {
        "colour": "black",
        "brand": "Tesla Model S",
        "license_plate": "TM900S",
        "owner": {
            "id": 11
        }
    }
  ```

### Для владельцев (users):

#### GET:

- `/users/list` - список всех пользователей

- `/users/list?name=s` - список пользователей в имени которого есть содежиться буква `s`. (`/users/list?name=Islam` - можно вести имя соответсвенно).

- `/users/find/9` - вернеть пользователя с `id` = `7`.

#### POST:

- `/users/create`  - создает запись в бд. В тело запроса передаеться `json` в след формате.
  
  ```json
  {
      "name": "Steve J",
      "rank": "CEO"
  }
  ```



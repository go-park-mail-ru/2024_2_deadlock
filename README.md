# 2024_2_deadlock
Репозиторий Бэкэнда проекта vc.ru команды deadlock

## Запуск проекта

Установите зависимости:

```shell
make init
```

Примените миграции:

```shell
make migrate
```

Запустите сервер:
```shell
make run
```

## Перед пушем в репозиторий 

Нужно прогнать тесты, линтер и форматтер:

```shell
make pre-commit
```

или

```shell
make lint
make test
```

## Работа с миграциями

Создание новой миграции:

```shell
make create_migration name="some_name"
```

Применение миграций:

```shell
make migrate
```

## Авторы

[Иван Павлов](https://github.com/darleet) - _Тимлид_

[Юрий Малхасян](https://github.com/ujognutsi) 

[Александр Новиков](https://github.com/AlexNov03)

## Менторы

[Кирпичов Владислав](https://github.com/) - _Frontend_

[Жиленков Илья](https://github.com/ilyushkaaa) - _Backend_

[Памужак Пётр](https://github.com/mars444) - _UX/UI_


## Ссылки

[Фронтенд проекта](https://github.com/frontend-park-mail-ru/2024_2_deadlock)



# Запись логов пользователей в файл
Программа производить асинхронную запись логов пользователей в файл.

За счет использование goroutine в программе - запись в файл происходит в **7х быстрее**:
```go
before using the goroutine time: 14.778768872s
after using the goroutine, time: 2.482559738s.
```


## Технологии 

- Golang
- библиотеки: 
    - sync



## Использование 
- Установите последнею версию **go**

- Запустите приложение  командой:

```go
    go run main.go
```

## Вывод в консоли
```
....
WRITING FILE FOR USER ID: 364
WRITING FILE FOR USER ID: 118
WRITING FILE FOR USER ID: 578
TIME ELAPSED: 2.482559738s
```

## Main 
![Текст описания](img/code.png)

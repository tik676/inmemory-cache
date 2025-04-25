# In-Memory Cache Library

Это простая библиотека кеша на языке Go с поддержкой TTL (Time To Live). Кеш хранит элементы в памяти с заданным временем жизни. Если элемент не был использован в течение указанного времени, он автоматически удаляется.

## Установка

Чтобы использовать эту библиотеку в своем проекте, выполните следующую команду:

```bash
go get github.com/tik676/inmemory-cache/cache
```

Пример использования
Создание кеша

Для создания нового кеша необходимо указать время жизни элементов (TTL):
```go
myCache := cache.NewCache(5 * time.Second) // Кеш с TTL 5 секунд
```
Добавление элементов в кеш

Для добавления элемента в кеш используйте метод Set:
```go
myCache.Set("user1", "John Doe")
myCache.Set("user2", "Jane Smith")
```
Получение элементов из кеша

Для получения значения из кеша используйте метод Get. Он возвращает два значения: значение элемента и булевое значение, которое указывает, найден ли элемент в кеше.
```go
value, found := myCache.Get("user1")
if found {
    fmt.Println("Value for user1:", value)
} else {
    fmt.Println("user1 not found")
}
```
Ожидание истечения TTL

Если элемент истекает по истечении времени TTL, то он автоматически удаляется из кеша. Например:
```go
time.Sleep(6 * time.Second) // Ждем больше времени, чем TTL

value, found = myCache.Get("user1")
if found {
    fmt.Println("Value for user1 after TTL:", value)
} else {
    fmt.Println("user1 not found after TTL")
}
```
Очистка кеша

Вы можете очистить весь кеш с помощью метода Clear:
```go
myCache.Clear()

value, found = myCache.Get("user2")
if found {
    fmt.Println("Value for user2 after clearing cache:", value)
} else {
    fmt.Println("user2 not found after clearing cache")
}
```
Удаление элемента из кеша

Для удаления элемента из кеша используйте метод Delete:
```go
myCache.Set("user3", "Alice")
myCache.Delete("user3")
value, found = myCache.Get("user3")
if found {
    fmt.Println("Value for user3 after deletion:", value)
} else {
    fmt.Println("user3 successfully deleted from cache")
}
```

Методы кеша
NewCache(ttl time.Duration) *Cache

Создает новый кеш с заданным временем жизни элементов (TTL).
Set(key string, value T)

Добавляет элемент в кеш. Значение сохраняется в кеше до тех пор, пока не истечет время TTL.
Get(key string) (T, bool)

Возвращает значение из кеша по ключу. Если элемент существует и не истек, возвращается значение и true. Если элемент не найден или истек, возвращается nil и false.
Clear()

Очищает все элементы в кеше.
Delete(key string)

Удаляет элемент из кеша по указанному ключу.
Snapshot() Cache

Возвращает копию текущего состояния кеша. Это позволяет сделать снимок кеша и работать с ним отдельно от оригинала.


Пример кода 
```go
package main

import (
    "fmt"
    "time"
    "github.com/tik676/inmemory-cache/cache"
)

func main() {
    // Создаем новый кэш с TTL 5 секунд
    myCache := cache.NewCache(5 * time.Second)

    // Устанавливаем элементы
    myCache.Set("user1", "John Doe")
    myCache.Set("user2", "Jane Smith")

    // Получаем элемент из кэша
    value, found := myCache.Get("user1")
    if found {
        fmt.Println("Value for user1:", value)
    } else {
        fmt.Println("user1 not found")
    }

    // Проверяем элементы после истечения TTL
    time.Sleep(6 * time.Second)
    value, found = myCache.Get("user1")
    if found {
        fmt.Println("Value for user1 after TTL:", value)
    } else {
        fmt.Println("user1 not found after TTL")
    }

    // Очистка кеша
    myCache.Clear()
    value, found = myCache.Get("user2")
    if found {
        fmt.Println("Value for user2 after clearing cache:", value)
    } else {
        fmt.Println("user2 not found after clearing cache")
    }

    // Удаление элемента
    myCache.Set("user3", "Alice")
    myCache.Delete("user3")
    value, found = myCache.Get("user3")
    if found {
        fmt.Println("Value for user3 after deletion:", value)
    } else {
        fmt.Println("user3 successfully deleted from cache")
    }
}
```
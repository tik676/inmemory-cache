# In-Memory Cache

Это простая реализация кэша в памяти на Go. Он позволяет хранить данные в памяти с ограничением по времени жизни (TTL) и автоматически удалять устаревшие элементы.

## Особенности

- Простая реализация кэша в памяти.
- Поддержка времени жизни (TTL) для элементов кэша.
- Возможность добавления, получения и удаления элементов.
- Потокобезопасность с использованием мьютексов.

## Установка

Для использования этого кэша в вашем проекте добавьте его как зависимость:

```bash
go get github.com/tik676/inmemory-cache
```
Пример использования

``` go
package main

import (
    "fmt"
    "time"
    "github.com/tik676/inmemory-cache"
)

func main() {
    // Создаем новый кэш с TTL 5 секунд 
    cache := cache.NewCache(5 * time.Second)

    // Устанавливаем элемент
    cache.Set("user1", "John Doe")

    // Получаем элемент 
    value, found := cache.Get("user1")
    if found {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key not found.")
    }

    // Ожидаем истечения TTL
    time.Sleep(6 * time.Second)

    // Проверяем снова
    value, found = cache.Get("user1")
    if found {
        fmt.Println("Value:", value)
    } else {
        fmt.Println("Key not found.")
    }
}
```
Методы
Set(key string, value interface{})

Добавляет элемент в кэш с указанным ключом и значением. Элемент автоматически удалится, если истечет время жизни (TTL).
Get(key string) (interface{}, bool)

Получает элемент из кэша по ключу. Возвращает значение и true, если элемент найден, или nil и false, если элемент не найден или его TTL истек.
Clear()

Удаляет все элементы из кэша.
Delete(key string)

Удаляет элемент с указанным ключом из кэша.



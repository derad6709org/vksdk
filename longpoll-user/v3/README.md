# Обработчик User Long Poll API v3

[![PkgGoDev](https://pkg.go.dev/badge/github.com/derad6709org/vksdk/v2/longpoll-user/v3)](https://pkg.go.dev/github.com/derad6709org/vksdk/v2/longpoll-user/v3)
[![VK](https://img.shields.io/badge/developers-%234a76a8.svg?logo=VK&logoColor=white)](https://vk.com/dev/using_longpoll)

Данный модуль поддерживает версию **3**.

```go
// import "github.com/derad6709org/vksdk/v2/longpoll-user/v3"

w := wrapper.NewWrapper(lp)

// event with code 4
w.OnNewMessage(func(m wrapper.NewMessage) {
  fmt.Printf("4 wrapper.NewMessage: %v\n", m)
})
```

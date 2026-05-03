![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)

# v2 - GoMorse with Cobra

```bash
go build -o gomorse main.go

sudo cp gomorse /usr/local/bin
```
---
```bash
# Перевод текста
gomorse encode "Hello!"

# Перевод с проигрышем
gomorse encode --text "I love pizza" --play 

# Справка
gomorse --help 
gomorse encode -h
```

Переводчик текста в азбуку Морзе с воспроизведением 
 - Пользователь вводит фразу
 - В файле morse.go происходит преобразование
 - player.go воспроизводит нужный звуковой файл


# Зависимости
 - Утилита ***ffmpeg*** для ffplay
 - go version: 1.21+
 - [Cobra](https://github.com/spf13/cobra) 

### Fedora
 - sudo dnf install ffmpeg-free

### Ubuntu/Debian
 - sudo apt install ffmpeg

# Запуск
 - git clone https://github.com/vladekx-s/GoMorse.git
 - cd GoMorse
 - go run main.go
---
![мемчик](https://media2.giphy.com/media/v1.Y2lkPTc5MGI3NjExbzhjM2JsNGJmYjdneHFrOGU5NHFhZjU0MHBoem1mYzM1cGwxMTR4ZCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/3o6nULePUoEGPRU9d6/giphy.gif)
### Дата: 04.05.2026
### Go version: 1.26.2
### Fedora version: 44


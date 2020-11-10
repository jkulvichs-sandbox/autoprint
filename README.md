![GitHub](https://img.shields.io/github/license/jkulvichs-sandbox/autoprint)
![GitHub top language](https://img.shields.io/github/languages/top/jkulvichs-sandbox/autoprint)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/jkulvichs-sandbox/autoprint)
![GitHub last commit](https://img.shields.io/github/last-commit/jkulvichs-sandbox/autoprint)

# AutoPrint
Сервис для автоматической отправки на печать файлов

# Компиляция
`make build` или `go build ./`

Первый вариант предпочтительнее, т.к. добавит время сборки.
Сборка при помощи **make** возможна на любом **linux** дистрибутиве.
Итоговый файл нацелен на запуск под Windows x32

# Конфигурация
Для генерации конфигурации по умолчанию выполните:  
`autoprint --generate-config`

В директории рядом появится шаблон конфигурации.
По умолчанию, стандартная конфигурация будет отслеживать файлы с
расширением **.pdf** в текущем каталоге и выводить новые файлы на печать
при помощи стандартного принтера.

## Ключи конфигурации
Конфиг в формате **json** по умолчанию ищется файл рядом **autoprint.json**

- `acrobatPath` - string - Команда запуска Adobe Acrobat или путь к нему
- `profiles` - []ConfigProfile - Массив профилей.
               Для работы с несколькими директориями, принтерами и форматами
    - `printerName` - string - Имя принтера. Оставьте пустым для принтера по умолчанию
    - `watchForDirectory` - string - Путь к отслеживаемой директории.
    По умолчанию: ./
    - `fileFilter` - string - Регулярное выражение для фильтра файлов.
    По умолчанию: ^.*\\.pdf$ (фильтр только .pdf)

# Запуск
Просто добавьте файл в автозагрузку или запустите как сервис.
В тестовых целях файл можно запустить через консоль или двойным кликом.

## Флаги
- `config` - Путь до файла конфигурации, по умолчанию **autoprint.json**
- `generate-config` - Сгенерировать шаблон конфигурации и завершиться

## Зависимости
Необходимо установить **Adobe Acrobat** и в конфигурации прописать
команду запуска (в случае если он в PATH) или путь до него. 

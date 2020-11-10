FILENAME=autoprint.exe
OS=windows
ARCH=386
WINE=wine

BUILDTIME=$(shell date)

# Компиляция
build:
	GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags "-X 'main.Build=$(BUILDTIME)'" -o $(FILENAME) ./
.PHONY: build

# Запуск сервиса нативно
run: build
	./$(FILENAME)
.PHONY: run

# Запуск сервиса из под wine
run-wine: build
	wine $(FILENAME)
.PHONY: run-wine

# Генерация конфигурации
generate-config: build
	wine ./$(FILENAME) --generate-config
.PHONY: generate-config

# Удаляет лишние файлы
clear:
	rm autoprint.json 2>/dev/null
	rm $(FILENAME) 2>/dev/null
.PHONY: clear
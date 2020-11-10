package main

import (
	"flag"
	"github.com/radovskyb/watcher"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"path"
	"regexp"
	"time"
)

// Build - Время компиляции
var Build string = "TEST"

func main() {
	flagConfig := flag.String("config", "autoprint.json", "path to autoprint's config file")
	flagGenerateConfig := flag.Bool("generate-config", false, "generates default config")
	flag.Parse()

	// Инициализация логировщика
	logger, _ := NewLogger()

	// Генерация стандартной конфигурации
	if *flagGenerateConfig {
		if err := ConfigGenerateDefault(*flagConfig); err != nil {
			logger.Fatalf("can't generate config: %s\n", err)
		}
		os.Exit(0)
	}

	// Загрузка конфигурации
	conf, err := ConfigLoad(*flagConfig)
	if err != nil {
		logger.Fatalf("can't load config: %s", err)
	}

	// Старт всех профилей
	for _, profile := range conf.Profiles {
		go func(prof ConfigProfile) {
			logger.Infof("starting profile for: \"%s\" with filter \"%s\"\n", prof.WatchForDirectory, prof.FileFilter)
			if err := RunProfile(logger, conf.AcrobatPath, prof); err != nil {
				logger.Errorf("can't run profile for folder %s: %s\n", prof.WatchForDirectory, err)
			}
		}(profile)
	}

	// Приветствие
	logger.Infof("autoprint %s by @jkulvich (TG or VK)\n", Build)
	logger.Infof("autoprint service is running...\n")

	// Ждём завершения от пользователя или системы
	closeChan := make(chan os.Signal)
	signal.Notify(closeChan, os.Kill, os.Interrupt)
	<-closeChan
	logger.Infoln("autoprint service stopped")
}

// Запуск работы отдельного профиля
func RunProfile(logger *logrus.Logger, acrobatPath string, profile ConfigProfile) error {
	// Начало слежения за файлами
	watch := watcher.New()

	watch.FilterOps(watcher.Create)
	rFilter, err := regexp.Compile(profile.FileFilter)
	if err != nil {
		logger.Fatalf("incorrect filter regexp: %s", err)
	}
	watch.AddFilterHook(watcher.RegexFilterHook(rFilter, false))

	// Функция ожидания и повторной попытки
	waitAndTryAgain := func(event watcher.Event, dur time.Duration) {
		<-time.After(dur)
		logger.Infof("continuing: %s\n", event.Name())
		watch.TriggerEvent(watcher.Create, event)
	}

	// Обработка событий вотчера
	go func() {
		for {
			select {
			case event := <-watch.Event:

				// Старт печати
				logger.Infof("print starting: %s\n", event.FileInfo.Name())

				// Путь к файлу
				filePath := path.Join(profile.WatchForDirectory, event.FileInfo.Name())

				// Проверка доступности файла
				logger.Infof("checking file ability for: %s\n", event.FileInfo.Name())
				if err := CheckFileAbility(filePath); err != nil {
					logger.Infof("can't get access to file: %s\n", err)
					logger.Infof("can't get access, waiting for 2s ...\n")
					go waitAndTryAgain(event, time.Second*2)
					continue
				}

				// Отправка на печать
				logger.Infof("sending to acrobat ...\n")
				_, err := PrintDoc(acrobatPath, filePath, profile.PrinterName)
				logger.Infof("sent to acrobat: %s\n", err)

			case err := <-watch.Error:
				logger.Errorf("watcher error: %s", err)
			case <-watch.Closed:
				return
			}
		}
	}()

	// Добавление директории которую отслеживаем
	if err := watch.Add(profile.WatchForDirectory); err != nil {
		logger.Fatalf("can't add directory %s: %s", profile.WatchForDirectory, err)
	}

	// Начинаем прослушивание событий директории
	if err := watch.Start(time.Millisecond * 100); err != nil {
		logger.Fatalf("can't watch for directory: %s", err)
	}

	return nil
}

// CheckFileAbility - Проверяет доступность файла к печати
func CheckFileAbility(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	_ = file.Close()

	return nil
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config - Конфигурация автопринтера
type Config struct {
	// Путь к Adobe Acrobat для печати
	AcrobatPath string `json:"acrobatPath"`
	// Профили для поддержки работы разных диекторий и принтеров
	Profiles []ConfigProfile `json:"profiles"`
}

type ConfigProfile struct {
	// Название принтера на печать которому выводить
	// Названия всех известных утилите принтеров можно получить флагом
	// Если оставить пустым - использует принтер по умолчанию
	PrinterName string `json:"printerName"`
	// Отслеживать эту директорию на добавление в неё файлов
	WatchForDirectory string `json:watchForDirectory`
	// Регулярное выражение сравнения файлов
	FileFilter string `json:fileFilter`
}

// ConfigGenerateDefault - Генерирует стандартную конфигурацию
func ConfigGenerateDefault(path string) error {
	conf := &Config{
		AcrobatPath: "acrord32",
		Profiles: []ConfigProfile{
			{
				PrinterName:       "",
				WatchForDirectory: "./",
				FileFilter:        `^.*\.pdf$`,
			},
		},
	}

	if err := ConfigStore(path, conf); err != nil {
		return err
	}

	return nil
}

// ConfigStore - Сохраняет конфигурацию
func ConfigStore(path string, conf *Config) error {
	data, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, data, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// ConfigLoad - Загружает конфигурацию из файла
func ConfigLoad(path string) (*Config, error) {
	conf := &Config{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, conf); err != nil {
		return nil, err
	}

	return conf, nil
}

package main

import (
	"os/exec"
)

// PrintDoc - Принимает имя принтера и путь до файла
func PrintDoc(acrobatPath, path string, printer ...string) (string, error) {
	prnt := ""
	if len(printer) > 0 {
		prnt = printer[0]
	}

	// Выполняем печать
	var cmd *exec.Cmd
	if prnt == "" {
		cmd = exec.Command(acrobatPath, "/h", "/t", path)
	} else {
		cmd = exec.Command(acrobatPath, "/h", "/t", path, prnt)
	}

	// Убиваем процесс на всякий случай чтобы он не повесил файл
	defer func() {
		_ = cmd.Process.Kill()
	}()

	// Информация о печати
	res, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(res), nil
}

package files

import (
	"os"

	"github.com/fatih/color"
)

type JsonDB struct {
	filepath string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		filepath: name,
	}
}

func (db JsonDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filepath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db JsonDB) Write(data []byte) {
	file, err := os.Create(db.filepath)
	if err != nil {
		color.Red("Не удаётся создать файл")
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		color.Red("Не удаётся записать файл")
	}
	color.Green("Запись успешна")
}

package encoding

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hive-bootcamp/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// 1 Читаю файл json
	data, err := os.ReadFile(j.FileInput)
	if err != nil {
		return fmt.Errorf("ошибка при чтении файла: %s", err.Error())
	}

	// 2 Вытаскиваю json данные
	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return fmt.Errorf("ошибка при парсинге JSON: %s", err.Error())
	}

	// 3 Из данных json делаю yaml
	yamlData, err := yaml.Marshal(jsonData)

	if err != nil {
		return fmt.Errorf("ошибка при попытке маршалирования: %s", err.Error())
	}

	// 4 Записаваю yaml данные в файл yaml

	err = os.WriteFile(j.FileOutput, yamlData, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при записи данных: %s", err.Error())
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	// 1 Читаю файл yaml
	data, err := os.ReadFile(y.FileInput)
	if err != nil {
		return fmt.Errorf("ошибка при чтении файла: %s", err.Error())
	}

	// 2 Вытаскиваю yaml данные
	var yamlData interface{}

	err = yaml.Unmarshal(data, &yamlData)
	if err != nil {
		return fmt.Errorf("ошибка при чтении данных yaml: %s", err.Error())
	}

	// 3 Из данных yaml делаю json
	jsonData, err := json.MarshalIndent(yamlData, "", "    ")
	if err != nil {
		return fmt.Errorf("ошибка при попытке маршалирования: %s", err.Error())
	}

	// 4 Записаваю json данные в файл json
	err = os.WriteFile(y.FileOutput, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("ошибка при записи данных: %s", err.Error())
	}

	return nil
}

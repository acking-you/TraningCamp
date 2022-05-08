package repository

import (
	"encoding/json"
	"os"
)

func Save(filePath string) error {
	if err := SaveTopicIndexMap(filePath); err != nil {
		return err
	}
	if err := SavePostIndexMap(filePath); err != nil {
		return err
	}
	return nil
}

func SaveTopicIndexMap(filePath string) error {
	open, err := os.OpenFile(filePath+"post", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	writer := open

	for _, v := range topicIndexMap {
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}
		data = append(data, '\n')
		writer.Write(data)
	}

	return nil
}

func SavePostIndexMap(filePath string) error {
	open, err := os.OpenFile(filePath+"post", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	writer := open

	for _, v := range postIndexMap {
		for _, post := range v {
			data, err := json.Marshal(*post)
			if err != nil {
				return err
			}
			data = append(data, '\n')
			writer.Write(data)
		}
	}
	return nil
}

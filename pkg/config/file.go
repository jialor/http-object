package config

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/tidwall/gjson"
)

// load a config file
func Load(name string) (err error) {
	path := Dir + name + ".json"
	bufferBytes, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	// config file must be json
	if !gjson.Valid(string(bufferBytes)) {
		err = errors.New("config file format error")
		return err
	}
	return nil
}

// save current config to json file
func Save() (err error) {
	name := "last"
	path := Dir + name + ".json"
	err = ioutil.WriteFile(path, bufferBytes, 0666) //存在就覆盖；不存在创建。
	if err != nil {
		return err
	}
	return nil
}

// list json files in config file dir
func List() (names []string, err error) {
	files, err := ioutil.ReadDir(Dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		name := strings.TrimRight(file.Name(), ".json")
		names = append(names, name)
	}
	return names, nil
}

// new a json config file
// func New(name string, data []byte) (err error) {
// 	path := Dir + name + ".json"
// 	err = ioutil.WriteFile(path, data, 0644) //存在就覆盖；不存在创建。
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// delete a json config file
// func Remove(name string) (err error) {
// 	path := Dir + name + ".json"
// 	err = os.Remove(path)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

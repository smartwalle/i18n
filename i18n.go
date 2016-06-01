package i18n

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"strings"
	"sync"
)

var ctx *context = &context{}

type message map[string]string

type context struct {
	sync.RWMutex
	langs    []string
	lang     string
	messages map[string]message
}

func (this *context) load(langPath string) (err error) {
	err = this.loadWithPath(langPath)
	return err
}

func (this *context) loadWithPath(langPath string) (err error) {
	this.RLock()
	defer this.RUnlock()

	var dir *os.File
	dir, err = os.Open(langPath)
	defer dir.Close()

	if err != nil {
		return err
	}

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat(langPath)
	if err != nil {
		return err
	}

	if !fileInfo.IsDir() {
		return errors.New("必须指定一个目录")
	}

	var dirNames []string
	dirNames, err = dir.Readdirnames(-1)
	if err != nil {
		return err
	}

	this.messages = make(map[string]message)
	this.langs = make([]string, 0, len(dirNames))

	for index, name := range dirNames {
		var fullPath = path.Join(langPath, name)
		fileInfo, err = os.Stat(fullPath)

		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			var key = strings.Split(fileInfo.Name(), ".")[0]
			if data, err := this.loadFile(fullPath); err == nil {
				if index == 0 {
					this.lang = key
				}

				this.messages[key] = data
				this.langs = append(this.langs, key)
			} else {
				return err
			}
		}
	}
	return err
}

func (this *context) loadFile(path string) (data message, err error) {
	var file *os.File
	file, err = os.Open(path)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	var jsonDecoder = json.NewDecoder(file)
	err = jsonDecoder.Decode(&data)

	return data, err
}

func (this *context) value(lang, key string) (value string) {
	if this.messages == nil || len(this.messages) == 0 {
		return key
	}

	var msg message
	var ok bool

	if msg, ok = this.messages[lang]; !ok {
		msg = this.messages[this.lang]
	}

	if value, ok = msg[key]; ok {
		return value
	}

	return key
}

func (this *context) exists(lang string) bool {
	var _, ok = this.messages[lang]
	return ok
}

func (this *context) setDefault(lang string) {
	this.lang = lang
}

////////////////////////////////////////////////////////////////////////////////
func Load(langPath string) error {
	return ctx.load(langPath)
}

func Exists(lang string) bool {
	return ctx.exists(lang)
}

func SetDefault(lang string) {
	ctx.setDefault(lang)
}

func TL(lang, key string) string {
	return ctx.value(lang, key)
}

func T(key string) string {
	return ctx.value(ctx.lang, key)
}

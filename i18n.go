package i18n

import (
	"github.com/smartwalle/ini4go"
	"sync"
)

var ctx *i18n
var once sync.Once

func init() {
	once.Do(func() {
		ctx = NewContext()
	})
}

type i18n struct {
	config *ini4go.Ini
	lang   string
}

func NewContext() *i18n {
	var c = &i18n{}
	c.config = ini4go.New(false)
	return c
}

func (this *i18n) Load(dir string) (err error) {
	err = this.config.Load(dir)
	if err == nil {
		var names = this.config.SectionNames()
		if len(names) > 0 {
			this.lang = names[0]
		}
	}
	return err
}

func (this *i18n) LoadFiles(files ...string) (err error) {
	err = this.config.LoadFiles(files...)
	if err == nil {
		var names = this.config.SectionNames()
		if len(names) > 0 {
			this.lang = names[0]
		}
	}
	return err
}

func (this *i18n) value(lang, key string) (value string) {
	return this.config.GetValue(lang, key)
}

func (this *i18n) exists(lang string) (ok bool) {
	return this.config.HasSection(lang)
}

func (this *i18n) setDefault(lang string) {
	if this.exists(lang) {
		this.lang = lang
	}
}

func (this *i18n) Reset() {
	this.config.Reset()
}

////////////////////////////////////////////////////////////////////////////////
func Load(dir string) error {
	return ctx.Load(dir)
}

func LoadFiles(files ...string) error {
	return ctx.LoadFiles(files...)
}

func Exists(lang string) bool {
	return ctx.exists(lang)
}

func SetDefault(lang string) {
	ctx.setDefault(lang)
}

func TL(lang, key string) string {
	if ctx.exists(lang) {
		ctx.value(lang, key)
	}
	return ctx.value(ctx.lang, key)
}

func T(key string) string {
	return ctx.value(ctx.lang, key)
}

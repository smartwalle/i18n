package i18n

import (
	"github.com/smartwalle/config"
)

var ctx *i18n = NewContext()

type message map[string]string

type i18n struct {
	config *config.Config
	lang   string
}

func NewContext() *i18n {
	var c = &i18n{}
	c.config = config.NewConfig()
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
	return this.config.MustValue(lang, key, key)
}

func (this *i18n) exists(lang string) (ok bool) {
	return this.config.HasSection(lang)
}

func (this *i18n) setDefault(lang string) {
	this.lang = lang
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
	return ctx.value(lang, key)
}

func T(key string) string {
	return ctx.value(ctx.lang, key)
}

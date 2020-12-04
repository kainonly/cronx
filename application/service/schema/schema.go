package schema

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"schedule-microservice/config/options"
)

type Schema struct {
	path string
}

func New(path string) *Schema {
	c := new(Schema)
	c.path = path
	return c
}

func (c *Schema) autoload(identity string) string {
	return c.path + identity + ".yml"
}

func (c *Schema) Update(option options.JobOption) (err error) {
	var bs []byte
	if bs, err = yaml.Marshal(option); err != nil {
		return
	}
	if err = ioutil.WriteFile(c.autoload(option.Identity), bs, 0644); err != nil {
		return
	}
	return
}

func (c *Schema) Get(identity string) (option options.JobOption, err error) {
	if _, err = os.Stat(c.path + identity + ".yml"); err != nil {
		return
	}
	var bs []byte
	if bs, err = ioutil.ReadFile(c.path + identity + ".yml"); err != nil {
		return
	}
	if err = yaml.Unmarshal(bs, &option); err != nil {
		return
	}
	return
}

func (c *Schema) Lists() (lists []options.JobOption, err error) {
	var files []os.FileInfo
	if files, err = ioutil.ReadDir(c.path); err != nil {
		return
	}
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext == ".yml" {
			var bs []byte
			if bs, err = ioutil.ReadFile(c.path + file.Name()); err != nil {
				return
			}
			var option options.JobOption
			if err = yaml.Unmarshal(bs, &option); err != nil {
				return
			}
			lists = append(lists, option)
		}
	}
	return
}

func (c *Schema) Delete(identity string) error {
	return os.Remove(c.autoload(identity))
}

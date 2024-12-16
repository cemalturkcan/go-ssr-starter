package templ

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
	"html/template"
	"projectesc/src/general/config"
)

const (
	LiveReloadScript = `<script type="text/javascript" src="https://livejs.com/live.js"></script>`
)

func New() *html.Engine {
	engine := html.New("./views", ".html")

	engine.AddFunc(
		"putValues", func(values ...interface{}) fiber.Map {
			if len(values)%2 != 0 {
				return fiber.Map{}
			}
			result := make(fiber.Map, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					log.Error("putValues: key is not a string")
					continue
				}
				result[key] = values[i+1]
			}
			return result
		},
	)

	if config.IsDevelopment {
		engine.Reload(true)
		engine.AddFunc(
			"liveReloadTag", func() template.HTML {
				return LiveReloadScript
			},
		)
	} else {
		engine.AddFunc(
			"liveReloadTag", func() template.HTML {
				return ""
			},
		)
	}
	return engine
}

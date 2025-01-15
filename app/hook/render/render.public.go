package render

import "github.com/gofiber/fiber/v2"

var defaultVaribal = fiber.Map{
	"Title": "Traveler Kota Baubau",
	"Menu":  "",
}

func RenderPublic(ctx *fiber.Ctx, filename string, data ...fiber.Map) error {
	var newData = fiber.Map{}

	if len(data) > 0 {
		newData = data[0]
	}

	for key, value := range defaultVaribal {
		if _, isExist := newData[key]; !isExist {
			newData[key] = value
		}
	}
	return ctx.Render(filename, newData, "layouts/public.template")

}

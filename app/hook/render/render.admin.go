package render

import "github.com/gofiber/fiber/v2"

var defaultNilai = fiber.Map{
	"Title":   "Default Title",
	"AppName": "Traveller Kota Baubau",
}

func RenderAdmin(ctx *fiber.Ctx, filename string, data ...fiber.Map) error {
	var newData = fiber.Map{}

	if len(data) > 0 {
		newData = data[0]
	}

	for key, value := range defaultNilai {
		if _, isExist := newData[key]; !isExist {
			newData[key] = value
		}
	}
	return ctx.Render("admin/"+filename, newData, "layouts/admin.template")

}

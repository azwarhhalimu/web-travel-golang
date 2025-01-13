Template User & admin : 
Gorm : https://gorm.io/

Dependensi 
Yang du butuhnan 
 1. Gofiber : go get github.com/gofiber/fiber/v2
 2. Format Number : go get golang.org/x/text/message
 3. Gorm MYSQL :
                     go get -u gorm.io/gorm
                     go get -u gorm.io/driver/mysql

4. Jet : go get github.com/gofiber/template/jet/v2
5. Air verse : go install github.com/air-verse/air@latest

    
--- embed map ----
<iframe width="600" height="400" style="border:0;"
                src="https://www.openstreetmap.org/export/embed.html?bbox={{Data.Lng}},{{Data.Lat}},{{Data.Lng}},{{Data.Lat}}&layer=mapnik&marker={{Data.Lat}},{{Data.Lng}}">
            </iframe>
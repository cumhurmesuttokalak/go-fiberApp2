//
- main.go Dosyası
Go dilinin fiber frameworkunu kullanacagım icin "github.com/gofiber/fiber/v2" package getleyip importluyorum.
Database işlemleri için olusturacagım database dosyasını ve handlerlarımın olacagı dosyayı importluyorum ("github.com/cumhurmesuttokalak/go-fiberApp2/database"
						"github.com/cumhurmesuttokalak/go-fiberApp2/routes")
-bir tane fiber.new funcı tipinde app degişkeni oluşturuyorum.İşlemlerimi bu değişken üzerinden yapacagım.
rotalarımı ayrı bir fonksiyonda kullanacagım icin bir setupfonksiyonu oluşturuyorum ve bu func app *fiber.App tipinde bir değişken alacak ki endpoint işlemlerini yapmam için.
Sonra sırasıyla endpointlerimi oluşturuyorum. 
app.Get("/api/users",rotues.Getusers) gibi
main funcı icinde oluşturdugum setupfonksiyonunu cagırıyorum ve app değişkenimi veriyorum.
app.Listen ile dinleyecegim portu belirliyorum ve işlemlerim yapılabilir halde artık.

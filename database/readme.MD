- database.go Dosyası
-Database bağlantısını kurmak için bir DBInstance tipinde struct oluşturuyorum
-Db adın bir field olusturup tipi *gorm.Db olacak
-Sonra Database adında bir değişken tanımlıyorum tipi de DbInstance
-DatabaseConnection adında bir func oluşturuyorum içine bir şey almıyor ve geri birşey dönmüyor.
Foksiyonun içi
-db,err:=gorm.Open(sqlite.Open("databaseIsmi"),&gorm.Config{})
sonra db.Automigrate işlemi yapıyoruz fakat bunun için bir modul dosyası oluşturmalıyız.
db. farklı işlemler yapabiliriz burada 
en son Database=DBInstance{Db:db} ye eşitliyoruz.
Artık DatabaseConnection fonksiyonumuzu diğer dosyalara importlayarak Database değişkenimiz ile işlemler yapabiliriz.
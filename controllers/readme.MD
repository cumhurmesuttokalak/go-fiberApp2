//
- controllers/user.go Dosyası

Önce ben bir tane user tipinde bir struct oluşturuyorum bu structum da models.user tipine bağlı olarak işlemler yapılacak fakat geri dönmesini istedigim fieldları kendim belirleyeceğim.

//CreateResponseUser fonksiyonu
Hemen geri dönmesini istediğim fieldlara ait bir CreateResponseUser fonksiyonu yazıyorum bu fonksiyon içine models.User tipinde bir değişken alıp bizim User tipimizde geri dönüş yapacak.İstedğimiz fieldları tanımlayıp models.Userdaki değişkenlerle eşliyoruz.User{ID:models.User,FirstName:models.FirstName}... gibi

//FindUser fonksiyonu
Bize delete,get,update gibi işlemlerimizde işimize yarayacağı icin bir findUser fonksiyonu yazıyorum.Bu fonksiyon databasede böyle bir user var mı diye kontrol ediyorum.Fonksiyonumuz int tipinde bir id değişkeni ve *models.User tipinde bir user değişkeni tutuyor.Sonra bakıyoruz databasemizin User tablosundaki idsi ile vereceğimiz id aynı mı diye aynı ise böyle bir user vardır işlemlerimizi yapabiliriz.Eğer userID 0 a eşitse de böyle bir user yoktur diyip nil dönüşümüzü yapıyoruz.

//CreateUser fonksiyonu
CreateUser endpointimiz için *fiber.Ctx tipinde bir değişken alıp err dönüyor.models.User tipinde bir nesne tanımlayıp bu tipi json verisi halinde göndereceğimiz için bodyParser yapmalıyız pointeri ile.Sonra bodyParser fonksiyonu bir err geri döndürdüğü için err kontrolu yapıyoruz.Sonra artık Db.Create(&user) ile crate yapıyoruz.Ve geriye sadece bizim istediğimiz değerler döneceği icin bir değişken oluşturup CreateResponseUser fonksiyonunu cagırıp user değişkeni veriyoruz.En son olarak bu değişkenimizi c.Status().JSON(responseUser) olarak dönüyoruz.

//GetUsers fonksiyonu
GetUsers endpointimiz için *fiber.Ctx tipinde bir değişken alıp err dönüyor.models.User dizisi tipinde bir nesne tanımlıyorum.Bir tane de User dizisi tipinde bir dizi nesnesi tanımlıyorum.Sonra db.Find(&users) ile databasemde bu nesneyi buluyorum.Sonra bir for range döngüsü kuruyorum.Bu döngü usersları dönecek.Bir tane responseUser değişkeni tanımlayıp CreateResponseUser fonksiyonunu cagırıp döngü değişkenini veriyorum.Ve en son User tipinde oluşturmuş oldugumu nesneye append ediyorum.Döngüyü bitirdikten sonra 
c.status().JSON(&responseUser) ile responsumu dönüyorum.

//GetUser fonksiyonu
GetUser endpointimiz için *fiber.Ctx tipinde bir değişken alıp err dönüyor.models.User tipinde bir nesne tanımlıyorum.Id ile verimi cekeceğim için id ve err değişkeni tanımlayıp c.ParamsInt("id")ile id mi yakalıyorum.Burada bir err fonksiyonu cagırıyorum id parametresinin int bir değer girip girilmediğini kontrol ediyor.Err den sonra int girildiğini kabul ediyorum ve findUser fonksiyonumu cagırıyorum.Databasemde böyle bir id ye sahip user var mı yok mu diye kontrol ettiriyorum.findUser fonksiyonum bir err döndügü icin erryi tekrar kontrol ettiriyorum.Bütün err leri kontrol ettirdim artık herhangi bir err yoksa responseUser değişkeni oluşturup CreateResponseUser fonksiyonumu cagırıp içine useri veriyorum.Ve return c.Status(400).JSON(responseUser) ı dönüyorum.

//DeleteUser fonksiyonu
DeleteUser endpointimiz için *fiber.Ctx tipinde bir değişken alıp err dönüyor.models.User tipinde bir nesne tanımlıyorum.Id ile verimi cekeceğim için id ve err değişkeni tanımlayıp c.ParamsInt("id")ile id mi yakalıyorum.Burada bir err fonksiyonu cagırıyorum id parametresinin int bir değer girip girilmediğini kontrol ediyor.Err den sonra int girildiğini kabul ediyorum ve findUser fonksiyonumu cagırıyorum.Databasemde böyle bir id ye sahip user var mı yok mu diye kontrol ettiriyorum.findUser fonksiyonum bir err döndügü icin erryi tekrar kontrol ettiriyorum.Bütün err leri kontrol ettirdim artık herhangi bir err yoksa Db.Delete(&user) ile o useri siliyorum.Ve return c.Status(200).SendString("Veri silinmiştir") mesajını döndürüyorum.

//UpdateUser fonksiyonu
UpdateUser endpointimiz için *fiber.Ctx tipinde bir değişken alıp err dönüyor.models.User tipinde bir nesne tanımlıyorum.Id ile verimi cekeceğim için id ve err değişkeni tanımlayıp c.ParamsInt("id")ile id mi yakalıyorum.Burada bir err fonksiyonu cagırıyorum id parametresinin int bir değer girip girilmediğini kontrol ediyor.Err den sonra int girildiğini kabul ediyorum ve findUser fonksiyonumu cagırıyorum.Databasemde böyle bir id ye sahip user var mı yok mu diye kontrol ettiriyorum.findUser fonksiyonum bir err döndügü icin erryi tekrar kontrol ettiriyorum.Artık böyle bir id ye sahip bir Userım var şimdi onu güncellemek için bir tane UpdateUser structu oluşturup sadece Fname ve Lname fieldlarını belirtiyorum.Hemen bu tipin nesnesini oluşturup.c.BodyParser(&UpdateData) yı parse ediyorum.Err kontrolü yaptıktan sonra user.FirstName=UpdateData.Fname ,user.LastName=UpdateData.Lname diye güncelliyorum.Ve db tarafında da güncellemem gerekiyor.Db.Save(&user).Yine bir responseUser değişkeni tanımlayıp CreateResponseUser fonksiyonu icin user i verip return ile responseUserimi dönüyorum.

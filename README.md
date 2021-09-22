# go-starter-kit-v2

Go starter kit serisinin ikinci versiyonu olan bu repo ile hızlıca Go dili ile rest api geliştirmeye başlayabilirsiniz.

## Projede neler var?

* Yüksek performanslı ve express.js'in Go diline uyarlanması olan [Fiber](https://github.com/gofiber/fiber) çatısını kullandım.
* Environment variable configurasyonu için [godotenv](https://github.com/joho/godotenv) kullandım.
* Veritabanı olarak MongoDB ve resmi [Driver eklentisini](https://github.com/mongodb/mongo-go-driver) kullandım.
* Public klasörü ile html ve statik dosyalarınızı sunabilirsiniz.

## Kurulum

Proje ana dizininde .env adında bir dosya oluşturun. .env.example dosyası içindeki ayarları kopyalayıp .env dosyasının içine yapıştırın.

"PORT" -> Uygulamayı hangi porttan çalıştırmak istiyorsanız burdan değiştirebilirsiniz.

"MONGO_URI" -> MongoDB bağlantısı için gerekli ayarı set edebilirsiniz. [MongoDB Cloud](https://www.mongodb.com/cloud/atlas) üzerinden ücretsiz bir MongoDB veritabanı oluşturabilirsiniz.

## Çalıştır

``` 
go run main.go
```

## License
MIT

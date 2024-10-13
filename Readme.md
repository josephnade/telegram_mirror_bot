# Telegram Mirror Bot

Bu proje, Golang kullanılarak geliştirilmiş basit bir Telegram botudur. Bu bot, kullanıcılardan gelen mesajları dinler ve onlara yanıt verir. Telegram botunuzu oluşturmak, ayarlamak ve kullanıcılarla iletişim kurmak için başlangıç seviyesi bir çalışma olarak tasarlanmıştır.

## Özellikler

- Kullanıcılardan gelen mesajları dinleme.
- Gelen her mesaja "Merhaba, mesajınızı aldım! '{KULLANICI MESAJI}'" şeklinde otomatik yanıt verme.

## Gerekli Araçlar ve Kurulum

1. **Golang Kurulumu**
   
   Golang henüz kurulu değilse, [Golang indirme sayfası](https://golang.org/dl/) üzerinden indirip kurun.

2. **Telegram Bot Token'ı Alın**
   
   Telegram'da `@BotFather` botunu kullanarak yeni bir bot oluşturun ve size verilen `API Token'ı` kaydedin.

3. **Proje Kurulumu**

   - Projeyi klonlayın:
     ```bash
     git clone https://github.com/josephnade/telegram-mirror_bot.git
     cd telegram-bot
     ```
   - Gerekli bağımlılıkları yükleyin:
     ```bash
     go mod tidy
     ```

4. **Telegram Bot API Kütüphanesini Yükleyin**

   Proje Golang'de `tgbotapi` kütüphanesini kullanmaktadır. Kütüphaneyi yüklemek için:
   ```bash
   go get github.com/go-telegram-bot-api/telegram-bot-api/v5
   ```

5. **Bot Token'ı Ayarlama**

   - Ana bot dosyasına gidin (`main.go`) ve `BOT_TOKENINIZ` yazan kısma kendi bot token'ınızı girin.

## Botun Çalıştırılması

```bash
go run main.go
```

### Değişiklikler ve İyileştirmeler

- Hataları daha anlaşılır hale getirmek için `log.Fatalf` kullanıldı.
- Botun hata ayıklama modunun ayarı değiştirildi (`bot.Debug = false`).
- Gereksiz mesaj işlemlerini atlamak için null kontrol eklendi (`if update.Message == nil { continue }`).

## Katkıda Bulunma

Eğer projeye katkıda bulunmak isterseniz, lütfen bir çekme isteği (pull request) oluşturun veya bir hata bildirimi açın.

## Lisans

Bu proje MIT Lisansı altında sunulmaktadır. Daha fazla bilgi için `LICENSE` dosyasına bakın.
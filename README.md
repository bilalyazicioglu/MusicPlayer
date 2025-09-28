# 🎵 Go Müzik Çalar

Terminal tabanlı basit müzik çalar uygulaması. MP3 dosyalarını çalar ve ses kontrolü sağlar.

## ✨ Özellikler

- MP3 dosya desteği
- Ses seviyesi kontrolü
- Duraklat/devam ettir
- Döngülü çalma

## 🚀 Kurulum

```bash
go mod tidy
go run main.go
```

## 🎮 Kontroller

| Komut | Açıklama |
|-------|----------|
| `p` | Duraklat/Devam Et |
| `+` veya `=` | Sesi Aç |
| `-` veya `_` | Sesi Kıs |
| `v [sayı]` | Belirli ses seviyesi (örn: `v -2`) |
| `q` | Çıkış |

## 📊 Ses Seviyeleri

- **0** = Normal ses
- **+1 → +5** = Yüksek ses
- **-1 → -10** = Düşük ses

## 📦 Bağımlılıklar

- [github.com/faiface/beep](https://github.com/faiface/beep) - Ses işleme kütüphanesi

## 📝 Notlar

MP3 dosyasını proje klasörüne koyup `main.go` içinde dosya adını güncelleyin.
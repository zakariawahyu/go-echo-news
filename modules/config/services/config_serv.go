package services

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-news/entity"
	"github.com/zakariawahyu/go-echo-news/modules/channel"
	"github.com/zakariawahyu/go-echo-news/modules/config"
	"github.com/zakariawahyu/go-echo-news/modules/region"
	"github.com/zakariawahyu/go-echo-news/modules/sub_channel"
	"github.com/zakariawahyu/go-echo-news/pkg/helpers"
	"github.com/zakariawahyu/go-echo-news/pkg/logger"
	"time"
)

type configServices struct {
	configRepo      config.ConfigRepository
	configRedisRepo config.ConfigRedisRepository
	channelRepo     channel.ChannelRepository
	subChannelRepo  sub_channel.SubChannelRepository
	regionRepo      region.RegionRepository
	zapLogger       logger.Logger
	contextTimeout  time.Duration
}

func NewConfigServices(configRepo config.ConfigRepository, configRedisRepo config.ConfigRedisRepository, channelRepo channel.ChannelRepository, subChannelRepo sub_channel.SubChannelRepository, regionRepo region.RegionRepository, zapLogger logger.Logger, timeout time.Duration) config.ConfigServices {
	return &configServices{
		configRepo:      configRepo,
		configRedisRepo: configRedisRepo,
		channelRepo:     channelRepo,
		subChannelRepo:  subChannelRepo,
		regionRepo:      regionRepo,
		zapLogger:       zapLogger,
		contextTimeout:  timeout,
	}
}

func (serv *configServices) GetAllConfig(ctx context.Context) (configs []entity.ConfigResponse) {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	redisData, err := serv.configRedisRepo.GetAllConfig(c, helpers.KeyRedis("config", ""))
	if redisData != nil {
		return entity.NewConfigArrayResponse(redisData)
	}

	res, err := serv.configRepo.GetAll(c)
	if err != nil {
		serv.zapLogger.Errorf("configServ.GetAllConfig.configRepo.GetAll, err = %s", err)
		panic(err)
	}

	configs = entity.NewConfigArrayResponse(res)

	if err = serv.configRedisRepo.SetAllConfig(c, helpers.KeyRedis("config", ""), helpers.Slowest, configs); err != nil {
		serv.zapLogger.Errorf("configServ.GetAllConfig.configRedisRepo.SetAllConfig, err = %s", err)
		panic(err)
	}

	return configs
}

func (serv *configServices) GetMetas(ctx context.Context, types string, key string) interface{} {
	c, cancel := context.WithTimeout(ctx, serv.contextTimeout)
	defer cancel()

	var data interface{}
	var err error

	if types == "channel" {
		data, err = serv.channelRepo.GetMetas(c, key)
		if err != nil {
			serv.zapLogger.Errorf("configServ.GetMetas.channelRepo.GetMetas, err = %s", err)
			panic(err)
		}
	} else if types == "subchannel" {
		data, err = serv.subChannelRepo.GetMetas(c, key)
		if err != nil {
			serv.zapLogger.Errorf("configServ.GetMetas.subChannelRepo.GetMetas, err = %s", err)
			panic(err)
		}
	} else if types == "region" {
		data, err = serv.regionRepo.GetMetas(c, key)
		if err != nil {
			serv.zapLogger.Errorf("configServ.GetMetas.regionRepo.GetMetas, err = %s", err)
			panic(err)
		}
	} else {
		data = OtherMeta(key)
		if data == nil {
			serv.zapLogger.Errorf("configServ.GetMetas.NotFound, err = %s", helpers.ErrNotFound)
			panic(helpers.ErrNotFound)
		}
	}

	return data
}

func OtherMeta(key string) interface{} {
	data := echo.Map{
		"homepage": echo.Map{
			"title":       "Berita Terkini dan Informasi Terbaru Hari Ini",
			"description": "iNews ID - Situs portal berita nasional dan daerah yang menyajikan informasi terkini dan terbaru seperti, Berita Politik, Hukum, Keuangan, Teknologi",
		},
		"multimedia": echo.Map{
			"title":       "Berita Hari ini, Berita Terbaru dan Terkini - Foto dan Video",
			"description": "Kumpulan Berita Video dan Foto Inews ID, Berupa Video dan Foto Peristiwa berita nasional dan internasional",
		},
		"video": echo.Map{
			"title":       "Berita Video Hari Ini - iNews Portal",
			"description": "Galeri Video iNews.id, kumpulan video terkini yang terjadi di Indonesia dan Dunia",
		},
		"photo": echo.Map{
			"title":       "Berita Foto Terbaru Hari Ini - iNews Portal",
			"description": "Galeri Foto iNews.id berisi kumpulan foto peristiwa yang terjadi di dalam dan luar negeri",
		},
		"infografis": echo.Map{
			"title":       "Berita Infografis Terbaru Hari Ini - iNews Portal",
			"description": "Galeri Infografis iNews.id berisi kumpulan Infografis peristiwa yang terjadi di dalam dan luar negeri",
		},
		"streaming": echo.Map{
			"title":       "Streaming iNews TV",
			"description": "Live Streaming I-News, Nonton Online Gratis, Nonton Video Gratis, Streaming TV Online hanya di Inews ID",
		},
		"streaming-rcti": echo.Map{
			"title":       "Live Streaming RCTI - iNews Portal",
			"description": "Live Streaming RCTI, Nonton Online Gratis, Nonton Video Gratis, Streaming TV Online hanya di Inews ID",
		},
		"streaming-mnctv": echo.Map{
			"title":       "Live Streaming MNCTV - iNews Portal",
			"description": "live streaming MNCTV, Nonton Online Gratis, Nonton Video Gratis, Streaming TV Online hanya di Inews ID",
		},
		"streaming-gtv": echo.Map{
			"title":       "Live Streaming GTV - iNews Portal",
			"description": "live streaming GTV, Nonton Online Gratis, Nonton Video Gratis, Streaming TV Online hanya di Inews ID",
		},
		"streaming-idx": echo.Map{
			"title":       "Live Streaming IDX Channel - iNews Portal",
			"description": "live streaming IDX Channel, Nonton Online Gratis, Nonton Video Gratis, Streaming TV Online hanya di Inews ID",
		},
		"streaming-mncnews": echo.Map{
			"title":       "Live Streaming MNC News - iNews Portal",
			"description": "live streaming MNC News, Nonton Online Gratis, Nonton Video Gratis, Streaming TV Online hanya di Inews ID",
		},
		"popular": echo.Map{
			"title":       "Berita Populer  - iNews Portal",
			"description": "Rangkuman berita terpopuler dan terlengkap di Portal iNews",
		},
		"tentang-kami": echo.Map{
			"title":       "Tentang Kami - iNews Portal",
			"description": "iNews.id hadir untuk menyajikan dan memberikan informasi berita terkini, peristiwa yang terjadi di dalam dan luar negeri",
		},
		"kode-etik": echo.Map{
			"title":       "Kode Etik - iNews Portal",
			"description": "Kode Etik iNews.id, informasi tentang pedoman pemberitaan media siber",
		},
		"term-of-services": echo.Map{
			"title":       "Term of Services - iNews Portal",
			"description": "Syarat dan Ketentuan yang mengatur penggunaan anda yang mengakses situs iNews.id",
		},
		"privacy-policy": echo.Map{
			"title":       "Privacy Policy - iNews Portal",
			"description": "Privacy policy iNews.id mengungkapkan kebijakan penanganan data-data pribadi Anda pada saat Anda mengkses situs kami",
		},
		"kontak-kami": echo.Map{
			"title":       "Kontak Kami - iNews Portal",
			"description": "Jika anda memiliki pertanyaan, saran atau pendapat mengenai iNews.id yang ingin disampaikan, silahkan hubungi kami melalui email/ no telp berikut.",
		},
		"redaksi": echo.Map{
			"title":       "Redaksi - iNews Portal",
			"description": "",
		},
		"term-of-service": echo.Map{
			"title":       "Term Of Service - iNews Portal",
			"description": "iNews ID - Situs portal berita nasional dan daerah yang menyajikan informasi terkini dan terbaru seperti, Berita Politik, Hukum, Keuangan, Teknologi",
		},
		"disclaimer": echo.Map{
			"title":       "Disclaimer - iNews Portal",
			"description": "",
		},
		"pedoman-media-siber": echo.Map{
			"title":       "Pedoman Pemberitaan Media Siber - iNews Portal",
			"description": "Kode Etik iNews.id, informasi tentang pedoman pemberitaan media siber",
		},
		"sitemap": echo.Map{
			"title":       "Sitemap - iNews Portal",
			"description": "",
		},
		"tag": echo.Map{
			"title":       "Berita Harian {{tag}}",
			"description": "Kumpulan berita dan informasi seputar {{tag}}",
		},
		"topic": echo.Map{
			"title":       "Berita {{topic}} Hari Ini",
			"description": "Menyajikan berita aktual dan informasi seputar {{topic}} hari ini",
		},
		"search": echo.Map{
			"title":       "Hasil pencarian untuk {{search}}",
			"description": "Kumpulan berita harian terkini dan terlengkap seputar {{search}} hanya di iNews.id",
		},
		"ini-indonesia": echo.Map{
			"title":       "Berita Populer Destinasi Wisata Indonesia - iNews Portal",
			"description": "Berita terkini yang terjadi di seluruh Indonesia, informasi terhangat dan terpercaya",
		},
		"update-korona": echo.Map{
			"title":       "Virus Korona - iNews Portal",
			"description": "Update Berita Informasi Terbaru Mengenai Virus Korona - Virus Corona (covid-19) secara lengkap dan terupdate",
		},
		"ramadan": echo.Map{
			"title":       "Jadwal Imsakiyah, Buka Puasa dan Sholat di Bulan Ramadhan {{date('Y')}} - iNews.id Portal",
			"description": "Berita Terkini Seputar Jadwal Imsakiyah, Jadwal Buka Puasa, Jadwal Sholat, Puasa, Ceramah Singkat Ramadhan  {{date('Y')}}",
			"keyword":     "berita terkini, jadwal imsakiyah hari ini, jadwal imsak hari ini, jadwal imsak, jadwal imsakiyah, jadwal buka puasa hari ini, jadwal buka puasa, ramadhan, jadwal sholat, tata cara, tarawih, doa, niat, tahajud, lailatul qadar",
		},
		"hikmah": echo.Map{
			"title":       "Hikmah Bulan Puasa {{date('Y')}} - iNews Portal",
			"description": "Hujan berkah dan ampunan menjadi hikmah Ramadhan yang selanjutnya. Seorang muslim yang berpuasa semata-mata karena keimanannya kepada Allah",
		},
		"tausiah": echo.Map{
			"title":       "Tausiyah dan Ceramah Ramadhan {{date('Y')}} - iNews.id",
			"description": "Berita Terkini Seputar Tausiyah, Puasa, dan Ceramah Singkat Ramadhan {{date('Y')}}",
			"keyword":     "tausiyah ramadhan, ceramah singkat ramadhan, puasa",
		},
		"kalkulator-zakat": echo.Map{
			"title":       "Kalkulator Zakat Maal dan Penghasilan Online - iNews.id",
			"description": "Berikut kalkulator zakat maal dan zakat penghasilan online hanya di iNews.id",
			"keyword":     "kalkulator zakat, zakat penghasilan, perhitungan berdasarkan hutang, nisab zakat",
		},
		"quran": echo.Map{
			"title":       "Al Quran Digital 30 Juz dan Terjemahan - iNews.id",
			"description": "Baca Al Quran Digital 30 Juz lengkap dengan terjemahannya, ayat-ayat suci Al-Qur'an hanya di iNews.id",
			"keyword":     "Al quran digital 30 juz, al quran digital, al quran dan terjemahan, ayat al quran dan terjemahan",
		},
		"murottal": echo.Map{
			"title":       "Murottal Al Quran Bulan Ramadhan {{date('Y')}} - iNews.id",
			"description": "Murottal Al Quran, Streaming Al-Qur'an 30 Juz, Streaming murottal Quran MP3, Murottal Al Quran 30 juz hanya di iNews.id",
			"keyword":     "Murottal Al Quran, Streaming murottal Quran MP3, Murottal Al Quran 30 juz",
		},
		"imsakiyah": echo.Map{
			"title":       "Jadwal Imsak Ramadhan {{date('Y')}} - Jadwal Sholat",
			"description": "Jadwal Sholat atau Waktu Sholat Jakarta dan Seluruh wilayah Indonesia dapat Anda akses di iNews Portal",
			"keyword":     "jadwal imsakiyah hari ini, jadwal imsak hari ini, jadwal imsak, jadwal imsakiyah, jadwal buka puasa hari ini, jadwal buka puasa, ramadhan, jadwal sholat, tata cara, tarawih, doa, niat, tahajud",
		},
		"religi": echo.Map{
			"title":       "Berita Terkini Seputar Ramadhan {{date('Y')}} - iNews.id",
			"description": "Berita Terkini Seputar Jadwal Imsakiyah, Jadwal Buka Puasa, Jadwal Sholat, Puasa, Ceramah Singkat Ramadhan {{date('Y')}}",
			"keyword":     " jadwal imsakiyah hari ini, jadwal imsak hari ini, jadwal imsak, jadwal imsakiyah, jadwal buka puasa hari ini, jadwal buka puasa, jadwal sholat, ramadhan, tata cara, tarawih, doa, niat, tahajud",
		},
		"doa-harian": echo.Map{
			"title":       "Doa Harian Seputar Ramadhan {{date('Y')}} - iNews.id",
			"description": "Informasi seputar doa harian, doa setelah sholat, keutamaan, tarawih, tahajud, dan dhuha hanya di iNews.id",
			"keyword":     "doa setelah tahajud, doa setelah sholat, ramadhan, tata cara, tarawih, niat, tahajud, keutamaan",
		},
		"sejarah-islam": echo.Map{
			"title":       "Informasi Sejarah Islam, Kisah Nabi, Hadits Seputar Ramadhan {{date('Y')}} - iNews.id",
			"description": "Informasi seputar doa harian, doa setelah sholat, keutamaan, tarawih, tahajud, dan dhuha hanya di iNews.id",
			"keyword":     "sejarah islam, kisah nabi, kisah rasul, sholat, tarawih, witir, lailatul qadar",
		},
		"resep": echo.Map{
			"title":       "Informasi Resep Makanan/Masakan Bulan Ramadhan {{date('Y')}} - iNews.id",
			"description": "Informasi seputar resep makanan dan minuman selama bulan suci Ramadhan",
			"keyword":     "menu buka puasa, menu ramadhan, resep masakan di bulan puasa, resep makanan untuk sahur dan buka puasa, resep minuman untuk sahur dan buka puasa, ramadhan {{date('Y')}}, inews",
		},
		"sehat": echo.Map{
			"title":       "Informasi Kesehatan Bulan Ramadhan - iNews Portal",
			"description": "Berita Hari ini dan Informasi terbaru, terlengkap seputar kesehatan Selama bulan Suci Ramadhan ",
		},
		"bmptn": echo.Map{
			"title":       "Pengumuman SBMPTN 2020",
			"description": "Pengumuman hasil Seleksi Bersama Masuk Perguruan Tinggi Negeri (SBMPTN) 2020 akan dilakukan pada hari ini 14 agustus 2020",
			"image":       "files/inews_new/2020/08/10/SBMPTN__ltmpt_.jpg",
		},
		"404": echo.Map{
			"title":       "Oops halaman tidak ditemukan (404) - iNews Portal",
			"description": "Oops halaman tidak ditemukan (404) - iNews Portal",
		},
		"piala-aff-2020": echo.Map{
			"title":       "Berita Piala AFF 2020 - iNews Portal",
			"description": "Berita Piala AFF 2020 - iNews Portal",
		},
		"piala-aff-2020-laga": echo.Map{
			"title":       "Berita Laga Piala AFF 2020 - iNews Portal",
			"description": "Berita Laga Piala AFF 2020 - iNews Portal",
		},
		"piala-aff-2020-serba-serbi": echo.Map{
			"title":       "Berita Serba Serbi Piala AFF 2020 - iNews Portal",
			"description": "Berita Serba Serbi Piala AFF 2020 - iNews Portal",
		},
		"piala-aff-2020-man-of-the-match": echo.Map{
			"title":       "Berita Man of The Match Piala AFF 2020 - iNews Portal",
			"description": "Berita Man of The Match Piala AFF 2020 - iNews Portal",
		},
		"kaleidoskop-2021": echo.Map{
			"title":       "Berita Kaleidoskop 2021 - iNews Portal",
			"description": "Berita Kaleidoskop 2021 - iNews Portal",
		},
		"outlook-2022": echo.Map{
			"title":       "Berita Outlook 2022 - iNews Portal",
			"description": "Berita Outlook 2022 - iNews Portal",
		},
		"tryout-2022": echo.Map{
			"title":       "Tryout 2022 - Tryout Online UTBK SBMPTN 2022",
			"description": "Tryout 2022 - Tryout Online UTBK SBMPTN 2022",
		},
		"pesta-bola": echo.Map{
			"title":       "Berita Seputar Piala Dunia Qatar 2022 - iNews Portal",
			"description": "Berita Seputar Piala Dunia Qatar 2022 - iNews Portal",
		},
		"pesta-bola-laga": echo.Map{
			"title":       "Berita Laga Piala Dunia Qatar 2022 - iNews Portal",
			"description": "Berita Laga Piala Dunia Qatar 2022 - iNews Portal",
		},
		"pesta-bola-serbadashserbi": echo.Map{
			"title":       "Berita Serba Serbi Piala Dunia Qatar 2022 - iNews Portal",
			"description": "Berita Serba Serbi Piala Dunia Qatar 2022 - iNews Portal",
		},
		"pesta-bola-bintang": echo.Map{
			"title":       "Berita Pemain Bintang Piala Dunia Qatar 2022 - iNews Portal",
			"description": "Berita Pemain Bintang Piala Dunia Qatar 2022 - iNews Portal",
		},
		"piala-aff-2022": echo.Map{
			"title":       "Berita Terkini Hari Ini Seputar Jadwal dan Hasil Piala AFF 2022 - iNews.id",
			"description": "Berita Bola Piala AFF, Terkini Hari Ini, AFF, Skor, Pemain Sepak Bola, Grup, live streaming, profil pemain, Final",
		},
		"laga-aff-2022": echo.Map{
			"title":       "Laga Terkini Seputar Piala AFF 2022 - iNews.id",
			"description": "Laga Sepak Bola Indonesia, Liga Indonesia dan Timnas Indonesia di Piala AFF 2022",
		},
		"bintang-aff-2022": echo.Map{
			"title":       "Pemain Bintang Terkini Seputar Piala AFF 2022 - iNews.id",
			"description": "Pemain Bintang Sepak Bola Indonesia, Liga Indonesia dan Timnas Indonesia di Piala AFF 2022",
		},
		"serba-serbi-aff-2022": echo.Map{
			"title":       "Serba Serbi Terkini Seputar Piala AFF 2022 - iNews.id",
			"description": "Serba Serbi Timnas Indonesia Mengenai Jadwal Indonesia, Hasil, Klasemen, Berita dan Kabar Timnas Terbaru di Piala AFF 2022",
		},
		"pemilu-rakyat": echo.Map{
			"title":       "Berita Terkini Terbaru Hari Ini Speutar Pemilu Pilpres Pileg 2024 - iNews.id",
			"description": "Berita Terkini Pemilu 2024, Kabar Terkini Pilpres Presiden - Wakil Presiden 2024, dan Quick Count Hasil Perhitungan, Pilpres 2024, Pileg 2024, Live Quickcount",
		},
		"pilpres-2024": echo.Map{
			"title":       "Berita Terkini Terbaru Hari Ini Seputar Profil Calon Presiden dan Wakil Presiden 2024 - iNews.id",
			"description": "Berita Terkini, Berita hari Ini, Berita Pemilu 2024, Berita Pilpres 2024, Berita Calon Presiden dan Calon Wakil Presiden, Profil Calon Presiden dan Calon Wakil Presiden 2024",
		},
		"pileg-2024": echo.Map{
			"title":       "Berita Terkini Terbaru Hari Ini Profil Pileg 2024 - iNews.id",
			"description": "Pileg 2024, berita Pileg 2024, Pileg 2024 terpercaya, Pileg 2024 hari ini, Pileg 2024 terkini",
		},
		"capres-2024": echo.Map{
			"title":       "Berita Terkini Terbaru Hari Ini Profil Calon Presiden dan Calon Wakil Presiden 2024- iNews.id",
			"description": "Berita Profil Calon Presiden dan Calon Wakil Presiden Terpercaya,Berita pileg 2024, Berita Pilpres 2024, Berita Pemilu 2024, Profil Capres dan Cawapres 2024",
		},
		"pemilupedia": echo.Map{
			"title":       "Berita Terkini Terbaru Hari Ini Pemilu, Pilpres, Pileg 2024 - iNews.id",
			"description": "Berita Terkini Pemilu 2024, Berita Terbaru Pemilu 2024, Pilpres 2024, Pileg 2024, Profil Capres dan Cawapres",
		},
		"quickcount-2024": echo.Map{
			"title":       "Berita Terkini Terbaru Hari Ini Hasil Perhitungan Pemilu Pilpres Pileg 2024 - iNews.id",
			"description": "Berita Terkini Pemilu 2024, Kabar Terkini Pilpres, dan Quick Count Hasil Perhitungan, Hasil Quickcount, Quickcount Pilpres 2024, Quickcount Pemilu, Quickcount Pileg, Live Quickcount",
		},
	}

	return data[key]
}

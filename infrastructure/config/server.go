package config

type ServerConfig struct {
	Host string // サーバーがバインド(待ち受け)するIPアドレス
	Port uint16 // サーバーがバインド(待ち受け)するポート番号
}

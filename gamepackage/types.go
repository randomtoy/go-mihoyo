package gamepackage

type Response struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	GamePackages []GamePackage `json:"game_packages"`
}

type GamePackage struct {
	Game        Game        `json:"game"`
	Main        Main        `json:"main"`
	PreDownload PreDownload `json:"pre_download"`
}

type Game struct {
	ID  string `json:"id"`
	Biz string `json:"biz"`
}

type Main struct {
	Major   Major   `json:"major"`
	Patches []Patch `json:"patches"`
}

type Major struct {
	Version    string     `json:"version"`
	GamePkgs   []Package  `json:"game_pkgs"`
	AudioPkgs  []AudioPkg `json:"audio_pkgs"`
	ResListURL string     `json:"res_list_url"`
}

type Patch struct {
	Version    string     `json:"version"`
	GamePkgs   []Package  `json:"game_pkgs"`
	AudioPkgs  []AudioPkg `json:"audio_pkgs"`
	ResListURL string     `json:"res_list_url"`
}

type Package struct {
	URL              string `json:"url"`
	MD5              string `json:"md5"`
	Size             string `json:"size"`
	DecompressedSize string `json:"decompressed_size"`
}

type AudioPkg struct {
	Language         string `json:"language"`
	URL              string `json:"url"`
	MD5              string `json:"md5"`
	Size             string `json:"size"`
	DecompressedSize string `json:"decompressed_size"`
}

type PreDownload struct {
	Major   *Major  `json:"major"`
	Patches []Patch `json:"patches"`
}

package database

type Database struct {
	Info        DatabaseServerInfo       `json:"info"`
	Posts       []DatabasePostItem       `json:"posts"`
	Playlists   []DatabasePlaylistItem   `json:"playlists"`
	Levels      []DatabaseLevelItem      `json:"levels"`
	Skins       []DatabaseSkinItem       `json:"skins"`
	Backgrounds []DatabaseBackgroundItem `json:"backgrounds"`
	Effects     []DatabaseEffectItem     `json:"effects"`
	Particles   []DatabaseParticleItem   `json:"particles"`
	Engines     []DatabaseEngineItem     `json:"engines"`
	Replays     []DatabaseReplayItem     `json:"replays"`
}

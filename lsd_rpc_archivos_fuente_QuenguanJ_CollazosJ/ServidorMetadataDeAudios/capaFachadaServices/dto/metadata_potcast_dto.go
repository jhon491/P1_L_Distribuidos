package dto

type MetadataPodcastDTO struct {
	Id             int    `json:"id"`
	NombrePodcast  string `json:"nombrePodcast"`
	TituloEpisodio string `json:"tituloEpisodio"`
	Anfitrion      string `json:"anfitrion"`
	EpisodioNumero int    `json:"episodioNumero"`
	Notas          string `json:"notas"`
	Clasificacion  string `json:"clasificacion"`
}

type RespuestaMetadataPodcastDTO struct {
	ObjPodcast MetadataPodcastDTO `json:"objPodcast"`
	Codigo     int                `json:"codigo"`
	Mensaje    string             `json:"mensaje"`
}

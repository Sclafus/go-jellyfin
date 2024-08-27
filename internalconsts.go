package jellyfin

type mediaItemType string

const (
	mediaTypeAlbum        mediaItemType = "MusicAlbum"
	mediaTypeArtist       mediaItemType = "MusicArtist"
	mediaTypeAudio        mediaItemType = "Audio"
	mediaTypePlaylist     mediaItemType = "Playlist"
	folderTypePlaylists   mediaItemType = "PlaylistsFolder"
	folderTypeCollections mediaItemType = "CollectionFolder"
	mediaTypeGenre        mediaItemType = "Genre"
	mediaTypeMovie        mediaItemType = "Movie"
)

const (
	errInvalidRequest       = "invalid request"
	errUnexpectedStatusCode = "unexpected statuscode"
	errServerError          = "server error"
	errNotFound             = "page not found"
	errUnauthorized         = "needs authorization"
	errForbidden            = "forbidden"
)

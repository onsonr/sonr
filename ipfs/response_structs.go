package ipfs

type DownloadResponse struct {
	Status    int
	FileData  []byte
	NameSpace string
	Path      string
}

type UploadResponse struct {
	Status int
	Cid    string
	Path   string
}

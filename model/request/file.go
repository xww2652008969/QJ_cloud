package request

type ReqFileUpload struct {
	Name       string
	Fileaddres string
	Filemd5    string
	File       []byte
}
type REqFileDelete struct {
	Name       string
	Fileaddres string
}
type ReqFileGetlist struct {
	FileAddres string
}

package model

type S3Conf struct {
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	Bucket          string `mapstructure:"bucket"`
}

type CDNConf struct {
	Domain string `mapstructure:"domain"`
}

type FileConf struct {
	S3  *S3Conf  `mapstructure:"s3"`
	CDN *CDNConf `mapstructure:"cdn"`
}

type UploadType string

const (
	UploadTypeUGC    UploadType = "ugc"
	UploadTypeSys    UploadType = "sys"
	UploadTypeAgency UploadType = "agency"
)

type FileFormat string

const (
	FileFormatJPG   FileFormat = "jpg"
	FileFormatPNG   FileFormat = "png"
	FileFormatGIF   FileFormat = "gif"
	FileFormatMP4   FileFormat = "mp4"
	FileFormatMP3   FileFormat = "mp3"
	FileFormatSVGA  FileFormat = "svga"
	FileFormatWEBP  FileFormat = "webp"
	FileFormatTXT   FileFormat = "txt"
	FileFormatExcel FileFormat = "excel"
	FileFormatZip   FileFormat = "zip"
	FileFormatGlb   FileFormat = "glb"
)

type UploadReq struct {
	UploadType UploadType `json:"upload_type" vd:"in($, 'ugc', 'sys')"`
	Format     FileFormat `json:"format" vd:"in($, 'jpg', 'png', 'gif', 'mp4', 'mp3', 'svga', 'webp', 'txt', 'excel', 'zip', 'glb')"`
	Filehash   string     `json:"filehash"` // 文件hash,如果为空,后端生成
}

type PartUploadReq struct {
	UploadType UploadType `json:"upload_type" vd:"in($, 'ugc', 'sys')"`
	Format     FileFormat `json:"format" vd:"in($, 'jpg', 'png', 'gif', 'mp4', 'mp3', 'svga', 'webp', 'txt', 'excel', 'zip', 'glb')"`
	Filehash   string     `json:"filehash"` // 文件hash,如果为空,后端生成
	PartNum    int64      `json:"part_num"` // 总分片数
}

type PartUploadInfo struct {
	Etag    string `json:"etag"`
	PartNum int64  `json:"part_num"`
}

type CompletePartUploadReq struct {
	FileName string            `json:"file_name"` // 文件hash,如果为空,后端生成
	Parts    []*PartUploadInfo `json:"parts"`
	UploadID string            `json:"upload_id"`
}

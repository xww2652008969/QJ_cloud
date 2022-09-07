package model

type File struct {
	ID          uint `gorm:"primaryKey"`
	Useruuid    string
	Name        string //文件名
	Md5         string //文件md5
	Fileaddress string //文件地址
	//Size        string //文件大小

}

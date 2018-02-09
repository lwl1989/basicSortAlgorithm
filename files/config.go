package files


type UploadConfig interface {
	SaveName() string
	SavePath() string
	Sha1()	string
	SaveBucket() string
	GetKey() string
}

type config struct {

}

type OmpConfig struct {

}

type UserConfig struct{

}




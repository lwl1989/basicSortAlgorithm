package image

import "strconv"

const FUNC_RATIO  = "RATIO"
const FUNC_RATIO_FILL  = "RATIO_FILLING"
const FUNC_CUT  = "CUT"
const FUNC_RATIO_EQUAL  = "RATIO_EQUAL"

type zoomOperate struct{
	MimeType string
	ZoomFunc string
	Suffix string
	Height uint64
	Width uint64
	S3Exists bool
}
/**
	根据请求产生缩图的对象指针
 */
func getOperate(fileInfo []string,length uint8) (*zoomOperate) {
	operate := newZoomOperate()
	switch length {
	case 6:
		operate.Suffix,operate.ZoomFunc = fileInfo[5],fileInfo[4]
		operate.Height,_ = strconv.ParseUint(fileInfo[3],10,64)
		operate.Width,_ = strconv.ParseUint(fileInfo[2],10,64)
		break
	case 5:
		operate.Suffix = fileInfo[4]
		operate.Height,_ = strconv.ParseUint(fileInfo[3],10,64)
		operate.Width,_ = strconv.ParseUint(fileInfo[2],10,64)
		break
	case 3:
		operate.Suffix = fileInfo[2]
		break
	case 0:
	default:
		operate.S3Exists = false
	}
	return operate
}


/**
产生一个默认的缩图的对象指针
 */
func newZoomOperate() *zoomOperate {
	return &zoomOperate{
		ZoomFunc:FUNC_RATIO_FILL,
		MimeType:"image/png",
		Width:0,
		Height:0,
		Suffix:"png",
		S3Exists:true,
	}
}


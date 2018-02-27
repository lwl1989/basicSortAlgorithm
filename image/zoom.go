package image

import (
	"fmt"
	"basicSortAlgorithm/awsS3"
	"github.com/aws/aws-sdk-go/service/s3"
	"gopkg.in/gographics/imagick.v1/imagick"
)

/**
根据请求产生缩图
 */
func DoZoom(file string) (code int, stream string) {
	fileInfo, length := getExp(file)
	operate := getOperate(fileInfo, length)

	if operate.S3Exists {
		fileName := fileInfo[1] + "." + fileInfo[length-1]
		fmt.Println(operate)
		out, err := awsS3.GetFileSteam(fileName)
		if err != nil {
			return
		}
		stream := ""
		operate.MimeType = *out.ContentType
		if operate.Width > 0 && operate.Height > 0 {
			switch operate.ZoomFunc {
			case FUNC_CUT:
				stream = doZoomCut(out, operate)
				break
			case FUNC_RATIO:
				stream = doZoomRatio(out, operate)
				break
			case FUNC_RATIO_FILL:
				stream = doZoomRatioFill(out, operate)
				break
			case FUNC_RATIO_EQUAL:
			default:
				stream = doZoomRatioEqual(out, operate)
			}
		} else {
			stream = string(getImageBuf(out)[:])
		}

		return 200, stream
	}

	return 404, ""
}

/**
等比缩放
 */
func doZoomRatioEqual(output *s3.GetObjectOutput, operate *zoomOperate) string {
	imagick.Initialize()
	defer imagick.Terminate()

	var err error
	mw := imagick.NewMagickWand()
	err = mw.ReadImageBlob(getImageBuf(output))
	if err != nil {
		panic(err)
	}

	srcWidth := mw.GetImageWidth()
	srcHeight := mw.GetImageHeight()
	hWidth := uint(operate.Width)
	hHeight := uint(operate.Height)

	var nWidth, nHeight uint
	if operate.ZoomFunc == FUNC_RATIO_FILL {
		if hWidth > hHeight {
			nWidth = uint(srcWidth * hHeight / srcHeight)
			nHeight = hHeight
		} else {
			nWidth = hWidth
			nHeight = uint(srcHeight * hWidth / srcWidth)
		}
	}

	err = mw.ResizeImage(nWidth, nHeight, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		panic(err)
	}

	err = mw.SetImageCompressionQuality(70)
	if err != nil {
		panic(err)
	}

	content := string(mw.GetImageBlob()[:])
	if err != nil {
		panic(err)
	}

	return content
}

/**
等比缩放，填充空白
 */
func doZoomRatioFill(output *s3.GetObjectOutput, operate *zoomOperate) string {
	imagick.Initialize()
	defer imagick.Terminate()

	var err error
	mw := imagick.NewMagickWand()
	err = mw.ReadImageBlob(getImageBuf(output))
	if err != nil {
		panic(err)
	}

	srcWidth := mw.GetImageWidth()
	srcHeight := mw.GetImageHeight()
	hWidth := uint(operate.Width)
	hHeight := uint(operate.Height)

	var nWidth, nHeight uint
	var startX, startY float64
	if operate.ZoomFunc == FUNC_RATIO_FILL {
		if hWidth > hHeight {
			nWidth = uint(srcWidth * hHeight / srcHeight)
			nHeight = hHeight
			startX = float64((hWidth - nWidth) / 2)
			startY = 0
		} else {
			nWidth = hWidth
			nHeight = uint(srcHeight * hWidth / srcWidth)
			startX = 0
			startY =  float64((hHeight - nHeight) / 2)
		}
	}

	err = mw.ResizeImage(nWidth, nHeight, imagick.FILTER_LANCZOS, 1)

	dw := imagick.NewDrawingWand()
	err = dw.Composite(imagick.COMPOSITE_OP_OVER, startX, startY,float64(nWidth), float64(nHeight), mw)
	if err != nil {
		panic(err)
	}
	newMw := imagick.NewMagickWand()
	newMw.DrawImage(dw)
	err = newMw.SetImageCompressionQuality(70)
	if err != nil {
		panic(err)
	}

	content := string(mw.GetImageBlob()[:])
	if err != nil {
		panic(err)
	}

	return content
}

/**
直接重置尺寸
 */
func doZoomRatio(output *s3.GetObjectOutput, operate *zoomOperate) string {
	imagick.Initialize()
	defer imagick.Terminate()

	var err error
	mw := imagick.NewMagickWand()
	err = mw.ReadImageBlob(getImageBuf(output))
	if err != nil {
		panic(err)
	}

	nWidth := uint(operate.Width)
	nHeight := uint(operate.Height)

	err = mw.ResizeImage(nWidth, nHeight, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		panic(err)
	}

	err = mw.SetImageCompressionQuality(70)
	if err != nil {
		panic(err)
	}

	content := string(mw.GetImageBlob()[:])
	if err != nil {
		panic(err)
	}

	return content
}

/**
直接裁剪缩放
 */
func doZoomCut(output *s3.GetObjectOutput, operate *zoomOperate) string {
	imagick.Initialize()
	defer imagick.Terminate()

	var err error
	mw := imagick.NewMagickWand()
	err = mw.ReadImageBlob(getImageBuf(output))
	if err != nil {
		panic(err)
	}

	nWidth := uint(operate.Width)
	nHeight := uint(operate.Height)

	err = mw.ResizeImage(nWidth, nHeight, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		panic(err)
	}

	err = mw.SetImageCompressionQuality(70)
	if err != nil {
		panic(err)
	}

	content := string(mw.GetImageBlob()[:])
	if err != nil {
		panic(err)
	}

	return content
}

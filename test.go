package main

import (
	"basicSortAlgorithm/files"
	"basicSortAlgorithm/image"
	"basicSortAlgorithm/awsS3"
)

func main()  {
	awsS3.Con = &files.Constant{Path:"/www/smart/app/stand/global/config/AwsS3Config.php"}
	awsS3.Con.ReadConstant()

	image.DoZoom("/dev-smart-app/599fe94303b1c1bf4b37436f/user/avatar/544/7/d/1/c7d1c96c5bcaa9829f5140f9e7206e2d8b63d0bd5_280x280.jpg")
}
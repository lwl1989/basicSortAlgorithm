package main

import (
	"basicSortAlgorithm/files"
	"fmt"
)

func main()  {
	con := &files.Constant{Path:"/www/smart/app/stand/global/config/AwsS3Config.php"}
	con.ReadConstant()
	fmt.Println(con.GetConstant("S3_REGION"))
}
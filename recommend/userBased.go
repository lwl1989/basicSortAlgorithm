package recommend

import "fmt"

//用户喜欢的物品集合
type UserLikeList struct {
	name string
	liked []string        //最大最近关注  最大长度50
}

//物品被用户喜欢的集合
type LikeUserList struct {
	name string
	liked []string			//最大最近关注  最大长度50
}

type VertexType rune // 顶点数值类型
type userLikedMatrix struct {
	UserVertex	[]string
	GoodsVertex []string  //顶点
	Vertex    [][]int	   //值
	goodsNum int
	userNum  int
}

func InitMatrix() {
	user := UserLikeList{
		"张三",
		[]string{"飞机","大炮"},
	}

	user1 := UserLikeList{
		"李四",
		[]string{"飞机","大炮222"},
	}


	matrix := &userLikedMatrix{
		UserVertex:[]string{},
		GoodsVertex:[]string{},
		Vertex:[][]int{},
		goodsNum:0,
		userNum:0,
	}
	matrix.AddUser(user)
	matrix.AddUser(user1)
	fmt.Println(matrix)
}
//所以要先排序
func (matrix *userLikedMatrix) AddUser(userLikeList UserLikeList)  {
	x := -1
	y := -1

	if matrix.userNum > 0 {
		for k, v := range matrix.UserVertex {
			fmt.Println(userLikeList.name, v)
			if userLikeList.name == v {
				x = k
				break
			}
		}
	}

	length := len(matrix.UserVertex)
	if x == -1 {
		x = length
		matrix.Vertex = append(matrix.Vertex, []int{0})
		matrix.userNum += 1
	}else{
		x -= 1
	}
	matrix.UserVertex = append(matrix.UserVertex, userLikeList.name)

	fmt.Println(y)
	for _, name := range userLikeList.liked {

		y = -1
		if matrix.goodsNum > 0 {
			for k, v := range matrix.GoodsVertex {
				if name == v {
					y = k
					break
				}
			}

			if y == -1 {

				for i:=len(matrix.Vertex[x]); i < matrix.goodsNum; i++ {
					matrix.Vertex[x] = append(matrix.Vertex[x], 0)
				}
				matrix.GoodsVertex = append(matrix.GoodsVertex, name)
				matrix.goodsNum += 1

				matrix.Vertex[x] = append(matrix.Vertex[x], 1)
			} else {
				fmt.Println(x,y,matrix.Vertex[x][y],name)
				matrix.Vertex[x][y] += 1
			}
		} else {
			matrix.goodsNum += 1
			matrix.GoodsVertex = append(matrix.GoodsVertex, name)
			matrix.Vertex[x][0] += 1
		}
	}

	//fmt.Println(matrix.Vertex)
	//fmt.Println(x,y)
	//if len(matrix.Vertex[x]) < y {
	//	matrix.Vertex[x] = append(matrix.Vertex[x], 0)
	//}
	//

}
//func (UserLikeList *UserLikeList) GetSimilar(otherList UserLikeList) []string {
//	list := make([]string, 50)
//	for k, v := range otherList.liked {
//		list[k] = v
//	}
//	return list
//}

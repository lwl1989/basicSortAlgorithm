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
	Vertex    [][]int32	   //值
	goodsNum int32
	userNum  int32
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
		Vertex:[][]int32{},
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
			if userLikeList.name == v {
				x = k
			} else {
				length := len(matrix.UserVertex)
				x = length
				matrix.UserVertex = append(matrix.UserVertex, v)
				matrix.userNum += 1
			}
		}
	} else {
		matrix.userNum = 1
		matrix.UserVertex = append(matrix.UserVertex, userLikeList.name)
	}

	if len(matrix.Vertex) < x || x == -1 {
		x = 0
		y = 0
		matrix.Vertex = append(matrix.Vertex, []int32{0})
	}
	if matrix.goodsNum > 0 {
		for k, v := range matrix.GoodsVertex {
			for _, name := range userLikeList.liked {
				if name == v {
					matrix.Vertex[x][k] += 1
				} else {
					if y == -1 {
						y = 0
					}
					length := len(matrix.GoodsVertex)
					y = length
					//if _,ok := matrix.Vertex[x][] {
					//
					//}

					matrix.GoodsVertex = append(matrix.GoodsVertex, v)
					matrix.goodsNum += 1
				}
			}
		}
	} else {
		matrix.goodsNum = 1
		matrix.GoodsVertex = userLikeList.liked
	}


	fmt.Println(matrix.Vertex)
	fmt.Println(x,y)
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

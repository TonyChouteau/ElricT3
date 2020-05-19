package https

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/TonyChouteau/elrict3/ai"
)

/*
Data : Struct
*/
/*type Data struct {
	Board     engine.Matrix3x3 `json:"board"`
	ListLegal []int            `json:"listLegal"`
	IsLegal   bool             `json:"isLegal"`
}*/

/*func test(c *gin.Context) {
	//countByType := storage.CountProjects(c)
	board := engine.CreateM()
	fmt.Println(board)

	board, _ = engine.Play(board, 5, CROSS)

	list := engine.ListLegal(board)
	fmt.Println(list)
	test1 := engine.IsLegal(board, 5)
	fmt.Println(test1)

	c.JSON(200, Data{
		board,
		list,
		test1,
	})
}*/

func moveAI(c *gin.Context) {
	board := c.Param("board")
	result := ai.GetBestMove(board)
	c.JSON(200, result)
}

/*
Serve function
*/
func Serve() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/ai/:board", moveAI)

	//err := http.ListenAndServe(":8082", r)
	err := http.ListenAndServeTLS(":8082", "/etc/letsencrypt/live/www.domain.com/fullchain.pem", "/etc/letsencrypt/live/www.domain.com/privkey.pem", r)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

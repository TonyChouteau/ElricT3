package https

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/TonyChouteau/elrict3/ai"
)

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

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://www.tonychouteau.fr", "https://www.thomaslepercq.fr/elrict3"},
		Methods: “GET, PUT, POST, DELETE”,
		ExposedHeaders: “”,
		Credentials: false,
		ValidateHeaders: false,
	}))
    

	r.GET("/ai/:board", moveAI)

	//err := http.ListenAndServe(":8082", r)
	err := http.ListenAndServeTLS(":8082", "/etc/letsencrypt/live/vps.tonychouteau.fr/fullchain.pem", "/etc/letsencrypt/live/vps.tonychouteau.fr/privkey.pem", r)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

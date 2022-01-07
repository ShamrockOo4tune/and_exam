package main
import "github.com/gin-gonic/gin"

const(hello_message = "Hello World 2 from golang webapp")
func main() {
    router := gin.Default()
    router.GET("/", func (c  *gin.Context) {c.String(200, hello_message)})
    router.Run(":8080")
}


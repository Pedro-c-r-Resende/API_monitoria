package API_monitoria

import (
	"loja-artesanato/routes"
	"net/http"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8080", nil)
}

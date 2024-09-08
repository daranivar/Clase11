package main

import "github.com/gin-gonic/gin"

//Generamos una estructura con los campos de la petición que recibiremos, para poder procesarla.
//Utilizaremos la etiqueta json, para especificarle cuáles serán los campos que recibiremos de la petición

type request struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

//Una variable de productos donde se guardarán los productos que enviemos
var products []request

//Una variable que guarde y vaya incrementando el ID, para siempre tomar el máximo
var lastID int

func main() {

	//Definiremos con Gin un servicio web mediante el método POST, el cual tendrá como path “productos”.
	r := gin.Default()
	//Creamos una agrupación para productos en el cual definiremos los diferentes endpoints. En nuestro caso, solo tendremos
	//el endpoint “Guardar”.
	pr := r.Group("/productos")
	pr.POST("/", Guardar())

	r.POST("/productos", func(c *gin.Context) {

		//Recibimos la petición y hacemos el traspaso de los datos a nuestra estructura con el método ShouldBindJSON (en
		//caso de no poder hacer el traspaso, nos devolverá error)
		var req request
		c.ShouldBindJSON(&req)
		//Tomamos el error que nos devuelve el bind y realizamos una validación.
		if err := c.ShouldBindJSON(&req); err != nil {
			//En caso de haber un error, lo retornamos. Utilizamos el método JSON para definir el código y el cuerpo del mensaje a
			//retornar.
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		//En caso de que la petición recibida sea correcta, agregamos un ID a nuestro producto.
		req.ID = 4
		//Enviamos el producto con el ID asignado como respuesta
		c.JSON(200, req)
	})

	r.Run()

}

//Función Guardar
func Guardar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		req.ID = 4
		c.JSON(200, req)
	}
}

//Headers en Go
//c.GetHeader("mi_header")

//En la función Guardar, lo primero que se hará es recibir el token que haya sido enviado en la petición:

//    func Guardar() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 	token := c.GetHeader("token")
// 	...
// 	}
// 	}

// 	//Validar token

// 	func Guardar() gin.HandlerFunc {
// 		return func(c *gin.Context) {
// 		token := c.GetHeader("token")
// 		if token != "123456" {
// 		c.JSON(401, gin.H{
// 		"error": "token inválido",
// 		})
// 		return
// 		}
// 		...
// 		}

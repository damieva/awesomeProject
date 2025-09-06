# NOTES:

- Carpeta cmd: puntos de entrada a nuestra aplicación
- Carpeta API con todo el codigo de la API.
- Carpeta internal: codigo interno entre ello las entidades, los servicios que constituyen la lógica de la aplicación, los puertos donde defino las interfaces para desacoplar accesos a componentes externos (como DBs) de mis servicios,
- Archivo .env para cargar todas las vbles del ambiente como por ejemplo las de conexion a DB
- Definir tipo player:
    * `json:"name"` Esto le indica a Gin (y al paquete estándar encoding/json) cómo se debe mapear este campo al hacer parseo (decodificación) o serialización (codificación) de JSON
    * `binding:"required"` Esto le dice a Gin que ese campo es obligatorio cuando haces un binding de JSON (usualmente en c.BindJSON(&player)). Si ese campo no viene en el JSON de entrada, Gin devolverá automáticamente un error 400 (Bad Request) indicando que faltan campos requeridos.
```
type Player struct {
	Name         string    `json:"name" binding:"required"`
	Age          int       `json:"age" binding:"required"`
	CreationTime time.Time `json:"creationTime"`
}
```
- Normalmente en el backend manejamos todos los CreationTime en UTC y en el Front traducimos a un tiempo concreto si es necesario.
- Antes de iniciar una operación que puede tardar (como conectarse a MongoDB, hacer una petición HTTP o lanzar una goroutine), creamos un contexto con un límite de tiempo:
```ctx, cancel := context.WithTimeout(context.Background(), time.Second)```
  Esto asegura que, si la operación no finaliza en 1 segundo, se cancele automáticamente.
- Usamos ```defer cancel()``` para garantizar que, al terminar la función, se liberen los recursos internos asociados al contexto (como timers). De este modo evitamos fugas de memoria.
- Problemas de tener todo el codigo en main.go:
  * Estamos abriendo una conexión nueva a la db cada vez q hay un request al endpoint de /players, la forma correcta de hacerlo sería reutilizando conexiones y no sobrecargar el uso de recursos.
- Un método con receptor (receiver method) en Go es una función asociada a un tipo (por ejemplo, un struct). El receptor se declara entre la palabra clave func y el nombre del método y determina sobre qué tipo se puede invocar ese método.
   * ```func (h Handler) CreatePlayer(c *gin.Context)```
   * Receptor por valor ((t Tipo)) → recibe una copia del valor, no modifica el original.
   * Receptor por puntero ((t *Tipo)) → recibe una referencia, puede modificar el valor original.
   * Esto permite dotar a los tipos de comportamiento propio, facilitando un estilo más orientado a objetos.
- En Go, un struct (o cualquier tipo) implementa una interfaz automáticamente cuando define todos los métodos de esa interfaz como métodos con receptor.
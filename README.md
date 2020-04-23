# Exercise #4: HTML Link Parser

[![exercise status: released](https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge)](https://gophercises.com/exercises/link)

## Detalles del ejercicio

En este ejercicio, su objetivo es crear un paquete que facilite analizar un archivo HTML y extraer todos los enlaces (`<a href="">...</a>` tags). Para cada enlace extraído, debe devolver una estructura de datos que incluya tanto el `href` como el texto dentro del enlace. Cualquier HTML dentro del enlace se puede eliminar, junto con cualquier espacio en blanco adicional, incluidas las nuevas líneas, los espacios consecutivos, etc.

Los enlaces estarán anidados en diferentes elementos HTML, y es muy posible que tenga que lidiar con HTML similar al código a continuación.

```html
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
```

En situaciones como estas, queremos obtener resultados que se vean más o menos así:

```go
Link{
  Href: "/dog",
  Text: "Something in a span Text not in a span Bold text!",
}
```

Una vez que tenga un programa en funcionamiento, intente escribir algunas pruebas para que practique el uso del paquete de prueba en marcha.


### Notas

**1. Use el paquete x/net/html **

Recomiendo revisar el  [x/net/html](https://godoc.org/golang.org/x/net/html) paquete para esta tarea, que necesitará 'go get'. Es proporcionado por el equipo de Go, pero no está incluido en la biblioteca estándar. Esto hace que sea un poco más fácil analizar archivos HTML.


**2. Ignore los enlaces anidados**

Puede ignorar cualquier enlace anidado dentro de otro enlace. Por ejemplo, con el siguiente HTML:

```html
<a href="#">
  Something here <a href="/dog">nested dog link</a>
</a>
```

Está bien si su código devuelve solo el enlace externo.

**3. hacer que algo funcione antes de enfocarse en casos extremos**

No te preocupes por tener un código perfecto. Es probable que haya muchos casos extremos aquí que serán un poco difíciles de manejar. Solo trate de cubrir los casos de uso más básicos primero y luego mejore eso.

**4. Se han proporcionado algunos ejemplos HTML.**

Creé algunos archivos HTML más simples y los incluí en este repositorio para ayudar con las pruebas. No cubrirán todos los casos de uso potenciales, pero deberían ayudarlo a comenzar a probar su código.


**5. El cuarto ejemplo lo ayudará a eliminar comentarios de su texto de enlace**

Es probable que su primera versión incluya el texto de los comentarios dentro de una etiqueta de enlace. El mío lo hizo. Use [ex4.html](ex4.html) para probar ese caso y corregir el error.
*Hint: See [NodeType](https://godoc.org/golang.org/x/net/html#NodeType) constants and look for the types that you can ignore.*


## Recursos externos

En la solución para este ejercicio termino usando un DFS, que es un algoritmo de teoría de grafos. Si quieres aprender un poco más sobre eso, lo he discutido en YouTube aquí:- <https://www.youtube.com/watch?v=zboCGDMnU3I>
Existe una serie completa sobre algoritmos y teoría de grafos, aunque en este momento es algo incompleta. Nunca tengo suficiente tiempo en el día 🙁. Ojalá un día y *Aprendamos Algoritmos* será su propia serie como *Gophercises*.

## Bonus

Los únicos bonos aquí son mejorar sus pruebas y la cobertura de casos límite.
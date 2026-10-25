# EvaDock: tres poses independientes — diseño

## Intención

Recuperar la mascota flotante de Eva para Windows con tres poses aprobadas, conservando el mecanismo de instalación y actualización exclusivamente desde el repositorio EvaDock.

## Restricciones acordadas

- Cada pose es una imagen PNG independiente; nunca se recorta una hoja colectiva.
- La identidad visual procede de las fotografías reales facilitadas el 25 de octubre de 2026. La pose sentada aprobada es la referencia común: pelaje moteado irregular café/gris/negro/ocre, pecho y patas blancos, ojos amarillo verdoso y cola estrecha oscura. Evitar las franjas artificiales y las colas de plumero.
- La silueta real era baja, robusta y con el vientre derecho visiblemente abultado. No generar ni integrar poses de pie, caminando o salto hasta que se proporcionen más fotografías de referencia anatómica.
- Ninguna oreja, pata ni cola debe quedar cortada por el archivo o el dibujo de la ventana.
- El ejecutable se conserva como aplicación autónoma de Windows; las actualizaciones siguen el canal `youteachtk/EvaDock`, con nombre de instalador fijo.
- No modificar otras aplicaciones; no usar generadores de video ni servicios de terceros en tiempo de ejecución.
- No publicar automáticamente a la rama principal sin validación de la presentación visual.

## Enfoques

1. **Tres PNG independientes embebidos (elegido):** conserva la arquitectura Go/Win32 y su motor de transparencia. Aumenta el tamaño del ejecutable, pero la mascota funciona sin instalar recursos externos.
2. Hoja de sprites: más compacta, pero reintroduce exactamente los recortes rechazados.
3. Animación en video: no corresponde al diseño actual y exige más recursos.

## Comportamiento

La interfaz permite seleccionar Sentada, Recostada y Dormida. El botón de rutina automática alterna únicamente esas tres poses. La vista previa y la ventana flotante muestran la misma pose. En **Recostada** (ojos cerrados) solamente la punta de la cola se mueve discretamente y de forma ocasional; la cabeza, los ojos, el torso y las patas permanecen quietos. Las demás poses permanecen explícitamente suspendidas para evitar deformaciones anatómicas.

## Validación y límites

Pruebas automatizadas verifican que se carguen exactamente tres PNG válidos, que cada uno mantenga transparencia real y que sus contenidos no sean recortes de una imagen compartida. La compilación cruzada Windows y la ejecución de pruebas de Go deben pasar. La apariencia final en un equipo Windows requiere verificación manual antes de publicar `0.11.0`.

# EvaDock — decisiones permanentes sobre Eva

Actualizado: 2026-10-26. Se conservan los cambios y antecedentes en GitHub; consultar este documento y el historial antes de pedir material otra vez.

## Referencias y apariencia real

- Eva es una gata de tamaño bajo y constitución robusta: **no alta ni delgada**.
- Por una antigua lesión en las costillas del lado derecho, su vientre se veía más abultado. Respetar esa silueta.
- Pelaje **carey/calico de manchas irregulares** marrón, gris oscuro y ocre; evitar las franjas repetidas de un atigrado típico.
- Pecho y partes de las patitas blancos; ojos verdosos.
- Cola de longitud y grosor contenidos, **no tipo plumero**. Se revisó una foto específica de la cola; se indicó ignorar la zona gris que pertenecía al perro.
- No inventar una quinta pata, ni cortar orejas, patas o cola por el borde del recurso.
- Cada pose debe proceder de **su propio PNG transparente completo**; jamás extraerla como recorte de una lámina con varias poses.

## Historial de elección de imágenes

- **DESCARTADA**: pose sentada de EvaDock 0.10.0 y 0.10.1. Objeto Git `1231ff9e780b470ce9520ee7088657c91bbd7a12` (680051 bytes). Se conserva sólo por el historial de Git; no debe volver a la aplicación publicada.
- **ELEGIDA como segunda alternativa generada**: `Eva, gata carey de pecho blanco.png`, generada el 2026-10-25. Coincidencia binaria comprobada con `assets/independientes/sentada.png` en la rama de trabajo. SHA-256 `02b5c462cd4a049f5e9a1e8fdd98da4746a994e81f195435df1252f3339fefb8`; objeto Git `66c18844ed3255224126894401cbcb5052d69710` (2141490 bytes, PNG RGBA 1254×1254). **Recurso de la versión publicada 0.10.2.**
- La otra propuesta `Gato atigrado con cola fina.png` también se recuperó, pero no es la alternativa elegida. No cambiar las imágenes aprobadas sin indicación expresa.
- Las imágenes independientes `recostada.png` (objeto Git `3dfa4ea4ac6e3bf882001dd7754c994842f66aeb`) y `dormida.png` (objeto Git `0769cdb7917bcdd819f59ad4799b7b5ff77723e9`) están en la rama `eva-poses-independientes-v011`. **No dar por aprobadas todas las poses ni publicarlas sin validación.**
- **APROBADA — pose horizontal**: `Eva, gata calicó en pose horizontal.png`, generada tras pedir rotar el cuerpo completo hasta dejarlo horizontal **sin voltear la cabeza**. El archivo se recuperó y comparó byte por byte: PNG RGBA 1536×1024, SHA-256 `72391725920009d6234574fbbbcde9a37fc38fb79b4cbd6d03af51a5a503d044`; objeto Git `521f49bf43c67320773ba488883cf7e81d102115`. Integrado en `assets/independientes/pose_horizontal_aprobada.png` para la versión 0.10.3. No recortar, regenerar, estirar ni girar la cabeza.
- **APROBADA — pose acurrucada**: `Eva, gata calicó acurrucada.png`, solicitada como la primera de las dos poses compartidas antes de corregir la orientación de la segunda. PNG RGBA 1677×938; SHA-256 `083264e6997024c5c9837a17ea08af102f1503116dd5cb181abb24881f5e25b4`; archivo `assets/independientes/pose_acurrucada_aprobada.png` para la versión 0.10.4. Mantener el PNG individual sin cambios.

## Poses y movimientos

- Mantener sólo **sentada**, **semiacostada/recostada** y **dormida** como posibilidades próximas.
- Para la imagen semiacostada, con los ojos cerrados, mantener quietos cabeza, ojos y cuerpo; la cola debe moverse de manera breve y ocasional.
- Se suspenden **de pie**, **caminando** y **saltando** hasta que se disponga de mejores fotografías. No tratar de compensar inventando anatomía o estrechando el cuerpo.
- **ESTADO VERIFICADO EL 26-10-2026 — NO APROBADO:** la implementación técnica 0.10.5 se publicó, pero el usuario **rechazó** los efectos de respiración, cola, deformación de PNG y el movimiento artificial de una imagen estática. **Ninguna secuencia de animación está aprobada todavía.** El requisito sigue siendo movimiento anatómico y realista de la gata, comparable a la experiencia FureverDock. Ver [ESTADO_ACTUAL.md](ESTADO_ACTUAL.md) e [HISTORIAL_ANIMACIONES.md](HISTORIAL_ANIMACIONES.md) antes de retomar cualquier prueba.
- **Implementación 0.10.5:** autonomía de cambios de postura pausados (sentada → horizontal → acurrucada → horizontal), microrespiración en sentada/acurrucada y una deformación mínima, ocasional y localizada en la punta de la cola visible de la pose horizontal. No se alteran los PNG fuente. El clic sobre Eva equivale a una caricia y la devuelve a sentada. La selección manual pausa la autonomía; botón «Rutina» la reanuda. Se conserva la inmovilidad del cuerpo durante el descanso horizontal y se elimina el antiguo desplazamiento artificial de «caminar».
- **Pendiente:** animación de parpadeo realista sin inventar rasgos, recordatorios y conversación local sin pagos, así como validación del aspecto del efecto de cola en el escritorio Windows. No presentar 0.10.5 como un clon funcionalmente completo de Furever Dock.

## Instalación y conservación

- **Herramienta descartada:** Kling AI (imagen a video) ya fue probado para EvaDock y el usuario confirmó que no sirve para el resultado requerido. No volver a proponer Kling como solución ni repetir la prueba.

- **Restricción confirmada el 2026-10-26:** CreativeClaw tiene los créditos agotados. No proponer ni usar CreativeClaw para animaciones de EvaDock hasta que el usuario confirme que tiene saldo y autorice su uso. Tampoco utilizar HeyGen en EvaDock.

- Una sola instalación permanente, `%LOCALAPPDATA%\EvaDock\EvaDock.exe`, con actualizaciones desde los Releases de `youteachtk/EvaDock`. No exigir descargar un nuevo .exe por cada versión.
- La imagen de la versión 0.10.2 se publica usando su objeto Git existente, no una imagen distinta regenerada.
- Conservar la ventana transparente que ya funciona y una forma de cerrar limpiamente desde la bandeja/programa.
- Toda decisión posterior debe anotarse aquí y sus archivos aprobados conservarse en el repositorio **antes** de dar por terminada una entrega. No solicitar de nuevo fotos ya compartidas mientras sigan recuperables.

## Materiales anteriores localizados

Se hallaron las fotografías originales compartidas en la sesión: `20250112_165601.jpg`, `20250112_165556.jpg`, `20241103_112810.jpg` y `20241123_003553.jpg`, además del video `20241123_003538.mp4` citado en la conversación. Los originales se conservan fuera del repositorio público por privacidad.

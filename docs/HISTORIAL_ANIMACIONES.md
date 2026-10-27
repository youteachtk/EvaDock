# EvaDock — historial de intentos y decisiones sobre animación

Corte: 27 de octubre de 2026. Este registro distingue lo que efectivamente se produjo, lo que sólo se investigó y lo que fue aprobado o rechazado. **Cero secuencias de animación aceptadas a esta fecha.**

## Línea temporal recuperada

| Fecha / fase | Trabajo o recurso | Resultado | Estado |
| --- | --- | --- | --- |
| 25 oct 2026, inicio | Referencia de FureverDock: mascota personalizada con comportamiento y movimientos naturales, interacción y descanso | El usuario solicitó replicar la experiencia de FureverDock con Eva y conservar sus características reales | Objetivo vigente |
| 25 oct | Versiones 0.10.0/0.10.1, pose sentada anterior | Pelaje/cola/silueta no fieles; reemplazo solicitado | Rechazada |
| 25 oct | Segunda imagen de Eva sentada, «Eva, gata carey de pecho blanco.png», v0.10.2 | Elegida y registrada mediante hash. Transparencia de ventana corregida previamente | **Pose estática aprobada** |
| 25/26 oct | Imagen horizontal RGBA 1536×1024, cuerpo recto y cabeza sin girar | Aprobada tras corregir orientación; no modificar cabeza | **Pose estática aprobada** |
| 25/26 oct | Imagen acurrucada individual 1677×938 | Aprobada como pose independiente | **Pose estática aprobada** |
| 25/26 oct | 0.10.4 con tres poses y ventana transparente | Release de GitHub, selector de poses; no equivale a animación realista | Publicado |
| 26 oct | 0.10.5 con cambios automáticos de pose, expansión respiratoria, cola ocasional, clic-caricia | Release de GitHub, pero los efectos que mueven una fotografía no satisfacen el requisito. No confundir con aprobación | **Publicado, animación NO aprobada** |
| 26 oct | Imágenes GIF tempranas de sentada/acurrucada/horizontal con movimiento mínimo | Ensayos mecánicos de escala/cola, no clips de actuación | No aprobadas |
| 26 oct | Ensayos de respiración APNG/Synfig/Python: abdomen («panza»), espalda, pecho, hombro, mayor amplitud | El usuario rechazó cada intento: se mueve la parte equivocada del cuerpo o parece deformación de foto | **Rechazados; prohibido repetir método** |
| 26 oct | «Eva_PixVerse_prueba_2_NO_APROBADA.mp4», ~3 s, 1024×688, 24 fps | Prueba generativa que no logró el resultado exigido; muestra marca de PixVerse | Rechazada |
| 26 oct | Visor privado web de respiración (APNG de prueba, versión v4) | Permitió revisar sin descargar, pero la animación visible movía el hombro, no mostraba respiración convincente | Rechazada; no integrar |
| 26 oct | Investigación: Live2D, Rive, Blender, Runway, Luma | Sólo exploración técnica, SIN aprobación ni ejecución de un nuevo pipeline de animación | Investigadas, no implementadas |
| 26 oct | CreativeClaw / Seedance 2.5 | Propuesto inicialmente, usuario informó que CreativeClaw no tiene créditos | **No disponible; no utilizar** |
| 26 oct | HeyGen | El usuario lo prohibió explícitamente para EvaDock tras experiencias previas | **No utilizar** |
| 26 oct | Kling AI image-to-video | Propuesta descartada porque el usuario indicó que **ya lo habían probado y no funcionó** | **No usar ni repetir prueba** |

## Archivos de ensayos NO aprobados (copia temporal examinada)

Los nombres siguientes identifican pruebas, no recursos válidos para la app final. Pueden faltar en una futura sesión porque no todos pertenecen a la rama principal de GitHub:

- Eva_Keyframe_prueba_NO_APROBADA.apng y Eva_Keyframe_transparente_PRUEBA_NO_APROBADA.apng.
- Eva_Respiracion_ANTERIOR_NO_APROBADA_1536x1024.apng.
- Eva_Respiracion_ANTERIOR_PANZA_NO_APROBADA_1536x1024.apng.
- Eva_Respiracion_ESPALDA_NO_APROBADA_1536x1024.apng.
- Eva_Respiracion_PECHO_1536x1024.apng y Eva_Respiracion_MAS_VISIBLE_1536x1024.apng: los nombres no significan que fueran aceptadas.
- Eva_Respiracion_Synfig_PRUEBA_1536x1024.apng, Eva_Synfig_Respiracion_SOLO_PRUEBA.sif y Eva_Synfig_Respiracion_MAS_VISIBLE_SOLO_PRUEBA.sif.
- Eva_Calibracion_Directa.sif y Eva_Calibracion_030.sif.
- Eva_PixVerse_prueba_2_NO_APROBADA.mp4.
- EvaDock-preview/Eva_sentada_respira.gif, Eva_acurrucada_respira.gif y Eva_horizontal_cola.gif.
- En el visor independiente: scripts/generate_chest_breathing.py, scripts/pack_apng.py y dist/assets/eva-animada-v4.png. **Todos los resultados siguen NO aprobados.**

No convertir estos materiales en archivos oficiales por el mero hecho de encontrarlos. No reconstruirlos en otro editor de deformación cambiando únicamente las coordenadas.

## Intentos de aplicación y problemas conocidos

- Hubo una fase 0.11 experimental reportada como guardada en GitHub, **no instalada**; el usuario informó que aparecía icono en bandeja sin ventana de Eva o sin mecanismo de cierre. Revalidar la rama antes de reutilizarla; no afirmar que el error quedó arreglado.
- La versión publicada 0.10.5 incluye código para alternar tres poses y deformación mínima. La copia local posterior examinada deshabilita comportamientos automáticos rechazados, pero sus archivos no estaban sincronizados con GitHub main. No sobrescribir cambios sin conciliación.
- Varias veces el usuario pidió que se mostraran todas las imágenes/acciones para aprobación antes de subirlas o integrarlas. **Cualquier «aprobado» requiere confirmación explícita y referida al recurso concreto.**


## 26-10-2026 — exploración adicional, sin aprobación ni publicación

- Se intentó **animación fotográfica por poses nuevas**: dos fotogramas RGBA separados de Eva sentada con cabeza girada aproximadamente 12° y 21° respecto a la pose sentada aprobada. No se reutilizó el antiguo efecto de respiración, y no se modificaron los PNG aprobados del repositorio.
- Para evaluar continuidad se montó temporalmente `preview-head-turn.mp4` mediante interpolación óptica de fotogramas (FFmpeg). La inspección de un fotograma intermedio mostró **difuminación/doble rasgo cerca de los ojos y deriva de pelaje**. **Resultado técnico insuficiente; NO aprobado por el usuario, NO integrado a EvaDock, NO publicado.** La prueba permanece sólo en espacio de trabajo efímero, no como material oficial.
- Se comprobó que el video original de Eva `20241123_003538.mp4`, citado en el historial, **no está disponible** ni en el espacio de trabajo actual ni en la búsqueda de archivos recuperables. No afirmar que puede usarse hasta tener los bytes verificables. Usar video REAL del animal y extraer sus cuadros con máscara alfa es una vía técnicamente superior a simular respiración de una sola foto, pero necesita el video fuente.
- Se investigó **Runway Gen-4.5 Image-to-Video** como vía alternativa de video real con movimiento completo a partir de la pose aprobada: documentación oficial https://help.runwayml.com/hc/en-us/articles/46974685288467-Creating-with-Gen-4-5. En una continuación se llegó a la pantalla de autenticación, pero **no hay cuenta Runway** para iniciar sesión; no se creó cuenta, no se subieron recursos y no se gastaron créditos. Queda descartado por ahora y no debe intentarse de nuevo sin cambio explícito de circunstancias.

## 27-10-2026 — mantenimiento de aplicación, sin cambios animados

- Prioridad expresa: Eva debe mantenerse visible por encima de otras ventanas en Windows; actualizar el programa actual y dejar las futuras animaciones para después.
- Versión objetivo: `0.10.6`, recuperación periódica de `HWND_TOPMOST` sin robar el foco, corrección de rutas de cierre que podían dejar la mascota oculta y el panel inaccesible.
- **Ninguna imagen fuente o movimiento fue editado o aprobado por esta intervención.** Las pruebas de animación rechazadas se mantienen registradas y pendientes de nueva solución realista.

## Regla de prevención

Al recibir «¿qué sigue?» o «continúa», leer [ESTADO_ACTUAL.md](ESTADO_ACTUAL.md) y este historial, verificar GitHub, identificar el último recurso **aprobado** y la última prueba **rechazada**. Nunca iniciar otra iteración de los métodos anteriores ni volver a preguntar por requisitos o archivos ya recuperables. Las futuras pruebas se agregan aquí, con fecha, herramienta, insumos, resultado, ubicación y decisión.

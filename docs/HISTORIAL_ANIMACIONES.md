# EvaDock — historial de intentos y decisiones sobre animación

Corte: 26 de octubre de 2026. Este registro distingue lo que efectivamente se produjo, lo que sólo se investigó y lo que fue aprobado o rechazado. **Cero secuencias de animación aceptadas a esta fecha.**

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

## Regla de prevención

Al recibir «¿qué sigue?» o «continúa», leer [ESTADO_ACTUAL.md](ESTADO_ACTUAL.md) y este historial, verificar GitHub, identificar el último recurso **aprobado** y la última prueba **rechazada**. Nunca iniciar otra iteración de los métodos anteriores ni volver a preguntar por requisitos o archivos ya recuperables. Las futuras pruebas se agregan aquí, con fecha, herramienta, insumos, resultado, ubicación y decisión.

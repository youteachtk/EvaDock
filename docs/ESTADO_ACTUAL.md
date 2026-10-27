# EvaDock — estado real y continuidad

**Corte:** 27 de octubre de 2026. Este documento se actualiza en cada intervención. Leer también [Historial de animaciones](HISTORIAL_ANIMACIONES.md), [Decisiones de Eva](DECISIONES_EVA.md) y [AGENTS.md](../AGENTS.md).

## Objetivo que NO debe reinterpretarse

Replicar para Eva la **experiencia de una mascota de escritorio realista y animada**, como FureverDock: movimientos naturales de un animal reconocible, acciones completas, pausas, descanso, interacción y transiciones creíbles. No basta con mostrar fotos y alternarlas. **Mover, inflar o retorcer una fotografía completa NO es una animación aceptable.** La identidad, pelaje, complexión, proporciones y poses reales de Eva tienen prioridad sobre cualquier efecto.

El usuario quiere resultados visibles y comprobables. Investigar es responsabilidad del asistente; sólo se consulta al usuario si se requiere autorización para IMPLEMENTAR o USAR un servicio/herramienta, enviar recursos a terceros, instalar software o gastar créditos. No subir ni integrar recursos antes de supervisión cuando ésta fue solicitada.

## Estado técnico verificado en GitHub

### Intervención del 27-10-2026 — visibilidad prioritaria

- El usuario solicitó **corregir de inmediato la aplicación actual** para que Eva permanezca encima de las demás ventanas y continuar las animaciones después. No modificar poses, recursos ni efectos rechazados en este cambio.
- Diagnóstico sobre `main.go` de v0.10.5: ya existía `WS_EX_TOPMOST` y una llamada inicial a `SetWindowPos(HWND_TOPMOST)`, pero no se renovaba la prioridad después; `petProc/WM_CLOSE` ejecutaba `showPet(false)` mientras el panel de controles permanecía oculto, dejando a Eva inaccesible. También se podía ocultar a Eva manualmente y cerrar el único panel recuperador.
- Corrección propuesta para 0.10.6: recuperar `HWND_TOPMOST` cada 2 segundos, sin tomar el foco; reactivar al mostrar/arrancar; al cerrar Eva abrir controles en vez de esconderla; no permitir cerrar el panel si la mascota está oculta sin antes recuperarla. Pruebas puras de política de visibilidad en `internal/windowvisibility/`.
- Confirmar en GitHub Actions que compila el EXE y existe release, y **verificar en Windows** la permanencia real por encima de otras ventanas, actualización automática y transparencia. No afirmar prueba real en Windows mientras no ocurra.
- No existe cuenta Runway disponible para iniciar sesión. La exploración web de autenticación no creó cuenta, no subió fotos ni generó clips; **suspender Runway** hasta autorización y cuenta real. CreativeClaw agotado, Kling inservible según prueba previa, HeyGen prohibido.

- Repositorio oficial: https://github.com/youteachtk/EvaDock (público).
- Rama de producción: main. Rama histórica de poses: eva-poses-independientes-v011. Consultar GitHub antes de suponer que existe otra rama.
- **Última release comprobada:** [Eva Dock 0.10.5](https://github.com/youteachtk/EvaDock/releases/tag/eva-dock-v0.10.5), con EvaDock.exe y EvaDock.sha256, publicada el 26 de octubre de 2026. version.txt en main indica 0.10.5.
- En main existen la ventana Win32, capa transparente, controles y actualizador que usa GitHub Releases. La transparencia fue confirmada funcional por el usuario en una fase anterior.
- El código remoto de 0.10.5 incluye rutina de alternar poses (sentada, horizontal, acurrucada), un efecto de respiración por expansión y cola ocasional. **Es una implementación PUBLICADA, pero NO una animación realista APROBADA.** No describirla como el objetivo cumplido.
- La ruta de instalación permanente acordada es %LOCALAPPDATA%\\EvaDock\\EvaDock.exe; un ejecutable de nombre fijo y actualizaciones desde GitHub. No producir instaladores sucesivos para cada versión.
- No se ha verificado en este corte el estado exacto del ejecutable instalado en la computadora del usuario; no afirmar que la actualización se instaló ni que una animación funciona allí.
- La copia local de trabajo examinada en /workspace/scratch/49f75133eb95/EvaDock-work estaba basada en el commit antiguo 3307ed8 (0.10.2) y tenía numerosos cambios SIN COMMIT: main.go, README.md, versión, workflow, poses y paquetes internal. En particular su behavior.go local desactiva efectos rechazados, a diferencia de main remoto. **NO sobrescribir main desde esa copia atrasada, NO confundir esas modificaciones con un release, NO borrar el trabajo sin inspeccionarlo.**

## Recursos de identidad aprobados, sin alteraciones

1. **Sentada:** assets/independientes/sentada.png. PNG RGBA 1254×1254. SHA-256 02b5c462cd4a049f5e9a1e8fdd98da4746a994e81f195435df1252f3339fefb8.
2. **Horizontal boca arriba:** assets/independientes/pose_horizontal_aprobada.png. PNG RGBA 1536×1024. SHA-256 72391725920009d6234574fbbbcde9a37fc38fb79b4cbd6d03af51a5a503d044. Cuerpo horizontal, cabeza de espaldas; NO girar la cabeza.
3. **Acurrucada:** assets/independientes/pose_acurrucada_aprobada.png. PNG RGBA 1677×938. SHA-256 083264e6997024c5c9837a17ea08af102f1503116dd5cb181abb24881f5e25b4.

Son **poses fuente** aprobadas, no secuencias de animación aprobadas. Los PNG se recuperaron de trabajos previos y sus huellas se verificaron localmente. La pose sentada antigua (v0.10.0/0.10.1) fue rechazada. Otras propuestas de 0.11 siguen sin aprobación.

Características de Eva: gata calicó/carey de cuerpo bajo y robusto, pecho y patas parcialmente blancos, ojos verdes, manchas irregulares marrón/oscuro/ocre, cola contenida no esponjosa. Conservar silueta original, incluida la asimetría abdominal mencionada por el usuario. Cada pose debe existir como archivo independiente; no recortes de collage.

## Restricciones firmes y pruebas rechazadas

- Ninguna animación de respiración producida hasta ahora está aprobada. Se intentó mover abdomen, espalda, hombro y pecho, y cada resultado resultó anatómicamente incorrecto o artificial.
- APNG, estiramientos de píxeles, Synfig y deformadores esféricos fueron pruebas fallidas; NO repetir variando el centro o la amplitud de la misma deformación.
- PixVerse produjo un MP4 de prueba **NO APROBADO** (también con marca visible en la muestra local). No presentarlo como éxito ni repetirlo automáticamente.
- Kling AI **YA SE PROBÓ y el usuario dijo que no sirve**. No volver a proponerlo ni gastar tiempo en otra prueba.
- CreativeClaw **está agotado**. No volver a ofrecerlo mientras el usuario no confirme nuevos créditos y autorice su uso.
- HeyGen **prohibido** en este proyecto. No usarlo.
- Investigar Blender/rigging, Runway o Luma NO equivale a autorización para instalar, suscribirse, gastar créditos o subir una imagen. Su utilidad para mantener la identidad de Eva NO está validada.
- **Runway no dispone de cuenta** para este proyecto: no volver a intentar iniciar sesión como si existiera. No utilizarlo para animaciones por ahora.
- La página independiente de revisión de respiración fue una prueba, NO parte de EvaDock ni una aprobación. No imponer descargas de APNG/ZIP/EXE para evaluar una animación.

## Interfaz y funcionamiento que se conservan

- Ventana transparente independiente en Windows, sin bordes/rectángulo opaco, con menú para cerrar.
- Una sola instalación y actualizaciones automáticas vía GitHub; nunca obligar a descargar un EXE nuevo por cada cambio.
- Mostrar a Eva sin recortar orejas, patas o cola; no estirar las imágenes.
- No inventar caminar, saltar, estar de pie ni una quinta pata para llenar una animación.
- Descanso horizontal: no girar cabeza para mostrar cara a cámara. Si se emplea una cola ocasional u otra acción, requerirá revisión de movimiento real.

## Próximo trabajo — orden de ejecución

**Primero cerrar y comprobar el arreglo de visibilidad de 0.10.6. Después retomar las animaciones.**

1. **No generar más deformaciones de PNG.** Tomar como referencia el material aprobado y la calidad de acciones reales de FureverDock. Analizar antes qué pipeline puede entregar un personaje que realmente actúe, con identidad consistente.
2. Identificar una vía viable para obtener animaciones articuladas/secuencias visuales nuevas sin CreativeClaw, HeyGen ni Kling. Comparar factibilidad, calidad y costes SIN suscripciones, instalaciones o cargas externas no autorizadas.
3. Para usar/instalar una herramienta distinta, subir la imagen de Eva a un tercero o incurrir en gasto, pedir **sólo esa autorización concreta**; no volver a preguntar requisitos ya documentados.
4. Preparar **una** prueba separada de Eva sentada con una acción claramente visible y anatómicamente creíble (por ejemplo, parpadear, desplazar levemente la cabeza y levantar/apoyar una pata); no basta un microbombeo en el pecho. Preservar el patrón de pelaje y el número de extremidades.
5. Mostrar la animación completa directamente en una página de revisión, sin instalación/descarga obligatoria. El usuario decide si aprueba. Si rechaza, registrar exactamente por qué, sin gastar en nuevas tomas sin consentimiento.
6. **Sólo después de aprobar el movimiento**, implementar reproducción real de fotogramas/acción en EvaDock, conservando la ventana transparente y el actualizador. Verificar en Windows, probar cierre y actualización. No declarar finalizada la etapa antes de esas pruebas.
7. Actualizar DECISIONES_EVA.md, HISTORIAL_ANIMACIONES.md, este estado y README, enlazar commit y release si la hubo.

## Fuente de verdad / criterios

- **GitHub main y Releases**: lo efectivamente publicado; no implica aprobación visual.
- **DECISIONES_EVA.md**: detalles estables, fotos/poses aprobadas, identidad.
- **HISTORIAL_ANIMACIONES.md**: cada ensayo y rechazo, sin volver a insistir.
- **Este documento**: dónde estamos y qué sigue.
- **AGENTS.md**: rutina obligatoria de lectura para futuras sesiones.
- Fotografías/videos originales y material privado: conservar fuera del repositorio público. NO pedirlos otra vez si pueden recuperarse; NO afirmar que un archivo privado sigue disponible sin verificarlo.

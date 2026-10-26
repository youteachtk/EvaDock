# EvaDock

Mascota independiente de Eva para el escritorio de Windows, con ventana flotante transparente y actualizaciones desde este repositorio.

## Versión 0.10.5 — primeros movimientos y rutina automática

- **Rutina tranquila activada al iniciar:** Eva alterna las tres poses aprobadas durante periodos de aproximadamente 60–120 segundos. No se desliza por el escritorio simulando caminar ni usa dibujos no aprobados.
- **Respiración muy discreta:** las poses sentada y acurrucada tienen un movimiento de expansión inferior al 1 %, manteniendo el PNG original intacto.
- **Cola de la pose horizontal:** movimiento local corto y ocasional, aproximadamente cada 40–65 segundos; el cuerpo, las patas y la cabeza permanecen quietos. El efecto se realiza al dibujar la imagen, sin cambiar el archivo aprobado.
- **Interacción:** arrastrar a Eva la mueve libremente; un clic sin arrastrar equivale a una caricia y la despierta a la pose sentada. Clic derecho abre los controles. La selección manual de una pose pausa la rutina hasta reactivarla con el botón correspondiente.
- **Escritorio despejado:** el panel de controles inicia oculto; Eva permanece como ventana transparente e independiente. Cerrar el panel sólo lo oculta y «Cerrar Eva Dock» finaliza el programa.
- **Actualizador permanente:** se publica de nuevo como `EvaDock.exe` y mantiene la instalación en `%LOCALAPPDATA%\EvaDock\EvaDock.exe`, sin instaladores sucesivos.

Esta versión **todavía no incorpora** parpadeo anatómico (requiere nuevas imágenes de referencia), conversación local ni recordatorios. Las poses y los movimientos anteriores están documentados a continuación.

## Versión 0.10.4 — tres poses aprobadas de Eva

- **Imagen sentada:** se sustituyó el archivo antiguo de 0.10.0/0.10.1 (cola excesivamente esponjada, expresión y proporciones rechazadas) por la imagen individual elegida el 25 de octubre: `Eva, gata carey de pecho blanco.png`. Está integrada exactamente, sin recortes ni regeneración, como `assets/independientes/sentada.png`.
- **Imagen horizontal:** se incorpora exactamente el PNG aprobado después de enderezar el cuerpo, manteniendo la cabeza de espaldas, como `assets/independientes/pose_horizontal_aprobada.png` (SHA-256 `72391725920009d6234574fbbbcde9a37fc38fb79b4cbd6d03af51a5a503d044`). Un botón independiente permite seleccionar esta pose.
- **Imagen acurrucada:** se incorpora el PNG individual solicitado junto con la pose horizontal, como `assets/independientes/pose_acurrucada_aprobada.png` (SHA-256 `083264e6997024c5c9837a17ea08af102f1503116dd5cb181abb24881f5e25b4`). No se modifica ni se recorta.
- **Transparencia:** se mantiene la corrección de `UpdateLayeredWindow` de 0.10.1 que el usuario confirmó que ya funciona.
- **Programa:** se conserva la ventana flotante y se puede alternar manualmente entre sentada, horizontal y acurrucada. No se mezclan otras poses que todavía requieren revisión.
- **Permanencia:** el ejecutable conserva el nombre `EvaDock.exe`. La copia instalada en `%LOCALAPPDATA%\EvaDock\EvaDock.exe` busca nuevas versiones en [GitHub Releases](https://github.com/youteachtk/EvaDock/releases) sin que sea necesario descargar manualmente un nuevo instalador por versión.

## Decisiones y materiales de referencia

El registro de decisiones de diseño, las poses aprobadas o pendientes y los identificadores de los recursos se encuentran en [docs/DECISIONES_EVA.md](docs/DECISIONES_EVA.md). Es la referencia permanente antes de modificar el diseño o solicitar nuevas fotografías.

Las imágenes de trabajo adicionales `recostada.png` y `dormida.png` están conservadas en la rama `eva-poses-independientes-v011`, para revisión; no se publicaron automáticamente en esta versión. Las fotografías de referencia originales no se divulgan en el repositorio público.

## Funcionamiento y limitaciones

La versión 0.10.4 incluye tres imágenes independientes aprobadas: sentada, acurrucada y horizontal. Los movimientos de la cola y las otras posiciones se preparan en la rama 0.11 y dependen de validación visual. La copia instalada comprueba nuevas versiones unos segundos después de abrirse y aproximadamente cada ocho horas; una actualización publicada se descarga, verifica mediante SHA-256 y sustituye automáticamente el ejecutable instalado.

La compilación se genera automáticamente en GitHub Actions para Windows 64 bits con:

`GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -tags eva_assets -ldflags="-H=windowsgui -s -w -X main.version=0.10.4" -o EvaDock.exe .`

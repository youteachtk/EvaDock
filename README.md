# EvaDock

Mascota independiente de Eva para el escritorio de Windows, con ventana flotante transparente y actualizaciones desde este repositorio.

## Versión publicada 0.10.2 — imagen sentada corregida

- **Imagen sentada:** se sustituyó el archivo antiguo de 0.10.0/0.10.1 (cola excesivamente esponjada, expresión y proporciones rechazadas) por la imagen individual elegida el 25 de octubre: `Eva, gata carey de pecho blanco.png`. Está integrada exactamente, sin recortes ni regeneración, como `assets/independientes/sentada.png`.
- **Transparencia:** se mantiene la corrección de `UpdateLayeredWindow` de 0.10.1 que el usuario confirmó que ya funciona.
- **Programa:** siguen disponibles los mismos controles y ventana flotante; esta actualización visual no mezcla las otras poses que todavía requieren revisión.
- **Permanencia:** el ejecutable conserva el nombre `EvaDock.exe`. La copia instalada en `%LOCALAPPDATA%\EvaDock\EvaDock.exe` busca nuevas versiones en [GitHub Releases](https://github.com/youteachtk/EvaDock/releases) sin que sea necesario descargar manualmente un nuevo instalador por versión.

## Decisiones y materiales de referencia

El registro de decisiones de diseño, las poses aprobadas o pendientes y los identificadores de los recursos se encuentran en [docs/DECISIONES_EVA.md](docs/DECISIONES_EVA.md). Es la referencia permanente antes de modificar el diseño o solicitar nuevas fotografías.

Las imágenes de trabajo adicionales `recostada.png` y `dormida.png` están conservadas en la rama `eva-poses-independientes-v011`, para revisión; no se publicaron automáticamente en esta versión. Las fotografías de referencia originales no se divulgan en el repositorio público.

## Funcionamiento y limitaciones

La versión 0.10.2 presenta una sola pose de Eva sentada. Los movimientos de la cola y las otras posiciones se preparan en la rama 0.11 y dependen de validación visual. La comprobación de nuevas versiones se realiza al abrir la copia instalada y aproximadamente cada ocho horas; la sustitución se efectúa al cerrar/reabrir conforme al flujo del programa.

La compilación se genera automáticamente en GitHub Actions para Windows 64 bits con:

`GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -tags eva_assets -ldflags="-H=windowsgui -s -w -X main.version=0.10.2" -o EvaDock.exe .`

# EvaDock

Mascota independiente de Eva para el escritorio de Windows. Proyecto Go/Win32 con ventana flotante, panel de control y actualizaciones desde este repositorio.

## Versión de trabajo 0.11.0 (pendiente de validación visual en Windows)

- Tres poses disponibles: **Sentada**, **Recostada** (con ojos cerrados) y **Dormida** (hecha bolita).
- Cada pose es un PNG independiente con fondo transparente, basado en fotografías reales de Eva. No se recorta una hoja colectiva.
- Patrón de manchas irregulares café, gris oscuro y ocre, pecho y patas blancos, ojos amarillo verdosos y cola estrecha oscura.
- En **Recostada**, sólo la punta de la cola se mueve dos veces de forma breve aproximadamente cada 25 segundos; la cabeza, los ojos cerrados y el cuerpo permanecen quietos.
- Selección manual y rutina automática que cambia entre las tres poses cada 12 segundos.
- Botones para ocultar, mostrar al frente, ajustar el tamaño y cerrar EvaDock.
- Las poses **De pie**, **Caminando** y **Salto** se suspendieron. Se retomarán con nuevas fotografías de referencia para conservar la anatomía real.

## Instalación y actualizaciones

El ejecutable publicado se llama siempre `EvaDock.exe`. Al ejecutarse desde otro lugar, se instala en `%LOCALAPPDATA%\EvaDock\EvaDock.exe` y crea el acceso directo `Eva Dock` en el menú Inicio.

La copia instalada consulta periódicamente [Releases de EvaDock](https://github.com/youteachtk/EvaDock/releases), descarga versiones posteriores y verifica el SHA-256 antes del reemplazo. El canal es exclusivo de EvaDock, no de WeekCalendarWidget. Las fotos originales del usuario no se alteran.

La rama `main` publica automáticamente un nuevo release al cambiar el código, los recursos o `version.txt`. Esta versión se prepara en una rama de trabajo; no se publica sin validación visual.

## Compilación y pruebas

```bash
go test ./internal/posecatalog
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -tags eva_assets -ldflags="-H=windowsgui -s -w -X main.version=0.11.0" -o EvaDock.exe .
```

## Prueba manual pendiente en Windows

1. Abrir el ejecutable y comprobar que aparezcan el panel y Eva flotante, o el modo alternativo visible si falla la transparencia.
2. Pulsar **Sentada**, **Recostada** y **Dormida**: verificar rostro, cola, patas y que no haya zonas cortadas.
3. Mantener **Recostada** al menos 8 segundos: observar dos movimientos de la cola sin cambios en el resto del cuerpo.
4. Activar/desactivar la rutina automática, ocultar/mostrar, traer al frente y ajustar tamaño.
5. Cerrar y reabrir desde el menú Inicio.

Si algo falla, revisar `%APPDATA%\EvaDock\eva.log`.

## Limitaciones

Las pruebas unitarias y la compilación cruzada no sustituyen la validación gráfica en Windows. El modo alternativo tiene fondo opaco si Windows rechaza `UpdateLayeredWindow`. La animación localizada de la cola debe revisarse visualmente antes de publicar.

## Historia

- `0.10.0`: primera pose sentada independiente y ventana alternativa de diagnóstico.
- `0.11.0` (en preparación): tres poses, controles manuales y pequeña oscilación localizada de la cola.

Las decisiones y el plan de implementación están documentados en `docs/superpowers/`.

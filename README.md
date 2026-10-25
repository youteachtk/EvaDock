# Eva Dock

Aplicación independiente de mascota para escritorio de Windows.

## Versión 0.10.1 — reparación de transparencia

- Corrige la llamada Win32 a `UpdateLayeredWindow`: ahora recibe un puntero válido a `BLENDFUNCTION`, que era el motivo del fallo que activaba la ventana alternativa con rectángulo beige.
- Se conserva la pose individual de 0.10.0 sin sustituir imágenes ni introducir poses no aprobadas.
- El ejecutable instalado en `%LOCALAPPDATA%\EvaDock\EvaDock.exe` conserva el mismo nombre y comprueba actualizaciones desde Releases de este repositorio; no hace falta descargar manualmente cada versión.
- Compilación y prueba unitaria verificadas. La transparencia en una PC Windows sigue pendiente de confirmación visual.

## Comportamiento heredado de 0.10.0

- La vista previa y la ventana flotante usan una imagen **independiente** de Eva, procedente del retrato individual aprobado. No se extraen poses de una hoja colectiva.
- Los archivos originales anteriores en la PC no se sobreescriben.
- La ventana flotante se coloca junto al panel, con botón para traerla al frente.
- Si Windows rechaza el dibujo transparente (UpdateLayeredWindow), se abre una ventana visible alternativa para diagnóstico.
- Se omiten las cinco poses anteriores de mala calidad hasta disponer de imágenes individuales apropiadas.
- Actualizaciones automáticas exclusivamente desde https://github.com/youteachtk/EvaDock/releases.

## Limitaciones actuales

La ventana alternativa tiene un fondo visible; la mascota sólo permanece sin fondo si Windows soporta la ventana transparente. La versión 0.10 no tiene movimientos corporales reales ni las otras cinco poses. Compilada para Windows 64-bit pero necesita validación visual en una máquina Windows.

## Compilación

`GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -tags eva_assets -ldflags="-H=windowsgui -s -w -X main.version=0.10.1" -o EvaDock.exe .`

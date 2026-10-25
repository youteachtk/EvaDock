# Eva Dock

Aplicación independiente para Windows que muestra a Eva, con seis poses aprobadas, sobre el escritorio.

## Estado actual
- Versión 0.9.0 preparada: ventana transparente, panel de control, seis poses y controles locales.
- El sistema de actualización consulta exclusivamente este repositorio: `youteachtk/EvaDock`.
- Las imágenes de Eva que se publican son solo los recortes generados y aprobados; no se publican las fotografías familiares originales.
- Limitación: las poses actuales son estáticas; el movimiento corporal fotorrealista aún no está implementado.

## Instalación y actualización
Se publicará `EvaDock.exe` en GitHub Releases con su comprobación SHA-256. Tras instalarlo una vez, la aplicación comprobará actualizaciones en este mismo repositorio. Conserva los archivos locales de Eva y la configuración al actualizar.

## Compilación (Go 1.22+)
```sh
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -tags eva_assets -ldflags="-H=windowsgui -s -w -X main.version=0.9.0" -o EvaDock.exe .
```

La compilación en Linux no sustituye las pruebas en Windows real.

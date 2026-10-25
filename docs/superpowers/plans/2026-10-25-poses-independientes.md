# EvaDock Independent Poses Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Mostrar tres poses individuales y una rutina automática sin romper el instalador y actualizador existente.

**Architecture:** Catálogo de recursos PNG en una unidad independiente de Go, con pruebas ejecutables en Linux. La UI Win32 existente consume las poses sin cambiar la arquitectura de ventanas. Se documenta cada decisión y validación.

**Tech Stack:** Go 1.23, Win32 nativo, `embed.FS`, PNG con alfa, GitHub Actions.

**Spec:** `docs/superpowers/specs/2026-10-25-poses-independientes-design.md`

## Global Constraints

- Sin recortes de hojas de sprites ni fotos reales incrustadas en el ejecutable.
- Conservar la única vía de actualización `youteachtk/EvaDock/releases`.
- No publicar aún en la rama `main` hasta validar visualmente las poses.
- Evitar las poses de pie, caminando y salto, suspendidas por falta de referencias anatómicas fiables.

## Review Focus

- Archivo PNG faltante: mostrar error legible; evitar ventana invisiblemente vacía.
- PNG dañado: no iniciar con catálogo parcial.
- Máscara alfa completamente opaca: rechazar recurso que muestre fondo rectangular.
- Selección de la última pose: no causar índice fuera de rango.
- Modo automático: alternar solamente tres poses, sin activar caminata ni salto.

---

### Task 1: Catálogo de poses verificable

**Files:** `internal/posecatalog/catalog.go`, `internal/posecatalog/catalog_test.go`, `main.go`.

**Interfaces:** `posecatalog.Load(readAsset func(string) ([]byte,error)) ([]posecatalog.Pose,error)` devuelve las tres imágenes en orden; `decode()` asigna ese resultado a `poseList`.

- [x] Escribir pruebas que fallen: catálogo contiene tres nombres ordenados; carga válida; fallo explícito para recursos faltantes o sin alfa.
- [x] Ejecutar `go test ./internal/posecatalog` y verificar fallo correcto antes de implementar.
- [x] Implementar `posecatalog.Load` en paquete portable y conectar `main.go` sin cambiar las llamadas Win32.
- [x] Incorporar las tres imágenes independientes en `assets/independientes/`.
- [x] Ejecutar pruebas del catálogo y compilación cruzada Windows.

### Task 2: Controles Win32 y rutina de tres poses

**Files:** `main.go`, `internal/posecatalog/catalog_test.go`.

**Interfaces:** controles `idSit,idLie,idSleep,idAuto`; `posecatalog.AutomaticIndex(tick int) int` para programación determinista, y `act(id)` para selección manual.

- [x] Escribir prueba de selección de tres poses y retorno periódico a primera pose; verificar fallo.
- [x] Añadir botones de las tres poses y selector automático, sin movimiento horizontal.
- [x] Escribir pruebas que fallen para `TailTwitch` y `TailShouldTwitch`; implementar la deformación sutil de una pequeña región de la cola en Recostada, sin alternar la imagen completa.
- [x] Ejecutar pruebas Go y compilación Win32.

### Task 3: Versionado, documentación y revisión

**Files:** `version.txt`, `README.md` y los anteriores.

- [x] Actualizar la documentación con controles, procedencia de imágenes y prueba manual en Windows.
- [x] Cambiar a `0.11.0` sólo cuando las pruebas confirmen que los tres recursos se enlazan.
- [x] Ejecutar todos los tests disponibles, generar binario de prueba y revisar diff antes de cualquier publicación.

## Registro de verificación local

- `go test -count=1 ./internal/...`: PASS.
- `go vet ./internal/...`: PASS.
- `GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -tags eva_assets ...`: PASS.
- Un equipo Windows todavía debe validar la ventana flotante, la apariencia y el movimiento de cola. No equivalen a una prueba de ejecución nativa.

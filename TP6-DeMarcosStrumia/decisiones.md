# Decisiones de testing

-------
Se ha implementado una estrategia de testing para el proyecto `ventas-app` / `ventas-frontend` que cubre:

- Unit y controller tests en backend (Go) usando `go test` y `github.com/stretchr/testify`.
- Mocks para la capa de persistencia (interfaz `database.DBHandler`) para aislar tests y simular errores y datos.
- Tests de frontend con `jest` y `@testing-library/react` (mock de axios en tests de componentes).
- Integración de ejecución de tests en CI (`.gitlab-ci.yml`) para ejecutar `go test` y `npm test`.

Principios y decisiones clave
-----------------------------

1) Abstracción de la DB mediante interfaz

- Se definió `database.DBHandler` (ya existente) y se adaptó `database.GetDB` para ser una variable-función que devuelve `DBHandler`.
- Esto permite inyectar mocks desde los tests sin tocar el código productivo.

2) Mocks controlables y realistas

- `mocks.MockDB` ahora soporta:
  - Flags finos para simular fallos por operación (`FailCreate`, `FailSave`, `FailFind`, `FailFirst`).
  - Relleno de datos en `Find`/`First` desde slices internos (`Productos`, `Usuarios`).
  - Simulación de `Create`/`Save` que actualiza el estado del mock (p. ej. stock, agregando registros).

3) Patrones de test: AAA y tests deterministas

- Todos los tests siguen AAA (Arrange, Act, Assert).
- Se busca que sean deterministas: no dependen de una base externa real, usan mocks controlados.

4) Cobertura y CI

- El pipeline (`.gitlab-ci.yml`) ejecuta build y tests backend y frontend.

  
Nota sobre MSW y Jest
---------------------

- Durante la integración se detectó una incompatibilidad puntual entre `msw` y el entorno de Jest usado (error de parseo por ESM en `node_modules/until-async`). Esto hacía que la inicialización del server MSW en `setupTests.ts` fallara en algunos entornos locales/CI.
- Decisión tomada: para evitar bloquear la entrega, los tests frontend se hicieron resilientes — si MSW puede arrancar, el test lo usa; si no, el test hace un mock de `api/axios` antes de importar el componente. Esto permite verificar el comportamiento del componente de forma determinista en ambos casos.

Archivos relevantes
------------------

- `ventas-app/mocks/db_mock.go` — mock mejorado de DB.
- `ventas-app/test/advanced_test.go` — tests avanzados de controllers (happy path + errores).
- `ventas-frontend/src/tests/Productos.test.tsx` — ejemplo de test de componente con mock de axios.
- `azure-pipelines.yml` — CI que ejecuta tests.

Cómo reproducir localmente
-------------------------

Backend:

```powershell
cd 'c:\Users\COTYS\Documents\Clases cuarto\Inge soft III\TP5-DeMarcosStrumia\ventas-app'
go mod tidy
go test ./... -v
```

Frontend:

```powershell
cd 'c:\Users\COTYS\Documents\Clases cuarto\Inge soft III\TP5-DeMarcosStrumia\ventas-frontend'
npm ci
npm test
```

Evidencias
----------
- Los resultados de ejecución de tests backend están en `evidence/backend_test_output.txt`.
 - Resultados y notas de los tests frontend están en `evidence/` (incluyen la estrategia MSW/fallback).


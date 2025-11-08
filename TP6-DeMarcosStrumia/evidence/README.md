Evidencias de ejecución de tests

Archivos:

- `backend_test_output.txt` - salida completa de `go test ./... -v` generada el 2025-10-28.

Cómo reproducir

1. Backend (Go)

```powershell
cd 'c:\Users\COTYS\Documents\Clases cuarto\Inge soft III\TP5-DeMarcosStrumia\ventas-app'
go mod tidy
go test ./... -v | tee ..\evidence\backend_test_output.txt
```

2. Frontend (Jest)

```powershell
cd 'c:\Users\COTYS\Documents\Clases cuarto\Inge soft III\TP5-DeMarcosStrumia\ventas-frontend'
npm ci
npm test -- --coverage
```

Nota: por motivos de velocidad y de entorno, no se ejecutaron tests frontend automáticos desde esta sesión; el archivo README incluye los pasos para reproducirlos localmente.

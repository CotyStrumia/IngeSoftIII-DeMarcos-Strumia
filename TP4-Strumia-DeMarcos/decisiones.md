#Instalacion
1. Instalamos el agente para hacer las pruebas
2. Para la instalacion generamos un access token, que debe tener como minimo el scope de agents pools (Read & manage)
3. Creamos un agents pool

#Pruebas
1.Usamos un pipeline simple para hacer pruebas, de esta manera verificamos que el pipeline este autorizado en nuestro 

#Carga inicial
1. Cargamos al repo desde main el proyecto con la estrcutura:
    / (raíz)
├─ /front        
├─ /back          
└─ azure-pipelines.yml

#Pipeline
1. Creamos un primer pipeline
2. Pasaje a multi-job (Front + Back)
    Acción: se definió un stage CI con dos jobs (build_front, build_back).
    Resultado: ambos fallaron en la primera corrida.
    Problemas detectados:
        Front: no existía frontend/dist al publicar artefactos.
        Back: el agente no tenía Go instalado.
    Soluciones aplicadas:
        Front: aseguramos scripts.build = "vite build", y publicamos frontend/dist.
        Back: instalamos Go 1.22+ en el agente y reiniciamos el servicio vstsagent....

 3. Front aún sin dist
    Acción: se revisó frontend/package.json y se agregó vite.config.js con outDir: 'dist'.
    Solución: job del front con npm ci && npm run build y publicación directa de frontend/dist.
    Resultado: el front empezó a pasar y generó el artefacto front.
4. Back: ruta del proyecto incorrecta
    Problema: el YAML asumía /back, pero el código estaba en /backend.
    Solución: se agregó un paso que detecta automáticamente el directorio del back:
        Usa /backend si existe con go.mod, si no, busca el primer go.mod fuera de /frontend.
        Guarda la ruta en $(BACK_DIR) para el resto del job.

5.  Back: error de cachés de Go (ruta hacia pip3.exe)
    Síntoma: go: could not create module cache: ...Python312\Scripts\pip3.exe.

    Causa: variables de entorno de Go (GOPATH/GOMODCACHE/GOCACHE) apuntaban a una ruta inválida.
    Solución: en el job forzamos:
        GOPATH=$(Agent.TempDirectory)\gopath
        GOMODCACHE=$(Agent.TempDirectory)\gopath\pkg\mod
        GOCACHE=$(Agent.TempDirectory)\gocache
        Creamos esas carpetas antes de go mod download/go build.
    Resultado: el back compiló y generó out/server(.exe).7
    De esta forma agregamos robutez

6.  Verificaciones y guardas del back
    Se añadieron:
        Listado del contenido del repo (debug) al inicio.
        Chequeo de existencia de go.mod y creación de out/.
        Detección del binario server.exe o server tras el build, con error claro si no aparece.

7.  Endurecimiento del pipeline
    Front: NodeTool@0 (Node 18) + Cache@2 de node_modules.
    Rutas Windows: se normalizaron a \ en pasos PowerShell.
    Logs útiles: impresión del contenido de frontend/dist y backend/out como evidencia.

8.  Pipeline final estable
    Stage: CI con dos jobs independientes (paralelizables).
    Artefactos publicados:
        front → contenido estático de frontend/dist (Vite).
        back → binario backend/out/server(.exe) (Go).
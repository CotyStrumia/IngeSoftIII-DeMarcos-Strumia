# 1. Forkear, clonar e identificar identidad

## 1. Fork
Hice fork del repositorio para poder trabajar desde mi github

## 2. Clonar
Clone el repositorio para poder trabajar localmente
- En la terminal:
    git clone https://github.com/margarita0912/2025_TP01_RepoBase.git

## 3. Configuración de identidad en Git
Para poder trabajar con commits identificables, configuré mi identidad en Git:

- En la terminal:
    git config user.name "margarita0912"
    git config user.email "2202199@ucc.edu.ar"

# 2. Desarrollar funcionalidad, rama feature

## 1. Creo la rama feature
- En la terminal, estando desde la rama main:
    git checkout -b feature

## 2. Commits atomicos
Creo al menos 2 commits atomicos
- En la terminal:
    1. git add decisiones.md
       git commit -m "creacion de decisiones.md"
    2. git add src/app.js
       git commit -m "mejora de mensaje en app.js"

## 3. Justificacion
Creé una nueva rama llamada `feature` para trabajar en la funcionalidad sin afectar `main`.  
Esto me permite seguir la práctica de GitHub Flow y mantener la rama principal siempre estable.
La estrategia de trabajar en una rama separada y con commits atómicos asegura trazabilidad, claridad y facilidad para futuras integraciones mediante Pull Requests.

# 3.Simulacion de error en main y arreglo con Hotfix

## 1. Me paso a la rama main y rompo app.js
- En la terminal: 
    git switch main
- Borro un parentesis en el "hola mundo"

## 2. Creo la rama hotfix y soluciono el error
- En la terminal: 
    git checkout -b hotfix
- Coloco nuevamente el parentesis
- En la terminal: 
    git add src/app.js
    git commit -m "arreglo sintaxis en app.js"
- Justificación: aislar la corrección y resolver con prioridad sin contaminar otras ramas.

## 3. Hago el merge de hotfix desde main
- En la terminal:
    git switch main
    git merge --no-ff hotfix -m "aplicar hotfix de sintaxis en app.js"
La linea del merge hace el commit tambien

# 4. Creo un pull request

## 1. Creo un pull request
- En la terminal:
    git push origin feature
- Abrí el PR en GitHub con base en `main` y compare en `feature-hola`.
- El título y la descripción explican claramente qué cambios se integran
Creé un Pull Request desde la rama `feature` hacia `main` en GitHub.  
El objetivo fue integrar la funcionalidad mejorada del hola mundo.

# 5. Versiono mi trabajo

## 1. Creo un tag
Desde mi rama main
- En la terminal:
    git tag v1.0 -m "Versión estable inicial v1.0"
    git push origin v1.0
Alli cree el tag y lo subi

## 2. Justificacion
Se creó el tag `v1.0` para marcar la primera versión estable del repositorio.
Se uso la convencion SameVer


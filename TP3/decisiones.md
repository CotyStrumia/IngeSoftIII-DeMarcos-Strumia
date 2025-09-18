# Decisiones del Proyecto - ProyectoStrumia-DeMarcos

## 1. Proceso
- Se utilizó el proceso **Basic**, pensado para proyectos simples.  
- En este proceso no existen Bugs como tipo de work item, por lo que se utilizaron **Issues** para cumplir ese rol.

## 2. Organización de Equipos y Áreas
- Se crearon dos equipos: **Backend** y **Frontend**.  
- Cada equipo cuenta con su propia área, lo que permite separar y organizar los work items.  
- Los miembros fueron asignados a sus equipos correspondientes.

## 3. Backlog y Work Items
- **Epic** creado en el área de Backend: agregar usuarios.  
- **User Stories**:  
  - Registro de usuario nuevo (general).  
  - Lista de usuarios (frontend).  
  - Función para listar usuarios (backend).  
- **Tasks**:  
  - Frontend: función para recibir datos del backend y diseño de la vista.  
  - Backend: endpoint para listar usuarios y otro para almacenar usuarios en la BD.  
- **Issues (Bugs)**:  
  - Error en username al permitir espacios.  
  - Error por contraseña poco segura.

## 4. Sprint
- Se creó un sprint de dos semanas (01/09/2025 – 15/09/2025).  
- En este sprint se incluyeron User Stories, Tasks e Issues del Frontend.  
- Ejemplo: diseño de la lista de usuarios y pruebas de listado.

## 5. Repositorio
- El repositorio se creó en Azure Repos.  
- Se utilizó un **Personal Access Token (PAT)** para clonar y pushear código desde local.  
- El código cargado corresponde a un mini proyecto en Go, con soporte para Docker.  

## 6. Pipelines
- Inicialmente se probó con un **agente local**, que requería configuración manual de PAT, permisos y ejecución como servicio.  
- Finalmente, se optó por **usar agentes hospedados de Azure**, simplificando la configuración.  
- El pipeline se definió en el archivo `azure-pipelines.yml` y corre automáticamente al hacer push.


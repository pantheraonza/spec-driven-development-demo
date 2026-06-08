# Spec-Driven Development Demo

Demo de Spec-Driven Development - Ecommerce sencillo con Go y React.

Este proyecto demuestra cómo las especificaciones, planes, tareas, pruebas y decisiones arquitectónicas guían la implementación de un producto pequeño pero realista.

## Estructura del proyecto

```
.
├── docs/                          # Documentación del proyecto
│   ├── 000-constitution.md       # Constitución del proyecto
│   ├── 001-specification.md      # Especificación funcional
│   ├── 002-plan.md               # Plan de implementación
│   ├── 003-tasks.md              # Tareas generadas
│   ├── 004-analysis.md           # Análisis y revisión
│   ├── adr/                      # Architecture Decision Records
│   └── specs/                    # Contratos API y modelos
├── backend/                       # Backend en Go
│   ├── cmd/
│   │   └── api/
│   │       └── main.go
│   ├── internal/
│   │   ├── catalog/
│   │   ├── cart/
│   │   ├── shared/
│   │   └── persistence/
│   ├── go.mod
│   ├── go.sum
│   └── Makefile
├── frontend/                      # Frontend en React
│   ├── src/
│   │   ├── features/
│   │   │   ├── catalog/
│   │   │   └── cart/
│   │   ├── shared/
│   │   └── App.tsx
│   ├── package.json
│   └── tsconfig.json
├── .gitignore
└── Makefile
```

## Inicio rápido

### Backend

```bash
cd backend
make run
```

### Frontend

```bash
cd frontend
npm install
npm start
```

## Documentación

Lee la documentación en el siguiente orden:

1. **[Constitución](./docs/000-constitution.md)** - Principios fundamentales del proyecto
2. **[Especificación](./docs/001-specification.md)** - Requisitos y historias de usuario
3. **[Plan](./docs/002-plan.md)** - Estrategia de implementación
4. **[Tareas](./docs/003-tasks.md)** - Tasks de desarrollo
5. **[Análisis](./docs/004-analysis.md)** - Revisión y validación

## Contratos API

Ver [OpenAPI spec](./docs/specs/openapi.yaml)

## Decisiones arquitectónicas

Ver [ADRs](./docs/adr/)

## Stack tecnológico

- **Backend:** Go
- **Frontend:** React + TypeScript
- **Persistencia:** SQLite (desarrollo) / PostgreSQL (producción)
- **API:** HTTP/JSON
- **Arquitectura:** Monolito modular

## Licencia

MIT

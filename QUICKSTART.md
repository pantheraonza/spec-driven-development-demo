# Quick Start Guide - Ecommerce Spec-Driven Development Demo

## Introducción

Esta es una aplicación ecommerce sencilla construida con **Go** (backend) y **React** (frontend) siguiendo principios de **Spec-Driven Development**.

La aplicación demuestra:
- Especificación clara antes de implementación
- Backend como fuente de verdad
- API HTTP/JSON con contratos estables
- Validación de reglas de negocio
- UX con estados claros (carga, vacío, error, éxito)

## Requisitos previos

- **Go 1.21+** - [Descargar](https://golang.org/dl/)
- **Node.js 18+** - [Descargar](https://nodejs.org/)
- **npm** - Viene con Node.js

Verificar instalación:
```bash
go version      # Go 1.21 o superior
node --version  # Node 18 o superior
npm --version   # npm 9 o superior
```

## Estructura del proyecto

```
spec-driven-development-demo/
├── backend/                 # Go API server
│   ├── cmd/api/
│   │   └── main.go         # Entry point
│   ├── internal/
│   │   ├── catalog/        # Módulo catálogo
│   │   ├── cart/           # Módulo carrito
│   │   ├── shared/         # Utilidades compartidas
│   │   └── persistence/    # Capa de datos
│   ├── go.mod
│   └── Makefile
├── frontend/                # React app
│   ├── src/
│   │   ├── features/       # Componentes
│   │   ├── api/            # Cliente HTTP
│   │   ├── types/          # TypeScript types
│   │   ├── styles/         # CSS
│   │   └── App.tsx
│   ├── index.html
│   ├── package.json
│   └── vite.config.ts
├── docs/                    # Documentación
│   ├── 000-constitution.md  # Principios del proyecto
│   ├── 001-specification.md # Especificación funcional
│   ├── 002-plan.md         # Plan de implementación
│   └── adr/                # Architecture Decision Records
└── README.md
```

## Inicio rápido

### Paso 1: Clonar repositorio (ya está hecho ✓)

```bash
cd spec-driven-development-demo
```

### Paso 2: Iniciar Backend

```bash
cd backend

# Descargar dependencias
go mod download

# Compilar y ejecutar
go run ./cmd/api/main.go
```

**Salida esperada:**
```
2026-06-08 Server starting on port 8080
```

El backend estará disponible en: `http://localhost:8080`

### Paso 3: Iniciar Frontend (en otra terminal)

```bash
cd frontend

# Instalar dependencias
npm install

# Ejecutar dev server
npm run dev
```

**Salida esperada:**
```
  VITE v5.0.0  ready in 234 ms

  ➜  Local:   http://localhost:5173/
```

El frontend se abrirá automáticamente en: `http://localhost:5173`

## Uso de la aplicación

### Pantalla principal: Catálogo

Al iniciar, verás:
- **5 productos de prueba** con nombre, descripción, precio
- **Productos disponibles** (con stock > 0) - botón "Add to Cart" habilitado
- **Productos sin stock** - etiqueta "Out of Stock", botón deshabilitado
- **Resumen del carrito** en la esquina superior derecha

### Agregar productos al carrito

1. Haz clic en "Add to Cart" en cualquier producto disponible
2. Verás un mensaje de éxito: "Product added to cart"
3. El resumen del carrito se actualiza con:
   - Número de items
   - Subtotal en EUR

### Intentar operaciones inválidas

Prueba estos escenarios:
- ❌ Haz clic en "Add to Cart" para producto sin stock → Error "Out of Stock"
- ❌ Intenta agregar el mismo producto dos veces → Incrementa cantidad en lugar de duplicar

### Refrescar página

El carrito persiste después de refrescar (guardado en servidor)

## Endpoints de API

### GET /api/products

Obtiene catálogo público de productos activos.

**Petición:**
```bash
curl http://localhost:8080/api/products
```

**Respuesta (200):**
```json
{
  "products": [
    {
      "id": "prod_001",
      "name": "Camiseta Básica",
      "description": "Camiseta de algodón 100%",
      "price": 19.99,
      "currency": "EUR",
      "stock": 12,
      "available": true
    }
  ]
}
```

### POST /api/cart/items

Añade producto al carrito.

**Petición:**
```bash
curl -X POST http://localhost:8080/api/cart/items \
  -H "Content-Type: application/json" \
  -d '{"productId":"prod_001","quantity":1}'
```

**Respuesta (200):**
```json
{
  "cart": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "items": [
      {
        "productId": "prod_001",
        "name": "Camiseta Básica",
        "unitPrice": 19.99,
        "currency": "EUR",
        "quantity": 1,
        "lineTotal": 19.99
      }
    ],
    "subtotal": 19.99,
    "currency": "EUR"
  }
}
```

**Errores posibles:**
- `400 INVALID_QUANTITY` - Cantidad < 1
- `404 PRODUCT_NOT_FOUND` - Producto no existe
- `409 PRODUCT_UNAVAILABLE` - Producto inactivo
- `409 PRODUCT_OUT_OF_STOCK` - Sin stock
- `409 INSUFFICIENT_STOCK` - Stock insuficiente

## Datos de prueba

El backend crea automáticamente 5 productos:

| ID | Nombre | Precio | Stock | Disponible |
|---|---|---|---|---|
| prod_001 | Camiseta Básica | €19.99 | 12 | ✓ |
| prod_002 | Sudadera Gris | €49.99 | 0 | ✗ |
| prod_003 | Pantalones Azules | €79.99 | 5 | ✓ |
| prod_004 | Calcetines Blancos | €9.99 | 50 | ✓ |
| prod_005 | Gorro Negro | €24.99 | 0 | ✗ (inactivo) |

## Compilación para producción

### Backend

```bash
cd backend
go build -o bin/api ./cmd/api
./bin/api
```

### Frontend

```bash
cd frontend
npm run build
npm run preview
```

## Troubleshooting

### Puerto 8080 ya en uso

```bash
# Backend en puerto diferente
PORT=3000 go run ./cmd/api/main.go
```

### Port 5173 ya en uso

Vite lo detectará automáticamente y usará el siguiente puerto disponible.

### Error: "Cannot connect to backend"

Verifica que:
1. Backend está corriendo en `http://localhost:8080`
2. CORS está habilitado en backend
3. No hay firewall bloqueando puerto 8080

### Error: "npm: command not found"

Instala Node.js desde https://nodejs.org/

## Estructura de documentación

📄 **Especificación:** `docs/001-specification.md`
- Historias de usuario
- Criterios de aceptación
- Reglas de negocio
- Expectativas de API

📄 **Plan:** `docs/002-plan.md`
- Arquitectura
- Modelos de datos
- Responsabilidades
- Fases de implementación

📄 **ADRs:** `docs/adr/`
- 0001: Arquitectura monolito modular
- 0002: Backend como fuente de verdad
- 0003: API HTTP/JSON
- 0004: Persistencia de carrito en servidor

## Desarrollo

### Ejecutar tests backend

```bash
cd backend
go test ./...
```

### Hot reload frontend

El dev server de Vite recarga automáticamente cambios en src/

### Ver cambios en vivo

1. Modifica archivo en `frontend/src/`
2. Guarda
3. El navegador se actualiza automáticamente

## Conceptos de Spec-Driven Development

Este proyecto demuestra:

✅ **Especificación primero**: Toda funcionalidad empieza con docs/001-specification.md

✅ **Backend como fuente de verdad**: El frontend NUNCA envía precios, solo productId y quantity

✅ **Validación en backend**: Todas las reglas de negocio se validan en servidor

✅ **Errores estables**: Códigos de error consistentes y predecibles

✅ **UX clara**: Estados de carga, vacío, error, éxito explícitos

✅ **Arquitectura documentada**: Decisiones registradas en ADRs

## Próximos pasos

Opciones para expandir la demo:

1. **Autenticación**: Agregar login de usuario
2. **Checkout**: Crear orden desde carrito
3. **Pagos**: Integrar procesador de pagos simulado
4. **Búsqueda**: Filtrar productos por nombre
5. **Categorías**: Agrupar productos
6. **Admin**: Panel para gestionar productos

Cada feature seguiría el mismo flujo SDD:
1. Escribir especificación
2. Crear ADR si es decisión arquitectónica
3. Implementar backend
4. Implementar frontend
5. Escribir tests

## Documentación adicional

- 📋 `docs/000-constitution.md` - Constitución y principios
- 📋 `docs/001-specification.md` - Especificación completa
- 📋 `docs/002-plan.md` - Plan técnico
- 📋 `docs/adr/` - Architecture Decision Records

## Contacto / Preguntas

Este es un proyecto de demostración de Spec-Driven Development.

Para preguntas sobre la implementación, revisa la documentación en `docs/`

## Licencia

MIT

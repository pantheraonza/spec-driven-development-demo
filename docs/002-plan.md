# Plan de implementación: Catálogo y Carrito

## Dirección técnica

Arquitectura de monolito modular:

**Backend (Go):**
- `catálogo` - productos y disponibilidad
- `carrito` - gestión de carrito anónimo
- `shared` - API y error handling
- `persistence` - capa de datos

**Frontend (React):**
- `features/catalog` - página y tarjetas
- `features/cart` - resumen
- `shared` - tipos, utilidades, UI
- `api` - cliente HTTP

## Estructura Backend

```
/backend
  /cmd/api
    main.go
  /internal
    /catalog
      product.go
      repository.go
      service.go
      handler.go
      dto.go
    /cart
      cart.go
      item.go
      repository.go
      service.go
      handler.go
      dto.go
    /shared
      errors.go
      response.go
      validation.go
      money.go
    /persistence
      db.go
      migrations/
      catalog_repository.go
      cart_repository.go
  go.mod
  go.sum
  Makefile
```

## Responsabilidades Backend

### GET /api/products
- Cargar productos activos
- Excluir inactivos
- Marcar disponibilidad por stock
- JSON estable
- Error: CATALOG_UNAVAILABLE

### POST /api/cart/items
- Validar productId
- Validar quantity
- Cargar producto
- Rechazar inexistente
- Rechazar inactivo
- Rechazar sin stock
- Rechazar cantidad > stock
- Crear carrito si no existe
- Añadir o incrementar
- Calcular totales
- Persistir
- Devolver carrito

## Modelo de datos

### Product
- ID string
- Name string
- Description string
- Price decimal
- Currency string
- Stock int
- Active bool
- CreatedAt time
- UpdatedAt time

Métodos:
- IsAvailable() bool
- CanAddToCart(quantity int) error

### Cart
- ID string
- Items []CartItem
- Currency string
- CreatedAt time
- UpdatedAt time

Métodos:
- AddItem(product Product, quantity int) error
- Subtotal() decimal
- ItemCount() int

### CartItem
- ProductID string
- Name string
- UnitPrice decimal
- Currency string
- Quantity int
- LineTotal() decimal

## Persistencia

Base de datos: SQLite (dev) / PostgreSQL (prod)

**Tabla: products**
```sql
CREATE TABLE products (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  price_amount DECIMAL NOT NULL,
  currency TEXT NOT NULL,
  stock INT NOT NULL,
  active BOOLEAN DEFAULT true,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Tabla: carts**
```sql
CREATE TABLE carts (
  id TEXT PRIMARY KEY,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

**Tabla: cart_items**
```sql
CREATE TABLE cart_items (
  id TEXT PRIMARY KEY,
  cart_id TEXT NOT NULL REFERENCES carts(id),
  product_id TEXT NOT NULL,
  quantity INT NOT NULL,
  unit_price_amount DECIMAL NOT NULL,
  currency TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

Importante:
- Precios guardados en cart_items para historial
- Sin reserva de stock
- IDs como UUIDs o strings generados

## Contratos API

Ver: `docs/specs/openapi.yaml`

## Estructura Frontend

```
/frontend
  /src
    /features
      /catalog
        CatalogPage.tsx
        ProductCard.tsx
        catalogApi.ts
        catalogTypes.ts
      /cart
        CartSummary.tsx
        cartApi.ts
      /shared
        ApiError.ts
        LoadingState.tsx
        ErrorMessage.tsx
        EmptyState.tsx
        money.ts
    /App.tsx
    /index.tsx
  package.json
  tsconfig.json
```

## Responsabilidades Frontend

### CatalogPage
- Cargar productos con GET /api/products
- Estados: loading, empty, error, success
- Renderizar ProductCard
- Pasar callbacks para añadir al carrito

### ProductCard
- Mostrar: nombre, descripción, precio, disponibilidad
- Botón "Añadir al carrito" deshabilitado si no disponible
- Feedback post-click
- Mostrar errores claros

### CartSummary
- Items actuales
- Cantidades
- Subtotal
- Moneda
- Actualizar post-operación

## Testing

### Unit tests Backend
- Disponibilidad de producto
- Validación de cantidad
- Añadir a vacío
- Incrementar cantidad
- Rechazar inactivo
- Rechazar sin stock
- Rechazar cantidad > stock
- Cálculo de totales

### Integration tests Backend
- GET /api/products devuelve activos
- GET /api/products excluye inactivos
- GET /api/products marca sin stock
- POST /api/cart/items añade disponible
- POST rechaza cantidad inválida
- POST rechaza inexistente
- POST rechaza inactivo
- POST rechaza sin stock
- POST devuelve carrito actualizado

### Frontend tests
- Catálogo carga
- Catálogo vacío
- Catálogo error
- Productos se renderizan
- Botón deshabilitado sin stock
- Añadir actualiza resumen
- Error del backend visible

## Gestión de errores

Códigos estables:
- CATALOG_UNAVAILABLE (500)
- INVALID_QUANTITY (400)
- PRODUCT_NOT_FOUND (404)
- PRODUCT_UNAVAILABLE (409)
- PRODUCT_OUT_OF_STOCK (409)
- INSUFFICIENT_STOCK (409)

Frontend mapea a mensajes legibles.

## Observabilidad

Logging:
- Fallo cargar catálogo
- Intento producto inactivo
- Intento sin stock
- Cantidad inválida
- Fallo persistencia carrito

## Seguridad

- Validar todas las entradas en backend
- Ignorar precios del cliente
- Códigos error sin detalles internos
- Validación backend es definitiva
- Monolito simplifica seguridad inicial

## ADRs a crear

1. Monolito modular (vs microservicios)
2. Backend fuente de verdad (precios, stock)
3. API HTTP/JSON (vs gRPC)
4. Carrito anónimo servidor (vs cliente)

## Fases implementación

1. **Setup inicial**
   - Estructura proyecto
   - Go module setup
   - React setup
   - Docker (opcional)

2. **Backend Catálogo**
   - Model Product
   - Repository
   - Service
   - Handler GET /api/products
   - Tests

3. **Backend Carrito**
   - Models Cart/Item
   - Repository
   - Service
   - Handler POST /api/cart/items
   - Tests

4. **Frontend Catálogo**
   - CatalogPage
   - ProductCard
   - Api client
   - Tests

5. **Frontend Carrito**
   - CartSummary
   - Integración con catálogo

6. **Pulido**
   - Error handling mejorado
   - UX refinamiento
   - Tests adicionales
   - Documentación

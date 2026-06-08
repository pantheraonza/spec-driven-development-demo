# Funcionalidad: Catálogo de productos y carrito de compra

## Objetivo de la funcionalidad

Permitir que los visitantes consulten el catálogo de productos y añadan productos disponibles a un carrito de compra.

Esta funcionalidad representa el primer flujo útil de un ecommerce:
1. El usuario abre el catálogo
2. El usuario ve los productos disponibles
3. El usuario añade un producto al carrito
4. El usuario ve el carrito actualizado

## Alcance

Incluye:
- Mostrar catálogo público de productos
- Mostrar disponibilidad de productos
- Añadir productos disponibles al carrito
- Devolver carrito actualizado
- Gestionar operaciones inválidas
- Estados claros en frontend: carga, vacío, éxito, error

Excluye:
- Autenticación/registro
- Checkout
- Procesamiento de pagos
- Administración de productos
- Búsqueda
- Categorías
- Descuentos
- Envío
- Reserva de inventario

## Actores

### Visitante

Usuario anónimo que puede:
- Ver catálogo público
- Consultar disponibilidad
- Añadir productos disponibles al carrito
- Ver resumen del carrito

No puede:
- Crear pedidos
- Pagar
- Gestionar productos
- Autenticarse

## Historia de usuario 1: Ver catálogo

Como visitante, quiero ver el catálogo de productos, para poder decidir qué comprar.

### Criterios de aceptación

**Escenario 1: Hay productos disponibles**
- Mostrar lista de productos activos
- Cada producto: nombre, descripción, precio, moneda, disponibilidad

**Escenario 2: Producto activo con stock**
- Aparece como disponible
- Visitante puede añadirlo al carrito

**Escenario 3: Producto activo sin stock**
- Aparece como no disponible
- Visitante no puede añadirlo

**Escenario 4: Producto inactivo**
- No aparece en catálogo público

**Escenario 5: Catálogo vacío**
- Mostrar estado vacío explicativo

**Escenario 6: Error al cargar**
- Mostrar error claro
- Sin datos parciales o inválidos

## Historia de usuario 2: Añadir al carrito

Como visitante, quiero añadir productos disponibles, para preparar compra futura.

### Criterios de aceptación

**Escenario 1: Añadir a carrito vacío**
- Producto aparece con cantidad 1
- Backend calcula subtotal

**Escenario 2: Añadir dos veces el mismo producto**
- Cantidad pasa a 2
- Backend recalcula subtotal

**Escenario 3: Añadir sin stock**
- Backend rechaza
- Frontend muestra mensaje claro

**Escenario 4: Añadir producto inactivo**
- Backend rechaza

**Escenario 5: Añadir inexistente**
- Error 404
- Carrito sin cambios

**Escenario 6: Cantidad inválida**
- Cantidad < 1 rechazada

**Escenario 7: Cantidad superior a stock**
- Backend rechaza
- Frontend muestra cantidad disponible

## Reglas de negocio

### Catálogo
1. Solo productos activos se muestran públicamente
2. Stock > 0 = disponible
3. Stock = 0 = no disponible
4. Precios del backend
5. Sin modificar precios desde frontend
6. Acceso sin autenticación

### Carrito
1. Anónimo en servidor
2. Creación implícita al añadir primer producto
3. Solo productos activos
4. Stock suficiente requerido
5. Cantidad > 0
6. Incrementar si producto existe
7. Backend calcula totales
8. Sin enviar precios desde frontend
9. No reserva stock
10. Checkout fuera de alcance

## Requisitos frontend

Página de catálogo debe:
- Cargar desde API backend
- Mostrar estado de carga
- Mostrar estado vacío
- Mostrar error claro
- Mostrar: nombre, descripción, precio, moneda, disponibilidad
- Diferenciar visualmente no disponibles
- Deshabilitar añadir para no disponibles
- Feedback claro tras añadir
- Resumen de carrito tras actualizar

No debe:
- Calcular precios
- Decidir stock final
- Confiar estado local sobre backend
- Mostrar errores internos

## Requisitos backend

Endpoints:
- Recuperar catálogo público
- Añadir producto al carrito
- Devolver carrito actualizado

Debe:
- Validar peticiones
- Filtrar inactivos
- Validar producto existe
- Validar activo
- Validar stock
- Validar cantidad
- Calcular totales
- Códigos error estables
- Sin exponer internos
- Reglas de negocio testables

## Expectativas de API

### GET /api/products

**Response 200:**
```json
{
  "products": [
    {
      "id": "prod_001",
      "name": "Camiseta básica",
      "description": "Camiseta algodón 100%",
      "price": 19.99,
      "currency": "EUR",
      "stock": 12,
      "available": true
    },
    {
      "id": "prod_002",
      "name": "Sudadera",
      "description": "Sudadera unisex",
      "price": 49.99,
      "currency": "EUR",
      "stock": 0,
      "available": false
    }
  ]
}
```

**Error 500:**
```json
{
  "error": {
    "code": "CATALOG_UNAVAILABLE",
    "message": "Catálogo no disponible temporalmente"
  }
}
```

### POST /api/cart/items

**Request:**
```json
{
  "productId": "prod_001",
  "quantity": 1
}
```

**Response 200:**
```json
{
  "cart": {
    "id": "cart_123",
    "items": [
      {
        "productId": "prod_001",
        "name": "Camiseta básica",
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

**Error 400 - Cantidad inválida:**
```json
{
  "error": {
    "code": "INVALID_QUANTITY",
    "message": "Cantidad debe ser mayor a cero"
  }
}
```

**Error 404 - No encontrado:**
```json
{
  "error": {
    "code": "PRODUCT_NOT_FOUND",
    "message": "Producto no existe"
  }
}
```

**Error 409 - No disponible:**
```json
{
  "error": {
    "code": "PRODUCT_UNAVAILABLE",
    "message": "Producto no disponible"
  }
}
```

**Error 409 - Sin stock:**
```json
{
  "error": {
    "code": "PRODUCT_OUT_OF_STOCK",
    "message": "Sin stock"
  }
}
```

**Error 409 - Stock insuficiente:**
```json
{
  "error": {
    "code": "INSUFFICIENT_STOCK",
    "message": "Stock insuficiente para cantidad"
  }
}
```

## Casos límite

1. Catálogo sin productos activos
2. Todos activos sin stock
3. Stock negativo (inconsistencia)
4. Precio inválido
5. Catálogo no cargable
6. Carrito no existe al añadir primero
7. Mismo producto varias veces
8. Cantidad cero
9. Cantidad negativa
10. Cantidad > stock
11. Producto existe pero inactivo
12. Producto no existe
13. Fallo persistencia carrito
14. Dos pestañas simultáneas
15. Datos obsoletos frontend

## Decisiones iniciales

1. No disponibles se muestran pero no se pueden añadir
2. Sin paginación v1
3. EUR
4. Precios incluyen IVA
5. Carrito anónimo
6. Servidor
7. Sin reserva stock
8. Resumen tras actualizar
9. IDs product strings
10. API devuelve stock (demo)

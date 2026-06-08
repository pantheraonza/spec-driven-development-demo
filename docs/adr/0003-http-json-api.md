# ADR 0003: API HTTP/JSON entre React y Go

## Estado
ACEPTADO

## Contexto

El frontend React y backend Go necesitan comunicarse. Se consideraron varios enfoques:
- HTTP/REST
- gRPC
- WebSockets
- GraphQL

Para una demo de Spec-Driven Development que debe ser entendible, se requiere:
- Simplicidad de implementación
- Facilidad de demostración
- Debugging trivial
- Sin dependencias complejas

## Decisión

Usamos **HTTP/JSON** con endpoints REST simples:
- `GET /api/products` - recuperar catálogo
- `POST /api/cart/items` - añadir al carrito

Los contratos están definidos en OpenAPI.

## Alternativas consideradas

1. **gRPC**: Más complejo, requiere herramientas especiales
2. **GraphQL**: Overhead innecesario para esta scope
3. **WebSockets**: No necesarios, operaciones son request/response
4. **HTTP/JSON**: Simple, estándar, observable ✓

## Consecuencias

**Positivas:**
- Fácil de debuggear con curl, Postman, navegador
- Estándar web conocido
- Sin dependencias adicionales
- Contratos claros en OpenAPI
- Obvio para aprender SDD

**Negativas:**
- Menos eficiente que gRPC (overhead)
- Tipado débil en respuesta (sin protobuf)

## Riesgos

- Bajo riesgo. Patrón universal.

## Validación

- Contratos definidos en OpenAPI spec
- Todos los tests usan HTTP
- Frontend solo hace llamadas HTTP

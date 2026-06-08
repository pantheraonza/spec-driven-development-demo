# ADR 0002: Backend como fuente de verdad para precios y validación de stock

## Estado
ACEPTADO

## Contexto

En una aplicación ecommerce, es crítico que los valores autoritativos (precios, disponibilidad de stock) no sean controlables por el cliente.

Se ha considerado confiar parcialmente en datos del cliente para mejorar la experiencia:
- El frontend podría calcular precios localmente
- El frontend podría validar stock antes de enviar al backend

Esto reduciría latencia pero comprometería seguridad y corrección del negocio.

## Decisión

El **backend Go es la única fuente de verdad** para:
- Precios de productos
- Disponibilidad y cantidad de stock
- Validación de cantidad en carrito
- Cálculo de totales

El frontend React **nunca** envía precios. Solo envía productId y quantity.

El backend siempre valida y rechaza datos inválidos.

## Alternativas consideradas

1. **Frontend como fuente de verdad**: Seguridad débil
2. **Validación dual**: Complejidad sin beneficio
3. **Backend como fuente de verdad**: Seguridad y simplicidad ✓

## Consecuencias

**Positivas:**
- Seguridad: imposible manipular precios o stock desde cliente
- Simplicidad: lógica centralizada
- Auditabilidad: todos los cambios en servidor
- Consistencia: datos siempre correctos

**Negativas:**
- Latencia: cada operación requiere servidor
- Debe optimizarse con caching si crece

## Riesgos

- Bajo riesgo. Patrón estándar en ecommerce.

## Validación

- Tests confirman que backend rechaza precios del cliente
- Tests confirman que frontend nunca envía precios
- Tests confirman que cálculos de carrito siempre vienen del backend

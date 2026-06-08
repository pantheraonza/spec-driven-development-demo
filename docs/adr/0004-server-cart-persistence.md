# ADR 0004: Carrito anónimo persistido en servidor

## Estado
ACEPTADO

## Contexto

El carrito debe sobrevivir a refrescos del navegador y actualizaciones de página. Se consideró almacenarlo en:
- LocalStorage del navegador
- SessionStorage
- Backend en servidor

## Decisión

El **carrito se persiste en el servidor** con ID enviado en cookie o header.

Frontend no es responsable de persistencia. Backend es la fuente de verdad.

## Consecuencias

**Positivas:**
- Carrito sobrevive refrescos
- Validación backend siempre correcta
- Seguridad mejorada
- Fácil añadir autenticación después

**Negativas:**
- Require sesión servidor
- Limpieza de carritos abandonados

## Validación

- Tests confirman que carrito persiste tras refresh
- Backend rechaza operaciones con carrito inválido

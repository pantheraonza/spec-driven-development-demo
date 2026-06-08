# ADR 0001: Usar arquitectura de monolito modular

## Estado
ACEPTADO

## Contexto

El proyecto es una demo de Spec-Driven Development para un ecommerce sencillo. Debe mantenerse comprensible para una audiencia técnica que aprende SDD.

Se ha considerado usar microservicios desde el inicio, pero esto añadiría complejidad innecesaria (orquestación, tolerancia de fallos distribuida, observabilidad avanzada).

## Decisión

Usamos un **monolito modular** donde el código se organiza por **capacidades de negocio** (catálogo, carrito), no por capas técnicas.

Esto proporciona límites claros sin introducir complejidad de sistemas distribuidos.

## Alternativas consideradas

1. **Microservicios**: Demasiada complejidad para una demo
2. **Monolito por capas**: Sin límites de negocio claros
3. **Monolito modular**: Equilibrio claridad-realismo ✓

## Consecuencias

**Positivas:**
- Fácil de entender y explicar
- Despliegue simple
- Debug sencillo
- Refactorización a microservicios posible después

**Negativas:**
- Límites menos estrictos que microservicios
- Escalabilidad limitada
- Base de datos compartida

## Riesgos

- Bajo riesgo. Si el proyecto crece, refactorizar a microservicios es viable.

## Validación

- Verificar que módulos tienen responsabilidades claras
- Comprobar que lógica de negocio no se filtra entre módulos
- Revisar que dependencias fluyen unidireccionalmente

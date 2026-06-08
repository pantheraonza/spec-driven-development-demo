# Constitución del proyecto: Ecommerce sencillo

## Contexto del proyecto

Este proyecto es una aplicación de ecommerce sencilla construida con:

- Go para el backend.
- React para el frontend.
- Comunicación HTTP/JSON entre frontend y backend.
- Una base de datos relacional simple para la persistencia.

El objetivo del proyecto es demostrar Spec-Driven Development de forma clara y práctica.

## Principios fundamentales

### 1. Primero la especificación

Toda funcionalidad debe empezar con una especificación clara antes de comenzar la implementación.

### 2. No inventar requisitos

El agente no debe inventar comportamiento que no haya sido especificado.

### 3. Arquitectura simple primero

La arquitectura inicial debe ser un monolito modular.

### 4. Separación clara entre frontend y backend

El frontend es responsable de la interacción y presentación.

El backend es responsable de reglas de negocio, validación, persistencia y cálculos.

### 5. Backend organizado por capacidades de negocio

El backend se organiza por:
- Catálogo de productos
- Carrito de compra
- Gestión compartida de API y errores
- Infraestructura de persistencia

### 6. La API como contrato

El backend expone contratos HTTP/JSON claros definiendo:
- Método HTTP
- URL
- Cuerpo de petición/respuesta
- Códigos de estado
- Códigos de error estables

### 7. Las reglas de negocio pertenecen al backend

Ejemplos:
- Productos inactivos no se muestran públicamente
- Productos sin stock no pueden añadirse
- Cantidad > 0
- Backend calcula totales
- Frontend no envía precios como datos autoritativos

### 8. La experiencia de usuario importa

Frontend debe proporcionar:
- Estados de carga
- Estados vacíos
- Estados de error
- Acciones deshabilitadas cuando corresponda
- Feedback claro

### 9. Requisitos de calidad

Cada funcionalidad incluye:
- Pruebas unitarias para reglas de negocio
- Pruebas de integración para endpoints
- Casos límite tratados
- Respuestas de error estables

### 10. Línea base de seguridad

- Validar todas las entradas en backend
- No confiar en valores del frontend
- No exponer errores internos
- Códigos de error estables

### 11. Línea base de observabilidad

Backend registra:
- Fallos al recuperar catálogo
- Intentos inválidos sobre carrito
- Operaciones inválidas

### 12. Architecture Decision Records

Se crean ADRs para decisiones que afecten:
- Estructura del sistema
- Persistencia
- Integración frontend-backend
- Diseño de API
- Validación
- Seguridad
- Mantenibilidad

### 13. Mantener la demo comprensible

Simple pero realista para demostrar por qué importan especificaciones y arquitectura.

## Restricciones tecnológicas

### Backend
- Lenguaje: Go
- API: HTTP/JSON
- Arquitectura: monolito modular
- Persistencia: base de datos relacional

### Frontend
- Framework: React
- Consume API del backend
- Gestiona estados: carga, vacío, éxito, error

### Persistencia
- Base de datos relacional
- SQLite o PostgreSQL
- Lógica de negocio no acoplada a SQL

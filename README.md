# Simple API Gateway in Go

Simple API Gateway desarrollado en Go para centralizar el acceso a múltiples microservicios mediante un único punto de entrada. El proyecto implementa enrutamiento de peticiones HTTP, validación de JWT, middleware de logging, manejo de errores y redirección hacia servicios internos como autenticación, inventario, pagos o usuarios.

Este proyecto tiene como objetivo demostrar el uso de Go en arquitecturas basadas en microservicios, aplicando buenas prácticas de backend, separación de responsabilidades y una estructura limpia y mantenible.

## Características principales

- API Gateway como punto único de entrada.
- Reverse proxy hacia múltiples microservicios.
- Middleware de autenticación con JWT.
- Middleware de logging de requests.
- Manejo centralizado de errores.
- Configuración mediante variables de entorno.
- Estructura basada en Clean Architecture.
- Preparado para Docker y Docker Compose.

## Tecnologías usadas

- Go
- net/http
- httputil ReverseProxy
- JWT
- Docker
- Docker Compose
- Clean Architecture
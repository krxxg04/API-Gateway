# simple-api-gateway-go

API Gateway simple en Go usando `net/http`, `httputil.ReverseProxy`, JWT y principios de Clean Architecture para separar configuración, middlewares y capa de enrutamiento/proxy.

## Características

- Punto único de entrada para microservicios.
- Rutas públicas y protegidas.
- Reverse proxy por prefijo de ruta.
- Validación JWT con `github.com/golang-jwt/jwt/v5`.
- Carga de variables de entorno con `godotenv`.
- Middleware de logging (método, ruta, status y latencia).
- Middleware de recuperación de pánicos.
- Endpoint de salud: `/health`.
- Ejecución local y con Docker Compose.

## Mapeo de rutas

- `/api/auth` -> `auth-service` (pública)
- `/api/users` -> `user-service` (protegida con JWT)
- `/api/inventory` -> `inventory-service` (protegida con JWT)
- `/api/payments` -> `payment-service` (protegida con JWT)

## Estructura

```text
simple-api-gateway-go/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── gateway/
│   │   ├── proxy.go
│   │   └── router.go
│   ├── middleware/
│   │   ├── auth_middleware.go
│   │   ├── logging_middleware.go
│   │   └── recovery_middleware.go
│   └── response/
│       └── error_response.go
├── mock-services/
│   ├── auth-service/
│   │   └── main.go
│   ├── user-service/
│   │   └── main.go
│   ├── inventory-service/
│   │   └── main.go
│   └── payment-service/
│       └── main.go
├── .env.example
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

## Variables de entorno

1. Crear archivo `.env`:

```bash
cp .env.example .env
```

2. Valores esperados:

```env
GATEWAY_PORT=8080
JWT_SECRET=dev-secret-key
AUTH_SERVICE_URL=http://localhost:8081
USER_SERVICE_URL=http://localhost:8082
INVENTORY_SERVICE_URL=http://localhost:8083
PAYMENT_SERVICE_URL=http://localhost:8084
```

## Ejecución local (Go)

Abrir 5 terminales en la carpeta del proyecto:

```bash
go run ./mock-services/auth-service
go run ./mock-services/user-service
go run ./mock-services/inventory-service
go run ./mock-services/payment-service
go run ./cmd/server
```

El gateway quedará en `http://localhost:8080`.

## Ejecución con Docker Compose

```bash
docker compose up --build
```

Gateway disponible en `http://localhost:8080`.

## Ejemplos con curl

### 1) Health check

```bash
curl -i http://localhost:8080/health
```

### 2) Login público y obtención de token simulado

```bash
curl -s -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}'
```

### 3) Consumir `/api/inventory/products` con JWT

```bash
TOKEN=$(curl -s -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}' | jq -r '.token')

curl -i http://localhost:8080/api/inventory/products \
  -H "Authorization: Bearer $TOKEN"
```

### 4) Intentar acceder sin JWT (debe responder 401)

```bash
curl -i http://localhost:8080/api/inventory/products
```

## Arquitectura aplicada

- `cmd/server`: punto de entrada de la aplicación.
- `internal/config`: carga y validación de configuración.
- `internal/gateway`: enrutamiento y construcción de reverse proxies.
- `internal/middleware`: concerns transversales (auth, logging, recovery).
- `internal/response`: respuestas de error consistentes en JSON.

## Notas

- `auth-service` firma tokens usando el mismo `JWT_SECRET` que valida el gateway.
- En Docker Compose, el gateway enruta por nombre de contenedor (`auth-service`, `user-service`, etc.).

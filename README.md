# 🏎️ F1 Tracker — Backend

API REST para la aplicación F1 Tracker, construida con **Go** y **PostgreSQL**.

🔗 **Repositorio del frontend**: https://github.com/ESJ0/f1-tracker-frontend.git  
🌐 **Aplicación en producción**: https://f1-tracker-backend.onrender.co

![screenshot](docs/screenshot.png)

---

## Tecnologías utilizadas

- **Lenguaje**: Go 1.21+
- **Router**: chi v5
- **Base de datos**: PostgreSQL
- **Librerías**: godotenv, lib/pq
- **Deploy**: Render

---

## Requisitos previos

- Go 1.21+
- PostgreSQL 14+

---

## Instalación local

### 1. Clonar el repositorio

```bash
git clone https://github.com/TU_USUARIO/f1-tracker-backend.git
cd f1-tracker-backend
```

### 2. Instalar dependencias

```bash
go mod tidy
```

### 3. Crear la base de datos

```bash
psql -U postgres
```

```sql
CREATE DATABASE f1tracker;
CREATE USER f1user WITH PASSWORD 'tupassword';
GRANT ALL PRIVILEGES ON DATABASE f1tracker TO f1user;
\q
```

### 4. Correr las migraciones

```bash
psql -U f1user -d f1tracker -f migrations/001_init.sql
```

### 5. Cargar datos de prueba

```bash
psql -U f1user -d f1tracker -f seed/001_seed.sql
```

### 6. Configurar variables de entorno

```bash
cp .env.example .env
```

Edita `.env` con tus credenciales:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=f1user
DB_PASSWORD=tupassword
DB_NAME=f1tracker
SERVER_PORT=8080
```

### 7. Correr el servidor

```bash
go run main.go
```

El servidor estará disponible en `http://localhost:8080`

---

## Endpoints de la API

### Pilotos

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/drivers` | Listar todos los pilotos |
| GET | `/drivers/:id` | Obtener piloto por ID |
| POST | `/drivers` | Crear piloto |
| PUT | `/drivers/:id` | Actualizar piloto |
| DELETE | `/drivers/:id` | Eliminar piloto |
| GET | `/drivers/:id/results` | Resultados de un piloto |

Parámetros disponibles en `GET /drivers`: `?q=`, `?sort=`, `?order=`, `?page=`, `?limit=`

### Carreras

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/races` | Listar todas las carreras |
| GET | `/races/:id` | Obtener carrera por ID |
| POST | `/races` | Crear carrera |
| PUT | `/races/:id` | Actualizar carrera |
| DELETE | `/races/:id` | Eliminar carrera |
| GET | `/races/:id/results` | Resultados de una carrera |

Parámetros disponibles en `GET /races`: `?q=`, `?sort=`, `?order=`, `?page=`, `?limit=`

### Resultados

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| POST | `/results` | Registrar resultado |
| DELETE | `/results/:id` | Eliminar resultado |

---

## CORS

Esta es una política de seguridad del navegador que bloquea peticiones entre orígenes distintos. Como el frontend y el backend corren en dominios diferentes, el servidor debe permitir explícitamente esas peticiones.

Configuración aplicada:

```
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Content-Type
```

---

## Challenges implementados

- ✅ Códigos HTTP correctos (201 al crear, 204 al eliminar, 404 si no existe, 400 en input inválido, 409 en conflicto)
- ✅ Validación server-side con respuestas de error en JSON descriptivas
- ✅ Paginación con `?page=` y `?limit=`
- ✅ Búsqueda por nombre con `?q=`
- ✅ Ordenamiento con `?sort=` y `?order=asc|desc`

---

## Reflexión

En este proyecto utilicé Go, chi y PostgreSQL para construir una API REST. Me gustó trabajar con Go porque muchos errores se detectan antes de ejecutar el programa gracias al tipado estático. También aprendí más sobre manejo de bases de datos, rutas REST y conexión entre frontend y backend.

Además, el proyecto me ayudó a entender temas importantes como CORS, migraciones, seeds y despliegue de aplicaciones. Si volviera a hacer algo similar, usaría una herramienta de migraciones como golang-migrate para manejar la base de datos de una forma más práctica.
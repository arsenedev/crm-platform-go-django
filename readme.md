# CRM Platform

A modern CRM platform built with **Go**, **Django**, **PostgreSQL**, **Nginx**, and **Docker Compose**.

---

# Overview

The platform consists of three main components:

- Customer API (Go)
- Customer Admin (Django)
- PostgreSQL

Nginx acts as a reverse proxy and exposes the application.

---

# Architecture

```text
                    ┌───────────────┐
                    │     Nginx     │
                    │  Ports 80/443 │
                    └───────┬───────┘
                            │
                  ┌─────────┴─────────┐
                  │                   │
                  ▼                   ▼

          Customer Admin         Customer API
              Django                 Go
             Port 8000           Port 8080

                         │
                         ▼

                    PostgreSQL
                     Port 5432
```

---

# Features

## Authentication

- JWT Authentication
- Session-Based Login
- Role-Based Access Control (RBAC)

## User Management

- Create Users
- Update Users
- Delete Users
- Role Management

## Customer Management

- Create Customers
- Update Customers
- Delete Customers

## Contact Management

- Create Contacts
- Update Contacts
- Delete Contacts

## Audit Logs

- User Activity Tracking
- Resource Change History
- Security Auditing

---

# Technology Stack

## Backend

- Go
- Gin
- PostgreSQL
- JWT

## Frontend

- Django
- Bootstrap 5
- Gunicorn

## Infrastructure

- Docker
- Docker Compose
- Nginx

---

# Project Structure

```text
.
├── customer-api/
│   ├── cmd/
│   ├── internal/
│   ├── Dockerfile
│   └── .env.example
│
├── customer-frontend/
│   ├── authentication/
│   ├── users/
│   ├── customers/
│   ├── contacts/
│   ├── audit/
│   ├── dashboard/
│   ├── templates/
│   ├── Dockerfile
│   └── .env.example
│
├── database/
│   └── init/
│       ├── 001_user-table.sql
│       ├── 002_customer-table.sql
│       ├── 003_contact-table.sql
│       └── 004_audit-table.sql
│
├── nginx/
│   └── conf.d/
│       └── crm.conf
│
├── docker-compose.yml
├── .env.example
└── README.md
```

---

# Environment Variables

Create a `.env` file at the project root.

```env
# PostgreSQL

POSTGRES_DB=postgres
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgresql16

POSTGRES_HOST=postgres
POSTGRES_PORT=5432

# JWT

JWT_SECRET=replace-with-secure-secret

# Initial Administrator

ADMIN_EMAIL=admin@company.com
ADMIN_PASSWORD=change-me

# Django

DJANGO_SECRET_KEY=replace-me

DJANGO_DEBUG=False

DJANGO_ALLOWED_HOSTS=localhost,127.0.0.1

# API URL

CRM_API_URL=http://customer-api:8080

# Gin

GIN_MODE=release
```

---

# Security

Generate secure secrets before production deployment.

## Django Secret Key

```bash
python -c "from django.core.management.utils import get_random_secret_key; print(get_random_secret_key())"
```

## JWT Secret

```bash
openssl rand -hex 32
```

## PostgreSQL Password

```bash
openssl rand -base64 32
```

Never commit:

```text
.env
.env.local
.env.production
```

Add to `.gitignore`:

```gitignore
.env
.env.*
!.env.example
```

---

# Database Initialization

PostgreSQL automatically executes all SQL files stored in:

```text
database/init/
```

during the first initialization.

```text
001_user-table.sql
002_customer-table.sql
003_contact-table.sql
004_audit-table.sql
```

To recreate the database:

```bash
docker compose down -v

docker compose up --build
```

---

# Build and Run

Build containers:

```bash
docker compose build
```

Start all services:

```bash
docker compose up
```

Start in detached mode:

```bash
docker compose up -d
```

Stop services:

```bash
docker compose down
```

Remove services and database volume:

```bash
docker compose down -v
```

---

# View Logs

All services:

```bash
docker compose logs -f
```

API:

```bash
docker compose logs -f customer-api
```

Django:

```bash
docker compose logs -f customer-admin
```

PostgreSQL:

```bash
docker compose logs -f postgres
```

Nginx:

```bash
docker compose logs -f nginx
```

---

# Application URLs

## Customer Admin

```text
http://localhost
```

or

```text
http://localhost:8000
```

depending on your Nginx configuration.

---

## Customer API

```text
http://localhost/api/
```

Internal Docker URL:

```text
http://customer-api:8080
```

---

# PostgreSQL Access

```text
Host: localhost
Port: 5432

Database: postgres
User: postgres
```

Connect manually:

```bash
docker exec -it crm-postgres psql -U postgres
```

---

# API Endpoints

## Authentication

```http
POST /auth/login
```

## Users

```http
GET    /api/v1/users
POST   /api/v1/users
PUT    /api/v1/users/{id}
DELETE /api/v1/users/{id}
```

## Customers

```http
GET    /api/v1/customers
POST   /api/v1/customers
PUT    /api/v1/customers/{id}
DELETE /api/v1/customers/{id}
```

## Contacts

```http
GET    /api/v1/contacts
POST   /api/v1/contacts
PUT    /api/v1/contacts/{id}
DELETE /api/v1/contacts/{id}
```

## Audit Logs

```http
GET /api/v1/audit
```

---

# Local Development

### Start PostgreSQL

```bash
docker compose up postgres
```

### Start API

```bash
cd customer-api

go run cmd/server/main.go
```

### Start Django

```bash
cd customer-frontend

python manage.py migrate

python manage.py runserver
```

Django will be available at:

```text
http://127.0.0.1:8000
```

---

# Production Notes

## Gin

Use release mode:

```env
GIN_MODE=release
```

## Gunicorn

Run Django using:

```bash
gunicorn \
config.wsgi:application \
--bind 0.0.0.0:8000 \
--workers 4 \
--timeout 120
```

## Nginx

Nginx acts as the public entry point and reverse proxy.

Only expose:

```text
80/tcp
443/tcp
```

Keep internal services private:

```text
5432/tcp
8080/tcp
8000/tcp

```
---
# Deployment

## Production Requirements

The platform requires:

- Docker 24+
- Docker Compose v2+
- Linux Server (Ubuntu 22.04 LTS recommended)
- Minimum 2 GB RAM
- Minimum 2 CPU cores

---

## Server Preparation

Update the server:

```bash
sudo apt update
sudo apt upgrade -y
```

Install Docker:

```bash
curl -fsSL https://get.docker.com | sh
```

Verify installation:

```bash
docker --version
docker compose version
```

---

## Clone Repository

```bash
git clone https://github.com/<your-username>/crm-platform.git

cd crm-platform
```

---

## Create Environment File

Create the `.env` file:

```bash
cp .env.example .env
```

Edit it:

```bash
nano .env
```

Example:

```env
POSTGRES_DB=postgres
POSTGRES_USER=postgres
POSTGRES_PASSWORD=change-me

POSTGRES_HOST=postgres
POSTGRES_PORT=5432

JWT_SECRET=replace-with-secure-secret

ADMIN_EMAIL=admin@company.com
ADMIN_PASSWORD=change-me

DJANGO_SECRET_KEY=replace-with-secure-secret

DJANGO_DEBUG=False

DJANGO_ALLOWED_HOSTS=crm.example.com

CRM_API_URL=http://customer-api:8080

GIN_MODE=release
```

---

## Build Containers

```bash
docker compose build
```

---

## Start Application

```bash
docker compose up -d
```

Verify:

```bash
docker compose ps
```

Expected services:

```text
crm-postgres
customer-api
customer-admin
nginx
```

---

## Check Logs

All services:

```bash
docker compose logs -f
```

API logs:

```bash
docker compose logs -f customer-api
```

Django logs:

```bash
docker compose logs -f customer-admin
```

PostgreSQL logs:

```bash
docker compose logs -f postgres
```

Nginx logs:

```bash
docker compose logs -f nginx
```

---

## Verify Database

Connect to PostgreSQL:

```bash
docker exec -it crm-postgres psql -U postgres
```

List tables:

```sql
\dt
```

Expected tables:

```text
users
customers
contacts
audit_logs
```

---

## Verify Services

### CRM Admin

```text
http://localhost
```

or

```text
http://<server-ip>
```

### API

Internal Docker URL:

```text
http://customer-api:8080
```

---

## Upgrade Application

Pull latest changes:

```bash
git pull
```

Rebuild:

```bash
docker compose up -d --build
```

---

## Database Reset

WARNING: This operation removes all application data.

Stop and remove containers, networks and volumes:

```bash
docker compose down -v
```

Recreate everything:

```bash
docker compose up --build
```

PostgreSQL will automatically execute:

```text
database/init/
├── 001_user-table.sql
├── 002_customer-table.sql
├── 003_contact-table.sql
└── 004_audit-table.sql
```

---

## Backup Database

Create backup:

```bash
docker exec crm-postgres \
pg_dump -U postgres postgres \
> backup.sql
```

---

## Restore Database

```bash
docker exec -i crm-postgres \
psql -U postgres postgres \
< backup.sql
```

---

## Recommended Production Architecture

```text
Internet
    │
    ▼

  Nginx
 80 / 443

    │

 ┌──┴──────────────┐
 │                 │

 ▼                 ▼

Django Admin    Customer API
    8000            8080

         │

         ▼

     PostgreSQL
        5432
```

Recommended exposed ports:

```text
80/tcp
443/tcp
```

Internal-only services:

```text
8000/tcp
8080/tcp
5432/tcp
```

---

## Health Check Commands

Verify PostgreSQL:

```bash
docker exec crm-postgres \
pg_isready -U postgres
```

Verify API:

```bash
curl http://localhost/api/v1/users
```

Verify Admin:

```bash
curl http://localhost/login/
```
# License

MIT

---

# Author

**Thierry ZONGO**

Cloud Engineer • Kubernetes Administrator (CKA) • DevOps Engineer • Full Stack Developer

Burkina Faso 🇧🇫

# Login Form Application

Una aplicación web simple de formulario de login construida en Go que captura credenciales de usuario y los muestra en el log del contenedor para poder auditar los intentos/ataques a los que está expuesta.

## 🚀 Características

- Servidor web HTTP (puerto 80) y HTTPS (puerto 443)
- Interfaz de login con HTML/CSS integrado
- Captura y registra credenciales en logs
- Certificados SSL autofirmados
- Health check endpoint
- Imagen Docker mínima basada en `scratch`
- CI/CD automático con GitHub Actions

## 📋 Requisitos

- Go 1.22 o superior
- Docker (opcional)
- OpenSSL (para generar certificados)

## 🔧 Instalación y Uso

### Ejecución Local

1. **Generar certificados SSL:**

   ```bash
   chmod +x gen-cert.sh
   ./gen-cert.sh
   ```

2. **Compilar la aplicación:**

   ```bash
   go build -o loginapp main.go
   ```

3. **Ejecutar:**

   ```bash
   ./loginapp
   ```

4. **Acceder a la aplicación:**
   - HTTP: http://localhost
   - HTTPS: https://localhost

### Usando Docker

#### Pull desde GitHub Container Registry

```bash
docker pull ghcr.io/jjavierrg/loggin-form:latest
```

#### Ejecutar el contenedor

```bash
docker run -d -p 80:80 -p 443:443 --name login-app ghcr.io/jjavierrg/loggin-form:latest
```

#### Construcción local (desarrollo)

```bash
# Generar certificados
./gen-cert.sh

# Compilar binario
CGO_ENABLED=0 go build -ldflags="-s -w" -o loginapp main.go

# Construir imagen
docker build -t login-app .

# Ejecutar
docker run -d -p 80:80 -p 443:443 --name login-app login-app
```

## 🔍 Health Check

La aplicación incluye un endpoint de health check:

```bash
curl http://localhost/health
```

O usando el binario:

```bash
./loginapp health
```

## 🔄 CI/CD Pipeline

El proyecto utiliza GitHub Actions para automatizar el proceso de compilación y publicación:

### Pipeline de Build y Push

El workflow `.github/workflows/build-and-push.yml` se ejecuta automáticamente en:

- Push a `main` o `master`
- Pull requests
- Ejecución manual (workflow_dispatch)

**Proceso:**

1. ✅ Checkout del código
2. ✅ Configuración de Go 1.22
3. ✅ Generación de certificados SSL
4. ✅ Compilación del binario Go
5. ✅ Login en GitHub Container Registry
6. ✅ Build y push de la imagen Docker

### Tags Generados

La pipeline genera automáticamente los siguientes tags:

- `latest` (solo en rama principal)
- `<branch-name>` (para cada rama)
- `<branch>-<sha>` (commit específico)
- Tags semánticos si se usan versiones

### Permisos Requeridos

El workflow necesita:

- `contents: read` - Para leer el código
- `packages: write` - Para publicar en GHCR

## 📦 Estructura del Proyecto

```
.
├── .github/
│   └── workflows/
│       └── build-and-push.yml    # Pipeline de CI/CD
├── Dockerfile                     # Imagen Docker mínima
├── gen-cert.sh                    # Script para generar certificados SSL
├── main.go                        # Aplicación principal
└── README.md                      # Este archivo
```

## 🐳 Dockerfile

El Dockerfile utiliza una imagen base `scratch` (sin sistema operativo) para crear una imagen extremadamente ligera:

- **Tamaño:** ~6-8 MB
- **Seguridad:** Superficie de ataque mínima
- **Rendimiento:** Inicio instantáneo

**Nota:** El build se realiza en la pipeline de CI/CD, no en el Dockerfile.

## 📝 Logs

Los logs capturan las credenciales enviadas:

```
LOGIN --> user='admin' password='pass123' ip='172.17.0.1'
```

**⚠️ ADVERTENCIA:** Esta aplicación es solo para fines educativos. No usar en producción sin las medidas de seguridad adecuadas.

## 🔒 Seguridad

- Los certificados SSL son autofirmados (no válidos para producción)
- Las credenciales se registran en texto plano
- No hay validación real de usuarios
- Diseñado para pruebas y demostración

## 🛠️ Desarrollo

### Hacer cambios

1. Clonar el repositorio
2. Realizar cambios en `main.go`
3. Hacer commit y push
4. La pipeline automáticamente construirá y publicará la nueva imagen

### Variables de Entorno

Actualmente la aplicación no requiere variables de entorno, pero los puertos están hardcodeados:

- HTTP: `80`
- HTTPS: `443`

## 📄 Licencia

Este proyecto es de código abierto y está disponible para fines educativos.

## 🤝 Contribuir

Las contribuciones son bienvenidas. Por favor:

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/amazing-feature`)
3. Commit tus cambios (`git commit -m 'Add amazing feature'`)
4. Push a la rama (`git push origin feature/amazing-feature`)
5. Abre un Pull Request

## 📞 Soporte

Para reportar problemas o sugerencias, por favor abre un issue en GitHub.

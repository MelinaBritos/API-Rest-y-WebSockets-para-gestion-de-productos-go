#  API REST y Websocket para gestión de productos.

[![Typing SVG](https://readme-typing-svg.demolab.com/?lines=Bsmart+Challenge!&color=FF00FF)](https://git.io/typing-svg)

> API RESTful y sistema de comunicación en tiempo real utilizando WebSockets para la gestión de productos y categorías. Programado en Golang y postgresSQL.

---

## 💡 1. Caracteristicas Principales

* **Autenticación y Autorización:** Implementada con **JWT** y proteccion de rutas segun roles (`admin`/`cliente`).
* **API RESTful:** Endpoints para gestión completa de **productos** y **categorías** (CRUD).
* **WebSockets:** Canal de comunicación bidireccional para notificaciones instantáneas (creacion, actualizacion y eliminacion de productos/categorias).
* **Persistencia:** Conexión a base de datos **PostgreSQL**.

---

## 🛠️ 2. Instrucciones de Instalación y Configuración

### 2.1. Prerrequisitos

Necesitas tener instalado en tu sistema:
* **Go:** Versión 1.20 o superior.
* **PostgreSQL:** Servidor de base de datos.
* **Git:** Para la clonación del repositorio.

### 2.2. Puesta en Marcha

1.  **Clona el repositorio** 

2.  **Configura el Entorno:**
    Crea un archivo llamado **`.env`** en la raíz del proyecto, basado en el `.env.example`, y rellena tus credenciales.

3.  **Instala Dependencias:**
    ```bash
    go mod tidy
    ```

4.  **Ejecuta la Aplicación:**
    ```bash
    go run main.go

## 3. Deployment y Acceso en la Nube
La API se encuentra hosteada para demostración

URL de conexion: https://api-rest-y-websockets-para-gestion-de.onrender.com

## 4. WebSockets: Endpoints y Formato
URL de Conexión: ws://api-rest-y-websockets-para-gestion-de.onrender.com/api/ws (puede tardar unos segundos)
 
### Ejemplo de Mensaje Recibido

```json
{
  "action": "Se creó un nuevo producto",
  "data": {
    "id": 25,
    "name": "Lampara Led",
    "description": "Luz Fenergy Velador 7w A Batería Flexible Calido, Neutro, Frio Dimerizable",
    "price": 14000,
    "stock": 70,
    "created_at": "2025-10-29T03:08:10.478211297Z",
    "updated_at": "2025-10-29T03:08:10.478211397Z"
  },
  "event": "PRODUCT_CREATED"
}
``` 

## 5. Ejemplos de Uso (Endpoints Clave)
| Método | Ruta | Requerimientos | Descripción |
| :--- | :--- | :--- | :--- |
| **POST** | `/api/login` | *Ninguno* | Inicia sesión y retorna el token JWT. |
| **GET** | `/api/products` | Auth | Listado de todos los productos. |
| **POST** | `/api/products` | Auth + **Admin** | Crea un nuevo producto. |
| **DELETE** | `/api/products/{id}` | Auth + **Admin** | Elimina un producto por ID. |
  
## 6. Documentación
En el siguiente enlace encontraran la documentacion detallada de los endpoints disponibles:
https://docs.google.com/document/d/1OizQNqU57X80ebuyWy4j4le7-wmRIKg1tO8YcuHYnj8/edit?tab=t.0

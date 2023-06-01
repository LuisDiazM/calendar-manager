## Definición caso de uso de la aplicación
* **Programación de reuniones**

    Los usuarios pueden crear nuevas reuniones virtuales especificando la fecha, la hora, la duración prevista y los invitados a la reunión.
    Pueden invitar a participantes introduciendo direcciones de correo electrónico.
    El formato de la reunión es videollamada por zoom.
* **Notificaciones a los usuarios**

    El sistema envía notificaciones automáticas a todos los participantes para informarles sobre la reunión programada, incluyendo detalles como la hora, la plataforma de reunión. Las notificaciones serán por correo electrónico o en su versión más simple sólo guardar el registro del evento en mongo.

## Diseño del modelo de datos
Cada microservicio tiene su modelo de datos,
- **CalendarEventManager** usa posgresql como motor de bd
- **CalendarNotificationManager** usa mongodb como motor de bd

![modelo](/docs/CalendarNotification-Modelo%20Datos.jpg)

## Diseño de la solución por microservicios
De manera simple (sin tener en cuenta capas de infraestructura nube),
La aplicación se compone de un front escrito en angular, un microservicio que expone un servidor HTTP para manejar las peticiones REST desde el front, un microservicio consumidor de eventos mediante un broker de mensajería.

![microservicios](/docs/CalendarNotification-Page-3.jpg)

## Definición del front

La experiencia de usuario se creó usando Angular Material.

### SSO de google para el registro de usuarios

Cuando un usuario hace clic en el botón de inicio de sesión en la aplicación Angular, se iniciará el flujo de autenticación de Google.

El usuario será redirigido a la página de inicio de sesión de Google para autenticarse.
Después de la autenticación exitosa, Google redirigirá al usuario de vuelta a tu aplicación Angular con un token de acceso.
En el lado del frontend.

Para las peticiones REST se debe enviar en los headers ``
Authorization: Bearer <token>``

### Módulos
1. Authenticación
Login de la aplicación con SSO google
/auth

2. Usuarios
Lee la información del SSO de google y le da la posibilidad al usuario de registrarse para empezar a crear eventos

3. Reuniones
De acuerdo al usuario registrado permite realizar el CRUD completo de reuniones, donde permite seleccionar la fecha, hora, duración de la reunión, los correos de los participantes para enviar la información de la reunión.


## Definición de las APIs
### Apis públicas
* GET /health   → status de la aplicación

### Apis privadas --> group: "/api/v1"
* POST 		/user
* GET 		/user/:id
* PUT		/user/:id

* POST 		/meetings
* GET 		/meeting/:id
* GET 		/meetings/user/:id
* UPDATE	/meeting/:id
* DELETE	/meeting/:id

## Contenedores Docker
Se tiene en total 8 contenedores agrupados de la siguiente forma:

* Interfaz de usuario

El proyecto de angular donde un servidor nginx se encarga de servir la aplicación para que el cliente pueda accederla de manera local http://localhost:4200 

* Bases de datos:

  * MongoDB (NoSQL)

  Mongo-Express (Gestor de base de datos para Mongo) → http://localhost:8081 

        user: luismigueldiazmorales
        password: fjalkshfklahs

  * Postgresql (SQL)

PgAdmin (Gestor de base de datos para postgres) http://localhost:8083

	    user: luismigueldiazmorales@gmail.com
	    password: 1232451fafr8

	Nota: para conectarse al servidor local en el host colocar  host.docker.internal

* Broker de mensajería

RabbitMQ (La imagen usada tiene el administrador, por lo tanto se puede usar en la maquina local http://localhost:8082)

	user: user
	password: gwohnfgj45*5

* Microservicio con servidor Web

Calendar Event Manager (se encarga de manejar todas las peticiones REST del sistema)
	

* Microservicio Consumidor de eventos

Calendar Notifications (se encarga de escuchar los eventos de crear reunión, en esta primera versión sólo escucha un evento “registryMeeting”, sin embargo, el código quedó estructurado para registrar muchos más eventos)

## Instrucciones para ejecución local
Requisitos:
Tener Docker y Docker compose instalado


**Pasos**
1. ``git clone https://github.com/LuisDiazM/calendar-manager.git``
2. ``cd calendar-manager``
3. ``docker-compose up --build -d``


*Nota*

El paso 3. sólo se ejecuta la primera vez para crear los contenedores, una vez estén creados la proxima vez para levantar los contenedores se usa 
``docker-compose up -d``

Como buena prácitca al final de usar el proyecto ``docker-compose down`` para remover estos contenedores, ya que en windows si no se remueven al iniciar la máquina se empiezan a ejecutar de nuevo consumiendo recursos innecesarios.

Algunas versiones de docker compose se usa sin el guion ``docker compose up -d``

Completados los pasos se debe esperar hasta que todos los contenedores estén arriba y se puede acceder a la aplicación http://localhost:4200 

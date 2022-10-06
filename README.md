# PersonalInfoMs

_Microservicio especializado en la gestion de la información personal de cada individuo perteneciente a una institución de educación_

## Comenzando 🚀

_Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu máquina local para propósitos de desarrollo y pruebas._

### Pre-requisitos 📋

_Necesitamos tener instalado Docker_

## Despliegue 📦

_Para desplegar el contenedor con Docker ejecutamos los siguientes comandos:_

```
docker build -t personal_info_ms .
docker run -it -d -p 3000:3000 -v $PWD:/app --name personalInfoMs personal_info_ms
```

## Autores ✒️

* **Cristian Vargas Morales** - *Developer* - [crvargas](https://github.com/crvargasm)

## Licencia 📄

Este proyecto está bajo la Licencia MIT - mira el archivo [LICENSE.md](LICENSE.md) para detalles

---
⌨️ con ❤️ por [crvargasm](https://github.com/crvargasm) 😊

# Crud utilizando Golang, MySQL, & Docker

---

## Comandos Docker

```bash
# comando para listar las imagenes
docker imagenes

# listar las maquinas
docker-mackine ls

# instalar MySQL
# sudo <optional>
# latest: se lo puede cambiar por una versión en especifico. Se recomienda siempre seleccionar una version para evitar errores con la compatibilidad de versiones
sudo docker pull mysql/mysql-server:latest
```

## Instalar driver de conexión a MySQL

```bash
go get -u github.com/go-sql-driver/mysql
```

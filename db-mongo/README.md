# Descripción

Proyecto para administrar la conexion a la BD

## Obtener todos los paquetes

go get -d -u gitlab.com/woh-group/woh-backend/db-mongo/...

## Crear modulo proyecto

go mod init gitlab.com/woh-group/woh-backend/db-mongo

## Parametros de filtro

Es una libreria que mejora el filtro de una api rest a una base de datos mongoDB.

Estos son los tipos de filtros validos:

* **contains**: devuleve todos los registros cuyo campo contenga XXX letra, palabra o frase.
* **startswith**: devuleve todos los registros cuyo campo comience con XXX letra, palabra o frase.
* **endswith**: devuleve todos los registros cuyo campo termine con XXX letra, palabra o frase.
* **between**: devuleve todos los registros cuyo campo estre entre XXX y YYY.
* **ne**: devuleve todos los registros cuyo campo sea diferente de XXX.
* **gt**: devuleve todos los registros cuyo campo sea mayor que XXX.
* **lt**: devuleve todos los registros cuyo campo sea menor que XXX.
* **gte**: devuleve todos los registros cuyo campo sea mayor o igual que XXX.
* **lte**: devuleve todos los registros cuyo campo sea menor o igual que XXX.

### Ejemplos

**Contains**: devuleve todos los registros cuyo campo contenga XXX letra, palabra o frase

```sh
{url}/api/getAll?field[contains]=palabra
{url}/api/getAll?field[contains]=palabra1,palabra2,palabra3
```

**startswith**: devuleve todos los registros cuyo campo comience con XXX letra, palabra o frase

```sh
{url}/api/getAll?field[startswith]=palabra
{url}/api/getAll?field[startswith]=palabra1,palabra2,palabra3
```

**endswith**: devuleve todos los registros cuyo campo termine con XXX letra, palabra o frase

```sh
{url}/api/getAll?field[endswith]=palabra
{url}/api/getAll?field[endswith]=palabra1,palabra2,palabra3
```

**between**: devuleve todos los registros cuyo campo estre entre XXX y YYY

```sh
{url}/api/getAll?field[between]=1,3
{url}/api/getAll?field[between]=2019-05-18T00:00:00,2019-05-18T23:59:59
```

**ne**: devuleve todos los registros cuyo campo sea diferente de XXX

```sh
{url}/api/getAll?field[ne]=1
{url}/api/getAll?field[ne]=1,2
```

**gt**: devuleve todos los registros cuyo campo sea mayor que XXX

```sh
{url}/api/getAll?field[gt]=1
```

**lt**: devuleve todos los registros cuyo campo sea menor que XXX

```sh
{url}/api/getAll?field[lt]=3
```

**gte**: devuleve todos los registros cuyo campo sea mayor o igual que XXX

```sh
{url}/api/getAll?field[gte]=1
```

**lte**: devuleve todos los registros cuyo campo sea menor o igual que XXX

```sh
{url}/api/getAll?field[lte]=3
```

**Create by:** Leonardo Solano Arguello

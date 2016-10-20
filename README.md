Yangee microservicio v0.1
=========================

(basado en )

Golang, mongodb, gorilla, jwt-middleware

Usar Mongodb_Address: "localhost:27017" en caso que se desee montar una base de datos en mongo propia

API: "localhost:3113"

Listar todos:

		# curl -i http://localhost:3113/api/todos

Show a single todo (replace {id} for the equivalent bson.ObjectIdHex):

		# curl -i http://localhost:3113/api/todos/{id}

Add todo:

		# curl -i -H "Content-Type: application/json" -X POST -d '{"name": "Task 14", "completed": false}'  http://localhost:3113/api/todos

		or

		# curl -i http://localhost:3113/api/todos -X POST -d @add.json

where add.json file is something like:

		{
			"name":   "Task 14",
			"completed":   false
		}

Update todo (replace {id} for the equivalent bson.ObjectIdHex):

		#  curl -i -H "Content-Type: application/json" -X PUT -d '{"name": "update task", "completed": false}'  http://localhost:3113/api/todos/{id}

		or

		# curl -i http://localhost:3113/api/todos/{id} -X PUT -d @update.json

where update.json file is something like:

		{
			"name":   "Task X",
			"completed":   true
		}

Delete todo:

		# curl -i http://localhost:3113/api/todos/{id} -X DELETE

Search todo by name (replace {name} for the equivalent search pattern):

		# curl -i http://localhost:3113/api/todos/search/byname/{name}

Search todo by status completed (replace {status} for true or false ):

		# curl -i http://localhost:3113/api/todos/search/bystatus/{status}

Log samples:

		2016/05/29 22:50:47 192.168.0.100:46340	GET	/api/todos	HTTP/1.1	200	815	962.763µs
		2016/05/29 22:51:06 192.168.0.100:46341	GET	/api/todos/574b4b92e561770001514888	HTTP/1.1	200	137	37.192966ms
		2016/05/29 22:51:53 192.168.0.100:46342	POST	/api/todos	HTTP/1.1	201	0	624.235µs
		2016/05/29 22:51:58 192.168.0.100:46343	GET	/api/todos	HTTP/1.1	200	979	799.181µs
		2016/05/29 22:53:37 192.168.0.100:46344	PUT	/api/todos5/74baac9cdc87225dc493c0b	HTTP/1.1	404	0	0
		2016/05/29 22:53:49 192.168.0.100:46345	PUT	/api/todos/574baac9cdc87225dc493c0b	HTTP/1.1	204	0	723.633µs
		2016/05/29 22:53:55 192.168.0.100:46346	GET	/api/todos	HTTP/1.1	200	982	610.816µs
		2016/05/29 22:54:25 192.168.0.100:46349	DELETE	/api/todos/574baac9cdc87225dc493c0b	HTTP/1.1	204	0	701.403µs
		2016/05/29 22:54:27 192.168.0.100:46350	GET	/api/todos	HTTP/1.1	200	815	615.476µs
		2016/05/29 22:55:01 192.168.0.100:46351	GET	/api/todos/search/byname/5	HTTP/1.1	200	163	579.613µs
		2016/05/29 22:55:05 192.168.0.100:46352	GET	/api/todos/search/byname/X	HTTP/1.1	200	166	569.061µs
		2016/05/29 22:56:51 192.168.0.100:46354	GET	/api/todos/search/byname/Tas	HTTP/1.1	200	815	684.036µs
		2016/05/29 22:57:37 192.168.0.100:46357	GET	/api/todos/search/bystatus/true	HTTP/1.1	200	163	616.616µs
		2016/05/29 22:57:41 192.168.0.100:46358	GET	/api/todos/search/bystatus/false	HTTP/1.1	200	654	623.589µs

Prueba de uso de token:

		# curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2p3dC1pZHAuZXhhbXBsZS5jb20iLCJzdWIiOiJhZHJpYW4uZGlhc2RhY29zdGFsaW1hQGdtYWlsLmNvbSIsIm5iZiI6MTQ3NTY3MDc2OSwiZXhwIjoxNDc1Njc0MzY5LCJpYXQiOjE0NzU2NzA3NjksImp0aSI6ImlkMTIzNDU2IiwidHlwIjoiaHR0cHM6Ly9leGFtcGxlLmNvbS9yZWdpc3RlciJ9.0p0BA2YzbpP1VxpckDUdLE4v86eir92ETH-SB4nThgI" http://localhost:3113/secured/ping

Agregar permisos:

		# curl -i -H "Content-Type: application/json" -X POST -d '{"permiso": "AgregarUsuario"}'  http://localhost:3113/api/permisos

Para loguear a la api estoy usando el plugin HttpRequester (firefox). Ejemplo de PUT ejemplos:

		METHOD: PUT
		URL: localhost:3113/api/ejemplos/5807af7041586016ef21b2e4
		CONTENT TO SEND: {"nombre":"ejemplo"}
		HEADERS: Name: Authorization Value: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NzY5NzQyOTEsImlkIjoiNTgwMTNlYzE0MTU4NjAxNWI3MDI4OTViIiwiaXNzIjoieWFuZ2VlYXBwQGdtYWlsLmNvbSJ9.b2rfvTOjKW2PwqXzTnW48XLEXZm6YXJGcdoZRtLB7-U

Donde "5807af7041586016ef21b2e4" es el ID del ejemplo a buscar. Notese que ese token es valido por solo una hora y está expirado. Utilizar /api/login con un usuario válido para obtener otro token.

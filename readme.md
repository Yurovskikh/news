# News microservice
## Start
Для запуска необходим устоновленный docker и docker-compose 
Выполнить```make start``` для запуска
## Example

Создать новость
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"header":"xyz"}' \
  http://localhost:8080/api/v1/news
````
Получить новость по идентификатору
```
curl --request GET \
   http://localhost:8080/api/v1/news/{id}
```  
  

## Todo
* Клиентская библиотека над nats streaming
* Multistage build
* Покрыть тестами
## Rotina de integração Sabium - Portal Fatura
> Busca faturas que foram baixadas no Sabium e muda o status para **Faturado** no Portal Fatura.

### Como executar
> Execute o script.
```shell
$ ./run.sh
```
> Dispare uma requisição **GET** na rota raiz **"/"** para iniciar a rotina.
```shell
$ curl http://localhost:3000/
```

### Variáveis de ambiente
- [x] DB_HOST 
- [x] DB_NAME
- [x] DB_USER 
- [x] DB_PASS 
- [x] DB_PORT
- [x] DB_SSLMODE
# rate-limiter

Suba o Redis com docker-compose up.

Rode o servidor com go run main.go.

chame a rota: localhost:8080/ e realize os testes.

Passe o header API_KEY e vc terá sucesso por 10 requests, após isso a app terá um tempo de descanso :)

NÃO Passe o header API_KEY e vc terá sucesso por 5 requests, após isso a app terá um tempo de descanso :)
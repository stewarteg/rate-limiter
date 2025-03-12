# rate-limiter

Suba o Redis com comando: docker-compose up

Rode a app com o comando: go run main.go

chame a rota: localhost:8080/ e realize os testes abaixo:

Ao passar o header "API_KEY" com um TOKEN JWT como por ex: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyZXF1ZXN0X2xpbWl0IjoiMTAifQ.ELoXRhil88grvYHOXeqk69XCsioGiajjM_mi95RSi_4
vc terá sucesso em 10 requests seguidas, após as 10 a app terá um tempo de descanso que é de BLOCK_TIME=300)

Para alterar o token use o site https://jwt.io/, com isso voce consegue editar o valor "request_limit" de dentro do token tirando de 10 para o valor que voce quiser. assim o limite de request passa a ser ele.

Ao NÃO passar o header "API_KEY" vc entra no fluxo de block por IP, nesse caso terá sucesso em 5 requests seguidas, após esse as 5 a app terá um tempo de descanso que é de BLOCK_TIME=300)
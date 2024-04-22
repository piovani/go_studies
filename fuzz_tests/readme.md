São testes que alteram o tipo entra durante execução para cobrir possiveis erros de entrada

importante observar:
ele irá ficar tesntando infinitamente até encontrar um erro ou uma interrupção manual, para isso colocamos a flag
"-fuzztime" para limitar o tempo de execução.

Para roda os testes:
```
go test -fuzz -fuzztime:30 . -v
```
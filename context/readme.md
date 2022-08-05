# Context

Facilita a passagem de valores com escopo de solicitação, sinais de cancelamento e prazos entre os limites da API para todas as goroutines envolvidas no tratamento de uma solicitação.

É uma interface que definido das seguinte forma:

```
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

<b>Done() <- struct{}</b> - retorna um canal que é fechado quando o contexto é cancelado ou expierada. Done pode retornar nil se o contexto nunca puder ser cancelado.

<b>Deadline() (deadline time.Time, ok bool)</b> - retorna a hora em que o contexto será cancelado ou expirado. O prazo retorna ok como false quando nenhum prazo é definido

<b>Err() (error)</b> - retorna um erro que explica por que o canal Done foi fechado. Se Done ainda não estiver fechado, ele retorná nil.

<b>Value(key any)</b> - any retorna o valor associado à chave ou nil se nnehum.

## Funções

### CancelFunc

A CancelFuncdiz a uma operação para abandonar seu trabalho e não espera que o trabalho pare. Se for chamado por várias goroutines simultaneamente, após a primeira chamada, as chamadas subsequentes para um CancelFuncnão fazem nada.

referencia
https://dev.to/karanpratapsingh/go-course-context-32oh

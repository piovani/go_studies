package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := todo()

	arrayCtx := []context.Context{
		background(),
		ctx,
		withValue(ctx, "key", "value"),
		withTimeout(ctx),
	}

	for _, c := range arrayCtx {
		fmt.Println(c)
	}
}

/**
retorna um contexto difernete de nil, que nunca é cancelado, não tem valores e não tem prazo

Normalmente é usado pela função principal, inicialização e testes e como o Contexto de nível superior para colicitações recebidas
*/
func background() context.Context {
	return context.Background()
}

/**
Tambem retorna um contexto dierente de vazio

No entando só deve ser usado quando não temos certeza de qual contexto usar ou se a fun'ão nao foi atualizada para receber um contexto. Isso significa que planejamos adicionar contexto à função no futuro.
*/
func todo() context.Context {
	return context.TODO()
}

/**
Recebe um valor base e associa os valores de chave e valor a esse contexto base, e retorna um novo contexto

isso significa que quem receber esse novo contexto tera acesso aos valores contidos nesse contexto

não é recomendao passar parâmetros cr;iticos usando valores de contexto, em vez disso as func'òes devem aceitar esses valores na assinatura tornando-o explicito
*/
func withValue(ctx context.Context, key any, value any) context.Context {
	return context.WithValue(ctx, key, value)
}

func withCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func withTimeout(ctx context.Context) context.Context {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	// VER O TIMEOUT SE ESGOTANDO DO CONTEXT
	// for {
	// 	fmt.Println("Executando %v", ctx)
	// 	time.Sleep(time.Second * 1)
	// }

	return ctx
}

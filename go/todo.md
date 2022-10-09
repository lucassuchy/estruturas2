# Projeto de implementação de bucket sort em go com multtreath

- Todo
  - Desenha a função inicial de sort
  - Desenhar a lista encadeada
    - Não tem orientação a objeto
    - 3 Campos
      - Nodo
      - Jogador
      - Numero(0-670)
  - Função bucket
    - Corre pela função valida o digito da centena
    - Cada bucket(novo nodo) vai receber um valor de centena
  - Função sort interna
    - Corre pelo bucket sorteando
  - Parte do multtreath
    - Executar cada um dos sorts do bucket em uma tread separada
    - O resultado do bucket 0 termina onde começa o bucket 1, assim por diante
  - Retorna a lista ordenada(não objeto)
  - Funções de testes
    - Importe para o benchmark
    - Alimenta a primeira lista
    - Forma aleatoria
    - Testar com elementos repetidos
    - Testar com menos elementos
    - Benchmarks

```go
type elemento struct = {
    nodo:1
    jogador:'String'
    numero:int
    proximo_nodo:2}
```

- Problema
  - Lista com muitos elementos ordenados ou não
- Custos
  - Correr a lista inteira pelo menos uma vez ja ordenando os elementos pelo MSD(maior digito significativo)
  - Dividir em listas menores
    - 3 Digitos:
      dividir pela centena
    - 2 digitos:
      dividir pela dezena
  - Ordernar as listas em ordem, considerando o MSD
    - Lista com 0 na centena primeiro
    - Lista com 1 na centena segundo
    ...
  - Quando finalizar a lista com 9 na centena, unimos as listas na sequencia e temos a lista ordenada
  - Custos
    - Primeira varredura
    - n ordenações de listas 'menores'(poder escolher adequado para o volume radix?)
- Soluções
  - Poder executar em paralelo
  - Algoritmo adequado para o volume
  - Recursivamente poder diminuir a lista até algo pequeno caso seja necessario
    - Validar se é realmente interessante isso
  - Se o volume aumenta o algoritmo mantem interessante?

- Exemplo
  - Figurinhas da copa
    - 670 figurinhas desordenadas
  - Divisão
    - 8 grupos(7 grupos de 100, 1 de 70)
  - ordena por grupo de centena
  - agrupa pela ordem dos grupos
  - 10 operações encadeadas
    - Separação - 1
    - Ordernação - 8
    - Agrupamento - 1
  - 3 operações paralelas:
    - divisão - 1
    - Ordernação - 1
    - Agrupamento - 1

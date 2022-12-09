## Executando o código: 

Docker garphQl: 

docker build . -t prova:graph

docker run -d -p 8080:8080 prova:graph


Docker rest: 

docker build . --build-arg run=rest -t prova:rest

docker run -d -p 8080:8080 prova:rest

## Ideia inicial

Ao ler a primeira vez a prova minha ideia inicial era montar um prjeto seguindo a arquitetura exagonal 
no qual: 

A) Teriamos uma função de core que:
  * Passaria uma regex para extrair os números romanos;
  * Converteria os números Romanos para inteiros
  * Descobriria se o número é primo. 
  * Essa função teria uma injeção de dependência de uma interface que armazenaria número
   primos já cálculados.

B) Uma função de armazenamento (comunicação com banco ou map em memória) para armazenar os primos (importada pelo core).

C) Dois pacotes que exporião o funcionamento para o mundo externo (injeção de dependência para o Core).

## Ideia implementada: 

Conforme fui pensando e analizando a ideia dos primos (o algoritmo mais rápido para descobrir se um número envolve dividir o número por outros primos.)

Os números romanos tem um limite de 3999 (A partir dai precisarimaos do traço em cima do V, o que não consta no exercício).

Logo não faz sentido trabalhar com os primos mais de uma vez, na inicialização do servidor vou fazer toda a computação de primos e valores romanos, para deixar a execução da chamada mais rápida. O custo de memória é quase 0 vez que apenas existem 550 primos inferiores a 3999.
 
Assim acabei alterando toda a ideia do projeto (e optando por uma extrutura de pastas muito mais simples).

A regra aplicada será: 


* Aplicar o Crivo de Eratóstenes para descobrir todos os primos menores de 3.999 (não desenvolvi a função, peguei uma pronta).
* Converter os números para números romanos e armazenar em um map[string]int 
* Passar uma função que leia a string recebida e tire todos os conjuntos de possíveis números romanos (retornará alguns falsos positivos, explicado melhor na função), vai performar melhor que fazer um REGEX geral para tudo com lookbehind e lookahead.
* Fazer uma comparação entre os valores dos conjuntos retornados com o map[string]int de maneira a remover os "falsos positivos" e permitir a comparação do primo.


### OBS: 

* Sendo o Crivo de Eratóstenes uma regra matemática conhecida não existe uma razão efetiva para reinventar a roda. Utilizei a função constante no repositorio: https://github.com/fxtlabs/primes a qual foi testada e possui um bom benchmark.
 Estou copiando a função em vez de importar o repositório pois o repo por tem uma func init que armazena em memória todos os primos menores de 10K.
* A função original de Ertóstenes foi alterada para evitar algumas coisas que não seriam usadas no projeto, as alterações tem um comentário // ALTERAÇÃO
* A performace poderia ser melhorada facilemnte aplicando a lógica de validar se é o maior primo dentro da função que le a string para pegar os romanso, todavia acredito que nesse caso a "melhoria" não compensa a perca de legibilidade / manutenibilidade.
* Como são poucos valores eu preferiria criar uma função que gerasse o texto do map[string]int e colar ele direto evitando a computação incial, mas acredito que ficaria feio para o teste, então estou criando os arquivos e fazendo a computação antes.

# Pacote: 

## Estruturação: 

/pkg: contem todo o código a ser utilizado

/pkg/helpers/*: contem funções que podem ser usadas por qualquer serviço e não importam nada.

/pkg/core/*: Contem a regra de negócio do produto. Por regra não deveria importar nada, no maximo alguns  helpers, não estou seguindo uma exagonal por causa que o projeto é bem pequeno.

/pkg/servers/*: contem as rotas e servidores


## Extras: 

Ele pede para mim encontrar números romanos em uma palavra com  o seguinte exemplo :
"A - XI - B - IV " encontrando XI e IV na palavra AXIBIV

Na palavra

DIDATICA:   Quais números eu devo considerar?
DI - ID - C  
DI - D - C  
D - ID - C
C  (DID não é um número romano válido)?

O Pior que a pergunta é relevante ainda mais nesse caso em que o "ID" é primo.

Estou partindo do pressuposto de que: Considere que as entradas serão sempre válidas.
engloba não ter um valor que não seja um número romano válido (vai ser ignorado).


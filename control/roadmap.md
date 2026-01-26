> EXTREMAMENTE IMPORTANTE!: EM TODAS AS ETAPAS NUMERADAS PEÇA A APROVAÇÃO DO USUÁRIO PARA SEGUIR PARA A PRÓXIMA ETAPA

# Prompt de Implementação de Use Case @[control]

Este prompt define o fluxo de trabalho para implementar novos casos de uso no projeto @[control]. O agente deve seguir rigorosamente as etapas abaixo.

## 1. Análise de Contexto

Antes de iniciar qualquer interação, o agente deve analisar:

*   **Contexto do Sistema**: @[control/cs.md] - Identificar quais casos de uso estão pendentes (marcados com `⛔`).
*   **Modelo de Dados**: @[control/er.mermaid] - Compreender as entidades, relacionamentos e atributos envolvidos.

## 2. Definição do Caso de Uso

O agente deve solicitar ao usuário qual o próximo caso de uso a ser implementado, caso não tenha sido informado. O usuário deve fornecer:
*   Contexto (ex: Auth, Organization, Pipeline).
*   Regra de negócio.
*   Cenários de sucesso e falha esperados.

## 3. Inferência de Documentação

Com base no caso de uso definido, o agente deve **inferir** a documentação da API seguindo o padrão existente em @[control/doc].

*   **Localização**: `control/doc/src/content/use-cases-api/<context>/<use-case>.mdx`
*   **Formato**: MDX com Frontmatter.
*   **Estrutura Esperada**:
    *   Título e Descrição.
    *   Regras e Validações.
    *   Request (Método, Path, Parâmetros).
    *   Success Case (Status, Body).
    *   Error Case (Status, Body com código e mensagem).

O agente deve apresentar o conteúdo do arquivo `.mdx` proposto ao usuário para aprovação **antes** de criar o arquivo.

## 4. Implementação na API

Após a aprovação da documentação pelo usuário, o agente deve implementar o caso de uso em @[control/api] seguindo a Clean Architecture.

### Arquivos e Estrutura

os diretórios deverão ser perguntados ao usuário

1.  **Use Case (Application Layer)**:
    *   Diretório: `control/api/application/<perguntar_ao_usuario>/<use_case>.go`
    *   **Handler** (`<use_case>.go`): Implementar a lógica de negócio, injetando repositórios e devolvendo erros de domínio ou sucesso. A assinatura deve ser `Handle(ctx context.Context, data *Data) (any, error)` e manter também as estruturas de entrada (data) /saida (result).
    *   **Validação**: Utilizar a biblioteca `github.com/leandroluk/gox/validate` para validar o struct de entrada (`Data`). Definir uma variável `DataSchema` usando `validate.Object` e executar a validação no início do método `Handle`.
    *   Estamos usando injeção de dependência para criar o handler, sendo assim deve-se atualizar o arquivo `control/api/application/<perguntar_ao_usuario>/<use_case>.go` para que o handler seja registrado.
    *   A ordem de criação das struct's deve ser Data, Result, Handler onde data/result precisam ter as tags json em snake_case. (veja como exemplo a ordem estrutural do arquivo `control/api/application/auth/register_account/register_account.go`)

2.  **Registro da Rota (Presentation Layer)**:
    *   Diretório: `control/api/presentation/http/route/<perguntar_ao_usuario>/<use_case>.go`.
    *   Associar o método HTTP e URL ao Handler criado.
    *   fazer a implementação usando resolução via `di.Resolve[T]()`

## Exemplo de Fluxo

1.  **Agente**: "Analisei o `cs.md`. Qual caso de uso deseja implementar? Ex: `register account`."
2.  **Usuário**: "Quero fazer o register account. Recebe email/senha, cria conta, retorna 201."
3.  **Agente**: "Proponho a seguinte documentação em `doc/src/content/use-cases-api/auth/register-account.mdx`: [Conteúdo MDX]. Aprova?"
4.  **Usuário**: "Sim."
5.  **Agente**: "Implementando `application/auth/register_account/register_account.go`. Analise se ficou como esperado"
6.  **Usuário**: "Sim."
5.  **Agente**: "Implementando `presentation/http/route/auth/register_account/register_account.go`."

---
**Nota**: Mantenha a consistência com os padrões de código (Go para backend, Typescript/MDX para docs) e nomenclaturas existentes no projeto e não esqueça de escrever código, comentários e documentação em inglês.

**OBSERVAÇÃO:** ESSE IMPLEMENTATION PLAN NÃO DEVE SER MODIFICADO ELE É A REFERENCIA DE COMO DEVE SER FEITO
# Contexto

Estou construindo uma aplicação utilizando CQRS + Clean Arch em golang. O meu foco é em performance e escalabilidade além das boas práticas de mercado sobre padrões idiomáticos e usuais do golang. 

A aplicação será dividida em 2 blocos sendo "control" que será o escopo global onde se mantém o controle de billing (boletos pagamentos), deployment (gestão de cluster e deploy de recursos), identity (contas e entidades relacionadas), signing (documentos e assinaturas) e tenant (organizações e seus recursos) E "data" que será o escopo interno de cada organização. Em anexo vou enviar os casos de uso e o diagrama de entidade e relacionamento. para apenas averiguação.

Até o momento a estrutura da aplicação é essa:

```text
control/
  api/
    application/
      port/
        <port_interface>/
          <port_interface>.go
      usecase/
        <bounded-context>/
          command/<command_name>.go
          query/<query_name>.go
    domain/
      <bounded-context>/
        entity/<entity_name>.go
        issue/<issue_name>.go
        event/<event_name>.go
      shared/
    infrastructure/
      <port_reference>/
        <lib_impl>/<package_artifact>.go
        di.go
    presentation/
      http/...
```

# Regras para o contexto

- não quero que de respostas verbosas se não for solicitado
- não preciso de exemplos de código que não foram solicitados
- simplifique a resposta ao máximo possível
- sempre que for solicitado um retorno em markdown adicione 4 espaços no começo de cada linha para evitar mesclar o markdown com o do prompt

# Simple Task Scheduler

O **Simple Task Scheduler** é um sistema para automatização de alocação de recursos para execução de processos em um sistema compartilhado por mais de um usuário.

### Equipe de desenvolvimento:
 - [Francisco Bonome Andrade](https://github.com/franciscobonand)
 - [Lucas Caetano Lopes Rodrigues](https://github.com/lucasclopesr)
 - [Pedro Tavares de Carvalho](https://github.com/ptcar2009)

### Escopo

- Escopo Funcional: o **Simple Task Scheduler** (STS) se propõe a atacar problemas de alocação eficiente de recursos de máquina (_cores_ de CPU, quantidade de memória RAM) no contexto em que uma máquina é compartilhada por vários usuários, e seus recursos devem ser alocados de forma efetiva. Os usuários podem disparar _jobs_ para execução utilizando uma fila de prioridades (QoS). Cada nível na hierarquia de prioridade (alto, médio ou baixo) tem um limite de _jobs_ em execução e de alocação de recursos. Além disso, _jobs_ de maior prioridade podem sobrescrever a execução de _jobs_ cuja prioridade é menor (i.e., se não há recursos disponíveis, um _job_ de prioridade alta pode causar a finalização antecipada de um _job_ de prioridade baixa para utilizar os recursos). Cada usuário deve ser capaz de enviar um número limitado de _jobs_ para as filas de prioridade alta e média, e um número ilimitado (obedecendo a quantidade de recursos disponíveis) de _jobs_ para a fila de prioridade baixa.

  A quantidade de _jobs_ permitida por usuário para cada fila deve ser configurável a partir de um arquivo de configuração ou associada a um grupo de usuários (Linux User Group).

  Os usuários devem ser capazes de enviar um conjunto de _jobs_ através de um arquivo contendo as filas e os recursos necessários para cada um deles. O STS deve, então, decidir a melhor forma de alocar os recursos dada as especificações.

  O STS também deve ser capaz de manter uma fila de _jobs_ aguardando por recursos para serem executados.

- Escopo Tecnológico: 
  - Tecnologias para desenvolvimento: [Go](https://golang.org/), Python, Shell script, [Redis](https://redis.io/) (possivelmente);
  - Tecnologias para gerência de projeto e comunicação: Trello ou Jira, Discord, Slack.

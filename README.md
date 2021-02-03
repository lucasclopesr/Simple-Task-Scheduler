# Simple Task Scheduler

O **Simple Task Scheduler** é um sistema para automatização de alocação de recursos para execução de processos em um sistema compartilhado por mais de um usuário.

### Equipe de desenvolvimento:
 - [Francisco Bonome Andrade](https://github.com/franciscobonand)
 - [Lucas Caetano Lopes Rodrigues](https://github.com/lucasclopesr)
 - [Pedro Tavares de Carvalho](https://github.com/ptcar2009)
 - [Arthur Alves Melo dos Santos Pacheco](https://github.com/arthur-pacheco)

### Escopo

- Escopo Funcional: o **Simple Task Scheduler** (STS) se propõe a atacar problemas de alocação eficiente de recursos de máquina (_cores_ de CPU, quantidade de memória RAM) no contexto em que uma máquina é compartilhada por vários usuários, e seus recursos devem ser alocados de forma efetiva. Os usuários podem disparar _jobs_ para execução utilizando uma fila de prioridades (QoS). Cada nível na hierarquia de prioridade (alto, médio ou baixo) tem um limite de _jobs_ em execução e de alocação de recursos. Além disso, _jobs_ de maior prioridade podem sobrescrever a execução de _jobs_ cuja prioridade é menor (i.e., se não há recursos disponíveis, um _job_ de prioridade alta pode causar a finalização antecipada de um _job_ de prioridade baixa para utilizar os recursos). Cada usuário deve ser capaz de enviar um número limitado de _jobs_ para as filas de prioridade alta e média, e um número ilimitado (obedecendo a quantidade de recursos disponíveis) de _jobs_ para a fila de prioridade baixa.

  A quantidade de _jobs_ permitida por usuário para cada fila deve ser configurável a partir de um arquivo de configuração ou associada a um grupo de usuários (Linux User Group).

  Os usuários devem ser capazes de enviar um conjunto de _jobs_ através de um arquivo contendo as filas e os recursos necessários para cada um deles. O STS deve, então, decidir a melhor forma de alocar os recursos dada as especificações.

  O STS também deve ser capaz de manter uma fila de _jobs_ aguardando por recursos para serem executados.

- Escopo Tecnológico: 
  - Tecnologias para desenvolvimento: [Go](https://golang.org/), Python, Shell script, [Redis](https://redis.io/) (possivelmente);
  - Tecnologias para gerência de projeto e comunicação: Trello ou Jira, Discord, Slack.
  
- Planejamento:

  - Nossas tarefas estão organizadas em um modelo Kanban, localizado no [GitHub Projects deste repositório](https://github.com/lucasclopesr/Simple-Task-Scheduler/projects/1).

## Estrutura do sistema

O Simple Task Scheduler é estruturado em três módulos principais: 

- Uma interface via linha de comando (CLI) para envio dos _jobs_, consultas aos _jobs_ em execução e na fila de _jobs_ em espera por recursos;
- Um módulo de gerência de recursos, cuja responsabilidade é verificar a disponibilidade de recursos, receber os _jobs_ e inserí-os na fila e enviar os _jobs_ para serem executados no sistema operacional.
- Um _daemon_ que é executado como um serviço do sistema, em segundo plano, cuja responsabilidade é prover uma API REST (utilizando um _socket_ UNIX) para comunicação entre a CLI e o gerenciador de recursos;

A imagem a seguir mostra um esquema da arquitetura do sistema.

<img src="./images/simp-arch.png"/>

## Execução

Para instalar direto no seu sistema (dentro do path) rode a seguinte linha de código:

```bash
curl https://raw.githubusercontent.com/lucasclopesr/Simple-Task-Scheduler/main/scripts/install.sh | bash
```

Como a arquitetura segue um modelo de cliente-servidor, a execução do sistema requer que o _daemon_ esteja sempre em execução, em segundo plano, para que o executável da CLI consiga fazer requisições. Portanto, são necessários dois executáveis, que podem ser gerados a partir do código neste repositório.

- Para gerar o executável do _daemon_, basta executar:

```bash
$ cd cmd/simpd
$ go build .
# O executável será gerado na pasta cmd/simpd, com o nome simpd
```

Em seguida, para executar o _daemon_, basta invocar o executável gerado na etapa anterior:

```bash
$ ./simpd
2021/01/31 14:30:40 Listening...
```

Detalhe: o serviço `simpd` permite que o usuário especifique quanto recurso (CPU, memória RAM) está disponível para que os _jobs_ submetidos através dele utilizem. Ele lê essas configurações do arquivo `~/.simp/config.json`. Caso ele não encontre o arquivo, ele gera um com valores padrão. O arquivo deve ter o seguinte formato:

```json
{
  "maxMemUsage": 100,
  "maxCPUUsage": 4
}
```



- Para gerar o executável da CLI, basta executar:

```bash
$ cd cmd/simp
$ go build .
# O executável será gerado na pasta cmd/simp, com o nome simp
```

Para enviar _jobs_ ao Simple Task Scheduler, utilize o executável `simp` gerado na etapa anterior. Consulte os comandos disponíveis utilizando a _flag_ `--help`:

```bash
$ ./simp --help
Usage:
   [command]

Available Commands:
  create              cria um novo job no STS
  deleteExecutingJob  deleta um job em execução no STS
  deleteExecutingJobs deleta todos os jobs em execução no STS
  deleteJobFromQueue  deleta um job na fila do STS
  deleteQueue         deleta todos os jobs na fila do STS
  getExecutingJob     retorna um job em execução no STS
  getExecutingJobs    retorna todos os jobs em execução no STS
  getJobFromQueue     retorna um job na fila do STS
  getQueuedJobs       retorna todos os jobs na fila do STS
  help                Help about any command

Flags:
  -a, --args stringArray   Argumentos do job que será criado (Array de Strings)
  -c, --cpu int            Mínimo de CPU disponível para a execução do job (0-100) (default 50)
  -h, --help               help for this command
  -i, --job_id string      ID do job que será criado (Inteiro) (default "no-id")
  -m, --mem int            Mínimo de memória disponível para execução do job (0-100) (default 50)
  -n, --name string        Caminho absoluto para o binário do job (String)
  -p, --priority int       Prioridade do job que será criado (0-1-2) (default 1)
  -w, --work_dir string    Diretório de trabalho do job (String)

Use " [command] --help" for more information about a command.
```




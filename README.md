# Programa Exibir Menu

**programa-exibir-menu** é um projeto desenvolvido em Go para monitorar o status de sites, exibindo o código de status HTTP de cada site monitorado. Ele também registra o histórico de monitoramento com a data e hora de cada verificação.

## Funcionalidades

- Monitora o status HTTP de uma lista de sites.
- Exibe se os sites estão online (status code 200) ou com problemas.
- Armazena os logs de monitoramento em um arquivo de registro.
- Permite visualizar o histórico de monitoramentos anteriores com data e hora.

## Como funciona

O programa exibe um menu com três opções:

1. **Iniciar Monitoramento**: Verifica o status HTTP dos sites listados no arquivo `sites.txt`.
2. **Exibir Logs**: Mostra o histórico de monitoramentos realizados, armazenados no arquivo `registro.txt`.
3. **Sair do Programa**: Encerra a execução do programa.

## Requisitos

- [Go](https://golang.org/dl/) instalado na sua máquina (se for compilar).
- Arquivo `sites.txt` contendo os URLs dos sites a serem monitorados (um por linha).
- **Versão compilada**: Para usuários do Windows, o programa já está compilado como `main.exe`.

## Como executar

### 1. Usando a versão compilada (Windows)

1. Baixe ou clone o repositório:
   
   ```bash
   git clone https://github.com/kaiquiBenevenutti/programa-exibir-menu.git

2. Navegue até o diretório do projeto:
   ```bash
   cd programa-exibir-menu

3. Execute o programa usando o arquivo main.exe:
    - Basta clicar duas vezes no main.exe ou rodar pelo terminal:
   ```bash
   ./main.exe

4. Escolha uma das opções do menu para iniciar o monitoramento ou visualizar os logs.
   
### 2. Compilando e rodando com Go

Se preferir compilar o código, siga os passos abaixo:

1. Clone o repositório do projeto:
  ```bash
git clone https://github.com/kaiquiBenevenutti/programa-exibir-menu.git

```
2. Navegue até o diretório do projeto:
  ```bash
cd programa-exibir-menu

```
3. Compile e execute o programa:
  ```bash
go build -o main.exe main.go
./main.exe

  ```
    
ou apenas rode o programa no terminal:
```bash
go run main.go

```

### Estrutura do Projeto

- **main.go:** Arquivo principal do programa que contém toda a lógica de monitoramento e exibição de logs.
- **sites.txt:** Arquivo com os URLs dos sites que serão monitorados.
- **registro.txt:** Arquivo onde os logs dos monitoramentos são salvos.
- **main.exe:** Versão compilada do programa para usuários do Windows.

### Observações

- O programa monitora cada site 3 vezes com um intervalo de 3 segundos entre as verificações.
- Para modificar esses valores, altere as constantes **monitoramentos** e **delay** no código.

### Repositório Original

Este projeto foi criado como parte do meu aprendizado em Go e faz parte do repositório [learning-Go-Alura](https://github.com/kaiquiBenevenutti/learning_Go_Alura), onde estou documentando meu progresso e estudos.

### Autor

- Kaiqui Benevenutti

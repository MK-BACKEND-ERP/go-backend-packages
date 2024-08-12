# go-backend-packages
Package para centralizar todos os objetos compartilhadas entre os projetos Go da MK Solutions

![](https://www.mksolutions.com.br/wp-content/uploads/2022/11/mk-positive.svg)

**Configuração Chave SSH GitHub**

Para que o Go consiga baixar o pacote é necessário realizar a configuração de Chave SSH.

https://docs.github.com/pt/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```
```bash
eval "$(ssh-agent -s)"
```
```bash
ssh-add ~/.ssh/id_ed25519
```
Adicionar chave no GitHub

https://docs.github.com/pt/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account

**Instalação**

Para permitir que o go get baixe módulos privados, você precisa configurar a variável do GOPRIVATE:

```bash
go env -w GOPRIVATE=github.com/MK-BACKEND-ERP/go-backend-packages
```
Em seguida deve baixar o pacote:
```bash
go get github.com/MK-BACKEND-ERP/go-backend-packages
```

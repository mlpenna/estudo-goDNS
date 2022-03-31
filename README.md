# Filtro web através de DNS utilizando Go

Este é um projeto de estudo da linguagem Go que consiste em desenvolver um filtro web baseado em DNS que pode atuar em uma rede local com o objetivo
de filtrar acesso a algumas URLs em nível de DNS. A ideia é retornar um endereço IP inválido ou página padrão sempre que o servidor DNS receber uma
URL presente na lista de URLs a serem bloqueadas.
O servidor DNS é implementado com a biblioteca: https://github.com/miekg/dns

# O que foi implementado e alguns pontos:

O servidor é capaz de atuar como servidor DNS em uma rede local (é necessário configurar manualmente o DNS nas estações presentes na rede). É possível filtrar
algumas URLs setadas manualmente dentro do código. Caso ela não seja uma URL a ser bloqueada, o DNS repassa a query a um de dois DNS públicos (google ou cloudflare)
e retorna a resposta ao cliente que fez a query original.

É necessário investigar e implementar uma forma mais eficiente e acurada para fazer a comparação entre a strings da URL recebida na query pelo cliente e as
contidas na lista de URL a serem bloqueadas.

Uma forma melhor de implementar esta lógica de bloqueio talvez seja usando uma API de clasificação de URL disponível no mercado (há algumas), e realizar a decisão com base no retorno dessa API.




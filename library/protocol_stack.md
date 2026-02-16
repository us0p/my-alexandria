---
id: 20260216-protocol_stack
tags:
  - architecture
  - networking
created: 2026-02-16
status: draft
---
# Protocol Stack
> Deve responder implicitamente: *“Qual problema essa nota resolve?”*
---
## ❓ Pergunta central
> Qual é a pergunta **real** que essa nota responde?

Ex:
- Por que a arquitetura hexagonal reduz acoplamento?
- Quando **não** usar esse padrão?
---
## 🧠 Explicação (com suas palavras)
> Anotação **generativa**. Nada de copiar fonte.
- Explique o conceito como se estivesse ensinando
- Use frases completas
- Foco em **causa → efeito**

Ex:
- Esse padrão existe porque sistemas tendem a acoplar regras de negócio a detalhes externos...
- A separação permite trocar interfaces sem reescrever regras centrais...
---
## ⚖️ Trade-offs / Limites
> Onde quebra? O que custa?
- Complexidade inicial
- Overengineering em sistemas simples
- Curva de aprendizado
---
## 🧪 Exemplos
### Exemplo típico
- Aplicação web com múltiplas interfaces (REST, CLI)
### Contraexemplo
- CRUD simples sem lógica relevante
---
## 🔁 Relações (Zettelkasten)
### Conecta com:
- [[Arquitetura em Camadas]]
- [[DDD – Separação de Domínios]]
- [[Complexidade Acidental vs Essencial]]
### Contrasta com:
- [[MVC Tradicional]]
> 💡 Links sempre com **verbo implícito**: “expande”, “contrasta”, “depende”
---
## 💭 Fricções / Dúvidas
> Pontos que ainda não estão claros (ouro puro)
- Ainda não entendo como isso afeta testes de integração
- Revisar impacto em performance
---
## 🧠 Síntese em 2–3 linhas
> Se você só lesse isso daqui a 1 ano, o que deveria lembrar?
---
## 🎯 Flashcards candidatos
> ⚠️ **Não escreva cards ainda. Só marque.**
- [ ] O que é arquitetura hexagonal?
- [ ] Qual problema ela resolve?
- [ ] Quando não usar?
- [ ] Qual trade-off principal?
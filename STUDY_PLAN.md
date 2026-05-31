# Study Plan

Topics and tools encountered while building this service.

---
## 2. PostgreSQL, pgx and transactions

### `pgx.Conn` vs `pgxpool.Pool`
`pgx.Conn` is a single connection. If you call `QueryRow` and never call `.Scan()` on the result, the connection stays busy and all subsequent queries on it fail with `conn busy`. A connection pool (`pgxpool.Pool`) mitigates this in concurrent scenarios but the root fix is always to consume your results.

Resources:
- [pgx README](https://github.com/jackc/pgx)
- [pgx — Querying](https://pkg.go.dev/github.com/jackc/pgx/v5#hdr-Querying)

### `QueryRow` and `.Scan()`
`db.QueryRow()` never returns an error directly. The error surfaces only when you call `.Scan()`. Always check the error from Scan, not from QueryRow.
---

## 3. Google ADK for Go

### Launchers vs Runners
`full.NewLauncher()` (and all launchers under `google.golang.org/adk/cmd/launcher`) are CLI entry points. They parse `os.Args`-style argument slices and spin up web servers, pub/sub listeners, or interactive consoles. They are **not** for programmatic invocation.

For calling an agent from your own code, use `runner.Runner`:

```
runner.New(runner.Config{
    AppName:           "...",
    Agent:             myAgent,
    SessionService:    session.InMemoryService(),
    AutoCreateSession: true,
})
```

Then call `r.Run(ctx, userID, sessionID, msg, agent.RunConfig{})` and iterate the returned event sequence.

### Sessions
Each call to `r.Run()` operates within a session. Sessions maintain conversation history. For independent, stateless invocations (like classifying N separate stories), use a unique session ID per story so they don't share history.

### Output schemas and `genai.Schema`
When you set `OutputSchema` on an `llmagent`, the model is constrained to return JSON matching that schema. Rules to remember:
- Any property with `Type: genai.TypeArray` **must** include an `Items` field describing the element type. Omitting it causes a 400 INVALID_ARGUMENT error from the API.

### Sequential agents
ADK provides `sequentialagent` for chaining multiple LLM reasoning steps where output from one feeds into the next via session state. Only use this when each step requires LLM reasoning. Pure data transformation steps (like calling an embedding model) belong in regular Go code, not in an agent.

Resources:
- [Google ADK for Go — GitHub](https://github.com/google/adk-go)
- [Google ADK docs](https://google.github.io/adk-docs/)

---

## 4. Google Generative AI SDK (`google.golang.org/genai`)

### Generating embeddings
Use `client.Models.EmbedContent()` directly — no agent needed. The call takes a model name, input content, and an optional config:

```go
client, _ := genai.NewClient(ctx, nil)
result, _ := client.Models.EmbedContent(ctx, "text-embedding-004", genai.Text("your text"), nil)
embedding := result.Embeddings[0].Values // []float32
```

`genai.NewClient` reads credentials from `GEMINI_API_KEY` by default.

Resources:
- [genai Go SDK — pkg.go.dev](https://pkg.go.dev/google.golang.org/genai)

---

## 5. pgvector — Semantic search in PostgreSQL

### Distance vs Similarity
pgvector's `<=>` operator computes **cosine distance** (0 = identical, 2 = opposite directions). Cosine similarity is the inverse:

```
similarity = 1 - cosine_distance
```

To filter by similarity >= 90%:

```sql
SELECT id, name, 1 - (embedding <=> $1) AS similarity
FROM genre
WHERE 1 - (embedding <=> $1) >= 0.9
ORDER BY similarity DESC
```

### Passing vectors from Go
Without extra dependencies you can format `[]float32` as a pgvector literal string and cast it in SQL (`$1::vector`). The cleaner option is `github.com/pgvector/pgvector-go`, which provides a `pgvector.Vector` type that pgx can serialize directly without manual casting.

Resources:
- [pgvector — GitHub](https://github.com/pgvector/pgvector)
- [pgvector-go — GitHub](https://github.com/pgvector/pgvector-go)
- [pgvector — Querying](https://github.com/pgvector/pgvector#querying)

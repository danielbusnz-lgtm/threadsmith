# threadsmith

A self-hostable AI gateway in a single Go binary. Point your app at it instead of OpenAI and it forwards the call through. The plan is to grow it into a full agent platform: request logging, multi-provider routing, a durable agent runtime, and a code sandbox.

Early days. Right now it proxies chat completions to OpenAI.

## Run

```bash
export OPENAI_API_KEY=sk-...
go run ./cmd/gateway
```

Listens on `:8080`.

## Endpoints

| Method | Path | What it does |
| --- | --- | --- |
| GET | `/healthz` | returns `ok` |
| POST | `/echo` | echoes the request body back |
| POST | `/v1/chat/completions` | forwards to OpenAI and returns the reply |

## Try it

```bash
curl -X POST http://localhost:8080/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{"model":"gpt-4o-mini","messages":[{"role":"user","content":"hi"}]}'
```

## License

MIT

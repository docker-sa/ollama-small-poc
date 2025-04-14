# 🦙 Ollama Small POC

## If you work on workstation on Windows or Linux with GPU

```bash
docker compose --file compose.gpu.yml up --build
```
> - Tested on Docker Cloud

## If you work on workstation on Windows or Linux without GPU or on MacOS

```bash
docker compose --file compose.no-gpu.yml up --build
```
> - Tested on macOS
> - Tested on a Windows VDI

## Be patient ⏳

The first time, it downloads ths LLMs

Then the result looks like this:
```bash
ollama-service-1       | time=2025-04-14T15:25:39.413Z level=INFO source=server.go:619 msg="llama runner started in 2.02 seconds"
generate-completion-1  | 🔴 starting the completion...
generate-completion-1  | 🟠 starting the completion...
generate-completion-1  | 🟢 starting the completion...
ollama-service-1       | [GIN] 2025/04/14 - 15:26:14 | 200 | 37.710368392s |   192.168.128.3 | POST     "/v1/chat/completions"
generate-completion-1  | Jean-Luc Picard is the captain of the USS Enterprise and one of the three founders of Starfleet. 
He serves as the central character across several episodes of "Star Trek: The Next Generation" 
and has developed into a beloved figure for many viewers over the years. Jean-Luc Picard was 
born on Corysta in Babelia, his childhood home being at the Tenth sector Academy on Utopian-Principality, 
from which he graduated with honors. His parents were not mentioned during my brief narrative 
of their deaths that occurred while he was a child.
generate-completion-1  | 🔵 done!
generate-completion-1 exited with code 0
```

## Remarks

I downladed 3 models:
- `qwen2.5:0.5b`
- `qwen2.5:1.5b`
- `qwen2.5:3b`

Depending on your GPU (or CPU), you can choose the model you want to use in the `.env` file:
```bash
OLLAMA_BASE_URL=http://ollama-service:11434
#LLM_CHAT=qwen2.5:0.5b
LLM_CHAT=qwen2.5:1.5b
#LLM_CHAT=qwen2.5:3b
```

## Troubleshooting

If you get errors about the memory, you need to increase the memory of your Docker Desktop.

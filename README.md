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

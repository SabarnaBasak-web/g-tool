# gtool 🚀

A lightweight CLI wrapper around Git to simplify everyday workflows.

## ✨ Why gtool?

Working with Git is powerful but often repetitive:

- `git add .`
- `git commit -m "message"`
- `git push`

**gtool** reduces this friction by combining common commands into simple, intuitive shortcuts.

---

## ⚡ Features

- 📦 One-command save (add + commit + push)
- 🔄 Sync branch with remote (pull --rebase + push)
- 🌿 Create branches quickly
- 🧠 Simple and minimal — no overengineering

---

## 📥 Installation

### Prerequisites

- Go installed
- Git installed

### Steps

```bash
git clone https://github.com/your-username/gtool.git
cd gtool
go build -o gtool

`Linux`
sudo mv gtool /usr/local/bin/


```

Now you can run:

```bash
gtool
```

---

## 🚀 Usage

### Save changes

```bash
gtool save "your commit message"
```

Runs:

```bash
git add .
git commit -m "your commit message"
git push
```

---

### Sync branch

```bash
gtool sync
```

Runs:

```bash
git pull --rebase
git push
```

---

### Create new branch

```bash
gtool new <branch-name>
```

Runs:

```bash
git pull origin
git checkout -b <branch-name>
```

---

## 📁 Project Structure

```
gtool/
 ├── main.go
 ├── git.go
```

- `main.go` → CLI entry point
- `git.go` → Handles Git command execution

---

## 🛠 Tech Stack

- Go (Golang)
- Native CLI (`os`, `exec`)
- Git (as underlying engine)

---

## 🎯 Roadmap

- [ ] Add commit message validation
- [ ] Interactive prompts
- [ ] Colored CLI output
- [ ] Support for flags (using Cobra)
- [ ] Better error handling

---

## 🤝 Contributing

Contributions are welcome!

If you have ideas to improve developer workflow, feel free to open an issue or submit a PR.

---

## ⚠️ Disclaimer

This tool is a wrapper around Git. It does not replace Git and relies on Git being installed on your system.

---

## 📄 License

MIT License

---

## 💡 Inspiration

Built to simplify daily Git workflows and reduce repetitive commands.

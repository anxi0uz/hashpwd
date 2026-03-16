# hashpwd

Простая CLI-утилита для хэширования паролей через bcrypt. Принимает строку, возвращает чистый хэш без спецсимволов.

A simple CLI tool for hashing passwords via bcrypt. Takes a string, returns a clean hash without special characters.

---

## Использование / Usage

```bash
hashpwd mypassword
# → J7bmcVNF1NQEGOcc3UOjp4Osc7w686V5UhBApN9Cln2pgvuYDbW
```

---

## Сборка / Build

Нужен Go 1.18+. / Requires Go 1.18+.

```bash
git clone https://github.com/yourname/hashpwd
cd hashpwd
go build -o hashpwd main.go
```

Windows:

```bash
go build -o hashpwd.exe main.go
```

---

## Установка в PATH / Add to PATH

### Linux

**bash / zsh**

```bash
mv hashpwd ~/.local/bin/

# если ~/.local/bin нет в PATH, добавить в ~/.bashrc или ~/.zshrc:
export PATH="$HOME/.local/bin:$PATH"
```

**fish**

```bash
mv hashpwd ~/.local/bin/
fish_add_path ~/.local/bin
```

Или вручную в `~/.config/fish/config.fish`:

```fish
fish_add_path ~/.local/bin
```

---

### macOS

**bash / zsh**

```bash
mv hashpwd /usr/local/bin/
```

Если нет прав:

```bash
sudo mv hashpwd /usr/local/bin/
```

**fish**

```bash
mv hashpwd /usr/local/bin/
fish_add_path /usr/local/bin
```

---

### Windows

Переместить `hashpwd.exe` в любую папку, например `C:\Tools\`, затем добавить её в PATH:

```
Пуск → "переменные среды" → Environment Variables
→ Path → Edit → New → C:\Tools
```

Или через PowerShell:

```powershell
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\Tools", "User")
```

После этого в новом терминале:

```bash
hashpwd mypassword
```

---

## Зависимости / Dependencies

```bash
go get golang.org/x/crypto/bcrypt
```

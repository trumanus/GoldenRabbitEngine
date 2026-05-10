# First push to a new public GitHub repository

The local repo is ready at:

`~/sviluppo/GoldenRabbitEngine`  
(on case-insensitive macOS this may appear as `~/Sviluppo/GoldenRabbitEngine`)

Optional absolute layout:

`/sviluppo/GoldenRabbitEngine` — create `/sviluppo` on your machine first (`mkdir`, permissions).

## 1. Create the empty repo on GitHub

1. Open GitHub → **New repository**  
2. Name: **`GoldenRabbitEngine`** (matches Go module `github.com/trumanus/GoldenRabbitEngine`)  
3. Visibility: **Public**  
4. **Do not** add README, `.gitignore`, or license (they already exist locally).

## 2. Authenticate

Either:

```bash
gh auth login
```

or ensure SSH keys / HTTPS credential helper are configured for `github.com`.

## 3. Push

```bash
cd ~/sviluppo/GoldenRabbitEngine   # or /sviluppo/GoldenRabbitEngine
git remote add origin https://github.com/trumanus/GoldenRabbitEngine.git
# oppure: git remote add origin git@github.com:trumanus/GoldenRabbitEngine.git
git branch -M main
git push -u origin main
```

If `origin` already exists:

```bash
git remote set-url origin https://github.com/trumanus/GoldenRabbitEngine.git
git push -u origin main
```

## 4. Tag release (optional)

```bash
git tag -a v0.1.0 -m "v0.1.0"
git push origin v0.1.0
```

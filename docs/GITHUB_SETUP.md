# Publishing this folder as a new GitHub repository

Use these steps if **`golden-rabbit-engine`** is published **standalone** (recommended for open source), separate from a private monorepo that contains books.

## 1. Create an empty repository on GitHub

1. GitHub → **New repository**  
2. Name: e.g. `golden-rabbit-engine`  
3. **Public**  
4. **Do not** add README, .gitignore, or license (this folder already has them)

## 2. Initialise Git inside this directory (first time only)

From your machine:

```bash
cd golden-rabbit-engine
git init
git add .
git commit -m "chore: initial public reference implementation v0.1.0"
git branch -M main
git remote add origin https://github.com/trumanus/GoldenRabbitEngine.git
git push -u origin main
```

Replace `trumanus` with your GitHub username or organisation if different.

## 3. Tag the first release

```bash
git tag -a v0.1.0 -m "v0.1.0"
git push origin v0.1.0
```

Then open **Releases** on GitHub and publish release notes (copy from `CHANGELOG.md`).

## 4. If the module path must change

If your repo URL is not `github.com/trumanus/GoldenRabbitEngine`:

1. Edit the first line of `go.mod`  
2. Replace every import path across `.go` files (same string as in `go.mod`)  
3. Update links in `README.md`, `CHANGELOG.md`, and this doc  

## SSH remote (optional)

```bash
git remote set-url origin git@github.com:trumanus/golden-rabbit-engine.git
```

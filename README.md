# toker
Toker is a lightweight app that clones GitHub projects and analyzes their codebase with Tokei, offering insightful statistics and details about project structure.

# Run
```bash
git clone https://github.com/milad-rasouli/toker.git
cd toker

# if you want to develop the code
make install-deps

cp config/config.example.toml ./config.toml 
make run
```
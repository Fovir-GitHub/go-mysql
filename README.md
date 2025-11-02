# Description

This repo stores source code of learning `Go` + `MySQL`.

# Getting Start

0. Set up the dev shell:

With `direnv`:

```bash
direnv allow
```

With `flake`:

```bash
nix develop
```

1. Initialize the database:

```bash
just init
```

2. Start the `MySQL` server:

```bash
just server
```

3. Enter the DBMS:

```bash
just start
```

4. Set up the database and exit:

```sql
SOURCE ./db.sql
exit
```

5. Run the program:

```bash
go run ./main.go
```

6. Stop the `MySQL` server:

```bash
just stop
```

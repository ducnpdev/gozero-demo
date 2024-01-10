# demo gozero

## k8s

## model generate

### postgresql
- need created table in database
- script generate model
```sh
goctl model pg datasource --url "postgres://paperless:1@127.0.0.1:5433/paperless" --cache true -t auth_user --dir sqlpg
```

### mysql

touch file sql: `prod.sql`

```sql
CREATE TABLE departments (
    dept_no CHAR(4)     NOT NULL,              -- in the form of 'dxxx'
    dept_name VARCHAR(40) NOT NULL DEFAULT '',
    PRIMARY KEY (dept_no),                     -- Index built automatically
    UNIQUE  KEY (dept_name)                    -- Build INDEX on this unique-value column
);
```
- gene
```sh
goctl model mysql ddl --src prod.sql --dir .  
```


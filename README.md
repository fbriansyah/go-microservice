# GO Microservice Template

## Project Structure
```
.
|-- cmd
|   `-- api
|-- docs
|   `-- ADL
|-- internal
|   |-- application
|   |   |-- commands
|   |   |-- domain
|   |   |   `-- user
|   |   `-- queries
|   |-- echo
|   |-- logging
|   `-- postgres
|-- migration
`-- util
```

### CMD Folder /cmd

Di dalam folder ini terdapat bootstrap aplikasi seperti setting database, config, repository dan juga application.

Kenapa main package terdapat di dalam folder **/api**?
Dengan memisahkan main package, bisa membuat projek fleksibel. Misal projek ingin membuat versi web atau desktop, kita bisa membuat folder khusus
seperti **/web** atau **/desktop**.

### Docs Folder /docs
Di dalam folder ini terdapat beberapa dokumentasi seperti ADL (Agnostic Development Language) yang berisi beberapa keputusan untuk menggunakan
teknologi atau disign pattern pada aplikasi.

### Internal Folder /internal
Di dalam folder ini berisi core aplikasi dan juga beberapa adapter seperti **echo**, **logging** dan **postgres**.

### Application Folder /internal/application
Folder ini berisi logic dan juga domain. Dalam aplikasi ini menerapkan pattern CQRS (detail bisa dibaca di folder dokumentasi **0002-use-a-ddd.md**).

### Migration Folder /migration
Folder ini berisi beberapa file migrasi untuk postgres sql.

### Utility Folder /util
Folder ini berisi beberapa utilities function, seperti generating dan verify password.

## Note:
- Folder internal di golang [/internal](https://go.dev/doc/go1.4#internalpackages).
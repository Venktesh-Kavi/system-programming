## Making PSQL Client Interactions Better


- To skip password prompts data can be stored ~/.pgpass, ensure permissions are provided.

```sh
// follow the below format
# hostname:port:database:user:password
localhost:5213:*:postgres:think@123
```

### psqlrc file

global: /etc/psqlrc
user: ~/.psqlrc

postgres's shell can be configured using /set or /pset. (pset is for changing the output format, set is for everything else).

histfile - by default psql stores the command history ~/.psql_history. (^r sometimes picks commands from here and replays them).
If you want a db specific query history, psql cli has some preset variables like :DBNAME, :PORT, :HOST etc.., 


### Pagers for psql client

- By default cli uses `less`
- https://github.com/okbob/pspg


References:

- https://thoughtbot.com/blog/improving-the-command-line-postgres-experience
- https://thoughtbot.com/blog/an-explained-psqlrc
- postgres explanation viewer (https://tatiyants.com/pev/#/plans/new)

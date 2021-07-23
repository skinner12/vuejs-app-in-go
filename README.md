# VUEJS APP IN GOLANG

> Spiegazione completa del programma su [rmazzu.com](https://www.rmazzu.com/articoli/inserire-vuejs-app-in-go/ "rmazzu.com")

Con l’introduzione in golang 1.16 del pacchetto `embed` è stata data la possibilità di inserire direttamente nel programma Go compilato un frontend, rendendo la pubblicazione di un server fullstack molto più semplice utilizzando solamente un file.

## Development

Eseguire questi comandi per attivare i servers per lo sviluppo in locale:

```bash

cd frontend
yarn watch

# In un altro terminale
go run .

```

## Production

Eseguire i seguenti comandi per la compilazione della build pronta da essere pubblicata:

```bash
cd frontend
yarn build
cd ..
go build -tags prod
```

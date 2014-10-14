#Godoc
```
$ godoc -http:=6060
```

#Recompile on save
```
$ ./run.sh
```

#Setup
1. Install Go
2. Install ArangoDB
3. Install [fresh][], i18n (see below) and kails

#Ambiente de Desarrollo
## Bases de Datos
* ArangoDB - multi-paradigm database

##Librerias
* martini
* aranGO

##Herramientas
### gocode
Provee compleciones inteligentes (para LightTable o Vim o sublime)

### [fresh][]
Compilacion automatica cuando detecta cambios en los archivos go o en los templates. (go get github.com/pilu/fresh)

### i18n
####Installation:
1. `$ go get -u github.com/nicksnyder/go-i18n/i18n`
2. `$ go get -u github.com/nicksnyder/go-i18n/goi18n`

####Usage:
1. Create the translation files in the respective language folder under translations/. Example:

```json
[
{
  "id": "d_days",
  "translation": {
    "one": "{{.Count}} day",
    "other": "{{.Count}} days"
  }
},
{
  "id": "my_height_in_meters",
  "translation": {
    "one": "I am {{.Count}} meter tall.",
    "other": "I am {{.Count}} meters tall."
  }
},
{
  "id": "person_greeting",
  "translation": "Hello {{.Person}}"
}
]
```

2. Compile the different translation files into one translation file per language under the translations/all directory.
    * Execute the next command while in the `translations` directory

```
$ goi18n -outdir all/ {english,spanish}/*.json
```

3. Use the strings in the templates, examples:

```
{{T "program_greeting"}}
{{T "your_unread_email_count" 1}}
```


[fresh]: https://github.com/pilu/fresh  "fresh"
[martini]: https://github.com/go-martini/martini/ "martini"

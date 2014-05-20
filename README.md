#Setup
1. Install Go
2. Install mongodb, Neo4j
3. Install [fresh][], i18n (see below) and kails

#Ambiente de Desarrollo
## Bases de Datos
* Neo4j - Base de datos de grafos
* MongoDB - NoSQL
* CouchDB - NoSQL
* MySQL - SQL

##Librerias
* mgo - Manejador de MongoDB para go.
* [martini][] - Paquete de desarrollo web. (go get github.com/go-martini/martini)

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

    :::json
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

2. Compile the different translation files into one translation file per language under the translations/all directory.
    * Execute  `$ goi18n -outdir all/ {english,spanish}/*.json` while in the `translations` directory.

3. Use the strings in the templates, example:
    * `{{T "program_greeting"}}`
    * `{{T "your_unread_email_count" 1}}`


[fresh]: https://github.com/pilu/fresh  "fresh"
[martini]: https://github.com/go-martini/martini/ "martini"

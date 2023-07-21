# Go vaja 1 - Osnove Go-ja potrebne za izdelavo REST API-ja

Osnove za tiste, ki si še niste prebrali.
 - https://golang.org/
 - https://tour.golang.org
 - https://talks.golang.org/2012/splash.article
 - https://golang.org/ref/spec#Keywords
 
Osnovno strukturo za vsak projekt in razne primere najdete [tukaj](https://github.com/IT-Team-Global/GoBasics/tree/master). Ta repo dobro preučite saj je v njem pokritih 90% pomembnih stvari. V vsako mapo sem dodal 
README z opisi in razlogi zakaj so stvari tako narejene. V datotekah pri [Strukturi Projekta](https://github.com/IT-Team-Global/GoBasics/tree/master/StrukturaProjekta) imate primer kode, ki je konkretno komentirana.

Navodila kako enostavno postavit bazo podatkov najdete [tukaj](https://github.com/IT-Team-Global/GoVaje/blob/master/Docker.md)

Vsakemu je bil na Skynet-u dodeljen task za to nalogo. Ko nalogo naredite prosim naložite zip z kodo pod datoteke na nalog in ga pošljite v potrditev.

Cilj ni da je naloga samo čim prej narejena tak kot na faksu ne glede na kvaliteto ampak da je napisana po navodilih, napake urejene, struktura programa ločena glede na to kaj koda dela da se vsi naučimo pisat Go projekte kvalitetno in na enak berljiv način da bomo čim manj časa porabili z iskanjem kdo je kaj napisal in kak program dela če pride do težav :)

## Naredite osnoven backend za TODO app

Frontenda ne rabite naredit.

Za testiranje lahko uporabite [Postman](https://www.postman.com/) ali kar koli drugega.

Zahteve:
 - HTTP REST API dostopen na portu 80
 - Podatki se naj shranjujejo v MariaDB (MySQL)
 - Kreiranje baze more bit samodejno (Znotraj programa/Go kode) - Če kdo želi lahko že pri prvi vaji uporabi [Migrator.go](https://gitlab.com/i.t.tim/go/-/tree/master/Ostalo/MigracijeBaze)
 - Konfiguracija je lahko za sedaj hardcoded ali pa jo dobite iz [Okoljskih Spremenjljivk (ENV)](https://github.com/IT-Team-Global/GoBasics/tree/master/Ostalo/Env)
 - Struktura more ustrezat strukturi [tukaj](https://github.com/IT-Team-Global/GoBasics/tree/master/StrukturaProjekta)
 - Vsaka funkcija v kateri se lahko zgodi error more ta error tudi vrnit
 - Pri vsaki klicani funkciji, ki vrne error more bit preverjanje če je prišlo do njega, ga izpisat in ustrezno reagirat
 - Za klicanje funkcij, ki komunicirajo z bazo more bit uporabljen interface in ne direkt objekt (primer je v Sktrukturi Projekta)
 - Pazite na memory leak pri povezavi z bazo (zapirajte vrstice z defer in Close / za poizvedbe uporabljajte Query za izvedbe pa Exec). Go ima Garbage Collector, ki poskrbi za večino stvari, ne more pa namesto vas zapret povezave z bazo.
 - Delovat morejo naslednje poti:
   - GET /api/v1/todo - Vrne JSON array vseh opravkov
   - POST /api/v1/todo - Prejme opravek v obiki JSON ter ga vstavi v bazo
   - GET /api/v1/todo/:todo_id - Vrne opravek z id-jem todo_id v JSON obliki
   - PUT /api/v1/todo/:todo_id - Prejme opravek v obiki JSON ter ga shrani na mesto opravka z id-jem todo_id (sprememba obstoječega)
   - DELETE /api/v1/todo/:todo_id - Izbiše opravek z id-jem todo_id
 
Vsako opravilo naj ima naslednje podatke:
  - ID (primary key v bazi)
  - Naslov
  - Opis
  - Datum dodajanja
  - Predviden datum dela
  
## IDE

Za razvojno okolje predlagam da tisti, ki še imate status študenta uporabite [GoLand](https://www.jetbrains.com/go/) saj je narejen prav za Go in ima tako vse podprto + ima dodatek za direktno povezavo na bazo in SQL highlighting v kodi glede na bazo (pride zelo prav saj javi napake in avtomatsko dopolni SQL stavke - imena stolpcev, stavke, joine itd.)

Za ostale pa bo najboljša izbira verjetno [Visual Studio Code + Go](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)

Ko bomo prešli iz vaj na dejanski razvoj v Go bomo najverjetneje na firmi nabavili GoLand za vse.
 

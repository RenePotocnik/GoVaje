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
 

----


# Go vaja 2 - Online beleženje napak, avtentikacija in migracija baze

### Cilji vaje:
- Sprememba baze z Migrator.go
- Online error handling
- Autentikacija z JWT Token-i
- Izdelava Gin middleware-a

Vsakemu je bil na Skynet-u dodeljen task za to nalogo. Ko nalogo naredite prosim naložite zip z kodo pod datoteke na nalog in ga pošljite v potrditev.
 
---

### Zahteve:
- Vsa konfiguracija za dostop do baze se naj pridobiva iz ENV (okoljske spremenjljivke)
- Vsak error v aplikaciji se more zabeležit v naš Sentry strežnik
- Z pomočjo [Migrator.go](https://github.com/IT-Team-Global/GoBasics/tree/master/Ostalo/MigracijeBaze) dodajte v bazo tabelo z uporabniki in v tabelo z todo-ji dodajte tuji ključ, ki kaže na id uporabnika. Uporabite tudi PostScript field definicije za migrator
  in v njega dodajte funkcijo, ki bo kreirala novega uporabnika z uporabniškim imenom demo in geslom demo, ga dodala v bazo in nato na vseh obstoječih taskih za user_id nastavila ID uporabnika demo (dodala vse taske/todo-je na njega).
- Gesla v bazi ne smejo biti shranjena kot plain-text. Uporabite bcrypt hashing algoritem.
- Dodajte naslednje poti v api:
   - /api/v1/login - Preveri uporabnika in geslo ter generira in nastavi JWT token (token more vsebovat časovni limit do kdaj je veljaven in ID uporabnika).
   - /api/v1/logout - Izbriše cookie z JWT tokenom (če ga shranjujete tu in ne v headeru)
   - /health - Naredi ping na bazo in vrne status 200 OK če je vse ok v nasprotnem primeru pa status 500 Internal server error
- Napišite svoj [Gin Middleware](https://github.com/gin-gonic/gin#custom-middleware), ki bo ob vsaki HTTP zahtevi preveril če je JWT token prisoten (Cookie ali Authorization header) in ga preveri. Ob napaki naj vrne status 401 Unauthorized in konča.
  Ob uspešnem preverjanju pa naj z c.Set() nastavi ID uporabnika, ki ga dobi iz JWT tokena, da ga lahko nato uporabite v handlerjih.
- Vsak uporabnik lahko vidi samo svoje stvari (mogli boste v svojih HTTP handlerjih dodat c.Get da dobite ID uporabnika, ki ga je middleware nastavil in ga dat v vse nadaljne funkcije da lahko iz baze poberete prave taske)

 ---

### Navodila:
- [Migrator.go](https://github.com/IT-Team-Global/GoBasics/tree/master/Ostalo/MigracijeBaze)
- Za Sentry so navodila vidna ko se prijavite v sistem.
- Navodila za [bcrypt](https://gowebexamples.com/password-hashing/)
- Delo z [JWT Tokeni v Go](https://github.com/dgrijalva/jwt-go)
- Kako naredit Gin middleware [primer 1](https://github.com/gin-gonic/gin#custom-middleware) / [primer 2](https://github.com/IT-Team-Global/GoBasics/tree/master/Ostalo/GinMiddleware)

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
 

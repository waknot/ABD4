# ABD4

Lancer le script php

```
php ./script [timer]
```

Ce script fait une request post sur `127.0.0.1:8000` et envoi un json de reservation

Lancer l'API

```
cd API/

go build

./API
```

Elle écoute et renvoi ok quand elle reçoit un request post sur `127.0.0.1:8000`

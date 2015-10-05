# Serveur de chat
* Chat avec salles de discussions, et sauvegarde des conversations
* Support de commandes et émoticônes
* Gestion des avatars des utilisateurs

### Groupe ViDuCha

* Morgane Vidal
* Geoffrey Dumas
* Manuel Chataigner

---

### Lancer l'application

#### Pré-requis

* Package websocket io
```bash
go get github.com/googollee/go-socket.io
```

#### Installation

* Récupérer les sources du projet : https://gitlab.info-ufr.univ-montp2.fr/HMIN302/go-viducha

* Se placer dans le dossier du projet

* Lancer la compilation
```bash
make
```

#### Utilisation

* Lancement du serveur
```bash
./serverChat
```

* Client
```bash
http://localhost:5000/
```

---

### Spécifications techniques

* Serveur Go, capable de gérer plusieurs clients simultanément
* Client web
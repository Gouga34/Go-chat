# Serveur de chat
* Chat avec salles de discussions, et sauvegarde des conversations
* Support de commandes et émoticônes
* Gestion des avatars des utilisateurs (Gravatar)

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
* Package gravatar
```bash
go get github.com/ftrvxmtrx/gravatar
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
http://localhost:1200/
```
* Commandes du chat

Commande pour obtenir l'heure courante
```bash
/time
```

Commande pour envoyer un message privé
```bash
/mp <member> <message>
```

---

### Spécifications techniques

* Serveur Go, capable de gérer plusieurs clients simultanément
* Client web Javascript
* Communication Websockets avec Socket.io

### Difficultés rencontrées

* Le choix des sockets

Nous pensions dans un premier temps utiliser des sockets web classiques. Nous avons perdu beaucoup de temps sur leur mise en place avec les routeurs httprouter, pour finalement nous rendre compte qu'il était plus judicieux d'utiliser les socket io, qui sont beaucoup plus simples d'utilisation et permettent une gestion efficace de la connexion des utilisateurs dans les salles de chat.

* Dépendances cycliques

Nous avons plusieurs fois eu des problèmes par rapport aux dépendances cycliques. En effet, pour l'enregistrement des données dans la base de données, nous avions besoin de connaitre les structures alors que parallèlement, les fichiers contenant les différentes structures avaient besoin d'inclure notre fichier de gestion de la base de données quand nous devions charger ou enregistrer des données.
Go empêchant les inclusions cycliques, nous avons dans un premier temps essayé de mettre toutes nos structures dans un même fichier, afin que celles-ci puissent être incluses par les autres fichiers. Mais nous rencontrions alors un autre problème quant à la définition des différentes fonctions appliquées aux structures, ainsi qu'un problème du à l'accès aux propriétés dû à la séparation structure et fonctions, qui nous obligeait à passer les membres d'une structure en public.

Pour pallier l'ensemble de ces problèmes, nous avons choisi de n'utiliser dans la base de données que des strings, et donc éviter la séparation entre définition des structures et définition des fonctions associées.

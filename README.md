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

*
```bash
go get github.com/googollee/go-socket.io
```

*
```bash
go get github.com/ftrvxmtrx/gravatar
```

*
```bash
go get github.com/boltdb/bolt
```

#### Installation

* Récupérer les sources du projet : https://gitlab.info-ufr.univ-montp2.fr/HMIN302/go-viducha

* Se placer dans le dossier du projet

* Lancer la compilation
```bash
make
```

#### Utilisation

* Lancement du serveur (avec port d'écoute en paramètre optionnel, le port par défaut est 1200)
```bash
./serverChat [port]
```

* Client
```bash
http://localhost:<port>/
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
* Emoticônes disponibles
    - :)
    - :(
    - ;)
    - :D
    - :'(
    - :o

---

### Spécifications techniques

* Serveur Go, capable de gérer plusieurs clients simultanément
* Client web Javascript
* Communication Websockets avec Socket.io

### Analyse

* Conception serveur

Notre serveur est représenté par une structure Server contenant l'objet serveur mis en place par la librairie, permettant la communication avec les clients, ainsi que l'ensemble des salles de chat disponibles. Pour gérer cet ensemble de salles, nous avons créé une structure contenant la liste, avec l'ensemble des méthodes de manipulation des données associées, les utilisateurs et les salons.
Une salle est représentée par une structure, gérant l'ensemble des utilisateurs et les messages postés. Chacune d'elle contient la liste des utilisateurs actuellement connectés.
Pour chaque utilisateur, le serveur stocke les données de connexion et informations personnelles, ainsi que la salle à laquelle il est connecté et la socket associée au client.

* Communication

Concernant la communication entre serveur et clients, celle-ci est effectuée à l'aide de websockets. Pour leur mise en place, nous avons finalement choisi le système socket io, permettant notamment une gestion efficace des groupes.
Pour la communication, nous avons complété les messages de base mis en place par la librairie (connection, disconnection...) par des messages personnalisés pour chacune des actions disponibles (changeRoom, message...). A chaque message est associée une strucure json contenant les données associées au message.

* Stockage des données

Afin de mettre en place la persistance des données dans notre application, nous avons choisi d'utiliser le système de stockage clé/valeur bolt. Un fichier chat.db contient donc l'ensemble des utilisateurs, salles et conversations de l'application.
La manipulation de l'ensemble de ces données est effectuée via un objet contenant l'ensemble des requêtes d'ajout et récupération.

### Difficultés rencontrées

* Le choix des sockets

Nous pensions dans un premier temps utiliser des sockets web classiques. Nous avons perdu beaucoup de temps sur leur mise en place avec les routeurs httprouter, pour finalement nous rendre compte qu'il était plus judicieux d'utiliser les socket io, qui sont beaucoup plus simples d'utilisation et permettent une gestion efficace de la connexion des utilisateurs dans les salles de chat.

* Dépendances cycliques

Nous avons plusieurs fois eu des problèmes par rapport aux dépendances cycliques. En effet, pour l'enregistrement des données dans la base de données, nous avions besoin de connaitre les structures alors que parallèlement, les fichiers contenant les différentes structures avaient besoin d'inclure notre fichier de gestion de la base de données quand nous devions charger ou enregistrer des données.
Go empêchant les inclusions cycliques, nous avons dans un premier temps essayé de mettre toutes nos structures dans un même fichier, afin que celles-ci puissent être incluses par les autres fichiers. Mais nous rencontrions alors un autre problème quant à la définition des différentes fonctions appliquées aux structures, ainsi qu'un problème du à l'accès aux propriétés dû à la séparation structure et fonctions, qui nous obligeait à passer les membres d'une structure en public.

Pour pallier l'ensemble de ces problèmes, nous avons choisi de n'utiliser dans la base de données que des strings, et donc éviter la séparation entre définition des structures et définition des fonctions associées.

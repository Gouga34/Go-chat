var socket = io();

socket.on('connect', function() {});

socket.on('login', function(data){
                                    userConnection(data);
                                });

socket.on('message', function(data){
                                      printMessage(data);
                                    });

socket.on('command', function(data){
                                      printCommand(data);
                                    });

socket.on('changeRoom', function(data){
                                        switchUsersRoom(data);
                                      });

socket.on('disconnect', function(){});

socket.on('newUser', function(data){
                                      addConnectedNewUser(data);
                                    });

socket.on('register', function(data){
  getResultOfInscription(data);
});

socket.on('mp', function(data){
  getAndPrintPersonalMessage(data);
})

socket.on('userLeft', function(data){
  deleteUserFromConnectedClients(data);
})

socket.on('newRoom', function(data){
  newRoom(data);
})


//Functions -------------------------------------------------------------------

//Traitement des messages reçus ------------------------------------------

/**
 * @action ajoute les salles
 */
function newRoom(data){
    datas=JSON.parse(data);
    addRoom(datas.Name);
}

/**
 * @action ajoute le client qui vient de se connecter à la room
 */
function addConnectedNewUser(data){
    datas=JSON.parse(data);
    addConnectedUserToList(datas.Login, datas.GravatarLink);
}

/**
 * @action supprime le client qui vient de se déconnecter de la room
 */
function deleteUserFromConnectedClients(data){
  datas=JSON.parse(data);
  deleteDiv(datas.Login);
}

/**
 * @param data le message perso
 * @action affiche le message perso reçu.
 */
function getAndPrintPersonalMessage(data){
  datas=JSON.parse(data);
  addMessage("From : "+datas.Author, datas.Time, datas.Content, datas.GravatarLink);
}

/**
 * @param data les données reçues
 * @action connecte l'utilisateur si l'inscription a bien marché, affiche le problème sinon
 */
function getResultOfInscription(data){
  datas=JSON.parse(data);
  if(datas.Success){
    $("#content").load('chat.html');
    connectUser(datas.Login, datas.GravatarLink, datas.RoomList);
  }
  else{
    if(!datas.LoginOk){
      alert('Le login est déjà pris');
    }
    else if(!datas.PasswordOk){
      alert('Les deux mots de passe sont différents');
    }
  }
}

/**
 * @param data les données reçues
 * @action affiche le message reçu dans la boite de chat
 */
function printMessage(data){
  datas=JSON.parse(data);
  addMessage(datas.Author, datas.Time, datas.Content, datas.GravatarLink);
}

/**
 * @param data les données reçues
 * @action affiche le résultat de la commande
 */
function printCommand(data){
  datas=JSON.parse(data);
  addCommand(datas.Content);
}

/**
 * @param data les données reçues
 * @action change l'utilisateur de salle, avec creation si besoin
 */
function switchUsersRoom(data){
  var datas=JSON.parse(data);
  if(datas.Success){
    if(datas.NewRoom){
      addRoom(datas.RoomName);
    }
    switchRoom(datas.RoomName, datas.ClientList, datas.MessageList);
  }
}

/**
 * @param data les données reçues
 * @action Traite le retour de connexion envoyé au serveur
 */
function userConnection(data){

  var datas=JSON.parse(data);
  if(datas.Success){
    $("#content").load('chat.html');
    connectUser(datas.Login, datas.GravatarLink, datas.RoomList);
  }
  else {
    printConnectionError(datas.LoginOk, datas.PasswordOk)
  }
}

//Envoi des messages ---------------------------------------------------

/**
 * @param login, password, passwordVerif, mail : strings, données à enregistrer
 * @action envoie un message d'inscritption au serveur.
 */
function sendInscriptionMessage(login, password, passwordVerif, mail){
  var messageToSend={Login:login, Password:password, VerifPassword:passwordVerif, Mail:mail};
  socket.emit("register", JSON.stringify(messageToSend));
}

/**
 * @action envoie au serveur les informations du formulaire de connexion
 */
function sendConnectionForm(){
  var login = document.getElementById('login').value;
  var password = document.getElementById('password').value;

  var messageToSend = {Login:login, Password:password};
  socket.emit("login", JSON.stringify(messageToSend));
}

/**
 * @action envoie un message pour changer de salle au serveur
 */
function changeRoom(data){
  var messageToSend={RoomName:data};
  socket.emit("changeRoom", JSON.stringify(messageToSend));
}

/**
 * @action envoie le message écrit par l'utilisateur au serveur
 */
function senddata() {
   var data = document.getElementById('textToSend').value;
   if (data.length != 0)
   {
   var time=new Date(Date.now());
   var messageToSend={Content:data, Author:"", Time: time.toLocaleDateString()+" "+time.toLocaleTimeString()};
   var serializedMessage = JSON.stringify(messageToSend);
   socket.emit('message', serializedMessage);}
}

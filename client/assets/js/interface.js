
/**
 * @param roomName nom de la room à ajouter
 * @action ajoute une room dans la liste des rooms de l'interface
 */
function addRoom(roomName){

  document.getElementById('listRooms').innerHTML +=
  "<div class=\"item\" id=\""+roomName+"\" >"
    +"<div class=\"right floated content\">"
        +"<div class=\"ui button tiny gray\" onclick=\"changeRoom('"+roomName+"')\">Connecté</div>"
      +"</div>"
      +"<div class=\"content room\">"
          +"<div class=\"ui label teal circular\" >5</div>"
          +roomName
      +"</div>"
  +"</div>";
}

/**
 * @action vide la fenêtre de chats
 */
function clearCommentsList(){
  if(document.getElementById('listComments')!=null){
    document.getElementById('listComments').innerHTML="";
  }
}

/**
 * @action change le bouton des salles à "rejoindre"
 */
function changeRoomsConnectionButtonToDisconnected(){
  if(document.getElementById('listRooms')!=null){
    var roomsConnection = document.getElementById('listRooms').getElementsByClassName('ui button tiny gray');
    if(roomsConnection.length>0){
      for (i=0; i<roomsConnection.length; i++){
        roomsConnection[i].innerHTML='Rejoindre'
      }
    }
  }
}

/**
 * @param roomName nom de la room
 * @action change le bouton de la salle a "connecté"
 */
function changeRoomConnectionButtonToConnected(roomName){
  if(document.getElementById(roomName)!=null){
    var roomConnected=  document.getElementById(roomName).getElementsByClassName('ui button tiny gray');
    if(roomConnected.length>0){
      roomConnected[0].innerHTML='Connecté';
    }
  }
}

/**
 * @param roomName
 * @action change le bouton de la room passée en paramètre à "Rejoindre"
 */
function changeRoomConnectionButtonToDisconnected(roomName){
  var roomConnected=  document.getElementById(roomName).getElementsByClassName('ui button tiny gray');
  if(roomConnected.length>0){
    roomConnected[0].innerHTML='Rejoindre';
  }
}

/**
 * @param roomName nom de la room
 * @param connectedClients liste des utilisateurs connectés à roomName
 * @action change la room à laquelle est connecté le client et affiche la liste des utilisateurs connectés à la salle.
 */
function switchRoom(roomName, connectedClients, messages){
  clearCommentsList();
  clearConnectedUsersList();
  changeRoomsConnectionButtonToDisconnected();
  changeRoomConnectionButtonToConnected(roomName);

  if(connectedClients!=null){
    for(var i=0; i<connectedClients.length; i++){
      addConnectedUserToList(connectedClients[i].Login, connectedClients[i].GravatarLink);
    }
  }

  if(messages!=null){
    for(var i=0; i<messages.length;i++){
      addMessage(messages[i].Author, messages[i].Time, messages[i].Content, messages[i].GravatarLink);
    }
  }
}

/**
 * @param content Contenu du résultat de la commande
 */
function addCommand(content){
  document.getElementById('listComments').innerHTML +=
  "<div class=\"comment\">"
      +"<div class=\"content\">"
          +"<div class=\"text\">"
            + content
          +"</div>"
      +"</div>"
  +"</div>";
}

/**
 * @param author auteur du message
 * @param time date du message
 * @param content contenu du message
 * @param image lien gravatar de l'image
 */
function addMessage(author, time, content, image){
  if(document.getElementById('listComments')!=null){
    document.getElementById('listComments').innerHTML +=
    "<div class=\"comment\" >"
        +"<a class=\"avatar\">"
            +"<img src=\""+image+"\">"
        +"</a>"
        +"<div class=\"content\" id=\""+time+"\">"
            +"<a class=\"author\">"+author+"</a>"
            +"<div class=\"metadata\">"
                +"<span class=\"date\">"+time+"</span>"
            +"</div>"
            +"<div class=\"text\">"
              + content
            +"</div>"
        +"</div>"
    +"</div>";
    document.getElementById('listComments').scrollTop = document.getElementById('listComments').scrollHeight;
  }
}

/**
 * @action affiche le formulaire de création d'une room
 */
function getRoomCreationForm(){
  document.getElementById('listRooms').innerHTML+=
  "<div class=\"ui fluid action input\" id=\"addRoom\">"
        +"<input id=\"newRoom\" type=\"text\" placeholder=\"Nom de la salle\">"
        +"<div class=\"ui button teal\" onclick=\"createRoom()\"><i class=\"send icon\" ></i>Créer</div>"
    +"</div>";
}

/**
 * @action crée une room
 */
function createRoom(){
  var roomName = document.getElementById('newRoom').value;
  document.getElementById('listRooms').removeChild(document.getElementById('addRoom'));
  changeRoom(roomName);
}

/**
 * @action vide la liste des utilisateurs connectés
 */
function clearConnectedUsersList(){
  if(document.getElementById("users")!=null){
    document.getElementById("users").innerHTML="";
  }
}

/**
 * @param login login de l'utilisateur à ajouter
 * @param image avatar de l'utilisateur
 * @action ajoute l'utilisateur dans la liste des utilisateurs connectés
 */
function addConnectedUserToList(login, image){
  if(document.getElementById("users")!=null){
    document.getElementById("users").innerHTML+=
    "<div class=\"item\" id=\""+login+"\">"
      +"<img class=\"ui avatar image\" src=\""+image+"\">"
      +"<div class=\"content\">"
        +"<a class=\"header\">"+login+"</a>"
      +"</div>"
    +"</div>";
  }
}

/**
 * @param login login de l'utilisateur connecté
 * @param image avatar de l'utilisateur
 * @param roomList liste des salons disponibles
 * @action affiche l'utilisateur courant comme connecté
 */
function connectUser(login, image, roomList){
  alert("Bienvenue "+login);
  for (var i = 0; i < roomList.length; i++){
    addRoom(roomList[i]);
  }
}

/**
 * @param loginOk
 * @param passwordOk
 * @action affiche une erreur de connexion à l'utilisateur
 */
function printConnectionError(loginOk, passwordOk){
  if(!loginOk){
    alert('Login inconnu');
  }
  else{
    alert('Mot de passe incorrect');
  }
}

/**
 * @action récupère les données pour l'inscription de l'utilisateur et envoie le message au serveur
 */
function inscription(){
  if(!checkIfPasswordsAreTheSame()){
    alert('Les deux mots de passe ne sont pas les mêmes !');
  }
  else{
    var login=document.getElementById('loginInsc').value;
    var password = document.getElementById('passwordInsc').value;
    var passwordVerif = document.getElementById('passwordInscrVerif').value;
    var mail = document.getElementById('mailInsc').value;

    sendInscriptionMessage(login, password, passwordVerif, mail);
  }
}

/**
 * @action verifie si le mot de passe donné est le même que la deuxième fois où le mot de passe est donné
 * @return true si les deux mots de passe sont les mêmes, false sinon
 */
function checkIfPasswordsAreTheSame(){
   var password = document.getElementById('passwordInsc').value;
   var passwordVerif = document.getElementById('passwordInscrVerif').value;
   if(password!==passwordVerif){
     return false;
   }
   return true;
}

/**
 * @param idDiv id de la div à supprimer
 * @action supprime la div ayant pour id idDiv
 */
function deleteDiv(idDiv){
  $("#"+idDiv ).remove();
}

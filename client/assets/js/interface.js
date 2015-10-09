
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
  document.getElementById('listComments').innerHTML="";
}

/**
 * @action change le bouton des salles à "rejoindre"
 */
function changeRoomsConnectionButtonToDisconnected(){
  var roomsConnection = document.getElementById('listRooms').getElementsByClassName('ui button tiny gray');
  if(roomsConnection.length>0){
    for (i=0; i<roomsConnection.length; i++){
      roomsConnection[i].innerHTML='Rejoindre'
    }
  }
}

/**
 * @param roomName nom de la room
 * @action change le bouton de la salle a "connecté"
 */
function changeRoomConnectionButtonToConnected(roomName){
  var roomConnected=  document.getElementById(roomName).getElementsByClassName('ui button tiny gray');
  if(roomConnected.length>0){
    roomConnected[0].innerHTML='Connecté';
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
function switchRoom(roomName, connectedClients){
  clearCommentsList();
  changeRoomsConnectionButtonToDisconnected();
  changeRoomConnectionButtonToConnected(roomName);
  if(roomName!="Defaut"){
    changeRoomConnectionButtonToDisconnected("defaut");
  }

  for(var i=0; i<connectedClients.length; i++){
    addConnectedUserToList(connectedClients[i].Login, connectedClients[i].GravatarLink);
  }
}

/**
 * @param author auteur du message
 * @param time date du message
 * @param content contenu du message
 * @param image lien gravatar de l'image
 */
function addMessage(author, time, content, image){
  document.getElementById('listComments').innerHTML +=
  "<div class=\"comment\">"
      +"<a class=\"avatar\">"
          +"<img src=\""+image+"\">"
      +"</a>"
      +"<div class=\"content\">"
          +"<a class=\"author\">"+author+"</a>"
          +"<div class=\"metadata\">"
              +"<span class=\"date\">"+time+"</span>"
          +"</div>"
          +"<div class=\"text\">"
            + content
          +"</div>"
      +"</div>"
  +"</div>";
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
  document.getElementById("users").innerHTML="";
}

/**
 * @param login login de l'utilisateur à ajouter
 * @param image avatar de l'utilisateur
 * @action ajoute l'utilisateur dans la liste des utilisateurs connectés
 */
function addConnectedUserToList(login, image){
  document.getElementById("users").innerHTML+=
  "<div class=\"item\" id=\""+login+"\">"
    +"<img class=\"ui avatar image\" src=\""+image+"\">"
    +"<div class=\"content\">"
      +"<a class=\"header\">"+login+"</a>"
    +"</div>"
  +"</div>";
}

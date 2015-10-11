var socket = io();

socket.on('connect', function() {});

socket.on('login', function(data){
                                    userConnection(data);
                                });

socket.on('message', function(data){
                                      printMessage(data);
                                    });

socket.on('changeRoom', function(data){
                                        switchUsersRoom(data);
                                      });

socket.on('disconnect', function(){});

socket.on('newUser', function(data){
                                      addConnectedUserToList(data.Login, data.GravatarLink);
                                    });


//Functions -------------------------------------------------------------------

//Traitement des messages reçus ------------------------------------------

/**
 * @param data les données reçues
 * @action affiche le message reçu dans la boite de chat
 */
function printMessage(data){
  datas=JSON.parse(data);
  addMessage(datas.Author, datas.Time, datas.Content, data.GravatarLink);
}

/**
 * @param data les données reçues
 * @action change l'utilisateur de salle, avec creation si besoin
 */
function switchUsersRoom(data){
  var datas=JSON.parse(data);
  alert("datas.success : "+datas.Success);
  alert("datas.newroom : "+datas.NewRoom);
  if(datas.Success){
    if(datas.NewRoom){
      addRoom(datas.RoomName);
    }
    alert('switch room');
    switchRoom(datas.RoomName);
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
   var time=new Date(Date.now());
   var messageToSend={Content:data, Author:"", Time: time.toDateString()+" "+time.getHours()+"h"+time.getMinutes()};
   var serializedMessage = JSON.stringify(messageToSend);
   socket.emit('message', serializedMessage);
}

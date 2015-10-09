var socket = io();

socket.on('connect', function() {});

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
  if(datas.Success){
    if(datas.NewRoom){
      addRoom(datas.RoomName);
    }
    switchRoom(datas.RoomName, datas.ConnectedClients);
  }
}

//Envoi des messages ---------------------------------------------------

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
   var messageToSend={content:data, author:"", time: time.toDateString()+" "+time.getHours()+"h"+time.getMinutes()};
   var serializedMessage = JSON.stringify(messageToSend);
   socket.emit('message', serializedMessage);
}

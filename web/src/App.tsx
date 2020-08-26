import React from 'react';

function App() {
  // Create WebSocket connection.
  const socket = new WebSocket('ws://localhost:8080/echo');

  // Connection opened
  socket.addEventListener('open', function (event) {
    socket.send(JSON.stringify({ id: 'getUserData', data: { patreon_user_id: 'HubertMoszka' } }));
  });

  // Listen for messages
  socket.addEventListener('message', function (event) {
    console.log('Message from server ', event.data);
  });

  return (
    <div className="App">
    </div>
  );
}

export default App;

<!-- Chat.svelte -->
<script>
    let messages = [];
    let newMessage = '';
  
    function sendMessage() {
      if (newMessage.trim() !== '') {
        const message = {
          text: newMessage,
          user: 'You',
          timestamp: new Date(),
        };
        messages = [...messages, message];
        newMessage = '';
  
        // Simulate receiving a message from another user (replace this with WebSocket communication)
        setTimeout(() => {
          const friendMessage = {
            text: 'Hello!',
            user: 'Friend',
            timestamp: new Date(),
          };
          messages = [...messages, friendMessage];
  
          // Scroll to the bottom to show the new message
          scrollChatToBottom();
        }, 1000);
      }
    }
  
    // Function to scroll the chat to the bottom
    function scrollChatToBottom() {
      const chatMessages = document.querySelector('.chat-messages');
      chatMessages.scrollTop = chatMessages.scrollHeight;
    }
  </script>
  
  <div class="chat-container">
    <div class="chat-header">
      <h2>Chat with Friend</h2>
    </div>
    <div class="chat-messages" on:scroll={scrollChatToBottom}>
      {#each messages as message (message.timestamp)}
        <div class="message">
          <p class="sender">{message.user}</p>
          <div class="message-content">
            <div class="message-text">{message.text}</div>
            <div class="message-timestamp">
              {message.timestamp.toLocaleTimeString()}
            </div>
          </div>
        </div>
      {/each}
    </div>
    <div class="message-input">
      <input
        type="text"
        placeholder="Type a message"
        bind:value={newMessage}
        on:keydown={(e) => e.key === 'Enter' && sendMessage()}
      />
      <button on:click={sendMessage}>Send</button>
    </div>
  </div>
  
  <style>
    /* Add your own CSS for styling */
    .chat-container {
      max-width: 400px;
      margin: 0 auto;
      padding: 20px;
      display: flex;
      flex-direction: column;
      height: 100vh;
    }
  
    .chat-header {
      text-align: center;
      margin-bottom: 10px;
    }
  
    .chat-messages {
      flex-grow: 1;
      border: 1px solid #ccc;
      padding: 10px;
      overflow-y: auto;
      background-color: #fafafa; /* Light gray background */
      border-radius: 5px;
      box-shadow: 0 0 5px rgba(0, 0, 0, 0.1); /* Add a subtle shadow */
      max-height: 70%; /* Limit chat box height */
    }
  
    .message {
      padding: 5px;
      margin: 5px;
    }
  
    .sender {
      font-weight: bold;
    }
  
    .message-content {
      background-color: #e0e0e0; /* Light gray for messages */
      padding: 10px;
      border-radius: 5px;
    }
  
    .message-text {
      margin: 0;
    }
  
    .message-timestamp {
      color: #777;
      font-size: 12px;
      text-align: right;
    }
  
    .message-input {
      display: flex;
      justify-content: space-between;
      padding: 10px;
      background-color: white;
      border-top: 1px solid #ccc;
    }
  
    input[type="text"] {
      flex-grow: 1;
      padding: 10px;
      border: none;
      border-radius: 25px;
    }
  
    button {
      padding: 10px 20px;
      background-color: #0084ff; /* Blue send button */
      color: white;
      border: none;
      border-radius: 25px;
      cursor: pointer;
    }
  </style>
  
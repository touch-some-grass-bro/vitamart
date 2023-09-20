<!-- Chat.svelte -->
<script>
  import axios from "axios";

  let messages = [];
  let newMessage = "";

  import API_URL from "../api";

  const WS_URL = API_URL.replace("http", "ws");

  let user;

  axios
    .get(API_URL + "/auth/is-authenticated", { withCredentials: true })
    .then((res) => {
      user = res.data.user;
    });

  const urlParams = new URLSearchParams(window.location.search);
  const isBuyer = urlParams.get("isBuyer") === "true";
  const isSeller = urlParams.get("isSeller") === "true";
  const transactionID = urlParams.get("transactionID");

  const ws_url = `${WS_URL}/chat?isBuyer=${isBuyer}&isSeller=${isSeller}&transactionID=${transactionID}`;

  const ws = new WebSocket(ws_url);

  ws.onclose = () => {
    console.log("WebSocket closed unexpectedly");
  };

  ws.onerror = (err) => {
    console.error("WebSocket errored: ", err);
  };

  function sendMessage() {
    if (newMessage.trim() !== "") {
      const message = {
        content: newMessage,
        name: user.name,
        email: user.email,
        room_id: transactionID,
      };

      // Send the message to the WebSocket server
      ws.send(newMessage);

      messages = [...messages, message];
      newMessage = "";
    }
  }

  ws.onmessage = (event) => {
    const recieve = JSON.parse(event.data);
    console.log(recieve)
    const message = {
      content: recieve.content,
      name: recieve.name,
      email: recieve.email,
      room_id: transactionID,
    };

    messages = [...messages, message];
  };

  // Function to scroll the chat to the bottom
  function scrollChatToBottom() {
    const chatMessages = document.querySelector(".chat-messages");
    chatMessages.scrollTop = chatMessages.scrollHeight;
  }
</script>

<div class="chat-container">
  <div class="chat-header">
    <h2>Chat with Friend</h2>
  </div>
  <div class="chat-messages" on:scroll={scrollChatToBottom}>
    {#each messages as message}
      <div class="message">
        <p class="sender">{message.name}</p>
        <div class="message-content">
          <div class="message-text">{message.content}</div>
        </div>
      </div>
    {/each}
  </div>
  <div class="message-input">
    <input
      type="text"
      placeholder="Type a message"
      bind:value={newMessage}
      on:keydown={(e) => e.key === "Enter" && sendMessage()}
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


const dialogDiv = document.getElementById("dialog");
const inputField = document.getElementById("input-field");
const sendButton = document.getElementById("send-button");
const themeToggle = document.getElementById("theme-toggle");
const themeIconDark = document.getElementById("theme-icon-dark");
const themeIconLight = document.getElementById("theme-icon-light");
const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
const host = window.location.host;
const ws = new WebSocket(`${protocol}//${host}/user/dialog`);

const setThemeMode = isDark => {
    if (isDark) {
        document.body.classList.add('dark-mode');
        themeIconDark.style.display = "none";
        themeIconLight.style.display = "block";
    } else {
        document.body.classList.remove('dark-mode');
        themeIconDark.style.display = "block";
        themeIconLight.style.display = "none";
    }
};

const checkSystemTheme = () => {
    const isDarkMode = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
    setThemeMode(isDarkMode);
};

const displayMessage = (message, type) => {
    const el = document.createElement("div");
    if (type === "system") {
        el.className = "message message-system";
    } else if (type === "sent") {
        el.className = "message message-sent";
    } else if (type === "received") {
        el.className = "message message-received";
    }
    el.innerText = message;
    dialogDiv.appendChild(el);
    dialogDiv.scrollTop = dialogDiv.scrollHeight;
};

const sendMessage = () => {
    const message = inputField.value.trim();
    if (!message) return;
    ws.send(message);
    displayMessage(message, "sent");
    inputField.value = "";
    inputField.focus();
};

ws.onopen = () => displayMessage("Connection established", "system");
ws.onmessage = event => {
    let data;
    try {
        data = JSON.parse(event.data);
    } catch (e) {
        console.error("JSON parse error", e);
        return;
    }
    displayMessage("Request Model: " + data.model, "system");
    for (let i = 0; i < data.messages.length; i++) {
        displayMessage(data.messages[i].role + ": " + data.messages[i].content, data.type);
    }

};
ws.onclose = () => displayMessage("Connection closed", "system");
ws.onerror = error => {
    console.error("WebSocket error", error);
    displayMessage("Connection error", "system");
};

sendButton.addEventListener("click", sendMessage);
inputField.addEventListener("keyup", event => {
    if (event.key === "Enter") sendMessage();
});
themeToggle.addEventListener("click", () => {
    const isDark = document.body.classList.contains("dark-mode");
    setThemeMode(!isDark);
});
window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', e => {
    setThemeMode(e.matches);
});
checkSystemTheme();
inputField.focus();

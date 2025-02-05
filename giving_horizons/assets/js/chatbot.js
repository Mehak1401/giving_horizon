const responses = {
    "hi":"Hi ,How can I help you?",
    "How can I donate?": "You can donate through our platform by selecting the cause you're passionate about, entering your donation amount, and completing the payment.",
    "What causes can I donate to?": "We support various causes including education, healthcare, environmental protection, animal welfare, and disaster relief.",
    "Is my donation secure?": "Yes, we use encrypted payment gateways to ensure that your donation is secure and safe.",
    "Can I track my donation?": "Absolutely! You can track the impact of your donation through our platform, and see the real-time effects of your generosity.",
    "How do I know where my money goes?": "Our platform ensures full transparency, and we provide detailed reports on how funds are allocated for each cause.",
    "What is the minimum donation amount?": "The minimum donation amount is $5, but any amount you contribute makes a significant difference.",
    "Can I get a tax receipt for my donation?": "Yes, we provide tax receipts for all donations made through our platform, which you can use for tax purposes.",
    "Who decides how donations are used?": "Each cause is managed by a dedicated team, and funds are allocated directly to their respective projects.",
    "How do I stop my recurring donation?": "You can manage your recurring donation settings by logging into your account and updating your donation preferences.",
};

async function sendMessage() {
    const userInput = document.getElementById("user-input").value;
    if (userInput.trim() !== "") {
        appendMessage(userInput, 'user');
        document.getElementById("user-input").value = '';

        // Check if we have a predefined response
        let botResponse = responses[userInput] || await getGeminiAPIResponse(userInput);
        
        setTimeout(() => appendMessage(botResponse, 'bot'), 500);
    }
}

function appendMessage(message, sender) {
    const chatBox = document.getElementById("chat-box");
    const messageDiv = document.createElement("div");
    messageDiv.classList.add(sender + "-message");
    messageDiv.innerHTML = `<p>${message}</p>`;
    chatBox.appendChild(messageDiv);

    chatBox.scrollTop = chatBox.scrollHeight; // Auto-scroll to the bottom
}

function toggleChatWindow() {
    const chatContainer = document.getElementById("chat-container");
    chatContainer.style.display = chatContainer.style.display === "none" ? "flex" : "none";
}

function closeChatWindow() {
    const chatContainer = document.getElementById("chat-container");
    chatContainer.style.display = "none";
}

async function getGeminiAPIResponse(userInput) {
    try {
        const response = await fetch('https://api.gemini.com/v1/ask', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer AIzaSyBaZmlDzLcFkhWVI_r43W010OQHFu5c8CM` // Replace with your actual API key
            },
            body: JSON.stringify({ query: userInput })
        });

        const data = await response.json();
        if (data && data.answer) {
            return data.answer;  // Return the answer from Gemini API
        } else {
            return "I'm sorry, I couldn't find an answer to that. Can you rephrase?";
        }
    } catch (error) {
        console.error("Error fetching response from Gemini API:", error);
        return "Sorry, I encountered an error while fetching the answer.";
    }
}

function askQuestion(question) {
    document.getElementById("user-input").value = question;
    sendMessage();
}

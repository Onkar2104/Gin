document.getElementById("signup-form").addEventListener("submit", async function (event) {
    event.preventDefault(); // Prevent form from reloading

    const formData = {
        first_name: document.getElementById("first_name").value,
        last_name: document.getElementById("last_name").value,
        email: document.getElementById("email").value,
        password: document.getElementById("password").value,
        confirm_password: document.getElementById("confirm_password").value
    };

    const messageBox = document.getElementById("message-box");
    messageBox.style.display = "none"; // Hide message before request

    try {
        const response = await fetch("/api/signup", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(formData)
        });

        const result = await response.json();
        messageBox.style.display = "block";

        if (response.ok) {
            // Success message
            messageBox.style.backgroundColor = "#d4edda"; // Light green
            messageBox.style.color = "#155724"; // Dark green text
            messageBox.innerHTML = "✅ " + (result.message || "Signup successful!");
        } else {
            // Error message
            messageBox.style.backgroundColor = "#f8d7da"; // Light red
            messageBox.style.color = "#721c24"; // Dark red text
            messageBox.innerHTML = "❌ " + (result.error || "Signup failed!");
        }
    } catch (error) {
        messageBox.style.display = "block";
        messageBox.style.backgroundColor = "#f8d7da"; // Light red
        messageBox.style.color = "#721c24"; // Dark red text
        messageBox.innerHTML = "❌ An unexpected error occurred.";
    }
});

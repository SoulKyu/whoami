document.addEventListener("DOMContentLoaded", function() {
    var loginButton = document.getElementById("login-button");
    loginButton.addEventListener("click", function(event) {
        event.preventDefault(); // Empêche le comportement par défaut du bouton de soumettre un formulaire

        var username = document.getElementById("username").value;
        var password = document.getElementById("password").value;

        // Envoie les informations de connexion au serveur
        // Tu devras remplacer cette logique par ton propre code d'authentification côté serveur

        // Exemple simple d'affichage du résultat de la connexion
        if (username === "admin" && password === "password") {
            alert("Connexion réussie !");
        } else {
            alert("Échec de la connexion. Veuillez vérifier vos informations.");
        }
    });
});
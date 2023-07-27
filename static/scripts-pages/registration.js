const inputPass = document.getElementById('form-pass');
const loginBar = document.getElementById('form-login');
const iconPass = document.getElementById('pass-icon');

inputPass.addEventListener('click', () => {
    document.getElementById("pass-icon").style.opacity = "1";
});

iconPass.addEventListener('mousemove', () => {
    document.getElementById("pass-icon").style.opacity = "1";
});

iconPass.addEventListener('mouseout', () => {
    document.getElementById("pass-icon").style.opacity = "0";
});

inputPass.addEventListener('blur', () => {
    document.getElementById("pass-icon").style.opacity = "0";
});

iconPass.addEventListener('click', () => {
    if (inputPass.getAttribute('type') === "password"){
        inputPass.setAttribute('type', 'text');
        iconPass.setAttribute('src', '../static/images/hide_password.ico')
    } else {
        inputPass.setAttribute('type', 'password')
        iconPass.setAttribute('src', '../static/images/show_password.ico')
    }
});

function registration() {

    document.getElementById("new_user_status").style.opacity = "0";

    var formData = {
        password: inputPass.value,
        username: loginBar.value
    };

    fetch("/registration", {
        method: "POST",
        headers: {
        "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        // Делайте что-то с ответом сервера, если необходимо
        if (data.redirect_url) {
            document.getElementById("new_user_status").style.opacity = 1;
            window.location.href = data.redirect_url;
        } else {
            // Делайте что-то с другими данными, если необходимо
            document.getElementById("new_user_status").style.opacity = 1;

            inputPass.value=null
            loginBar.value=null

        }
    })
    .catch(error => {
        console.error("Ошибка при отправке POST-запроса:", error);
    });
}
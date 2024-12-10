const apiUrl = "http://localhost:8080";

// Элементы DOM
const authContainer = document.getElementById("auth-container");
const registerContainer = document.getElementById("register-container");
const deviceContainer = document.getElementById("device-container");
const loginForm = document.getElementById("login-form");
const registerForm = document.getElementById("register-form");
const deviceList = document.getElementById("device-list");
const showRegister = document.getElementById("show-register");
const showLogin = document.getElementById("show-login");

// Переключение между формами
showRegister.addEventListener("click", () => {
    authContainer.style.display = "none";
    registerContainer.style.display = "block";
});

showLogin.addEventListener("click", () => {
    authContainer.style.display = "block";
    registerContainer.style.display = "none";
});

// Авторизация
loginForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    try {
        const response = await fetch(`${apiUrl}/auth/sign-in`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ username, password }),
        });

        if (!response.ok) {
            throw new Error("Неверный логин или пароль");
        }

        const data = await response.json();
        localStorage.setItem("token", data.token);
        showDevices(); // Переход на список устройств
    } catch (error) {
        alert(error.message);
    }
});

// Регистрация
registerForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const username = document.getElementById("reg-username").value;
    const password = document.getElementById("reg-password").value;

    try {
        const response = await fetch(`${apiUrl}/auth/sign-up`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ username, password }),
        });

        if (!response.ok) {
            throw new Error("Ошибка при регистрации");
        }

        alert("Регистрация успешна! Теперь вы можете войти.");
        showLogin.click(); // Возвращаемся к форме авторизации
    } catch (error) {
        alert(error.message);
    }
});

// Показать список устройств
async function showDevices() {
    authContainer.style.display = "none";
    registerContainer.style.display = "none";
    deviceContainer.style.display = "block";

    try {
        const token = localStorage.getItem("token");
        const response = await fetch(`${apiUrl}/api/devices`, {
            headers: { Authorization: `Bearer ${token}` },
        });

        if (!response.ok) {
            throw new Error("Ошибка загрузки устройств");
        }

        const devices = await response.json();
        renderDevices(devices);
    } catch (error) {
        console.error(error.message);
    }
}

// Отображение устройств
function renderDevices(devices) {
    deviceList.innerHTML = ""; // Очистка списка

    devices.forEach((device) => {
        const deviceEl = document.createElement("div");
        deviceEl.className = "device";
        deviceEl.innerHTML = `
            <span>${device.name} (${device.type}) - ${
            device.status ? "Включено" : "Выключено"
        }</span>
            <button onclick="toggleDevice('${device.id}')">
                ${device.status ? "Выключить" : "Включить"}
            </button>
        `;
        deviceList.appendChild(deviceEl);
    });
}

// Переключение состояния устройства
async function toggleDevice(deviceId) {
    try {
        const token = localStorage.getItem("token");
        const response = await fetch(`${apiUrl}/api/devices/${deviceId}/toggle`, {
            method: "POST",
            headers: { Authorization: `Bearer ${token}` },
        });

        if (!response.ok) {
            throw new Error("Ошибка переключения устройства");
        }

        showDevices(); // Обновить список устройств
    } catch (error) {
        console.error(error.message);
    }
}
